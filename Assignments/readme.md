### NTUC Golang internal Assignment 01

### Concurrency assignment 1:
### Creating a worker pool using buffered channels for a program which receives JSON data from an 
### interface.
### The interface will be xkcd website. This website gives data in the form of JSON.
### The xkcd website has over 2500 comics to download. To do this sequentially, it would take a long time. 
### Hence it would not make any sense to download this resource sequentially.
### A solution will be using a concurrent model, by implementing a Worker pool to handle multiple HTTP 
### requests at a time, and getting multiple results in a very short time.

### Unit testing assignment 1:
### Create a test suite using Ginkgo and GoMock for the above program

# to run App
>> go run main.go

# to check response refer below link
>> https://xkcd.com/571/info.0.json