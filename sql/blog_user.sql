CREATE TABLE user
(
    id         int AUTO_INCREMENT
        PRIMARY KEY,
    user_name  varchar(64)                         NOT NULL,
    password   varchar(255)                        NOT NULL,
    nick_name  varchar(64)                         NULL,
    email      varchar(64)                         NULL,
    avatar     varchar(255)                        NULL,
    deleted    tinyint   DEFAULT 0                 NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT user_user_name_uindex
        UNIQUE (user_name)
);

