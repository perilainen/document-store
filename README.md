# Document Store

## Start Application

```
go run main.go
```

Application runs on port 8080

## Test Application

Documents have the current form and must be supplied in the body when making post and put requests.

```JSON
{
"title": "Employee contract",
"content": {
    "header": "Employee contract",
    "data": "I agree to the terms"
},
"signee": "John McClane"
}
```

### Find documents

#### All documents

GET /documents

#### Document by id

GET /document/{id}

### Create new document

POST /document/{body}

example creating document:

```
curl -d '{
    "title": "Employee contract",
        "content": {
            "header": "Employee contract",
            "data": "I agree to the terms"
        },
    "signee": "John McClane"}' -H "Content-Type: application/json" -X POST http://localhost:8080/document
```

### Update document

PUT /document/{body}
example updating document with id=1:

```
curl -d '{
    "title": "Employee contract",
        "content": {
            "header": "Employee contract",
            "data": "I agree to the terms"
        },
    "signee": "John McClane"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/document/1
```
