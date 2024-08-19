CREATE TABLE configurations (
    configuration_id INT PRIMARY KEY NOT NULL,
    asoiza_login VARCHAR(2048) NOT NULL,
    asoiza_password VARCHAR(2048) NOT NULL,
    collecting_interval INT NOT NULL DEFAULT 1200, -- 20 minutes
    deleting_interval INT NOT NULL DEFAULT 2592000, -- 30 days
    deleting_threshold INT NOT NULL DEFAULT 94608384, -- 3 years
    disabling_interval INT NOT NULL DEFAULT 21600, -- 6 hours
    disabling_threshold INT NOT NULL DEFAULT 604800 -- 7 days
);

INSERT INTO configurations (configuration_id, asoiza_login, asoiza_password)
VALUES (1, 'login', 'password')