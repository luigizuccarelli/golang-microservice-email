# Email microservice using 3rd party API's

A simple golang microservice that interfaces to an email api.

## Status 
[![Quality Gate Status](http://ec2-52-19-191-18.eu-west-1.compute.amazonaws.com/api/project_badges/measure?project=myportfolio-authinterface&metric=alert_status)](http://ec2-52-19-191-18.eu-west-1.compute.amazonaws.com/dashboard?id=myportfolio-authinterface)


## Usage 

```bash
# cd to project directory and build executable
$ go build -o bin/microservice .

```

## Docker build

```bash
docker build -t <your-registry-id>/email-sender:1.13.1 .

```

## Curl timing usage
```
curl -w "@curl-timing.txt" -o /dev/null -s "http://site-to-test

```

## Executing tests
```bash
# clear the cache - this is optional
go clean -testcache
GOCACHE=off go test -v validate.go validate_test.go schema.go handlers.go middleware.go middleware_test.go handlers_test.go -coverprofile tests/results/cover.out
go tool cover -html=tests/results/cover.out -o tests/results/cover.html
# run sonarqube scanner (assuming sonarqube server is running)
# NB the SonarQube host and login will differ - please update it accordingly 
 ~/Programs/sonar-scanner-3.3.0.1492-linux/bin/sonar-scanner  -Dsonar.projectKey=myportfolio-authinterface  -Dsonar.sources=.   -Dsonar.host.url=http://<add-host-here>   -Dsonar.login=<add-login-token-here> -Dsonar.go.coverage.reportPaths=tests/results/cover.out -Dsonar.exclusions=vendor/**,*_test.go,main.go,connectors.go,tests/**

```
## Testing container 
```bash

# start the container
# curl the isalive endpoint
curl -k -H 'Token: xxxxx' -w '@curl-timing.txt'  http://127.0.0.1:9000/api/v1/sys/info/isalive

```

```
export MAIL_HOST=https://api.sendgrid.com
export MAIL_URL=/v3/mail/send

```
