package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"

)
//自定义中间件
/*func myHandler(gin.HandlerFunc)  {
	return func(context*gin.Context) {
		//通过自定义的中间件，设定的值，后续只要调用这个中间件的都可以拿到这里的参数
		context.Set("usersession","userinfo_1")
		context.Next()//允许
		context.Abort()//不允许
	}
}*/
func main()  {
	//创建一个服务
   ginServer := gin.Default()
   ginServer.Use(favicon.New("./twt.png"))
   //增加初始环境

	//访问地址，处理请求  request response
	//接收请求，响应数据  返回函数
	//RestFul
   /*ginServer.GET("/user", func(context *gin.Context) {
     context.JSON(200,gin.H{"message":"返回给前端的信息"})
   })
  ginServer.POST("./user", func(context *gin.Context) {
	   context.JSON(200,gin.H{"message":"post.user"})
   })*/
   //响应一个页面给前端
   ginServer.LoadHTMLGlob("templates/*")//全局加载
   //ginServer.LoadHTMLFiles("templates/index.html")//加载指定文件
   ginServer.GET("/index", func(context *gin.Context) {
	   context.HTML(http.StatusOK,"index.html",gin.H{
	   	"message":"来自后台的数据",
	   })
   })

	//ginServer.LoadHTMLFiles("templates/index.html")//加载指定文件
   //加载静态资源文件
   ginServer.Static("/static","./static")
	//获取前端传值并返回给前端
	/*ginServer.GET("/userInfo", func(context *gin.Context) {
		userId:=context.Query("userId")
		userName:=context.Query("userName")
		context.JSON(http.StatusOK,gin.H{
			"userId": userId,
			"userName": userName,
		})
	})*/
	//使用RestFul风格获取前端传值并返回给前端
	/*ginServer.GET("/userInfo/:userId/:userName", func(context *gin.Context) {
		userId:=context.Param("userId")//传递单个值
		userName:=context.Param("userName")
		context.JSON(http.StatusOK,gin.H{
			"userId":userId,
			"userName":userName,
		})
	})*/
	//前端传给后端json
	ginServer.POST("/userInfo", func(context *gin.Context) {
		dataInfo,_:=context.GetRawData()//从请求体里获取对象
		var m map[string]interface{}//go语言传过来的对象一般用空接口接收
		_=json.Unmarshal(dataInfo,&m)//解析Json //指针赋值 返回结果
		context.JSON(http.StatusOK,m)//返回信息
	})
	//获取表单信息
	ginServer.POST("/user/add", func(context *gin.Context) {
		userName:=context.PostForm("userName")
		userId:=context.PostForm("userId")
		context.JSON(http.StatusOK,gin.H{
			"message":"获取成功",
			"userName":userName,
			"userId":userId,
		})
	})
	//路由重定向
	ginServer.GET("/RuiBo", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently,"https://www.baidu.com/")
	})
	//响应404页面
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound,"404.html",nil)
	})
	//路由组  /user/select/add/delete
	UserGroup:=ginServer.Group("/user")
	{
		UserGroup.GET("/select")
		UserGroup.GET("/add")
		UserGroup.GET("/delete")
	}
    
	//服务器端口
	ginServer.Run(":8082")
}