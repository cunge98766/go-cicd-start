package ginService

//后端服务

import(
"fmt"
"github.com/gin-gonic/gin"
_ "go-cicd-start/docs"
swaggerfiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"
)

func InitGin(){
	fmt.Println("This is main in packageB111")  
	ginService := gin.Default()

	ginService.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//
	ginService.GET("/json", Aaccc)

	ginService.Run(":8082")
}


// @Tags 测试用
// @Router /json [get]
// @Description 描述，我是干嘛的
func Aaccc(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "hellp"})
}