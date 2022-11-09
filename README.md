# avito_balance
## Микросервис для работы с балансом пользователей

**Запуск:**

```bash
docker-compose up --build balance
^C
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```
Если приложение запусается повторно:
```bash
docker-compose up balance
```
Swagger доступен по ссылке:
http://localhost:8000/swagger/index.html#/

Адрес для Postman:  
localhost:8000/bill/

**Примеры:**



**Сценарии использования:**
