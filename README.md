# go-mysql-rest-api
a set of APIs to be used by front-end engineers to develop an application that store and display tax amounts, Tech stack : Go, Docker-Compose, MySQL, HTML, CSS

This API serve 3 routes at Order Collection [/order/{store_id}], Cart Collection [/cart], Tax Code Collection [/tax_code] .
- You can perform fetch List Tax Code [GET] from Tax Code Collection
- Add to cart or Create Tax Object Add Cart [POST] from Cart Collection
- And can get/view the bill List All Order Item [GET] from Order Collection, from here we use store_id=1 hardcode at GUI add cart
so we can call [/order/1] to fetch my bill

The frontend show some GUI for add cart / add tax object, status update add cart, view my bill

See [API Documentation](https://github.com/boantp/go-mysql-rest-api/blob/master/apiary.apib) on how to use it.

## Directory Structure
```
go-mysql-rest-api
    |--config                   - to initialize template and database
        |--db.go                - for initialize mysql database connection
        |--tpl.go               - for template view configuration
    |--controllers              - to store package controllers
        |--cart.go              - to handle Cart Collection [/cart]
        |--order.go             - to handle Order Collection [/order/{store_id}]
        |--tax_code.go          - to handle Tax Code Collection [/tax_code]
        |--web.go               - to handle Web/Frontend for GUI add cart["/"], GUI view bill ["order_view/:store_id]
    |--docker                   - Dockerfile for Golang at folder web, Dockerfile for MySQL at folder db
        |--db
            |--Dockerfile
        |--web
            |--Dockerfile
    |--models                    - to store package models for object and mysql query
        |--order_details.go      - for table order_details
        |--orders.go             - for total amount, total tax amount, grand total object, and for table orders
        |--tax_code.go           - for table tax_code
    |--mysql_init                - init Table shopee
        |--shopee.sql
    |--templates                 - to store html file for golang *gohtml
    |--apiary.apib               - json file docs from APIARY for API DOCS
    |--database_design.png       - DB structure and explanation
    |--docker-compose.yml        - for docker-compose service and config
    |--main.go                   

  
```

## Setup

**Steps**
1. git clone [git@github.com:boantp/go-mysql-rest-api.git](git@github.com:boantp/go-mysql-rest-api.git)
2. install docker and docker-compose 
3. open terminal and run docker-compose build (service are build), docker-compose up(builds, recreates, attaches to container for service), docker-compose down (stop containers and remove containers etc) See [Docker Documentation](https://docs.docker.com/compose/reference/build/) on how to use it.
4. now your server ready for http:localhost:3000/
