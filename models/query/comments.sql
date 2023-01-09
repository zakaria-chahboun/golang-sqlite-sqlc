-- name: GetComment :one
SELECT *
FROM comments
WHERE id = ?
LIMIT 1;

-- name: ListCommentsInPost :many
SELECT *
FROM comments
WHERE post_id = ?
ORDER BY created_at;

-- name: CreateComment :one
INSERT INTO comments (id, text, post_id)
VALUES (new_id('c'), ?, ?)
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = ?;