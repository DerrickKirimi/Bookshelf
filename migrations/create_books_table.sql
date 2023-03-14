--+goose up
CREATE TABLE IF NOT EXISTS books 
(
    id              INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    title           VARCHAR(255)    NOT NULL,
    author          VARCHAR(255)    NOT NULL,
    published_date  DATE            NOT NULL,
    image_url       VARCHAR(255)    NULL,
    created_at      TIMESTAMP       NULL,
    modified_at     TIMESTAMP       NULL, 
    deleted_at      TIMESTAMP       NULL,
    PRIMARY KEY (id)
);

--goose down 
DROP TABLE IF EXISTS books;