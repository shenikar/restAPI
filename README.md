# User Management API

Простой HTTP сервер для управления пользователями на языке Go.

## Функциональность

- Middleware авторизация по JWT токену
- Получение информации о пользователе
- Просмотр таблицы лидеров
- Выполнение заданий
- Ввод реферального кода

## API Endpoints

### GET /users/{id}/status
Получение информации о пользователе

### GET /users/leaderboard
Получение таблицы лидеров (топ 10 пользователей)

### POST /users/{id}/task/complete
Выполнение задания

Тело запроса:
```json
{
    "task_id": 1
}
```

### POST /users/{id}/referrer
Ввод реферального кода

Тело запроса:
```json
{
    "referrer_id": 1
}
```

## Запуск приложения

1. Убедитесь, что у вас установлены Docker и Docker Compose
2. Клонируйте репозиторий
3. Запустите приложение:
```bash
docker-compose up --build
```

Приложение будет доступно по адресу: http://localhost:8080

## Тестирование

Для тестирования API можно использовать Postman или curl. Примеры запросов:

```bash
# Получение информации о пользователе
curl -X GET http://localhost:8080/users/1/status \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Получение таблицы лидеров
curl -X GET http://localhost:8080/users/leaderboard \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Выполнение задания
curl -X POST http://localhost:8080/users/1/task/complete \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"task_id": 1}'

# Ввод реферального кода
curl -X POST http://localhost:8080/users/1/referrer \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"referrer_id": 2}'
``` 