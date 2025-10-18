-- +goose Up
-- +goose StatementBegin
-- Create a sample table for event store
CREATE TABLE IF NOT EXISTS event_store (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    aggregate_id VARCHAR(36) NOT NULL,
    aggregate_type VARCHAR(255) NOT NULL,
    event_type VARCHAR(255) NOT NULL,
    event_data JSON NOT NULL,
    version BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_aggregate_id (aggregate_id),
    INDEX idx_aggregate_type (aggregate_type),
    UNIQUE KEY unique_aggregate_version (aggregate_id, version)
);

-- Create a sample table for projections/read models
CREATE TABLE IF NOT EXISTS read_models (
    id VARCHAR(36) PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    data JSON NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_type (type)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS read_models;
DROP TABLE IF EXISTS event_store;
-- +goose StatementEnd
