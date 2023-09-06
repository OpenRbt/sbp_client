CREATE USER wash_sbp_service WITH PASSWORD 'wash_sbp_password';
CREATE DATABASE wash_sbp;
GRANT ALL PRIVILEGES ON DATABASE wash_sbp TO wash_sbp_service;