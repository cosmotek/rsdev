# RecSpot Dev Tool

Authentication Proxy Example:
```sh
$ rsdev proxy http://34.218.48.156:5000
whats your phone number?: 7408772320
 
Pincode sent successfully, please check your phone (pincode may take as much as 10 minutes to arrive). 
whats the pincode: 8750
 
A new session has been created: 
{
  "requestSMSSessionToken": {
    "accountAssociated": false,
    "cookieKey": "X_RECSPOT_SESSION_TOKEN",
    "expires": "2021-01-22T20:02:16Z",
    "identity": {
      "createdAt": "2021-01-22T16:50:47Z",
      "id": "6608e66f-505a-44ef-b68e-dca50cfa8749",
      "lastLoginAt": "2021-01-22T17:02:16Z",
      "linkedPhoneNumber": "17408772320",
      "personas": []
    },
    "tokenValue": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJtb2JpbGUudGhlcmVjc3BvdC5jb20iLCJleHAiOiIyMDIxLTAxLTIyVDIwOjAyOjE2LjkyMDAxMjgxOFoiLCJpYXQiOiIyMDIxLTAxLTIyVDE3OjAyOjE2LjkyMDAxMDMwNloiLCJqdGkiOiJjMjk1OTFlNC02ZWFiLTQyN2QtYTFkMy02MWU0YzBmNDA5NTIiLCJuYmYiOiIyMDIxLTAxLTIyVDE3OjAyOjE2LjkyMDAxMTYxM1oiLCJyb2xlIjoidXNlciIsInN1YiI6IjE3NDA4NzcyMzIwIiwic3ViaWQiOiI2NjA4ZTY2Zi01MDVhLTQ0ZWYtYjY4ZS1kY2E1MGNmYTg3NDkiLCJzdWJ0eXAiOiJwaG9uZV9udW1iZXIiLCJ0dmVyIjoxfQ.OJK2M2PdJbZVsS0Zeh1_1vTaca0uDTNH2zE3Z4PHBb59UgXbbp3-iOGpeQ-tSVBfeIlMRnlzWA0rdWVQmAo7jJt6cEKH6M78w-V5GLIohY8xg3Z9n3korHpH0YXzZxOn4rwWoikpGFlGV_Q2tjK_9WHr5rveHKMgCXusO2d6nRd-EVl1B4mS437Cx3gqB12lK09mZYR1j7DflUtXZjd0bIFDS1LZaNzlvDoY0ZjpYt9YKKmDh1MWAgBghbu3yQ9onyTtlAzg1paI-0yNk_O8-lPAA0UT3-PP1Og7mQ-KfaVenvZcJOOnPnLWYL6rWxEV5_ruyRLaoMi78tyNQZtUNRsZfY4yAeRoCvQ16UpgVOY7l75gEKpcfMi1lnpNRkGkVvbvYGbh9CGOQalXgjDwqoGZ0e3_ftEwESNs40npHJ-0zRYJLac8yx2Nm5N5jjSSqpcrzLPweBqH4atcjKWqncCpPx1J7W37mKQwoK20I4bXcObu7jodcF2ztxgXb731j220-v-cpJN0ZZpLdlPHLglMC_571vPOsp29S1Eod2Hmyr_fX1afb_00fUddJ0OZGjLjMpa3gfg-l9PNa8ztVsnCuKTYsYhrtLgu99NwvfmTDtp1Fiih4y7WCcsYgFUp7RwbTFLiRYNKOaaaRYFRKpgyvVs-VXI4iaS4MA_TU38"
  }
}
 
Copy and paste the following into the GraphQL Playground to make authenticated queries: 
Starting authentication proxy on http://0.0.0.0:44239
```

# Git Config (Required for all private repos)
Add the following to your global `.gitconfig` file:
```toml
[url "git@gitlab.com:"]
	insteadOf = https://gitlab.com/
```

## Installation

_Please ensure your `$GOBIN` is set and in your `$PATH`._
```sh
go get -u gitlab.com/therecspot/rsdev
```
