package spider

import (
	"fmt"
	"net/http"
)

func GetHttp() {
	resp, error := http.Get("https://baidu.com")
	if error != nil {
		fmt.Print(error)
	}
	fmt.Print(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Print(resp.Body)
	}
}
