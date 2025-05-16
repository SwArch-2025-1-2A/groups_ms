-- name: GetUserById :one
SELECT *
FROM "User"
WHERE "id" = $1;

-- name: CreateUser :one
INSERT INTO "User" ("id", "name", "profile_pic")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ChangeUserProperties :one
UPDATE "User"
SET "name" = $2,
    "profile_pic" = $3
WHERE "id" = $1
RETURNING *;

-- name: ChangeUserName :one
UPDATE "User"
SET "name" = $2
WHERE "id" = $1
RETURNING *;

-- name: ChangeUserProfilePic :one
UPDATE "User"
SET "profile_pic" = $2
WHERE "id" = $1
RETURNING *;

-- name: AddUserInterest :one
INSERT INTO "UserInterests"("user_id", "interest_id")
VALUES ($1, $2)
RETURNING *;

-- name: RemoveUserInterest :exec
DELETE FROM "UserInterests"
WHERE "user_id" = $1
  AND "interest_id" = $2;

-- name: GetUserInterests :many
SELECT c.*
FROM "UserInterests" as ui
  JOIN "Category" as c
  ON ui.interest_id = c.id
WHERE ui.user_id = $1
  AND c.deleted_at IS NULL;

