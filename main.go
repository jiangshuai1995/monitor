package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"monitor/common"
	"monitor/initialization"
	"monitor/middlewares"
	"monitor/utils"
)


func main(){
	conf:="conf"
	initialization.SetConfig(conf)
	initialization.SetDB()
	app.Use(middlewares.Crs())
	app.Get("/version",Versionhandler)
	app.Get("/log",Loghandler)

	app.Run(iris.Addr(":8080"))
}

//example: 127.0.0.1:8080/version
func Versionhandler(ctx iris.Context){
	rows,err:=common.DB.Query("SELECT version()")
	var status string
	if utils.CheckErr(err) {
		status="ERROR"
	}else{
		status="OK"
	}
	var result string
	for rows.Next(){
		err= rows.Scan(&result)
		utils.CheckErr(err)
	}
	ctx.JSON(iris.Map{
		"version":&result,
		"status":&status,
	})
	/*{
		"status": "OK",
		"version": "PostgreSQL 9.4.20 (Greenplum Database 6.0.0-beta.1 build dev) on x86_64-unknown-linux-gnu, compiled by gcc (GCC) 4.8.5 20150623 (Red Hat 4.8.5-36), 64-bit compiled on Mar 21 2019 13:12:25"
	}*/
}

//example: 127.0.0.1:8080/log?level=WARNING&starttime=2019-04-27 00:50&endtime=2019-04-27 00:55
func Loghandler(ctx iris.Context){
	 level := ctx.URLParam("level")
	 starttime := ctx.URLParam("starttime") //format yyyy-mm-dd
	 endtime := ctx.URLParam("endtime")
	 var query =fmt.Sprintf("SELECT logtime,logseverity,logstate,logmessage from  public.log_alert_history where logseverity='%s' AND logtime>'%s' AND logtime < '%s' limit 2;",level,starttime,endtime)
	 row,err:=common.DB.Query(query)
	 var status string
	 if utils.CheckErr(err){
	 	status = "ERROR"
	 }else{
	 	status = "OK"
	 }
	 var result []common.Log

	for row.Next(){
		l:= &common.Log{}
		err = row.Scan(&l.Logtime,&l.Logseverity,&l.Logstate,&l.Logmessage)
		utils.CheckErr(err)
		result=append(result,*l)
	}

	 data,_:=json.Marshal(result)

	 var strdata = string(data)


	 ctx.JSON(&iris.Map{
		 "version": &strdata,
		 "status":  &status,
	 })
	/*{
		"status": "OK",
		"version": "[{\"Logtime\":\"2019-04-27T00:52:28.760062+08:00\",\"Logseverity\":\"WARNING\",\"Logstate\":\"01000\",\"Logmessage\":\"time constraints added on superuser role\"},{\"Logtime\":\"2019-04-27T00:52:28.987812+08:00\",\"Logseverity\":\"WARNING\",\"Logstate\":\"01000\",\"Logmessage\":\"time constraints added on superuser role\"}]"
	}*/

	 //ctx.WriteString(strdata)
	 // [{"Logtime":"2019-04-27T00:52:28.680727+08:00","Logseverity":"WARNING","Logstate":"01000","Logmessage":"time constraints added on superuser role"},{"Logtime":"2019-04-27T00:52:28.826267+08:00","Logseverity":"WARNING","Logstate":"01000","Logmessage":"time constraints added on superuser role"}]

}
