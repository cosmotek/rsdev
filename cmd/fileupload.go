package cmd

import (
	"context"
	"fmt"
	"os"

	pjson "github.com/hokaccha/go-prettyjson"
	"github.com/machinebox/graphql"
)

const uploadQuery = `
mutation($file: Upload!) {
	setUserAvatar(userID: "3b81c770-2108-4207-a53f-843e56b2cb60", imageFile: $file) {
		id
		username
		avatarImageUrl
	}
}
`

func justTest() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: rsdev [GRAPHQL_QUERY_ENDPOINT]")
		os.Exit(1)
	}

	// create a client (safe to share across requests)
	client := graphql.NewClient(args[1], graphql.UseMultipartForm())
	ctx := context.Background()

	req := graphql.NewRequest(uploadQuery)
	file, err := os.Open("/home/cosmotek/Pictures/me.png")
	if err != nil {
		fmt.Printf("file open: %s\n", err.Error())
		os.Exit(1)
	}

	req.File("file", "me.png", file)
	req.Header.Add("X_RECSPOT_SESSION_TOKEN", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJtb2JpbGUudGhlcmVjc3BvdC5jb20iLCJleHAiOiIyMDIxLTAxLTIwVDIxOjQ3OjUwLjQxOTM5NTQ3OVoiLCJpYXQiOiIyMDIxLTAxLTIwVDIxOjE3OjUwLjQxOTM5NTI2NFoiLCJqdGkiOiI3ZDI2YTc5MS00OGZlLTRmNTUtYmVkNy0yMWQ5Y2ZiODk0NGUiLCJuYmYiOiIyMDIxLTAxLTIwVDIxOjE3OjUwLjQxOTM5NTM5WiIsInJvbGUiOiJ1c2VyIiwic3ViIjoiMTc0MDg3NzIzMjAiLCJzdWJpZCI6IjQ5ZTRlNmEwLTQzOGQtNDc4OC1hZjc4LTdiOGI3ODdlYWQ3ZSIsInN1YnR5cCI6InBob25lX251bWJlciIsInR2ZXIiOjF9.Ue9EzKOja17Gr3zPS_yo9T3G_xbReuqzlBgjVBDGfY-u-ipfRMX4rO2mNEP_55k1Q5FJrMxu3fHmDjcrArqOF7idaj0qwAYA2CfpFegAXNwZoIcDeLxZj3D_GpEIJg0-izJHS_kMrnPEpI1_9UHvTB-HV8xvEbLinAfP1dwbnR4d93NF_MJq8bTkta5eS9c7zpaLAHB1m0UGHE4zVH0wXIKvNoATELnMAq-onqQvNnB-XC_61-ZQsv0TUVVZl4V6lMp6vqouilzKgWMcV5eNmE9k9VXIdcvsoJqCO0Nh1UJQyeD5qFU6NWxLRx-DrG2rmeNh8ukTs-LUOqIX2F6ewuVCBl_34E3gBhwyK5OEHA9Ta3Sbn0zpm2DCbEdmkKTxfEmub5qlLoZiAyUAOJm08To6Oz9x8abJgA_4uKfgZ9ACZTGmvxFKiE5auw1W1i-WsZbq_Y3gZkfJZrsT7nN7WIvX-sK04DHeYca_MoWv0mZZTgHw-5yA-MkZ8Ck2MNroEBvbtjmi60H9dOPd-wcdCN-5k3UFG60RqI15hqnSQ8nkn2pxK5JNtFwUMcITBpbtzip98ZeM8KWnrwVKHoLklbJdnbXmftOXHYZGhEA_0lPm5miYlxengBFdedEn51zRT-C6mOChMXdYmU6qBzjWlV_3NzJhmtOXtvX1sSDyn0U")

	// run it and capture the response
	var res interface{}
	if err := client.Run(ctx, req, &res); err != nil {
		fmt.Printf("res: %s\n", err.Error())
		os.Exit(1)
	}

	jsonData, err := pjson.Marshal(res)
	if err != nil {
		fmt.Printf("marshal: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
