# Counters API

# TODO
- Description of methods
- Tests with server mock
- Swagger API

## **GET /:key**

Возвращает текущее значение счетчика с указанным `key` (по умолчанию: 0).

## **POST /:key/increment**

Увеличивает значение счетчика на 1.

## **POST /:key/decrement**

Уменьшает значение счетчика на 1.

## **POST /:key/reset**

Сбрасывает значение счетчика.

## Websocket **/subscribe**

