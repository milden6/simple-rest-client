package restapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"simple-rest-client/restutil"
)

type Header struct {
	//name  string
	//value string
}

type Request struct {
	Method string
	//headers []Header
	Url string
	//body    string
}

func (r *Request) SendRequest() error {

	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.Url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", restutil.PrettifyJSON(body))
	fmt.Printf(resp.Status)

	return nil
}
