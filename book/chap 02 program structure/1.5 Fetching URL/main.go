package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		var urlTransformed = ""
		if strings.HasPrefix(url, "http://") == false {
			urlTransformed = "http://" + url
		} else {
			urlTransformed = url
		}
		fmt.Println("url: ", urlTransformed)
		resp, err := http.Get(urlTransformed)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Println("resp.Status: ", resp.Status)
	}
}