package comms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cyclo9/gotils"
)

// Send an HTTP GET request
func GET(baseUrl string, params url.Values) []byte {
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	res, err := http.Get(fullUrl)
	gotils.HandleErr(err, "sending GET request")

	body, err := io.ReadAll(res.Body)
	gotils.HandleErr(err, "reading body message")
	return body
}

func POST(baseUrl string, params url.Values, data map[string]any) []byte {
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	dataBytes, err := json.Marshal(data)
	gotils.HandleErr(err, "marshalling data")

	res, err := http.Post(fullUrl, "application/json", bytes.NewBuffer(dataBytes))
	gotils.HandleErr(err, "making POST request")

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	gotils.HandleErr(err, "reading body message")
	return body
}
