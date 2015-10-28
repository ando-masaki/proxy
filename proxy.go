package proxy

import (
	"fmt"
	"io"
	"net"
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
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.80 Safari/537.36")
	resp, err := client.Do(req)
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
	return &http.Transport{
		Proxy: http.ProxyURL(URL),
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}, nil
}

func (p Proxy) String() string {
	return fmt.Sprintf("%s://%s:%s", p.Scheme, p.IP, p.Port)
}
