package http

/*
Support for HTTP clients is included out the box in the net/http packages
*/

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue HTTP get
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print HTTP response
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	// Print first 5 lines of response body
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
