# About

**Basic REST API** is simple CRUD rest API for managing _Orders_ entitiy

## Environment setup

This application are built with **Gin** and the latest version of **GORM** from gorm.io

- Postgres SQL with a DB named `basic_rest`
- Port `6969` available
- Gin
- GORM

## Installation

`git clone -b basic-rest https://github.com/wahyusa/go-learn.git
`

I am using GORM auto migration so DB table should be auto created but if you need the SQL file I also provide it on this branch.

## Usage

**Running gin server in port 6969 on debug mode**

```bash
go run main.go
```

## Testing

While I don't really implement testing, for those who want to test this API, you may try to use basic CURL that usually already available on any OS.

Some request and response are sent around `application/json` header.

Request and response requirements are based on this https://anotepad.com/notes/3sjp4bg3

CURL Testing command are expected to run in shell that support `backslash` for multiline command. (Git bash should works)

### Step 1 GET Orders (First time)

Expect: No data []

```bash
curl -X GET http://localhost:6969/orders
```

### Step 2 POST Orders

Expect: The required responses

```bash
curl -X POST http://localhost:6969/orders \
-H "Content-Type: application/json" \
-d '{
    "orderedAt": "2021-10-06T16:53:27.675931+07:00",
    "customerName": "Test",
    "items": [
        {
            "itemCode": "Item1",
            "description": "ItemDescription",
            "quantity": 1
        },
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        }
    ]
}'
```

### Step 3 GET Orders (Check if POSTed request are saved to DB)

Expect:
- At least one data that we've POST above will appear here
- The result usually start with ID 1

```bash
curl -X GET http://localhost:6969/orders
```

### Step 4 Update Order PUT

Expect:
- The required response
- If Order ID not found, sent NOT FOUND 404

```bash
curl -X PUT http://localhost:6969/orders/1 \
-H "Content-Type: application/json" \
-d '{
    "orderedAt": "2021-10-06T16:53:27.675931+07:00",
    "customerName": "Test updated from put",
    "items": [
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        },
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        }
    ]
}'
```

### Step 5 Check if Order data was updated by GET Order again

Expect: At least the `customeName` was updated

```bash
curl -X GET http://localhost:6969/orders
```

### Step 6 DELETE Order

Expect: Success delete

```bash
curl -X DELETE http://localhost:6969/orders/1
```

### Step 7 Check if Order data DELETED

Expect: No data []

```bash
curl -X GET http://localhost:6969/orders
```