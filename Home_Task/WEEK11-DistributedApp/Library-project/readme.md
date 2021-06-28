TO RUN - `docker compose up`  in cmd or shell
TO DEPLOY using swarm - `docker stack deploy -c docker-compose.yml library_stack`


Created library-project using 3 microservices library-service, users-service, books-service

Using mysql db with beego-orm on top to persist data

*MYSQL hardcoded for now, user - root, pass - root*

Works for all rest endpoints mentioned at - https://kb.epam.com/display/GDOKB/JPOP2+Case+study

GET     `localhost:5002/library/users` 
POST    `localhost:5002/library/users`
GET	    `localhost:5002/library/users/([0-9]+)`
PUT	    `localhost:5002/library/users/([0-9]+)`
DELETE	`localhost:5002/library/users/([0-9]+)`
GET 	`localhost:5002/library/books` 
POST	`localhost:5002/library/books`
GET	    `localhost:5002/library/books/([0-9]+)`
PUT 	`localhost:5002/library/books/([0-9]+)`
DELETE	`localhost:5002/library/books/([0-9]+)`
POST	`localhost:5002/library/users/([0-9]+)/books/([0-9]+)`
DELETE	`localhost:5002/library/users/([0-9]+)/books/([0-9]+)`

GET     `localhost:5001/users` 
POST    `localhost:5001/users`
GET	    `localhost:5001/([0-9]+)`
PUT	    `localhost:5001/([0-9]+)`
DELETE	`localhost:5001/([0-9]+)`

GET 	`localhost:5000/books` 
POST	`localhost:5000/books`
GET	    `localhost:5000/books/([0-9]+)`
PUT 	`localhost:5000/books/([0-9]+)`
DELETE	`localhost:5000/books/([0-9]+)`


POST `localhost:5003/login` --> json body --> 

{
    "username":"admin",
    "password":"admin12345"
}

SECURITY Using JWT:  

UseCase : Only for admin login we should allow to add/delete users and books.

Solution: Once user is `authenticated` as `admin` we will recieve a jwt `<token>` in response headers.
        Now, While accessing add/delete users and books api we have to use this token as request header in the below way.
        `"Authorization": Bearer <token>`

Implementation Details:
The Signing certificate private key to create the jwt token only lies with the authentication service, in this case the login service.
Any other service needing authorization of user can just validate the jwt token using public key.


New: 1. Added `localhost:5003/refresh` endpoint to generate a new access token by using refresh token.
     2. Added `localhost:5003/logout` endpoint to invalidate tokens at logout.
     2. Added cache for generated tokens.
     3. Using Docker swarm with multiple instances of a books-service, url resolution/discovery managed by swarm.

Note: Ideally the refresh endpoint is to be used by client logic to fetch a new access token every 15 mins(expiry time of access token) using long expiry(24 hours) refresh tokens, but here need to do it manually in postman. Find the refresh token in http-only cookie in response of `/login`.






