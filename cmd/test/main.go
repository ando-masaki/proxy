package main

import(
	"fmt"
	"io"
	"log"
	"os"
	"net/http"
	// "net/url"
	// "time"

	// "github.com/k0kubun/pp"
	"github.com/ando-masaki/proxy"
)

const URL = "http://inet-ip.info/ip"

func main() {
	client := http.DefaultClient
	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("")

	proxies, err := proxy.CyberSource()
	if err != nil {
		log.Fatal(err)
	}
	for i, proxy := range proxies {
			log.Printf("%6d/%6d\t%#v\n", i, len(proxies), proxy)
			proxy.Test(client, URL)
	}
}


