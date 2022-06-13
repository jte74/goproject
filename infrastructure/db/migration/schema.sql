CREATE TABLE Users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  firstname  TEXT NOT NULL,
  age INTEGER NOT NULL,
  password TEXT NOT NULL,
  token TEXT,
  datecreated DATE
);