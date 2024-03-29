
## 點餐速 Order-Fast
---
# 點餐！！速度！！

This project is design for link-base order system via scan QRcode.

 I'm inspired while ordering foods from a damn busy resturant, it's time consuming before I got my food, so I thought the Order Fast might be helpful for each resturants to speed up their order flows and save everyone time ;).


> ***Chief focus on the cuisine present. Customer waiting for the tasty food with finger-licking ;).***

**合作請求**: 如果你有前端的經驗並對這項專案有興趣的話請和我聯繫，我們可以討論一起完成這個專案:)

**collabrate Request**: If you're frontend engineer and interesting to participant this project, please contact me for further discussion :).

聯繫信箱 Contact Email: harryuwang@gmail.com


# ***Backend Development Environmet Setup***

 >*Backend Server is base on **Golang Gin Framework**, Frontend Server is base on **React**.*

## *Install required package*
```
go mod download
```
## *Running the Order-Fast backend service**
*Launch the db service*
```
docker-compose up
```
*DB migration up*
```
make migrateup
```
*DB migration down*
```
make migratedown
```
*Generate psql connecting interface for golang*
```
make sqlc
```
*Running tests*
```
make test
```
*Launch the gin backend server*
```
make server
```


- HTTP Gateway is listen on port 8081
- GRPC is listen on port 8082
- Gin Server is listen on port 8083
> you can go to **localhost:8081/swagger/grpc** to review the API docs

#

*If you change the content of db/sql/., running the following cmd which generate the Querier(db mock) for api testing*
```
make mock
```

### How to generate code
- Create a new db migration:

	```bash
	migrate create -ext sql -dir db/migration -seq <migration_name>
	```
###
# ***Frontend Development Setup***
TBD
