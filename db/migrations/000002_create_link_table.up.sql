CREATE TABLE IF NOT EXISTS links (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    short varchar(16) UNIQUE NOT NULL,
    long_url varchar(256) NOT NULL,
    active BOOLEAN DEFAULT TRUE NOT NULL,
    visited INT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);
