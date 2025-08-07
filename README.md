# Payment-Tansaction-golang-api

This is an application that demonstrates how to build a REST API using GoLang.

## Run Locally

To run, from the root of the repo

Building the image

```
docker build -t go-docker .
```

Running the Docker image

Type the following command to run the docker image -

```
docker run -d  go-docker
```

## Access the App 

The App has a few Endpoints

All api endpoints are prefixed with `/api/v1`

To reach any endpoint use `localhost:8080/api/v1/payment/transaction`

```text
Get transactions filtered by provider
"/payment/transaction?provider={provider}" 


Get transactions by statusCode
"/payment/transaction?statusCode={authorised,decline,refunded}" 


Get Book by Currency
"/payment/transaction?currency={AUD}" 


```

## Tools Used

 - Go
 - Docker


## Inside Image
   - `docker run -it go-docker  sh`
 
## Postman collection
   https://www.getpostman.com/collections/19b2f872b7fa75c27c91

## To Run Test
 ```
go test .
```




