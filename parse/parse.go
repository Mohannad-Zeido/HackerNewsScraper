package parse

import (
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strconv"
)

func GetData(page int) (*html.Node, error) {
	return readDataFromWebsite("https://news.ycombinator.com/news?p=" + strconv.Itoa(page))
	//return readDataFromFile("data" + strconv.Itoa(page) + ".helper")

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
