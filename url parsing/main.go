package main

import (
	"fmt"
	"net/url"
)

const urll = "https://www.mocky.io/v2/5ecfd5dc3200006200e3d64b"

func main() {
	fmt.Println("Hey welcome to the url parsing chapter")
	response, err := url.Parse(urll)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("your output is : ", response)
	fmt.Println("your Scheme is : ", response.Scheme)
	fmt.Println("your Scheme is : ", response.Host)
	fmt.Println("your Scheme is : ", response.Port())

}
