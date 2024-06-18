package comms

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	utils "github.com/cyclo9/gotils"
)

// Send an HTTP GET request
func GET(baseUrl string, params url.Values) []byte {
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	res, err := http.Get(fullUrl)
	utils.HandleErr(err, "sending GET request")

	body, err := io.ReadAll(res.Body)
	utils.HandleErr(err, "reading body message")
	return body
}
