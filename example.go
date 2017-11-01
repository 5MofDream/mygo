package main

import (
	_ "apollo/router"
	"apollo/moltencore"
	_ "apollo/rpcrouter"

)

func main() {

	moltencore.Moltencore().Fire()
}
