# GO-market

Цель проекта "GoMarket" - предоставить жителям удобный инструмент для управления и отслеживания своих покупок и продаж через консольный интерфейс. Пользователи смогут регистрироваться, добавлять товары, просматривать список товаров, 
совершать покупки, а также управлять своим аккаунтом.

Используемые технологии:
- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса)
- Swagger (для документации API)
- Gin (веб фреймворк)
- golang-migrate/migrate (для миграций БД)
- pgx (драйвер для работы с PostgreSQL)
- golang/mock, testify (для тестирования)

Сервис был написан с Clean Architecture, что позволяет легко расширять функционал сервиса и тестировать его.
Также был реализован Graceful Shutdown для корректного завершения работы сервиса

# Начало работы

Для запуска сервиса необходимо предварительно:
- Опционально, настроить `congig/config.yaml` под себя

# Использование

Запустить сервис можно с помощью команды `make compose-up`

Документацию после завпуска сервиса можно посмотреть по адресу `http://localhost:8080/swagger/index.html`
с портом 8080 по умолчанию

Для запуска тестов необходимо выполнить команду `make test`, для запуска тестов с покрытием `make cover`

Для запуска линтера необходимо выполнить команду `make linter-golangci`

## Примеры

Некоторые примеры запросов
- [Регистрация](#sign-up)
- [Аутентификация](#sign-in)
- [Добавление продукта](#add-product)
- [Получение информации о всех продуктах](#get-products)
- [Получение информации о продукте по его ID](#get-product)
- [Обновление информации о продукте по его ID](#update-product)
- [Удаление информации о продукте по его ID](#delete-product)
- [Совершение покупки](#operations-report-link)
- [Получение всех покупках пользователя по его ID](#operations-report-file)
- [Получение всех покупках продукта по его ID](#operations-report-file)

### Регистрация <a name="sign-up"></a>

Регистрация сервиса:
```curl
curl -X 'POST' \
  'http://localhost:8080/auth/sign-up' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "example@gmail.com",
  "password": "Qwerty1!",
  "username": "test"
}'
```
Пример ответа:
```json
{
  "id": 3
}
```

### Аутентификация <a name="sign-in"></a>

Аутентификация сервиса для получения токена доступа:
```curl
curl -X 'POST' \
  'http://localhost:8080/auth/sign-in' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "password": "Qwerty1!",
  "username": "test"
}'
```
Пример ответа:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4"
}
```

### Добавление продукта <a name="add-product"></a>

Добавление продукта:
```curl
curl -X 'POST' \
  'http://localhost:8080/api/v1/products/add-product' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4' \
  -H 'Content-Type: application/json' \
  -d '{
  "description": "Вода негазированная 0.5л",
  "name": "Святой источник",
  "price": 59.9,
  "quantity": 100
}'
```
Пример ответа:
```json
{
  "id": 4
}
```

### Получение информации о всех продуктах <a name="get-products"></a>

Получение информации о всез продуктах:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/v1/products/get-products' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4'
```
Пример ответа:
```json
{
  "Products": [
    {
      "ID": 4,
      "Name": "Святой источник",
      "Description": "Вода негазированная 0.5л",
      "Price": 59.9,
      "Quantity": 100
    }
  ]
}
```

### Получение информации о продукте по его ID <a name="get-product"></a>

Получение информации о продукте по его ID:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/v1/products/get-product/4' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4'
```
Пример ответа:
```json
{
  "product": {
    "ID": 4,
    "Name": "Святой источник",
    "Description": "Вода негазированная 0.5л",
    "Price": 59.9,
    "Quantity": 100
  }
}
```

### Обновление информации о продукте по его ID <a name="update-product"></a>

Обновление информации о продукте по его ID:
```curl
curl -X 'PUT' \
  'http://localhost:8080/api/v1/products/update-product/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4' \
  -H 'Content-Type: application/json' \
  -d '{
  "description": "Вода газированная 1.0л",
  "name": "Bon Aqua",
  "price": 69.9,
  "quantity": 200
}'
```
Пример ответа:
```json
{
  "message": "succes"
}
```

### Удаление информации о продукте по его ID <a name="delete-product"></a>

Удаление информации о продукте по его ID:
```curl
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/products/delete-product/4' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4'
```
Пример ответа:
```json
{
  "message": "succes"
}
```

### Совершение покупки <a name="make-purchase"></a>

Совершение покупки:
```curl
curl -X 'POST' \
  'http://localhost:8080/api/v1/purchase/make-purchase' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4' \
  -H 'Content-Type: application/json' \
  -d '{
  "product_id": 4,
  "quantity": 20,
  "user_id": 3
}'
```
Пример ответа:
```json
{
  "id": 2
}
```

### Получение всех покупках пользователя по его ID <a name="get-user-purchase"></a>

Сервис формирует отчёт и возвращает его в виде csv файла:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/v1/purchase/get-user-purchase/3' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4'
```
Пример ответа:
```json
{
  "Purchases": [
    {
      "ID": 2,
      "UserID": 3,
      "ProductID": 4,
      "Quantity": 20,
      "Timestamp": "2024-09-05T09:41:07.810032Z"
    }
  ]
}
```

### Получение всех покупках продукта по его ID <a name="get-product-purchase"></a>

Получение всех покупках продукта по его ID:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/v1/purchase/get-product-purchase/4' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU1MzkyNDMsImlhdCI6MTcyNTUyODQ0MywiVXNlcklEIjozLCJVc2VybmFtZSI6InRlc3QifQ.VJS-PK2xp1Ks7_KjCulyKVDpBT5Q8g7f_Wlj-rTrcg4'
```
Пример ответа:
```json
{
  "Purchases": [
    {
      "ID": 2,
      "UserID": 3,
      "ProductID": 4,
      "Quantity": 20,
      "Timestamp": "2024-09-05T09:41:07.810032Z"
    }
  ]
}
```