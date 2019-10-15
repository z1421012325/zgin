package main

import (
	"fmt"
	_ "mygo/config"				// 加载配置文件
	_ "mygo/pools"				// 加载连接池

	_ "mygo/server"	// server开启 默认开启7999 配置文件修改
)



func main(){
	fmt.Println("开启服务!")
}
