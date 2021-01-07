package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	pjson "github.com/hokaccha/go-prettyjson"
	"github.com/machinebox/graphql"
	"github.com/rucuriousyet/monolog"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: rsdev [GRAPHQL_QUERY_ENDPOINT]")
		os.Exit(1)
	}

	// create a client (safe to share across requests)
	client := graphql.NewClient(args[1])
	ctx := context.Background()
	var phoneNumber string

	err := monolog.New(nil, nil).
		Add(func(p *monolog.Prompter) monolog.Cmd {
			p.Write("whats your phone number?: ")
			phoneNumber = strings.ToLower(p.Read())

			// make a request
			req := graphql.NewRequest(`
mutation ($phoneNumber: PhoneNumber!) {
	requestVerificationPincode(input: $phoneNumber)
}
`)

			// set any variables
			req.Var("phoneNumber", phoneNumber)

			// run it and capture the response
			var res interface{}
			if err := client.Run(ctx, req, &res); err != nil {
				fmt.Printf("%s\n", err.Error())
				return monolog.ExitChain
			}

			jsonData, err := pjson.Marshal(res)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				return monolog.ExitChain
			}

			fmt.Println(string(jsonData))
			return monolog.Continue
		}).
		Add(func(p *monolog.Prompter) monolog.Cmd {
			p.Write("whats the pincode: ")
			pinCode := strings.ToLower(p.Read())

			req := graphql.NewRequest(`
mutation ($phoneNumber: PhoneNumber!, $pinCode: String!) {
	requestSMSSessionToken(input: {
		phoneNumber: $phoneNumber,
		pincode: $pinCode
	}) {
		cookieKey
		tokenValue
		expires
			accountAssociated
			identity {
				id
				personas
				linkedPhoneNumber
				lastLoginAt
				createdAt
			}
	}
}
`)

			// set any variables
			req.Var("phoneNumber", phoneNumber)
			req.Var("pinCode", pinCode)

			// run it and capture the response
			var res2 interface{}
			if err := client.Run(ctx, req, &res2); err != nil {
				fmt.Printf("%s\n", err.Error())
				return monolog.ExitChain
			}

			jsonData, err := pjson.Marshal(res2)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				return monolog.ExitChain
			}

			fmt.Println(string(jsonData))
			return monolog.Continue
		}).Do()

	if err != nil {
		panic(err)
	}
}
