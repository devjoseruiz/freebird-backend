# FreeBird

An OpenSource Twitter clone made with GOLANG

## Description

This is a social network, Twitter-like made with GOLANG for learning purposes.

It has the following features:

- User login/registration
- User profile info
- User avatar and banner
- Editable profile
- "Tweets" posting
- "Tweets" removal
- Follow/unfollow other users
- "Tweets" timeline

## Getting Started

### Requirements

- GOLANG >= 1.12.x
- MongoDB

### Installation

Go to project folder and download the dependencies:

```bash
cd project_folder
go mod tidy
```

Then build the project:

```bash
go build
```

### How to run it

Copy the `.env.example` file as `.env`:

```bash
cp .env.example .env
```

Then fill the following key-value pairs:

```ini
# The Mongo server that the app will connect with
MONGODB_URI=mongodb+srv://<username>:<password>@<cluster>.<server>/<params>
DATABASE_NAME=freebird
# The port the server will listen to
HOST_PORT=8080
# The private key the server will use to encrypt the user JWT
JWT_PRIVATE=<super_random_and_secret_phrase>
```

Finally, run the server:

```bash
go run main.go
```

## Authors

- José Ruiz (joseruiz@keemail.me)

## Contributing

Any help is always welcomed. Just send PR or contact me.

## License

Licensed under GPLv3.
