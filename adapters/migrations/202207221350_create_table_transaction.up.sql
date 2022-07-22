CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY,
    wallet_id TEXT, 
    status TEXT NOT NULL,
    amount INTEGER DEFAULT 0,
    reference_id TEXT NOT NULL,
    transaction_type TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    updated_by TEXT NOT NULL,
    FOREIGN KEY (wallet_id)
       REFERENCES wallet (id),
    UNIQUE(reference_id, transaction_type)
);