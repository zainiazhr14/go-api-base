-- +goose Up
-- +goose StatementBegin
CREATE TYPE purpose AS ENUM ('email_verification', 'password_reset');

CREATE TABLE one_time_passwords (
    id SERIAL PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code VARCHAR(10) NOT NULL,
    purpose purpose NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE one_time_passwords;
-- +goose StatementEnd
