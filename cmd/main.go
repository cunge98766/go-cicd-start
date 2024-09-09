package main

import(
"fmt" 
"go-cicd-start/gin"
)

func main() {  
    fmt.Println("This is main in packageB")  
		//创建服务
		ginService.InitGin()
		//运行服务后，啥都不执行了，后面就不执行了
		fmt.Println("This is main in packageB end11111")  
}