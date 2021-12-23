CREATE TABLE IF NOT EXISTS ports (
    id INTEGER PRIMARY KEY,
    name varchar(60),
    is_active   bool,
    company    varchar(60),
    email      varchar(60),
    phone      varchar(60),
    address    varchar(60),
     about      varchar(500),
    registered varchar(60),
     latitude   float,
      longitude  float
);

