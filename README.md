# gift-sats

Pet project inspired by Lightsats

This is a minimalistic draft, using Echo[https://echo.labstack.com/], Sqlite[https://www.sqlite.org/] and Gorm[https://gorm.io/]

Endpoints for auth (signup, signin), users (list), tips (list, create) and wallets (list)

A POST /tips receives a tipperId and amount, creates a user and wallet with LNBits API and stores the tip and wallet in DB.

Use Postman collection to test requests:

- Use signup to create user first
- use GET /users to list user
- use POST /tips to create tip
- use GET /tips and GET /wallets to list tips and lnbits wallets
