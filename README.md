# potential-octo-guacamole

Running tests:

`make test`

Installing:

`make build`

Starting:
`./counting-request-server`

Running with docker
  * Build: `docker build -t counting-request-server .`
  * Run: ```docker run -it -p 3000:3000 -v `pwd`:/tmp counting-request-server:0.1```

Or get a Image from docker Hub:
  * Pull: `docker push fabriciojean/counting-request-server:0.1`
  * Run: ```docker run -it -p 3000:3000 -v `pwd`:/tmp counting-request-server:0.1```
