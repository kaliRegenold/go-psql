# Go Post
A very simple docker composition of a postgres container and a Go web app container.

## Build
`docker-compose up --build -d`

## Test
`curl 0.0.0.0:8080/ping`

should return:

`Success pinging database`

## Use
`curl 0.0.0.0:8080/user/5`

should return:

`{"UserId":5,"Lastname":"Lyman","Firstname":"Josh","State":"CT","City":"Westport"}`
