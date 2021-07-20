package conf

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const Port = ":1234"

// Authentication 用于实现用户名和密码的认证
type Authentication struct {
	User     string
	Password string
}

// GetRequestMetadata 返回的认证信息包装user和password两个信息
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

// RequireTransportSecurity 方法return false表示不要求底层使用安全链接
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx) // 从ctx上下文中获取元信息
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	// 取出相应的的认证信息进行认证
	var appID string
	var appKey string

	if val, ok := md["user"]; ok {
		appID = val[0]
	}
	if val, ok := md["password"]; ok {
		appKey = val[0]
	}
	// 认证失败，返回一个odes.Unauthenticated类型的错误
	if appID != "gopher" || appKey != "password" {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}
