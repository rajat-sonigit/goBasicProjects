package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://httpbin.org/get"

func main() {
	fmt.Println("Hey its an web request handling")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("your output is : ", response)

	defer response.Body.Close()
	data , err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	consat :=  string(data)
	fmt.Println(consat)
}
