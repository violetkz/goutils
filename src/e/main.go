package main

import (
        "./config"
        "fmt"
       )

func main() {

	con, err := config.Config("config.txt", ":")
	if err != nil {
		fmt.Printf("%s\n", err)
        return
	}

	c, err := con.Parse()
	if err == nil {
		fmt.Printf("==: %s\n", c["logfile"])
        if "" == c["logfiles"] {
            fmt.Printf("nil\n");
        }
		fmt.Printf("==: %s\n", c["logfiles"])
	}
}
