package main

import (
	"os"
	"ozone/internal/pkg/config"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	//

	config.MustLoad()
	return nil

}
