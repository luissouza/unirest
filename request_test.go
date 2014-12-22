package unirest

import (
	"io"
	"net/http"
	"os"
)

func ExampleRequest_Auth_doNotSendImmediately() {
	req := Get("http://httpbin.org/get").Auth("username", "p@$$w0r|>", false)
}

func ExampleRequest_Auth() {
	req := Get("http://httpbin.org/get").Auth("username", "P2ssw0rd", true)
}

func ExampleRequest_Header() {
	headers := map[string]string{
		"Accept": "application/json",
	}
	err := Get("http://httpbin.org/get").Header(headers).End(func(resp *http.Response) error {
		_, err := io.Copy(os.Stdout, resp.Body)
		return err
	})
}

func ExampleRequest_Query() {
	params := map[string]string{
		"name": "nijiko",
	}
	err := Post("http://httpbin.org/get").Query(params).End(func(resp *http.Response) {
		_, err := io.Copy(os.Stdout, resp.Body)
		return err
	})
}
