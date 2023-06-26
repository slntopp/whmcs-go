package main

import (
	"fmt"
	"os"
	whmcs "whmcs-go/pkg"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	api := os.Getenv("api")
	user := os.Getenv("user")
	pass := os.Getenv("pass")
	dangerMode := os.Getenv("dangerMode") != "0"
	
	client, err := whmcs.NewClient(api, user, pass, dangerMode)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	res, err := client.Support.GetAnnouncements(nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	spew.Dump(res)
}
