package RPC

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"
)

func RunLocalRPCSev(port string) (ExitCode int) {
	ExitCode = 1
	rpc.Register(new(Judger))

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("fatal error: \n", err)
		ExitCode = 2
	}

	fmt.Fprintf(os.Stdout, "\n%s\n=== Runing on localhost:%s\n\n", "=== SERVER START SUCCESSFUL ===", port)

	for {
		conn, err := lis.Accept()
		if err != nil {
			ExitCode = 3
			continue
		}

		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s %s %s\n", time.Now().Format("2006/1/2 15:04:05"), conn.RemoteAddr().String(), "New connection initialized.")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
	return ExitCode
}
