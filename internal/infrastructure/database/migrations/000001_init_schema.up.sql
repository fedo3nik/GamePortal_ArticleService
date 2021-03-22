CREATE TABLE IF NOT EXISTS articles(
    ID serial PRIMARY KEY,
    userID varchar(255),
    title varchar(255),
    game varchar(255),
    article_text text
);

CREATE TABLE IF NOT EXISTS rating(
    ID serial PRIMARY KEY,
    articleID serial,
    rating double precision,
    CONSTRAINT fk_article FOREIGN KEY(articleID) REFERENCES Articles(ID)
);