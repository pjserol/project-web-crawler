package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type node struct {
	Title string  `json:"title"`
	URL   string  `json:"url"`
	Nodes []*node `json:"nodes"`
}

func webCrawler(url string, depth int) (node, error) {
	parent := node{
		URL: url,
	}

	err := analyseDocument(url, &parent, depth)
	if err != nil {
		return node{}, err
	}

	return parent, nil
}

func analyseDocument(url string, newNode *node, depth int) error {
	if depth == 0 {
		return nil
	} else {
		depth--
	}

	doc, err := getDocument(url)
	if err != nil { // continue the process to get the other children
		log.Println(err)
		return nil
	}

	newNode.Title = getTitleDocument(doc)
	links := getLinksDocument(nil, doc)

	children := []*node{}
	for _, link := range links {
		fURL := formatURL(url, link)

		newChild := node{
			URL: fURL,
		}

		children = append(children, &newChild)

		analyseDocument(fURL, &newChild, depth)
	}

	newNode.Nodes = children

	return nil

}

func getDocument(url string) (*html.Node, error) {
	body, err := getBodyFromURL(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return doc, nil
}

func getBodyFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("wrong status code")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil
}

func getLinksDocument(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = getLinksDocument(links, child)
	}

	return links
}

func getTitleDocument(node *html.Node) string {
	var title string

	if node.Type == html.ElementNode && node.Data == "title" {
		if node.FirstChild != nil {
			return node.FirstChild.Data
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		title = getTitleDocument(child)
		if title != "" {
			break
		}
	}

	return title
}

func formatURL(base, url string) string {
	base = strings.TrimSuffix(base, "/")
	switch {
	case strings.HasPrefix(url, "/"):
		return base + url
	case strings.HasPrefix(url, "#"):
		return base + "/" + url
	default:
		return url
	}
}
