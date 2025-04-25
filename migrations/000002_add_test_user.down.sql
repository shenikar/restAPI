-- Удаляем тестового пользователя
DELETE FROM users
WHERE
    username = 'testuser';