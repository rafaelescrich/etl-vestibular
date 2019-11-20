# etl-vestibular

Parse, modify and store a structured CSV in a SQL database

## Run Docker

```bash
docker run --name etl -e MYSQL_ROOT_PASSWORD=etl -e MYSQL_DATABASE=etl -p 3306:3306 -d mysql:latest
```
