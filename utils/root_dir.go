package utils

import "os"

func GetRootDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic("cannot get working dir: " + err.Error())
	}
	return wd
}
