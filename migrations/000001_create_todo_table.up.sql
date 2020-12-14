CREATE TABLE IF NOT EXISTS todos (
  todo_id serial PRIMARY KEY,
  name VARCHAR(60) NOT NULL,
  details TEXT,
  done BOOLEAN DEFAULT false
);
