CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    preview TEXT NOT NULL,
    text TEXT NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id), -- предполагается, что есть таблица users
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    categories TEXT[] DEFAULT '{}', -- массив категорий
    alias VARCHAR(255) UNIQUE NOT NULL,
    CONSTRAINT alias_check CHECK (alias ~ '^[a-z0-9-]+$') -- только латиница, цифры и дефисы
);
