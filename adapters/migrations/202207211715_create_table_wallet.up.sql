CREATE TABLE IF NOT EXISTS wallet (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL UNIQUE,
    owner_id TEXT NOT NULL,
    status TEXT NOT NULL,
    balance INTEGER DEFAULT 0,
    status_updated_at TEXT NOT NULL,
    FOREIGN KEY (account_id)
       REFERENCES account (id) 
);