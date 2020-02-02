# Golang practice

## Description

練習Golang，提供簡易 `RESTful API`。

## Quick Start

### Develop with Docker

1. Start environment
```
$ docker-compose up -d
```

2. Install dependencies which needed
```
$ docker-compose exec api govendor sync
```

3. Start to run api service
```
$ docker-compose exec api go run main.go
```

4. Stop environment
```
$ docker-compose down
```
