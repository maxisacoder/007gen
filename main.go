package main

import (
	"flag"
	"fmt"
	"main/pk"
)

func main() {
	prefix := flag.String("prefix", "007", "前缀")
	flag.Parse()
	fmt.Println(*prefix)
	res, err := pk.GenPK("0x" + *prefix)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
