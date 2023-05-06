## Testing out Hex/Clean architecture

Running redis:
```
docker run --name temp-redis --rm -p 6379:6379 -d redis
```

Running Postgres:

```
docker run --name temp-postgres --rm -e POSTGRES_PASSWORD=pass1234 -p 5432:5432 -d postgres
```
