package main

import (
	"crypto/tls"
	"io"
	"net/http"
)

func doHttpRequest(urlPath, method string, body io.Reader, cookies []*http.Cookie, header http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, urlPath, body)
	if err != nil {
		return nil, err
	}
	req.Header = header
	for _, ck := range cookies {
		//fmt.Println(ck.Name,ck.Value)
		req.AddCookie(ck)
	}
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		}, // ignore expired SSL certificates
	}}
	//fmt.Println(req)
	return client.Do(req)
}
