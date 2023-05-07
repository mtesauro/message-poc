## Testing out Hex/Clean architecture

Running redis:
```
docker run --name temp-redis --rm -p 6379:6379 -d redis
```

Running Postgres:

```
docker run --name temp-postgres --rm -e POSTGRES_PASSWORD=pass1234 -p 5432:5432 -d postgres
```

{"body": "Redis message 5"}

**API Calls**

Create a message:

```
curl -k -v -X POST -d '{"body": "Redis message 6"}' http://127.0.0.1:5000/messages
```

Less verbose:
```
curl -k -X POST -d '{"body": "Redis message 6"}' http://127.0.0.1:5000/messages
```

Create a ton of messages:
```
for i in `seq 8 4000`; do curl -k -X POST -d "{\"body\": \"Redis message $i\"}" http://127.0.0.1:5000/messages; echo ""; done
```
Note: ideally, the sequence (seq) will start at the next message number and increase to the value you want - in the above case 8 to 4,000 to create a total of 4k messages

Time to create 4k messages on my local system:
```
real	0m16.001s
user	0m6.326s
sys	0m6.442s
```

Read all messages:

```
curl -k -X GET http://127.0.0.1:5000/messages
```

Time to read 12k messages on my local system:
```
$ time curl -k -s -X GET http://127.0.0.1:5000/messages 1>/dev/null

real	0m0.044s
user	0m0.000s
sys	0m0.005s
```

Read a specific message:

```
curl -k -X GET http://127.0.0.1:5000/messages/c3dbbcc3-3a9f-491b-a8ec-fb0dd4da7258
```
Note: You must replace the UUID in the call below with a valid one for the API call to be successfully
