-- +goose Up
-- +goose StatementBegin\
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id  uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  full_name VARCHAR(100) not_null,
  email VARCHAR (100) not_null,
  verified_at DATETIME,
  
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
