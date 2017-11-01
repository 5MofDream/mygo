package rpcrouter

import (
	"apollo/moltencore"
	"apollo/rpcnodes"
)

func init() {
	rxs := moltencore.Moltencore().RpcxServer()
	rxs.RegisterNode("Arith", new(rpcnodes.Arith), "")
}
