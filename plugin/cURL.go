package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type Responses struct {
	Status  int
	Header  http.Header
	Body    map[string]interface{}
	Payload io.Reader
}

func Body(payload interface{}) io.Reader {
	payloadBytes, _ := json.Marshal(payload)
	body := bytes.NewReader(payloadBytes)

	return body
}

func Init(method, url string, header map[string]string, body io.Reader) (Responses, error) {

	var n Responses
	req, _ := http.NewRequest(method, url, body)
	if v := reflect.ValueOf(header); v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			req.Header.Set(key.String(), strct.String())
		}
	}
	fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		n.Status = http.StatusInternalServerError
		n.Header = nil
		n.Body = nil
		return n, err
	}
	n.Status = resp.StatusCode
	n.Header = resp.Header

	json.NewDecoder(resp.Body).Decode(&n.Body)
	resp.Body.Close()
	return n, nil
}
