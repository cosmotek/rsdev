# RecSpot Dev Tool

Right now this tool only has one function, dev login. Here's an example:
```sh
$ rsdev http://localhost:5000/query
whats your phone number?: 7408772320
{
  "requestVerificationPincode": true
}
whats the pincode: 3005
{
  "requestSMSSessionToken": {
    "accountAssociated": true,
    "cookieKey": "X_RECSPOT_SESSION_TOKEN",
    "expires": "2021-01-07T19:45:28Z",
    "identity": {
      "createdAt": "2021-01-06T20:59:28Z",
      "id": "74da7343-6134-4a3d-ad1f-63e930ba21a1",
      "lastLoginAt": "2021-01-07T19:15:28Z",
      "linkedPhoneNumber": "7408772320",
      "personas": [
        "eea4a2ef-ade4-4528-a49b-43500531ca7f"
      ]
    },
    "tokenValue": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJtb2JpbGUudGhlcmVjc3BvdC5jb20iLCJleHAiOiIyMDIxLTAxLTA3VDE5OjQ1OjI4LjMxNTAzNjk1MloiLCJpYXQiOiIyMDIxLTAxLTA3VDE5OjE1OjI4LjMxNTAzNjgzNVoiLCJqdGkiOiIzNzNiNzVhYS1kYzFjLTQ0NWItOTI4My03MDI2ZGJkNDVkMGQiLCJuYmYiOiIyMDIxLTAxLTA3VDE5OjE1OjI4LjMxNTAzNjkwMloiLCJyb2xlIjoidXNlciIsInN1YiI6Ijc0MDg3NzIzMjAiLCJzdWJpZCI6Ijc0ZGE3MzQzLTYxMzQtNGEzZC1hZDFmLTYzZTkzMGJhMjFhMSIsInN1YnR5cCI6InBob25lX251bWJlciIsInR2ZXIiOjF9.CP2lhhkADvCSI8Zo-3QrhQKo8CwL8MOO1S5P9fnVFuJHnaBELsH8ywrUTkKAHY0fR9hw5jqMYIbRn9DvcLpxKUqvB5PcDAU6gx48tOHsqZDlpKEulikmoWxs9uD-ugDR2np6o8N0S5i2oHfbsHl4CAk1kxfMD4RgDttUG0QNikHhd9hDzOh_ojUImaECz3CLWg5I5bLuPkmWfO76EcTkx2f3YHXR6sLW9C2kpP0j-nu5epuIBOojlRBnYgILpKvyJ75TVYv-LSAfp56Q73UyEhGRy1PdmQk_JcX6a6oAE-iBQDeMEahn2sScknDyKNEiWoXBwqSYJMYVPzBjL-FH4gUoeQWdR6kMJQYwxRGscxg6Z6Qj8UrutmKK4uOmL1POSavN2PJxoGN_16Nqp2Kd5IQ2dtJ9rQUQPWskGIXNv6k9_FIQbifkbQeAI308LlNB2APiIjkO8ERqQ_erwg4XkcBnYM0uOSIbIz_Kj99vmMZSdYnNzwCYEszhBdSnTwx9CiRChs09w4t1v5pQRNx3_RakZowyGIzO8ZCBuL23mnQe8r7VJBxfSiXeG2Y71qiS_pypK59f7V6HrLTOg_MNzF3WaPAKi5P_hzYXTx7INqtDL7XT8yIGlKh2Y5poJsxDYwn_BzT6MGlvKgaF_0zRQxv9C1dCHneJCl8ABfIcjVA"
  }
}
```

## Installation

_Please ensure your `$GOBIN` is set and in your `$PATH`._
```sh
go get gitlab.com/therecspot/rsdev
```