-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id             BIGSERIAL PRIMARY KEY,              -- Уникальный идентификатор пользователя
    username       VARCHAR(255) NOT NULL,              -- Имя пользователя
    status         BOOLEAN   DEFAULT FALSE,            -- Статус (онлайн/оффлайн)
    last_active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Время последней активности
);

CREATE TABLE IF NOT EXISTS chats
(
    id         BIGSERIAL PRIMARY KEY,              -- Уникальный идентификатор чата
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Временная метка создания чата
);

CREATE TABLE IF NOT EXISTS messages
(
    id           BIGSERIAL PRIMARY KEY,                                 -- Уникальный идентификатор сообщения
    sender_id    BIGINT NOT NULL,                                          -- ID отправителя
--     recipient_id BIGINT,                                                   -- ID получателя (NULL для групповых чатов)
    chat_id      BIGINT,                                                   -- ID чата (NULL для личных сообщений)
    text         TEXT,                                                  -- Текст сообщения (макс. 2000 символов)
    images       TEXT[],                                                -- Список изображений (макс. 3 пути к изображениям)
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                   -- Временная метка отправки
    CONSTRAINT chk_text_length CHECK (LENGTH(text) <= 2000),            -- Ограничение длины текста
    CONSTRAINT chk_images_count CHECK (array_length(images, 1) <= 3),   -- Ограничение количества изображений
    FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE,    -- Связь с отправителем
--     FOREIGN KEY (recipient_id) REFERENCES users (id) ON DELETE CASCADE, -- Связь с получателем
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE       -- Связь с чатом
);



CREATE TABLE IF NOT EXISTS chat_members
(
    chat_id   BIGINT NOT NULL,                                        -- ID чата
    user_id   BIGINT NOT NULL,                                        -- ID участника
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                 -- Временная метка присоединения
    PRIMARY KEY (chat_id, user_id),                                -- Составной первичный ключ
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE, -- Связь с таблицей чатов
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE  -- Связь с таблицей пользователей
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS chat_members;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
