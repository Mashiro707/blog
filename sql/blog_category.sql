CREATE TABLE category
(
    id            bigint AUTO_INCREMENT PRIMARY KEY,
    name          varchar(64)                         NOT NULL,
    deleted       tinyint   DEFAULT 0                 NOT NULL,
    created_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT category_name_uindex
        UNIQUE (name)
);
