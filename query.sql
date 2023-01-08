-- name: GetPost :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id;

-- name: CreatePost :one
INSERT INTO posts (
  title, text
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = ? LIMIT 1;

-- name: ListCommentsInPost :many
SELECT * FROM comments WHERE post_id = ?
ORDER BY id;

-- name: CreateComment :one
INSERT INTO comments (
  text, post_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = ?;

-- name: GetNewName :one
SELECT new_name();
