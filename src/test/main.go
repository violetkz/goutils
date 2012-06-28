package main

import (
	"config"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type request struct {
	method string
	url    string
	user   string
	passwd string
}

func new_req(method string, url string, user string, pwd string) (req *request) {
	return &request{method, url, user, pwd}
}

//FIXME: reutrn value should be Boolean represented success or fail.
func (req *request) send() ([]byte, error) {
	client := &http.Client{}
	http_req, err := http.NewRequest(req.method, req.url, nil)
	http_req.SetBasicAuth(req.user, req.passwd)
	resp, err := client.Do(http_req)
	if err != nil {
		fmt.Printf("send_req failed\n")
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func main() {

	con, err := config.Config("config.txt", ":")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	c, err := con.Parse()
	if err == nil {
		fmt.Printf("\nConnect: %s, User: %s\n", c["host"], c["user"])
	}

	//send_req()
	r := new_req("HEAD", c["host"], c["user"], c["passwd"])
	resp, err := r.send()
	if err == nil {
		//fmt.Printf("%s", resp)
		fmt.Printf("\n    OK!     ")
	} else {
		fmt.Printf("\n    NG!     ")
	}

	time.Sleep(4000 * time.Millisecond)
}
