package main

import (
	"fmt"
	"os"

	"github.com/leo-alvarenga/to-go/cli"
)

func main() {
	fmt.Println("ok")
	cli.Entrypoint(os.Args)
}
