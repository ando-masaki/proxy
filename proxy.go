package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Proxy struct {
	Scheme   string
	IP       string
	Port     string
	ConnTime time.Duration
}

func (p *Proxy) Test(client *http.Client, URL string) error {
	transport, err := p.Transport()
	if err != nil {
		return err
	}
	client.Transport = transport
	resp, err := client.Get(URL)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("")
	return nil
}

func (p *Proxy) Transport() (*http.Transport, error) {
	URL, err := url.Parse(p.String())
	if err != nil {
		return nil, fmt.Errorf("can't parse proxy url %s,%v", p.String(), err)
	}
	return &http.Transport{Proxy: http.ProxyURL(URL)}, nil
}

func (p Proxy) String() string {
	return fmt.Sprintf("%s://%s:%s", p.Scheme, p.IP, p.Port)
}
