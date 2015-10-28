package main

import (
	"log"

	"github.com/ando-masaki/proxy"
	"github.com/k0kubun/pp"
)

func main() {
	proxies, err := proxy.CyberSource()
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(proxies)
}
