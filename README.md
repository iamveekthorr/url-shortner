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
> The project also makes use of [migrate](https://github.com/golang-migrate/migrate) to run migrations in the [migrations](./migrations/) directory.
> The project also makes use of [pgx](https://github.com/jackc/pgx) as a driver for `postgresql`.

- Golang
- Postgresql
- pgx
- golang-migrate

## Running The Application Locally

> ⚠️You would require to have golang installed as well as migrate-cli
> to run the migrations.
