CREATE TABLE wallets (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         balance NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Пример данных
INSERT INTO wallets (id, balance) VALUES ('123e4567-e89b-12d3-a456-426614174000', 1000.00);