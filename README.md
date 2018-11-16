# PubSub Server

This server implements a dummy pub/sub server exposing an HTTP API

It has one route only

``` BASH
/
```

Which handles both a GET and a POST request.
The GET request adds the IP of the requester to a list of subscribers for the standard empty topic
while the post request returns the number of subscribers accepting the message
It is possible to specify an X-FORWARDED-FOR header to the GET request in order to test it.
The messages are not sent to any subscriber for real, this is implemented in another branch:

```publish-requests```

another branch with possibility to specify topics uses a different url for GET:

```/{topic}```

and protobuf for the messages, the branch is called:

```protobuf```
In order to make the request there is an utility whch can be used this way:
``` BASH

cd makerequest
go run main.go -url http://localhost:9090 -topic /topic1 -message HI

```

The address to run the server is set to ```:9090``` by default but can be changed with the flag:

```-addr :80```
