CREATE TABLE IF NOT EXISTS prices
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255)   NOT NULL,
    category    VARCHAR(255)   NOT NULL,
    price       NUMERIC(10, 2) NOT NULL,
    create_date DATE           NOT NULL
);