CREATE TABLE xkcd (
  id BIGSERIAL PRIMARY KEY,
  month CHAR(2) NOT NULL,
  num INT NOT NULL UNIQUE,
  link VARCHAR(255),
  year CHAR(4) NOT NULL,
  news VARCHAR(255),
  safe_title VARCHAR(255),
  transcript TEXT NOT NULL,
  alt TEXT,
  img VARCHAR(255),
  title VARCHAR(255),
  day CHAR(2) NOT NULL
);

CREATE INDEX transcript ON xkcd (transcript);