package config

import (
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)

var config *ini.File

func init() {
	c, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}

	config = c
}

func Instance() *ini.File {
	return config
}

func Get(key string) string {
	needle := "."
	if strings.Contains(key, needle) {
		keyArr := strings.Split(key, needle)
		return config.Section(keyArr[0]).Key(keyArr[1]).String()
	}
	return ""
}

func GetToBool(key string) bool {
	rBool, _ := strconv.ParseBool(Get(key))
	return rBool
}
