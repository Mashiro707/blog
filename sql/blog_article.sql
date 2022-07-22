CREATE TABLE article
(
    id           bigint AUTO_INCREMENT PRIMARY KEY,
    title        varchar(255)                        NOT NULL,
    content      text                                NULL,
    description  text                                NULL,
    is_top       boolean   DEFAULT false             NOT NULL,
    is_comment   boolean   DEFAULT false             NOT NULL,
    category_id  bigint                              NOT NULL,
    deleted      boolean   DEFAULT false             NOT NULL,
    create_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_time  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);
