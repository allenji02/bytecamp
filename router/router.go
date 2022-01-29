package router

import (
	"bytedance/auth"
	"bytedance/member"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	// 成员管理
	g.POST("/member/create", member.Create)
	g.GET("/member", member.GetMember)
	g.GET("/member/list")
	g.POST("/member/update")
	g.POST("/member/delete")

	// 登录

	g.POST("/auth/login", auth.Login)
	g.POST("/auth/logout", auth.Logout)
	g.GET("/auth/whoami", auth.Whoami)

	// 排课
	g.POST("/course/create")
	g.GET("/course/get")

	g.POST("/teacher/bind_course")
	g.POST("/teacher/unbind_course")
	g.GET("/teacher/get_course")
	g.POST("/course/schedule")

	// 抢课
	g.POST("/student/book_course")
	g.GET("/student/course")

}