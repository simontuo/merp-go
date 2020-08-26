package main

import (
	"fmt"

	"github.com/simontuo/merp-go/helper/password"
)

func main() {
	fmt.Println(password.Hash("11111111"))
}
