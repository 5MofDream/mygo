package rpcrouter

import (
	"github.com/5MofDream/apollo/moltencore"
	"github.com/5MofDream/apollo/rpcnodes"
)

func init() {
	rxs := moltencore.Moltencore().RpcxServer()
	rxs.RegisterNode("Arith", new(rpcnodes.Arith), "")
}
