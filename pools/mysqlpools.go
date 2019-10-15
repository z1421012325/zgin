package pools

import (
	"fmt"
	"gorm"
	"log"
	_ "mysql"
	"os"
)

var DB *gorm.DB

/**
mysql连接池
 */
func init(){
	// "user:password@/tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	name := os.Getenv("MYSQL_NAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	tbname := os.Getenv("MYSQL_DB")


	db,err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			name,password,host,port,tbname))
	if err != nil {
		panic(err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)


	db.LogMode(true)
	db.SetLogger(gorm.Logger{})
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	DB = db

}


