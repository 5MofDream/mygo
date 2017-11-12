package main

import (
	"apollo/moltencore"
	_ "apollo/router"
	_ "apollo/rpcrouter"
)

func main() {
	moltencore.Moltencore().Fire()
}
