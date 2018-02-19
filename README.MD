# MARKETPLACE CATALOG API
## Install
```
cd $GOPATH/src/github.com
mkdir VRlectorium
git clone https://github.com/VRlectorium/catalog-service
cd VRlectorium/catalog-service
```
or
```
git clone https://github.com/VRlectorium/catalog-service folder_name
cd VRlectorium/folder_name
```
## Before running
Current golang version: 1.9.3
Use this command to install project dependencies
```
go get ./...
```
or
```
go get github.com/gorilla/mux
go get github.com/go-pg/pg
```
## DEVELOPMENT
```
go build
./folder_name
```
or
```
go build -o main
./main
```
## PRODUCTION
```
docker-compose up -d
```
go to the http://localhost:9090/api/v1/courses and you see the list of avaliable courses