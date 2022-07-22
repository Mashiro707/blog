CREATE TABLE category
(
    id           bigint AUTO_INCREMENT PRIMARY KEY,
    name         varchar(64)                         NOT NULL,
    deleted      boolean   DEFAULT false                 NOT NULL,
    create_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT category_name_uindex
        UNIQUE (name, deleted)
);
