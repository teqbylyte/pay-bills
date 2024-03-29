package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"sort"
)

type HttpClient struct{}

func (h HttpClient) Get(url string, headers map[string]string) (*resty.Response, error) {
	fmt.Println("Making request to:: ", url)

	client := resty.New()
	res, err := client.R().
		EnableTrace().
		SetHeaders(headers).
		Get(url)

	if err != nil {
		fmt.Println("error in send req: ", err)
		return nil, err
	}

	return res, nil
}

func (h HttpClient) Post(url string, headers map[string]string, body any) (*resty.Response, error) {
	fmt.Println("Making request to:: ", url)

	// request payload
	var payload map[string]interface{}
	b, err := json.Marshal(body)
	_ = json.Unmarshal(b, &payload)
	fmt.Println("Payload: ", string(b))

	// sort payload for hashing
	keys := make([]string, 0, len(payload))
	for k := range payload {
		if k == "meta" {
			continue
		}

		keys = append(keys, k)
	}
	sort.Strings(keys)

	// compute request hash
	payloadString := ""
	for _, key := range keys {
		payloadString += fmt.Sprintf("%s=%v;", key, payload[key])
	}
	payloadString = payloadString[:len(payloadString)-1] + headers["api-key"]

	crypto := md5.New()
	crypto.Write([]byte(payloadString))
	hash := hex.EncodeToString(crypto.Sum(nil))
	headers["teqilla-hash"] = hash

	client := resty.New()

	res, err := client.R().
		EnableTrace().
		SetHeaders(headers).
		SetBody(body).
		Post(url)

	if err != nil {
		fmt.Println("error in send req: ", err)
		return nil, err
	}

	return res, nil
}
