## Interview stages:

- [x] 1-hour session technical talk
- [x] 1-hour session live coding
- [ ] CTO meeting
- [ ] HR meeting

## 1-hour session technical talk:
---
### Database:
1. What is your criteria for selection db in project?
1. How can you reduce fetch time of this query ”select * from user”?
1. How can you handle transaction in older version of mongo db(<4)?
1. Which command in sql give us information about query that we do?
1. Which command in sql give us information about table details?
1. What is the "in-memory database"?
1. Do you work with redis/elastic?
1. How do you store paginated data?
1. Which command did you use most in redis?
1. If you want send otp verification ,how do you store it?

### Git:
1. which git flow do you work?
1. If you have bug in release branch ,how fix it?(assume we have master,dev,release)
1. What is different between reset and revert in git?
1. How we can have one commit from another branch that we have not in current branch?
1. What is different between merge and rebase in git?

### Authentication:
1. Which mechanism do you use for authentication in your app?
1. Where are you store auth token(server or client)?
1. Is it ok that store user id or product id directly in token?
1. Can we modify payload of jwt token?
1. How can we prevent a smart user from making more requests than expected to our service?
1. reasonable expire time of jwt token,how much can be?

### Golang:
1. Why did you choose Golang?
1. Is it ok that we have global variable in our app?
1. Why do we use empty struct in go program?
1. What is difference between concurrency and parallelism?
1. What is GOPATH and GOROOT?
1. How to recover panic?
1. How to control unrecovered panic in the program?
1. How do we know that a map/slice is empty?
1. What is a default value of bool/int/string variable?
1. How to define type of object in runtime?

### Microservices:
1. How can we have a global middleware in Microservices structure?
1. Which protocol do you use for communication between services?
1. How to manage transaction in microservices?

### Programming:
1. Do you know solid?explain one of them.
1. Which design pattern did you work with?explain it.
1. In HTTP verb, What is “PUT” for?what is “OPTION” for?

## 1-hour session live coding
---
Refactor this code .[link](https://github.com/snappmarket/Challenge-Go-1)