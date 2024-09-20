# Stori

Stori API to process monthly balance of customers' cards.

## Installation

Use the golang version [Go v1.22.4](https://go.dev/doc/go1.22#introduction) to run the project.

## Local run

Download go modules
```bash
go mod download
```
Execute stori api
```bash
go run main.go
```
API running successfully
````
  ____ _____ ___  ____  ___ 
 / ___|_   _/ _ \|  _ \|_ _|
 \___ \ | || | | | |_) || | 
  ___) || || |_| |  _ < | | 
 |____/ |_| \___/|_| \_\___|

Stori API - Brandon Jaime 2024
Server: http://localhost:8080
````

### Usage
Import cURL into postman or talend

Send id values for example:

```json
"id": 1234
"id": 12345
```
```bash
curl --location 'http://localhost:8080/api/user/summary' \
     --header 'Content-Type: application/json' \
     --data-raw '{
          "email": "jcastillo.brandon@gmail.com",
          "id": 1234
     }'
```

## Aws execution

### Usage
Import cURL into postman or talend

Send id values for example:

```json
"id": 1234
"id": 12345
```

```bash
curl --location 'https://6uhsdbbg7d.execute-api.us-east-2.amazonaws.com/stori/api/user/summary' \
     --header 'Content-Type: application/json' \
     --data-raw '{
          "email": "gfb-mpd@hotmail.com",
          "name": "Brandon Jaime",
          "id": 12345
     }'
```


## Author
Brandon Jaime - 2024
