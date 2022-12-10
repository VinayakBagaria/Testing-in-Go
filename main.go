package main

import (
	cli "testing-go/poker/cmd/cli"
	poker "testing-go/poker/cmd/webserver"
)

const TYPE = "web"

func main() {
	if TYPE == "web" {
		poker.StartPlayerServer()
	} else if TYPE == "cli" {
		cli.StartCLIServer()
	}
}
