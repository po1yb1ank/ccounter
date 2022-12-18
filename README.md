# Counters API

Service that provides API for creating and updating custom-labeled counters.
Each counter is a tuple of `key` and `value`

## API

`/swagger/index.html` - API description

## Handlers

### /{key}

**Schema**: `HTTP`

**Methods**: `GET`

Returns counter `key` value. If there no counter - sets a new one with default value `0`

### Ok Response

***Status***: `200 Ok`

```json
{
    "message": "Successful get key value action",
    "status": 200,
    "key": "{key}",
    "value": 123
}
```
### Error Responses

***Status***: `500 Internal`

###
```json
{
    "message": "failed to get counter value",
    "status": 500
}
```

### /{key}/increment

**Schema**: `HTTP`

**Methods**: `POST`

Increments counter value by `1`. Returns `404 Not found` if this counter not exists.

### Ok Response

***Status***: `200 Ok`

```json
{
    "message": "Successful increment key value action",
    "status": 200,
    "key": "{key}",
    "value": 123
}
```
### Error Responses

***Status***: `404 Not Found`

```json
{
    "message": "key not found",
    "status": 404
}
```

***Status***: `500 Internal`

```json
{
    "message": "failed to increment value",
    "status": 500
}
```

### /{key}/decrement

**Schema**: `HTTP`

**Methods**: `POST`

Decrements counter value by `1`. Returns `404 Not found` if this counter not exists.

### Ok Response

***Status***: `200 Ok`

```json
{
    "message": "Successful decrement key value action",
    "status": 200,
    "key": "{key}",
    "value": 123
}
```
### Error Responses

***Status***: `404 Not Found`

```json
{
    "message": "key not found",
    "status": 404
}
```

***Status***: `500 Internal`

```json
{
    "message": "failed to decrement value",
    "status": 500
}
```

### /{key}/reset

**Schema**: `HTTP`

**Methods**: `POST`

Resets counter value. If no counter presents - sets to `0` by default.

### Ok Response

***Status***: `200 Ok`

```json
{
    "message": "Successful reset key value action",
    "status": 200,
    "key": "{key}",
    "value": 123
}
```
### Error Responses

***Status***: `500 Internal`

```json
{
    "message": "failed to reset counter value",
    "status": 500
}
```

### /subscribe

**Schema**: `WS`

After establishing ws-connection server will notify about all counters updates.

```json
{
    "key": "some-key",
    "value": 123
}
```


