package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	pjson "github.com/hokaccha/go-prettyjson"
	"github.com/machinebox/graphql"
	"github.com/mgutz/ansi"
	"github.com/rucuriousyet/monolog"
)

var (
	lime  = ansi.ColorCode("green+h")
	cyan  = ansi.ColorCode("cyan+h")
	reset = ansi.ColorCode("reset")
)

type AuthHeader struct {
	HeaderKey string    `json:"headerKey"`
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func StartPromptedSMSVerification(ctx context.Context, graphqlEndpointURL string) (AuthHeader, error) {
	client := graphql.NewClient(graphqlEndpointURL)

	var phoneNumber string
	var tokRes RequestSMSSessionTokenResponse

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

			fmt.Println(lime+"\nPincode sent successfully, please check your phone (pincode may take as much as 10 minutes to arrive).", reset)
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

			if err := client.Run(ctx, req, &tokRes); err != nil {
				fmt.Printf("%s\n", err.Error())
				return monolog.ExitChain
			}

			jsonData, err := pjson.Marshal(tokRes)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				return monolog.ExitChain
			}

			fmt.Println(lime+"\nA new session has been created:", reset)
			fmt.Println(string(jsonData))

			fmt.Println(cyan+"\nCopy and paste the following into the GraphQL Playground to make authenticated queries:", reset)
			jsonData, err = pjson.Marshal(map[string]string{
				tokRes.Response.CookieKey: tokRes.Response.TokenValue,
			})
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				return monolog.ExitChain
			}

			return monolog.Continue
		}).Do()
	if err != nil {
		return AuthHeader{}, err
	}

	return AuthHeader{
		HeaderKey: tokRes.Response.CookieKey,
		Value:     tokRes.Response.TokenValue,
		ExpiresAt: tokRes.Response.Expires,
	}, nil
}

type RequestSMSSessionTokenResponse struct {
	Response struct {
		AccountAssociated bool      `json:"accountAssociated"`
		CookieKey         string    `json:"cookieKey"`
		TokenValue        string    `json:"tokenValue"`
		Expires           time.Time `json:"expires"`
		Identity          struct {
			CreatedAt         time.Time `json:"createdAt"`
			ID                string    `json:"id"`
			LastLoginAt       time.Time `json:"lastLoginAt"`
			LinkedPhoneNumber string    `json:"linkedPhoneNumber"`
			Personas          []string  `json:"personas"`
		} `json:"identity"`
	} `json:"requestSMSSessionToken"`
}

type RefreshSessionTokenResponse struct {
	Response struct {
		AccountAssociated bool      `json:"accountAssociated"`
		CookieKey         string    `json:"cookieKey"`
		TokenValue        string    `json:"tokenValue"`
		Expires           time.Time `json:"expires"`
		Identity          struct {
			CreatedAt         time.Time `json:"createdAt"`
			ID                string    `json:"id"`
			LastLoginAt       time.Time `json:"lastLoginAt"`
			LinkedPhoneNumber string    `json:"linkedPhoneNumber"`
			Personas          []string  `json:"personas"`
		} `json:"identity"`
	} `json:"refreshSessionToken"`
}

func RefreshToken(ctx context.Context, graphqlEndpointURL, oldToken string) (AuthHeader, error) {
	client := graphql.NewClient(graphqlEndpointURL)

	req := graphql.NewRequest(`
mutation c($sessionToken: String!) {
	refreshSessionToken(sessionToken: $sessionToken) {
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

	req.Var("sessionToken", oldToken)

	var res RefreshSessionTokenResponse
	if err := client.Run(ctx, req, &res); err != nil {
		return AuthHeader{}, err
	}

	return AuthHeader{
		HeaderKey: res.Response.CookieKey,
		Value:     res.Response.TokenValue,
		ExpiresAt: res.Response.Expires,
	}, nil
}
