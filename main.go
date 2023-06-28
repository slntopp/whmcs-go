package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	whmcs "whmcs-go/pkg"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	client, err := whmcs.NewClient(getCred("/home/eternal/examples/whmcs_cred.json"))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	res, err := client.System.SendEmail(&whmcs.SendEmailRequest{
		MessageName: "instance_created",
		Id:          48,
		CustomVars: map[string]any{
			"client_name":      "name test",
			"nocloud_instance": "inst-1234",
			"signature":        "sign test",
		},
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	spew.Dump(res)
}

func getCred(file string) (string, string, string, bool) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	creds := struct {
		Api        string `json:"api"`
		User       string `json:"user"`
		Pass       string `json:"pass"`
		DangerMode bool   `json:"dangerMode"`
	}{}
	err = json.Unmarshal(data, &creds)
	if err != nil {
		panic(err)
	}

	return creds.Api, creds.User, creds.Pass, creds.DangerMode
}
