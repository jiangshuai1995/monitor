package utils

import "log"

func CheckErr(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return false
}
