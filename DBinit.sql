CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    balance NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Пример данных
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM wallets WHERE id = '123e4567-e89b-12d3-a456-426614174000') THEN
        INSERT INTO wallets (id, balance) VALUES
            ('123e4567-e89b-12d3-a456-426614174000', 1000.00),
            ('223e4567-e89b-12d3-a456-426614174001', 500.00),
            ('323e4567-e89b-12d3-a456-426614174002', 0.00);
    END IF;
END $$;