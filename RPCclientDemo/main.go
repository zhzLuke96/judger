package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// Judger program Request
type JudgerRequest struct {
	Code         string
	LangType     string
	TestCaseJSON string
}

// Judger Ret
type JudgerResponse struct {
	JSONContent string
}

func RPCCall(req JudgerRequest, port string) (res *JudgerResponse) {
	conn, err := jsonrpc.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	err = conn.Call("Judger.DoJudger", req, &res)
	if err != nil {
		log.Fatalln("Judger Error: ", err)
	}
	return res
}

var jsCode = `const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.on('line', function (input) {
    console.log(eval(input));
    rl.close();
});

rl.on('close', function () {
    process.exit(0);
});`

var pythonCode = "print(eval(input()))"

func main() {
	//
	// res := RPCCall(JudgerRequest{
	// 	Code:         pythonCode,
	// 	LangType:     "python",
	// 	TestCaseJSON: `{"timeout":500,"mem":1024,"data":[["1 + 2","3"],["9 - 8","1"],["100000 + 100000","200000"],["7 / 8","0.875"],["4 << 2","16"],["8 % 5","3"]]}`},
	// 	"8088")
	// fmt.Printf("[LOG] res = %v\n", res)

	res := RPCCall(JudgerRequest{
		Code:         jsCode,
		LangType:     "javascript",
		TestCaseJSON: `{"timeout":500,"mem":1024,"data":[["1 + 2","3"],["9 - 8","1"],["100000 + 100000","200000"],["7 / 8","0.875"],["4 << 2","16"],["8 % 5","3"]]}`},
		"8088")
	fmt.Printf("[LOG] res = %v\n", res)
}
