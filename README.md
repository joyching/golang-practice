# Golang practice

## Description

練習Golang，提供簡易 `RESTful API`。

<img src="https://i.imgur.com/GnpyN3K.png" width="100%" />

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

## Web API

支援以下API:

* **GET**      `/books`        取得所有書籍
* **GET**      `/books/{id}`   取得特定書籍
* **POST**     `/books`        新增書籍
* **PUT**      `/books/{id}`   更新書籍資料
* **DELETE**   `/books/{id}`   刪除書籍
* **GET**      `/health`       確認服務正常運作

### GET /books

取得所有書籍

**Response body**

``` json
{
    "data": [
        {
            "id": 16621,
            "tilte": "無瑕的程式碼：敏捷軟體開發技巧守則",
            "published_at": "2013-03-22"
        },
        {
            "id": 18627,
            "tilte": "單元測試的藝術",
            "published_at": "2017-09-29"
        },
        {
            "id": 20490,
            "tilte": "Modern PHP: New Features and Good Practices",
            "published_at": "2015-03-01"
        }
    ]
}
```

### GET /books/{id}

取得特定書籍

**Response body**

``` json
{
    "id": 16621,
    "tilte": "無瑕的程式碼：敏捷軟體開發技巧守則",
    "published_at": "2013-03-22"
}
```

### POST /books

新增書籍

**Request body**

| name      | type   | description  |
|-----------|--------|--------------|
| title | string | 書名 |
| published_at | date | 出版日期 |

**Response body**

``` json
{
    "id": 16621,
    "tilte": "無瑕的程式碼：敏捷軟體開發技巧守則",
    "published_at": "2013-03-22"
}
```

### PUT /books/{id}

更新書籍資料

**Request body**

| name      | type   | description  |
|-----------|--------|--------------|
| title | string | 書名 |
| published_at | date | 出版日期 |

**Response body**

``` json
{
    "id": 16621,
    "tilte": "無瑕的程式碼：敏捷軟體開發技巧守則",
    "published_at": "2013-03-22"
}
```

### DELETE /books/{id}

刪除書籍

**Response body**

``` json
{
    "message": "Delete Book ID:16621 Success"
}
```

### GET /health

確認服務正常運作

**Response body**

``` json
{
    "status": "ok"
}
```
