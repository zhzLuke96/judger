package main

import (
	"fmt"
	"os"
)

func clearCompliceFile() (err error) {
	os.RemoveAll("./__pycache__/")
	return nil
}

func main() {

	fmt.Println("THX 4 USEING!")
}
