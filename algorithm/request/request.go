package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Request() {
	resp, err := http.Get("https://m.hupu.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
