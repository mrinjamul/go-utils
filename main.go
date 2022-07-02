package main

import (
	"fmt"

	"github.com/mrinjamul/go-utils/utils"
)

func main() {
	hash, _ := utils.Sha256sum("Injamul")
	fmt.Println(hash)
}
