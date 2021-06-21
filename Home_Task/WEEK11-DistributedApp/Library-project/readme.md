TO RUN - `docker compose up`  in cmd or shell


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



