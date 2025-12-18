-- name: CreateBucket :one
INSERT INTO buckets (name, account_id, access_key, secret_key)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListBuckets :many
SELECT * FROM buckets
ORDER BY created_at DESC;

-- name: GetBucketByID :one
SELECT * FROM buckets
WHERE id = $1 LIMIT 1;

-- name: CreateSnapshot :one
INSERT INTO snapshots (bucket_id, obj_count, size_byte)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListSnapshots :many
SELECT * FROM snapshots
WHERE bucket_id = $1
ORDER BY recorded_at DESC;