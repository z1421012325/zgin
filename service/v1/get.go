package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"mygo/pools"
	"redigo/redis"
	"regexp"
)


// 心跳检测
func Ping(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"ping",
	})
}



const (
	//regular = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\\\d{8}$"
	regular = "^1[0-9]{10}$"
)

/**
发送短信
需要参数 电话号码
*/
func PhoneCaptcha(c *gin.Context){
	//验证
	phone := c.Query("phone")
	reg := regexp.MustCompile(regular)
	if len(phone)==0 || !reg.MatchString(phone){
		c.JSON(203,gin.H{
			"msg":"手机号码不正确",
		})
		return
	}


	// redis 验证
	// 如果在redis中寻找到该电话号号码出现的次数,如果超过5次或者几次则返回信息进行限制
	in,err := redis.Int(pools.RD.Get().Do("get",phone))
	if err != nil {
		c.JSON(203,gin.H{
			"msg":err.Error(),
		})
	}
	if (in>=5 && in<10){
		c.JSON(203,gin.H{
			"msg":"尝试次数太多,请稍后再试",
		})
		return
	}

	// 增加redis中的电话号码数值次数
	if (in+1>=5 && in<10){
		_,err = pools.RD.Get().Do("set",phone,in+1,"ex",60*60*0.5)
		if err != nil {
			log.Print("redis 存储出错")
		}
	}else {
		_,err = pools.RD.Get().Do("set",phone,in+1,"ex",60*60*0.2)
		if err != nil {
			log.Print("redis 存储出错")
		}
	}

	// send phone 第三方发送
	// TODO
	// phone+"验证码" 保存在redis中
	// 假设为123456  6位数验证码  已经发送了
	captcha := 123456
	_,err = pools.RD.Get().Do("set",phone+string(captcha),captcha,"ex",60*5)
	if err != nil {
		log.Print("redis 存储出错")
	}


	c.JSON(200,gin.H{
		"msg":"短信发送成功!",
	})


	// 测试使用sql语句返回结果
	type ceshi struct {
		Name 	string
		Eamil 	string
	}
	rows,err := pools.DB.Exec("select name,email from users where id in (?)",[]int{1,2,3,4,6,7}).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	if rows.Next(){
		var cs ceshi
		pools.DB.ScanRows(rows,&cs)
	}
	return
}



