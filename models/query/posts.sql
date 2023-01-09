-- name: GetPost :one
SELECT *
FROM posts
WHERE id = ?
LIMIT 1;

-- name: ListPosts :many
SELECT *
FROM posts
ORDER BY created_at;

-- name: CreatePost :one
INSERT INTO posts (id, title, text)
VALUES (new_id('p'), ?, ?)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?;
