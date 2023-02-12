# gift-sats

Pet project inspired by Lightsats

This is a minimalistic first draft, using Fiber[https://docs.gofiber.io/], Sqlite[https://www.sqlite.org/] and Gorm[https://gorm.io/]

Endpoints for auth (signup, signin), users (list) and tips (list, create)

A POST /tips receives a tipperId and amount, creates a user and wallet with LNBits API and stores the tip and wallet in DB.
