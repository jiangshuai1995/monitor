package initialization

import (
	"github.com/spf13/viper"
	"monitor/common"
)

func SetConfig(name string) {
	viper.SetConfigName(name)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var C common.Conf
	C.DbUser = viper.Get("db_user").(string)
	C.DbPass = viper.Get("db_pass").(string)
	C.DbHost = viper.Get("db_host").(string)
	C.DbPort = viper.Get("db_port").(string)
	C.DbName = viper.Get("db_name").(string)

	C.Port = ":" + viper.Get("port").(string)

	common.C = C
}