package RPC

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func RunLocalRPCSev(port string) (ExitCode int) {
	ExitCode = 1
	rpc.Register(new(Judger))

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("fatal error: \n", err)
		ExitCode = 2
	}

	fmt.Fprintf(os.Stdout, "\n%s\nRuning on localhost:%s\n", "=== SERVER START SUCCESSFUL ===\n", port)

	for {
		conn, err := lis.Accept()
		if err != nil {
			ExitCode = 3
			continue
		}

		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
	return ExitCode
}
