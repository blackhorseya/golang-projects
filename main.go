package main

import (
	"flag"
	"fmt"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "bot token")
	flag.Parse()
}

func main() {
	fmt.Println(token)
}
