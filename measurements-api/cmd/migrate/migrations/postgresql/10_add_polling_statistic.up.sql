CREATE TABLE polling_statistics (
    polling_statistic_id SERIAL    NOT NULL PRIMARY KEY,
    datetime             TIMESTAMP NOT NULL,
    duration             INTERVAL  NOT NULL,
    post_count           INT       NOT NULL,
    received_count       INT       NOT NULL
);
