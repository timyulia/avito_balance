# avito_balance

## Запуск

```bash
docker-compose up --build balance

docker-compose up balance
```

Swagger доступен по ссылке:  
http://localhost:8000/swagger/index.html#/

Адрес для Postman:    
localhost:8000/bill/

## Примеры

**1. Запрос истории пользователя**  
GET  
localhost:8000/bill/info/history/14/amount?page=5&limit=2

Ответ: 
```bash
{  
"data": [  
{  
"reason": "write-off for the service without name",  
"amount": 7000,  
"date": "2022-11-09"  
},  
{  
"reason": "account replenishment: string",  
"amount": 10000,  
"date": "2022-11-09"  
}  
]  
}  
```

Параметры пагинации не предусмотрены в Swagger, но работают в Postman.

**2. Запрос баланса пользователя**
GET  
localhost:8000/bill/14

Ответ:  
983560

**3. Запрос отчета**  
GET  
localhost:8000/bill/info/report/2022/11  

Ответ:  
```bash
{  
"link": "/bill/info/report"  
}
```

Можно перейти по ссылке, кликнув на нее в Postman. Затем после отправки запроса в поле ответа появится файл.

**4. Запрос на добавление названия услуги**  
POST    
localhost:8000/bill/info/specify  

```bash
{  
"name": "med",  
"service_id": 2  
}
```

Ответ: 
```bash 
{  
"status": "ok"  
}
```
Для удобства составления отчетов создана специальная функция для определения названий сервисов.

**5. Запрос на разрезервирование денег**  
POST    
localhost:8000/bill/return

```bash
{
  "order_id": 2,
  "user_id": 14
}
```

Ответ:
```bash 
{  
"status": "ok"  
}
```

**6. Запрос на резервирование денег**  
   POST    
   localhost:8000/bill/reserve

```bash
{
  "amount": 7000,
  "order_id": 2,
  "service_id": 2,
  "user_id": 14
}
```

Ответ:
```bash 
{  
"status": "ok"  
}
```

**7. Запрос на списание денег**  
   POST    
   localhost:8000/bill

```bash
{
  "amount": 7000,
  "order_id": 2,
  "service_id": 2,
  "user_id": 14
}
```

Ответ:
```bash 
{  
"status": "ok"  
}
```

Метод проверяет соответствие всех полей заказа с зарезервированным счетом. Если что-то не совпадает, списание не произойдет.
