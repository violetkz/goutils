package  main;

import  "net/http"


func send_req() {
	client := &http.Client{}
	//resp, err := client.Get("http://localhost:8080/oneTest/task")
	req, err := http.NewRequest("GET", "http://localhost:8080/oneTest/task", nil)
	req.SetBasicAuth("zhoukaiz@cn.ibm.com", "123456")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("send_req failed\n")
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}

type loginFWInofo struct {
	user    string
	psw     string
	logfile string
}

