# redditClone

## User guide
A running server can be accessed at http://ec2-54-169-136-167.ap-southeast-1.compute.amazonaws.com:8080/

Type in the topic and click submit to submit the topic.

Click the upvote and downvote button to vote for the corresponding topic

### Deployment
Require Go 1.9

To install:
```
go get vietthang207/redditClone
```

To run the application, we need the template folder from $GOPATH/src/github.com/vietthang207/redditClone
```
cd $GOPATH/src/github.com/vietthang207/redditClone
```

To run the application
```
go build server.go && ./redditClone &
```
API server will run at localhost:8080

A running server can be accessed at http://ec2-54-169-136-167.ap-southeast-1.compute.amazonaws.com:8080/

### Api reference

GET  /                   :return the home page with top 20 topics

POST /submit             : submit a new topic with the content from form data

POST /upvote/{topicId}   : upvote the topic with topicId

POST /downvote/{topicId} : downvote the topic with topicId

### Architecture

The app contains 4 packages:

main:       contain server.go which has the main function. 

            routes.go which defines the rest api and corresponding handler functions. 
      
            handlers.go which implements handlers functions
      
databases:  implements a database which contains a topic table and an indexer

models:     specify object models

indexers:   implement a min max heap to allow us to query top topics

utils:      utility functions.

### Detail of implementations:

databases:  
            
            use a mutex variable to make sure that only 1 thread can read or write to the database

heap:       
            
            heap is a binary heap that allows us to add {key, value} pair, query or remove the member with max(or min) key.               
            However, this heap is modified so that it keeps a backpointer of item value to its current location on the heap,             
            thus allow us to update a heap member by its value without the need of it to be the top elements.
            
minMaxHeap: 
            
            is a composition of 2 heaps. In the case that we need get top n max elements, we maintain a queryHeap which is a 

            min heap of size n, this heap will store n max element, all other elements will be stored in a bigHeap             
            which is a max heap. In other words, the invariant of the datastructure is that queryHeap size if always n, and               
            queryHeap always contains top n max elements in a min-heap manner. When we add to a minMaxHeap, we always try to             
            add to the queryHeap if it is bigger than the top elements of the queryHeap, and move this top element to the 
            
            other heap, thus make sure that the invariant is satisfied.
