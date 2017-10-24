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
