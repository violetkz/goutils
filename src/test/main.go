package main

import (
        "fmt"
        "config"
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
		fmt.Printf("==: %s\n", c["logfiles"])
	}
}
