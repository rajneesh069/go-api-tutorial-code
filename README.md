# GO API tutorial project

Tech Stack

- Gin
- JWT
- PostgreSQL

## To run the project just run the following commands:

1. Run the docker container (Make sure you have Docker installed and the daemon is running in the background):

```shell
docker compose up -d
```

- Ports at which the services are running are as follows:
  - Postgres: 5432 - accessible at localhost:5432 on the host
  - PgAdmin: 8000 - accessible at localhost:8000 on the host
  - Server: 8080 - accessible at localhost:8080 on the host

2. Migrate the database

- Make sure that you're inside the repository's folder on your local setup.
- Then run:

```shell
make migrateup
```

3. To use pgAdmin, do the following:

- Go to `http://localhost:8000/` in your browser.
- Credentials:
  - Email: admin@gmail.com
  - Password: admin123
- Todo DB connection params:
  - Host: postgres (because pgadmin will talk to postgres inside the docker network)
  - Port: 5432
  - Database: todo_db
  - User: postgres
  - Password: mysecretpassword
