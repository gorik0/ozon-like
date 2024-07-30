package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {

	dirs := strings.Split("address auth cart category comments config hub metrics middleware notifications order products profile promo recommendations search ticker utils", " ")
	_, file, _, _ := runtime.Caller(0)
	for _, dir := range dirs {

		err := os.Mkdir(filepath.Join(filepath.Dir(file), dir), 0666)
		if err != nil {
			log.Println(err)
		}

	}
}
