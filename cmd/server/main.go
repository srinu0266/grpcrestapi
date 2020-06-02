package main

import (
	"fmt"
	"os"

	"github.com/srinu0266/grpcrestapi/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
