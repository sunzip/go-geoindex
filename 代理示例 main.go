package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"uav/account/controller"
	"uav/common/jwt"
	"uav/common/tools/form"

	// "uav/account/form"
	"uav/account/service"
	"uav/common/api"
	"uav/common/enum"
	"uav/common/model"
	"uav/common/tools"

	"git.hiscene.net/gokit/grpc-tool/client"
	"git.hiscene.net/gokit/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	// validator "github.com/go-playground/validator.v8"
)

func main() {
	/* err := config.Load()
	if err != nil {
		fmt.Printf("init config failed, err is %v\n", err)
		return
	} */
	/* c := flag.String("c", "", "config path")
	flag.Parse() */
	/* if len(*c) > 0 {
		fmt.Printf("%s %s\n", *c, *c)
		return
	} */
	tools.GlobalInit()
	jwt.DBClient.RedisConnect(tools.GetServerConfig().RedisInfo)

	// init logger
	if err := logger.InitGlobal(logger.LogOpt{
		Level:  tools.ServerConfigVar.Server.LogLevel,
		Caller: true,
	}); err != nil {
		log.Fatal("init logger error:", err)
	}
	defer logger.Sync()

	// grpc connection
	service.RegistGRPC()
	// close grpc connections
	defer client.CloseAll()

	gin.SetMode(tools.ServerConfigVar.Server.RunMode)

	if err := form.TransInit("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// redis 初始化
	controller.DBClient.RedisConnect(tools.GetServerConfig().RedisInfo)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("nameValidator", form.NameValidator)
		v.RegisterValidation("emailValidator", form.EmailValidator)
		v.RegisterValidation("mobileValidator", form.MobileValidator)
		v.RegisterValidation("idNumberValidator", form.IdNumberValidator, false)
	}
	e := gin.Default()
	ginRouter := e.Use(controller.CheckLicense, api.CacheBody)
	ginRouter.POST("/api/v1.0/account/login", controller.Signin) // 账号密码登录

	/* account := ginRouter.Group("/account")
	{
		// 用户管理
		// account.Use(controller.CheckToken)
		// account.POST("/login", controller.Signin)   // 账号密码登录
		account.POST("/create", controller.UserAdd) // 账号添加
	} */
	/* userRouter := ginRouter.Group("/user")
	{
		// 用户管理
		userRouter.Use(controller.CheckToken, controller.CheckPrivlege)
		userRouter.POST("/add", controller.UserAdd) // 添加企业下用户

	} */

	vi := e.Group("/api/v1.0/platform")
	{
		vi.Use(api.CheckJwtToken)
		vi.Any("/*action", WithHeader)
	}
	moduleNameCn := "账户服务"
	tools.UriTitleMap["/api/v1.0/account/login"] = model.LogUrlInfo{UrlTitle: moduleNameCn + "登陆", UrlLogType: enum.LogLogin}
	tools.UriTitleMap["/api/v1.0/platform/enterprise/v0.1/organiztion/searchOrg"] = model.LogUrlInfo{UrlTitle: moduleNameCn + "部门查看", UrlLogType: enum.LogView}
	tools.UriTitleMap["/api/v1.0/platform/enterprise/v0.1/privilege/addGroupUser"] = model.LogUrlInfo{UrlTitle: moduleNameCn + "角色加人", UrlLogType: enum.LogCreate}
	tools.UriTitleMap["/api/v1.0/platform/enterprise/v0.1/privilege/deleteGroupUser"] = model.LogUrlInfo{UrlTitle: moduleNameCn + "角色移人", UrlLogType: enum.LogDelete}
	tools.UriTitleMap["/api/v1.0/platform/enterprise/v0.1/account/signout"] = model.LogUrlInfo{UrlTitle: moduleNameCn + "退出登陆", UrlLogType: enum.LogLogout}
	// todo:添加账户管理、权限管理的相关接口
	e.Run(tools.ServerConfigVar.Server.Port)
}

var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		host := tools.ServerConfigVar.Platform.Host
		url := strings.Replace(req.RequestURI, "v1.0/platform/", "", -1)
		// /api/enterprise/v0.1 这个是正常的api
		if url == "/api/enterprise/v0.1/user/export" {
			req.Method = "GET"
		} /* else if !strings.Contains(url, "/api/enterprise") { //处理头像
			url = strings.Replace(url, "/api", "", -1)
			// req.Method = "POST"
		} */
		req.URL.Scheme = "http"
		req.URL.Host = host
		req.URL.Path = url
		req.RequestURI = url
		req.Host = host
	},
	ModifyResponse: func(r *http.Response) error {
		b, _ := ioutil.ReadAll(r.Body)
		buf := bytes.NewBufferString("Monkey")
		buf.Write(b)
		r.Body = ioutil.NopCloser(buf)
		r.Header["Content-Length"] = []string{fmt.Sprint(buf.Len())}
		return nil
	},
}
var ModuleName = "account"

func WithHeader(c *gin.Context) {
	cookieName := tools.ServerConfigVar.Platform.AuthTokenCookie
	// ctx.Request.Header.Add("requester-uid", "id")
	// token, err := ctx.Cookie(tools.ServerConfig.InfoHik.Platform.AuthTokenCookie)
	userClaims := api.GetUserClaims(c)
	c.Request.Header[cookieName] = []string{userClaims.Token}
	// controller.SetLoginCookie(c, userClaims.Token, 100000)
	maxAge := tools.ServerConfigVar.Platform.CookieMaxAge
	cookie1 := &http.Cookie{Name: cookieName, Value: userClaims.Token, MaxAge: maxAge, HttpOnly: true}
	c.Request.AddCookie(cookie1)

	simpleHostProxy.ServeHTTP(c.Writer, c.Request)
}
