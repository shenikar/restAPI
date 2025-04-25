-- Добавляем тестового пользователя с хешированным паролем (bcrypt hash для 'password')
INSERT INTO
    users (username, email, password, points)
VALUES
    (
        'testuser',
        'test@example.com',
        '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
        0
    ) ON CONFLICT (username) DO NOTHING;