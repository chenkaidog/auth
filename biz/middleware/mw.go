package middleware

import (
	"auth/biz/middleware/accesslog"
	"auth/biz/middleware/cors"
	"auth/biz/middleware/recovery"
	"auth/biz/middleware/session"
	"auth/biz/middleware/trace"

	"github.com/cloudwego/hertz/pkg/app"
)

func Suite() []app.HandlerFunc {
	return []app.HandlerFunc{
		recovery.New(),  // panic handler
		trace.New(),     // 链路ID
		accesslog.New(), // 接口日志
		cors.New(),      // 跨域请求
		session.New(),   // 会话
	}
}
