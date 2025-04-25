CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        points INTEGER NOT NULL DEFAULT 0,
        referrer INTEGER REFERENCES users (id),
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        points INTEGER NOT NULL,
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS completed_tasks (
        id SERIAL PRIMARY KEY,
        user_id INTEGER REFERENCES users (id),
        task_id INTEGER REFERENCES tasks (id),
        completed_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            UNIQUE (user_id, task_id)
    );

-- Создаем индекс для оптимизации запросов к таблице completed_tasks
CREATE INDEX IF NOT EXISTS idx_completed_tasks_user_id ON completed_tasks (user_id);

-- Создаем индекс для оптимизации запросов к таблице users по points
CREATE INDEX IF NOT EXISTS idx_users_points ON users (points DESC);