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

2. Migrate the database

- Make sure that you're inside the repository's folder on your local setup.
- Then run:

```shell
 make migrateup
```
