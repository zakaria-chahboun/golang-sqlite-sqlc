CREATE TABLE posts (
  id   INTEGER PRIMARY KEY,
  title text    NOT NULL,
  text text NOt NULL,
  created_at datetime default CURRENT_TIMESTAMP
);

CREATE TABLE comments (
  id   INTEGER PRIMARY KEY,
  text text    NOT NULL,
  post_id  INTEGER NOT NULL,
  created_at datetime default CURRENT_TIMESTAMP,
  FOREIGN KEY(post_id) REFERENCES posts(id) ON UPDATE CASCADE ON DELETE CASCADE
);


