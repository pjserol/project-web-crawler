package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	url := os.Args[1]
	if url == "" {
		log.Println("you must provide a url")
		os.Exit(1)
	}

	depth, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println("you must provide an integer for the depth")
	}

	node, err := webCrawler(url, depth)
	if err != nil {
		log.Println(err)
	}

	res, err := json.Marshal(node)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(res))
}
