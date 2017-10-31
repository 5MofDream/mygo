package main

import (
	_ "apollo/router"
	"apollo/boot"
)

func main(){
	boot.Moltencore().Fire()
}
