package admin

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/wencan/fastrest/resterror"
	"github.com/wencan/go-service-demo/client/mydb/admindb"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/client/rds"
	"github.com/wencan/go-service-demo/model/adminmodel"
	"github.com/wencan/go-service-demo/model/protocolmodel"
)

// TokenTTL  Token生存时间。
const TokenTTL time.Duration = time.Minute * 5

// UserBusiness admin用户业务逻辑。
type UserBusiness struct {
	// DBx MySQL对象。
	DBx dbinterface.Execer

	// Rds redis客户端。
	Rds *redis.Client
}

// CreateUser 创建用户。
func (userBusiness UserBusiness) CreateUser(ctx context.Context, req protocolmodel.CreateUserRequest) (*protocolmodel.CreateUserResponse, error) {
	// 计算摘要
	salt := uuid.NewString()
	sum := sha256.Sum256([]byte(req.Password + salt))
	passwdSha256 := fmt.Sprintf("%X", sum)

	user := &adminmodel.User{
		UserName:       req.UserName,
		NiceName:       req.NiceName,
		PasswordSha256: passwdSha256,
		SecureSalt:     salt,
	}
	id, ok, err := admindb.InsertUser(ctx, userBusiness.DBx, *user)
	if err != nil {
		return nil, err
	}
	if !ok { // username 冲突
		return nil, resterror.ErrorWithStatus(fmt.Errorf("username conflict"), resterror.StatusAlreadyExists)
	}

	user, _, err = admindb.GetUser(ctx, userBusiness.DBx, id)
	if err != nil {
		return nil, err
	}

	return &protocolmodel.CreateUserResponse{
		User: *user,
	}, nil
}

// Login 登录。
func (userBusiness UserBusiness) Login(ctx context.Context, req protocolmodel.LoginRequest) (*protocolmodel.LoginResponse, error) {
	// 根据username，查询用户信息，得到salt
	user, _, err := admindb.GetUserByUserName(ctx, userBusiness.DBx, req.UserName)
	if err != nil {
		return nil, err
	}
	if user == nil { // 无效的username
		return nil, resterror.ErrorWithStatus(fmt.Errorf("invalid username"), resterror.StatusUnauthenticated)
	}

	// 计算摘要
	sum := sha256.Sum256([]byte(req.Password + user.SecureSalt))
	passSha256 := fmt.Sprintf("%X", sum)

	// 获得用户信息
	user, _, err = admindb.GetUserForLogin(ctx, userBusiness.DBx, req.UserName, passSha256)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, resterror.ErrorWithStatus(fmt.Errorf("failed in login"), resterror.StatusUnauthenticated)
	}

	// 应该踢掉现有token

	// 生成token
	token := uuid.NewString()
	err = rds.SaveToken(ctx, userBusiness.Rds, token, user.ID, TokenTTL+time.Second+10)
	if err != nil {
		return nil, fmt.Errorf("failed in save token, error: [%w]", err)
	}

	return &protocolmodel.LoginResponse{
		Token:      token,
		ExpireTime: time.Now().Add(TokenTTL),
	}, nil
}

// UserInfoByToken 根据token，获得用户信息，包括权限。
func (userBusiness UserBusiness) UserInfoByToken(ctx context.Context, req protocolmodel.UserInfoByTokenRequest) (*protocolmodel.UserInfoByTokenResponse, error) {
	userID, _, err := rds.UserByToken(ctx, userBusiness.Rds, req.Token)
	if err != nil {
		return nil, fmt.Errorf("failed in get user by token, error: [%w]", err)
	}
	if userID == 0 {
		return nil, resterror.ErrorWithStatus(fmt.Errorf("failed to authentication"), resterror.StatusUnauthenticated)
	}

	user, _, err := admindb.GetUser(ctx, userBusiness.DBx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed in get user %d, error: [%s]", userID, err)
	}
	if user == nil {
		return nil, resterror.ErrorWithStatus(fmt.Errorf("failed to authentication"), resterror.StatusUnauthenticated)
	}

	userRoleIDs, err := admindb.GetRoleIDsByUser(ctx, userBusiness.DBx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed in get role by user %d, error: [%w]", userID, err)
	}
	permissionIDs, err := admindb.GetPermissionIDsByRoles(ctx, userBusiness.DBx, userRoleIDs)
	if err != nil {
		return nil, fmt.Errorf("failed in get permission ids by role %v, error: [%w]", userRoleIDs, err)
	}
	permissionCodes, err := admindb.GetPermissionCodes(ctx, userBusiness.DBx, permissionIDs)
	if err != nil {
		return nil, fmt.Errorf("failed in get permission code by role %v, error: [%w]", userRoleIDs, err)
	}

	return &protocolmodel.UserInfoByTokenResponse{
		User:            user,
		PermissionCodes: permissionCodes,
	}, nil
}
