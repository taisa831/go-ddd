# go-ddd

## LIST

```
curl -X GET http://localhost:8081/users -H "content-type:application/json"
```

## GET

```
curl -X GET http://localhost:8081/users/ec51aa4e-08d8-42be-a9d6-9ec0af54c83d -H "content-type:application/json"
```

## POST

```
curl -X POST http://localhost:8081/users -H "content-type:application/json" -d '{ "name": "test" }'
```

## PATCH

```
curl -X PATCH http://localhost:8081/users/ec51aa4e-08d8-42be-a9d6-9ec0af54c83d -H "content-type:application/json" -d '{ "name": "test3", "address": "address" }'
```

## DELETE

```
curl -X DELETE http://localhost:8081/users/ec51aa4e-08d8-42be-a9d6-9ec0af54c83d -H "content-type:application/json"
```

# usecase

## Create User

```
curl -X POST http://localhost:8081/users -H "content-type:application/json" -d '{ "name": "test" }'
```

## Create Circle

```
curl -X POST http://localhost:8081/circles?userId={userId}&circleName={circleName} -H "content-type:application/json" -d '{}'
```

## Join Circle

```
curl -X POST http://localhost:8081/circles/{circleId}/join?userId={userId} -H "content-type:application/json" -d '{}'
```
