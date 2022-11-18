
## 點餐速 Order-Fast
---
# 點餐！！速度！！

This project is design for link-base order system via scan QRcode.

 I'm inspired while ordering foods from a busy resturant, it's crazy and really time consuming to get my food, so I hope this project can help each resturants to speed up their order flows and save everyone time ;).

 Backend Server is base on **Golang Gin Framework**, Frontend Server is base on **React**.

## Running the OF service
Launch the service
```
docker-compose up
```
DB migration up
```
make migrateup
```
DB migration down
```
make migratedown
```
Generate psql connecting interface for golang
```
make sqlc
```
