package parse

import (
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func GetData(url string) (*html.Node, error) {
	return readDataFromWebsite(url)
	//return readDataFromFile(url)

}

func readDataFromFile(filepath string) (*html.Node, error) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return html.Parse(file)
}

func readDataFromWebsite(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}
	return html.Parse(resp.Body)
}
