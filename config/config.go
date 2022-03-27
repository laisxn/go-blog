package config

import "gopkg.in/ini.v1"

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
