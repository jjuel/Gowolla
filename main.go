package main

import (
	"fmt"
	"log"

	"github.com/jjuel/Gowolla/auth"
)

func main() {
	req := auth.Request{
		ClientID:     "6tudYSTEAmso8mKxeAV7hvIlfxnwUnRJzE1S2ppYMXUqcSIVeV",
		ClientSecret: "GNDOmkju2x40rGPN4hWmTH0M9N0mlsbz0CYC10bVvVAvF5yXDs",
		//ClientSecret: "",
		GrantType: "grant_type=client_credentials",
	}

	resp, err := auth.GetToken(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
