package server

import (
	"fmt"
	"gin/gw/middleware/global"
	"gin/gw/middleware/locality"
	"gin/handler/common"
	"gin/handler/login"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type service struct {
	engine    *gin.Engine
	commonCgi *gin.RouterGroup
}

func newService(r *gin.Engine, common *gin.RouterGroup) *service {
	return &service{
		engine:    r,
		commonCgi: common,
	}
}

func InitService(r *gin.Engine) {
	// 设置全局中间件
	//commonCgi := r.Group("/v2",global.MiddleWare())
	commonCgi := r.Group("/v2")
	srv := newService(r, commonCgi)

	srv.initCommonWeb()
	srv.initTutorial()
}



func (s *service) initCommonWeb() {

	s.commonCgi.POST("/upload/single", common.UploadSingle)
	s.commonCgi.POST("/upload/multi", common.UploadMulti)
}

func (s *service) initTutorial() {
	s.commonCgi.POST("/form", common.Form)

	s.commonCgi.POST("/loginJson", login.LoginJson)
	s.commonCgi.POST("/loginForm", login.LoginForm)
	s.commonCgi.GET("/loginUri/:user/:password", login.LoginUri)

	// todo 不同的响应格式
	// 1.json
	s.commonCgi.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	// 2. 结构体响应
	s.commonCgi.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// 3.XML
	s.commonCgi.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc", "status": 200})
	})
	// 4.YAML响应
	s.commonCgi.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan", "status": 200})
	})

	// 重定向
	s.commonCgi.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// todo 同步异步
		s.commonCgi.GET("/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	s.commonCgi.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	// todo 中间件
	// 全局中间件在InitService中设置即可
	s.commonCgi.GET("/middle/global", func(c *gin.Context) {
		// 获取了中间值设置的request
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})

	// 局部中间件则可以在路由中设置
	s.commonCgi.GET("/middle/locality", locality.MiddleWare(), func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})

	// todo Cookie的获取和设置

	s.commonCgi.GET("/cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否只能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})

	s.commonCgi.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	s.commonCgi.GET("/home", global.AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	// todo session

}


