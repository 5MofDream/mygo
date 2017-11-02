package rcpclient

import (
	"context"
	"log"
	"apollo/conf"
	example "apollo/rpcnodes"
	"github.com/smallnest/rpcx/client"
	"apollo/moltencore"
	"apollo/lib"
)

var config *conf.ConfigImp

func init() {
	config = moltencore.Moltencore().YamlConf()
}

//client example
func ExampleArithMul(a int ,b int) int {
	addr2, error := config.Get("example_cli_arith")
	lib.PanicError(error)
	d := client.NewPeer2PeerDiscovery("tcp@"+addr2, "")
	xclient := client.NewXClient("Arith", "Mul", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	//args := &example.Args{
	//	A: 10,
	//	B: 20,
	//}
	args := &example.Args{
		A: a,
		B: b,
	}
	reply := &example.Reply{}
	call, err := xclient.Go(context.Background(), args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
	return reply.C
}
