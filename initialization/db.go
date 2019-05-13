package initialization

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"monitor/common"
)

func SetDB(){
	var err error
	connStr :=fmt.Sprintf("user=%s dbname=%s password=%s host=%s  port=%s sslmode=disable",
		common.C.DbUser,
		common.C.DbName,
		common.C.DbPass,
		common.C.DbHost,
		common.C.DbPort)
	fmt.Println(connStr)
	DB,err:=sql.Open("postgres",connStr)
	if err != nil{
		log.Println("数据库连接出现问题...")
		panic(err)
	}
	log.Println("数据库连接成功")
	common.DB =DB


}