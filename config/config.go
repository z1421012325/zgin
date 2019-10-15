package config



import (
	"godotenv"
)

func init(){
	err := godotenv.Load("./config.env")
	if err != nil {
		panic(err.Error())
	}
}