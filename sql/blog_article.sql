CREATE TABLE article
(
    id            bigint AUTO_INCREMENT PRIMARY KEY,
    title         varchar(64)                         NOT NULL,
    cover         varchar(255)                        NOT NULL,
    content       text                                NULL,
    description   text                                NULL,
    status        tinyint   DEFAULT 0                 NOT NULL,
    is_top        tinyint   DEFAULT 0                 NOT NULL,
    is_comment    tinyint   DEFAULT 0                 NOT NULL,
    category_id   bigint                              NOT NULL,
    deleted       tinyint   DEFAULT 0                 NOT NULL,
    created_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);
