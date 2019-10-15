package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygo/pools"
)


type registeruser struct {
	name string 			`form:name json:name bind:required`
	password string  		`form:password json:password bind:required`
	phone string			`form:phone json:phone bind:required`
	phonecaptcha string 	`form:captcha json:captcha bind:required"`
}

// 手机号码注册
func UserRegister(c *gin.Context){
	// 验证
	var user registeruser
	err := c.ShouldBind(&user)
	fmt.Println(user)
	if err != nil {
		c.JSON(203,gin.H{
			"msg":"参数不完整!",
			"err":err.Error(),
		})
		return
	}

	// captcha数据库验证 一般在redis中验证
	_,err = pools.RD.Get().Do("get",user.phone+user.phonecaptcha)
	if err != nil {
		c.JSON(203,gin.H{
			"msg":"验证码不正确!",
			"err":err.Error(),
		})
		return
	}
	// mysql数据库查询是否存在


	// 不存在则存入数据库中

	//返回结果
	c.JSON(201,gin.H{
		"msg":"用户注册成功!",
		"data":user.name + user.password,
	})
}
