package main

import(
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/k0kubun/pp"
	"github.com/ando-masaki/proxy"
)

var URL *string = flag.String("url", "http://inet-ip.info/json/indent", "URL")

func main() {
	flag.Parse()
	proxies, err := proxy.CyberSource()
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Timeout: 10*time.Second,
	}
	for i, proxy := range proxies {
			pp.Println(fmt.Sprintf("%6d/%6d\t%#v\n", i, len(proxies), proxy))
			proxy.Test(client, *URL)
			time.Sleep(10 * time.Second)
	}
}
