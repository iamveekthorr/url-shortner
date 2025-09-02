# URL Shortening Service

---

## The Challenge

**You are required to create a simple RESTful API that allows users to shorten
long URLs. The API should provide endpoints to create, retrieve, update,
and delete short URLs. It should also provide statistics
on the number of times a short URL has been accessed.
Follow the link to see full description [project url](https://roadmap.sh/projects/url-shortening-service)**

## The Requirements

- [x] Create a new short URL
- [x] Retrieve an original URL from a short URL
- [x] Update an existing short URL
- [x] Delete an existing short URL
- [x] Get statistics on the short URL (e.g., number of times accessed)

## Tech Stack

> The project is made of using [Golang](https://go.dev/) as the programming language,
> and [postgres](https://www.postgresql.org/) as the database.
> The project also makes use of [migrate](https://github.com/golang-migrate/migrate)
> to run migrations in the [migrations](./migrations/) directory.
> The project also makes use of [pgx](https://github.com/jackc/pgx)
> as a driver for `postgresql`.

- Golang
- Postgresql
- pgx
- golang-migrate

## Developing The Application Locally

### ‼️ Make sure you run the migrations before sending any request

> ⚠️ You are required to have golang installed as well as migrate-cli
> to run the migrations as well as the [air](https://github.com/air-verse/air)
> package to have hot reload while in development.
> For local development, you will need to install [godotenv](https://github.com/joho/godotenv)
> to load all environment variables into `os.Getenv()`.
> If you fail to do so, the project will not start‼️.
> An alternative is to setup [doppler](https://doppler.com/) so you
> do not have to worry about installing another dev-dependency.

### Using Air

```bash
# if you have air installed just run the air command
air
```

### Using Pure Golang

```bash
# This assumes you have environment variables set
go run main.go
```

### Using Doppler

```bash
# Run doppler login to authenticate
doppler login

# Run setup command
doppler setup

# Assuming the above commands are successful
doppler run -- air
```

## Environment Variables

---

| Name                 | Required |
| -------------------- | -------- |
| DB_CONNECTION_STRING | true     |
| PORT                 | false    |
