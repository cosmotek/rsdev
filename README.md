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
    "expires": "2021-01-07T20:47:06Z",
    "identity": {
      "createdAt": "2021-01-06T20:59:28Z",
      "id": "74da7343-6134-4a3d-ad1f-63e930ba21a1",
      "lastLoginAt": "2021-01-07T20:17:06Z",
      "linkedPhoneNumber": "7408772320",
      "personas": [
        "eea4a2ef-ade4-4528-a49b-43500531ca7f",
        "cb2f27f5-34ef-4e3f-bedc-f850dc65bb06"
      ]
    },
    "tokenValue": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJtb2JpbGUudGhlcmVjc3BvdC5jb20iLCJleHAiOiIyMDIxLTAxLTA3VDIwOjQ3OjA2LjEyODk3NTE4OVoiLCJpYXQiOiIyMDIxLTAxLTA3VDIwOjE3OjA2LjEyODk3NTA3WiIsImp0aSI6IjU1MmZkYjE2LWE4NmUtNGI3OC05MzMzLTQzYzk0OTFhMGFjNyIsIm5iZiI6IjIwMjEtMDEtMDdUMjA6MTc6MDYuMTI4OTc1MTM0WiIsInJvbGUiOiJ1c2VyIiwic3ViIjoiNzQwODc3MjMyMCIsInN1YmlkIjoiNzRkYTczNDMtNjEzNC00YTNkLWFkMWYtNjNlOTMwYmEyMWExIiwic3VidHlwIjoicGhvbmVfbnVtYmVyIiwidHZlciI6MX0.jNF84zzCjEX_CnfTkCqYMZVC9R72ijjbJjjjcZiFMh7QXjnjhh_mQSQHKQtA6c4hUVDABLTcJaBdoN5mLqkouEia72xyOjbG-3S_hbTM71MlMzTparu_c6M8_E5C7Hppv5GY1SjyOAk05Npo6uBE23q3QOCbcIGG01DqoOLVL0CPQml3R7TRG7IlVbuy9CEOQ0KrfDq4ExCAtmLNp942bnP6ifmBsu_Ymq6NaXWZDlDN3oNAHvtfLCgQvYqL7Xoc1L5ln3WPcV8m76BPuK65bGQhdSWlOd_xUrQQ2eOxC35OGq8G2z-cWl9bZFdA0rIAOpF4eKIbi_IVZBuRG7sylLvkb4r5Kw_DpHVWoVQ47Nce2h8nyik1ye-NiG3i5smnZG3LQ-HAy0QGt62oFiZmRVSnN3G2-ayQi26tuB9V_3j0GhA5ugZuuu_lc9FYbYa5plNhWQ_Etdpw144mtw5wrZqQGdW_a2skx3Dh8eqnABvGlBu5WzjmhINx5KAskrQnQreAFmJda8rQvE8ilKuZp75LvYmoxFP11AXf3ZdnC990GoCbj8Nm9rM9x4rNhHIJX6kmKSzTukPrwk18N9CDX-m2UCBqozlkfeDbnqk5xyr0-Uc8nlgVldhBFdLgDNPebeQxm1vvZeFfWIecEiVQVGOqt3WEPDuPAq0o2CnaMbQ"
  }
}

Copy and paste the following into the GraphQL Playground to make authenticated queries:
{
  "X_RECSPOT_SESSION_TOKEN": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJtb2JpbGUudGhlcmVjc3BvdC5jb20iLCJleHAiOiIyMDIxLTAxLTA3VDIwOjQ3OjA2LjEyODk3NTE4OVoiLCJpYXQiOiIyMDIxLTAxLTA3VDIwOjE3OjA2LjEyODk3NTA3WiIsImp0aSI6IjU1MmZkYjE2LWE4NmUtNGI3OC05MzMzLTQzYzk0OTFhMGFjNyIsIm5iZiI6IjIwMjEtMDEtMDdUMjA6MTc6MDYuMTI4OTc1MTM0WiIsInJvbGUiOiJ1c2VyIiwic3ViIjoiNzQwODc3MjMyMCIsInN1YmlkIjoiNzRkYTczNDMtNjEzNC00YTNkLWFkMWYtNjNlOTMwYmEyMWExIiwic3VidHlwIjoicGhvbmVfbnVtYmVyIiwidHZlciI6MX0.jNF84zzCjEX_CnfTkCqYMZVC9R72ijjbJjjjcZiFMh7QXjnjhh_mQSQHKQtA6c4hUVDABLTcJaBdoN5mLqkouEia72xyOjbG-3S_hbTM71MlMzTparu_c6M8_E5C7Hppv5GY1SjyOAk05Npo6uBE23q3QOCbcIGG01DqoOLVL0CPQml3R7TRG7IlVbuy9CEOQ0KrfDq4ExCAtmLNp942bnP6ifmBsu_Ymq6NaXWZDlDN3oNAHvtfLCgQvYqL7Xoc1L5ln3WPcV8m76BPuK65bGQhdSWlOd_xUrQQ2eOxC35OGq8G2z-cWl9bZFdA0rIAOpF4eKIbi_IVZBuRG7sylLvkb4r5Kw_DpHVWoVQ47Nce2h8nyik1ye-NiG3i5smnZG3LQ-HAy0QGt62oFiZmRVSnN3G2-ayQi26tuB9V_3j0GhA5ugZuuu_lc9FYbYa5plNhWQ_Etdpw144mtw5wrZqQGdW_a2skx3Dh8eqnABvGlBu5WzjmhINx5KAskrQnQreAFmJda8rQvE8ilKuZp75LvYmoxFP11AXf3ZdnC990GoCbj8Nm9rM9x4rNhHIJX6kmKSzTukPrwk18N9CDX-m2UCBqozlkfeDbnqk5xyr0-Uc8nlgVldhBFdLgDNPebeQxm1vvZeFfWIecEiVQVGOqt3WEPDuPAq0o2CnaMbQ"
}
```

## Installation

_Please ensure your `$GOBIN` is set and in your `$PATH`._
```sh
go get gitlab.com/therecspot/rsdev
```