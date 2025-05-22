-- name: CreateGroup :one
INSERT INTO "Group" ("name", "description", "profile_pic", "isOpen" )
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: VerifyGroup :one
UPDATE "Group"
SET "isVerified" = true
WHERE "id" = $1
RETURNING *;

-- name: SoftDeleteGroup :one
UPDATE "Group"
SET "deleted_at" = now()
WHERE "id" = $1
RETURNING *;

-- name: GetGroups :many
SELECT * FROM "Group"
WHERE "deleted_at" is NULL
ORDER BY "name";