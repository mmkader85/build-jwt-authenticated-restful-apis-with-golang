# Build JWT authenticated Restful APIs with Golang

Based on udemy course - https://naspers.udemy.com/course/build-jwt-authenticated-restful-apis-with-golang/learn/lecture/12509338?start=0#overview

### Run the program

```
$ ./user_jwt
```

### Curl for end-points

##### Signup
```
curl --location --request POST 'http://localhost:8000/signup' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "email@example.com",
    "password": "password"
}'
```

##### Login
```
curl --location --request POST 'http://localhost:8000/login' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "email@example.com",
    "password": "password"
}'
```

##### Get all users
```
curl --location --request GET 'http://localhost:8000/get_all_users' \
--header 'Authorization: Bearer eyJhbG...fT2XZs'
```
