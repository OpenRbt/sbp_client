CREATE USER wash_sbp_service WITH PASSWORD 'wash_sbp_password';
CREATE DATABASE wash_sbp_db;
GRANT ALL PRIVILEGES ON DATABASE wash_sbp_db TO wash_sbp_service;