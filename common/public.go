package common

import (
	"database/sql"
)

type Conf struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
	Port   string
}

type Log struct {
	Logtime string
 	Logseverity string
	Logstate string
	Logmessage string
}

var (
	C  Conf
	DB *sql.DB
)