package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var baseUrl string = "http://tinyurl.com/api-create.php?url="

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}

	_, err := url.ParseRequestURI(os.Args[1])

	if err != nil {
		log.Fatal("Oops , Submitted URL is not valid")
		panic(err)
	}

	urlArg := os.Args[1]
	getReqUrl := baseUrl + urlArg

	fmt.Println("Thus the URL ", getReqUrl)

	response, err := http.Get(getReqUrl)

	if err != nil {
		log.Fatal("Oops, url parse did not work")
		panic(err)
	}

	defer response.Body.Close()

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal("Url Shortener did not work")
		panic(err)
	}
}
