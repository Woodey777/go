CREATE TABLE IF NOT EXISTS projects 
(
  id INT NOT NULL,
  name CHAR(200) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY id
);

CREATE INDEX IF NOT EXISTS projects_id ON projects(id);
