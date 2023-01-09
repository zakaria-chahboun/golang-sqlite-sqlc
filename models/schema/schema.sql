CREATE TABLE
  posts (
    id VARCHAR PRIMARY KEY,
    title text NOT NULL,
    text text NOt NULL,
    created_at datetime default CURRENT_TIMESTAMP
  );

CREATE TABLE
  comments (
    id VARCHAR PRIMARY KEY,
    text text NOT NULL,
    post_id VARCHAR NOT NULL,
    created_at datetime default CURRENT_TIMESTAMP,
    CONSTRAINT "comments_post_id_posts_fkey" FOREIGN KEY (post_id) REFERENCES posts (id) ON UPDATE CASCADE ON DELETE CASCADE
  );

CREATE INDEX comments_post_id_idx ON comments(post_id);
