package main

import (
	"config"
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//var enc = base32.NewEncoding(base32.encodeHex)
	s := base32.HexEncoding.EncodeToString([]byte("violetsss~~"))
	fmt.Printf("E: %s", s)
	r, e := base32.HexEncoding.DecodeString(s)
	if e != nil {
		fmt.Print("error")
	}
	fmt.Printf("D: %s", r)
}

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

func doAuth(c *map[string]string) {

	password, e := base32.HexEncoding.DecodeString(c["passwd"])
	if e != nil {
		fmt.Print("decode password error")
		return
	}

	//send_req()
	r := new_req("HEAD", c["host"], c["user"], password)

	//resp, err := r.send()
	_, err = r.send()
	if err == nil {
		//fmt.Printf("%s", resp)
		fmt.Printf("\n    OK!     ")
	} else {
		fmt.Printf("\n    NG!     ")
	}

}

func doSetPassword(c *ConfigFile,  p *string) {
    c.Set("passwd",base32.HexEncoding.EncodeToString([]byte(p))
}

func main() {

	con, err := config.Config("config.txt", ":")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer con.Close()

	c, err := con.Parse()
	if err == nil {
		fmt.Printf("\nConnect: %s, User: %s\n", c["host"], c["user"])
	}

    var new_pass = flag.String("s", nil, "new password");
    flag.Parse()
    var l = flag.NArg()
    if  0 == 1 {
        doAuth(c);
    } else if  1 ==  l {
        doSetPassword(c, new_pass);
    }

	time.Sleep(4000 * time.Millisecond)
}
