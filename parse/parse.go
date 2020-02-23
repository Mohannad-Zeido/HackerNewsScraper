package parse

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strconv"
)

func GetData(page int, state types.RunState) (*html.Node, error) {

	if state.TestState {
		return readDataFromFile(state.TestFile + strconv.Itoa(page) + ".html")
	}
	return readDataFromWebsite("https://news.ycombinator.com/news?p=" + strconv.Itoa(page))

}

func readDataFromFile(filepath string) (*html.Node, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return html.Parse(file)
}

func readDataFromWebsite(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}
	return html.Parse(resp.Body)
}
