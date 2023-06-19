package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request() {
	resp, err := http.Get("https://m.hupu.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		panic(error)
	}
	fmt.Printf("%s", body)
}
