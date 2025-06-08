-- name: ListNotes :many
SELECT * FROM notes ORDER BY created_at DESC;

-- name: GetNote :one
SELECT * FROM notes WHERE id = ?;

-- name: CreateNote :one
INSERT INTO notes (title, content) values (?, ?) RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = ?;
