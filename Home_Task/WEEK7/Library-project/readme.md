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

TODO : 
1. Use config files for port and domain information
2. Add logs
3. Use 3rd party router
4. Document API using swagger
5. Achieve database consistency using microservice patterns for transactions


