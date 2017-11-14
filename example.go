package main

import (
	"github.com/5MofDream/apollo/moltencore"
	_ "github.com/5MofDream/apollo/router"
	_ "github.com/5MofDream/apollo/rpcrouter"
)

func main() {
	moltencore.Moltencore().Fire()
}
