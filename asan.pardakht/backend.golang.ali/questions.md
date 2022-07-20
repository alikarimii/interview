## Interview stages:

- [x] ~2-hour session technical talk
- [ ] HR

## OOP
1. Is Golang object oriented?
2. Which one of OOP feature is not in Golang?
3. What is a difference between inheritance and composition?
4. Can we have Multiple inheritance in Go?
5. What is SOLID?
6. explain open closed principle.


## Golang
1. Which types are mutable and immutable in the Go Language?
2. Where we use mutable and where immutable?
3. Context is mutable or immutable?
4. Why we use immutable data?
5. How many ways are there to use a stack in multiple goroutines?
6. Why array index begins with zero?
7. What is a deadlock?
8. What is a memory leak?


## Data Structure and Algorithm
1. What is a difference between array and linkedlist?
2. If we have a sorted list of items which data structure do you prefer for searching one item in it?
3. explain binary search
4. What is a time complexity of linear/binary search?
5. What factors other than time do you consider to choose an algorithm?
6. Is it an advantage that algorithm run in parallel?

## Authentication
1. How JWT work?
2. Why do we use JWT instead of  generating some random token an giving it to the user?
3. If the user account is hacked, how do you prevent the attacker from accessing the resource?
4. Hash, Encode, Encrypt, Signing, explain these?


## GRPC
1. Why do you use GRPC?
2. What is a difference between REST and GRPC?

## Microservice
1. How do we scale system in Microservices architecture?
2. Are you familiar with macroservice patterns?
3. When we have multiple instance of one service, this can be challenging.explain a few things.
4. Based on what criteria do you separate the services?
5. Can we differentiate services based on functionality?


## System Design
1. What is a Cohesion and loose Coupling?
2. If multiple node access to shared resource(we have one third party service that need signin and get token,all node use one token,if one of them get new token older token expired),How can we handle lock between multiple node?
3. How can we handle single point of failure?
4. What rate limiting pattern are you working with?
5. If one of service have problem ,How do you find bottleneck?
6. What is a Hexagonal architecture?

## Distributed Data
1. What is CAP?
2. What is an eventually consistent?
3. How can we have availability and consistency?
4. In a Payment systems,Which one do you prefer, sql or nosql,why?
5. Why do you use RabbitMQ?
6. What is a Database normalization?

### Tips for study
1. lock free data structure
2. distributed lock manager
3. ACID and CAP
4. Database normalization
5. System profiler
6. distributed tracing (Jaeger,Zipkin)