-- name: GetEvents :many
SELECT * FROM "Event" 
WHERE "deleted_at" IS NULL
ORDER BY "starts_at";

-- name: GetEventById :one
SELECT * FROM "Event"
WHERE "id" = $1
  AND "deleted_at" IS NULL;

-- name: GetEventsByCreator :many
SELECT * FROM "Event"
WHERE "user_creator_id" = $1
  AND "deleted_at" IS NULL
ORDER BY "starts_at";

-- name: GetEventsByGroup :many
SELECT * FROM "Event"
WHERE "group_creator_id" = $1
  AND "deleted_at" IS NULL
ORDER BY "starts_at";

-- name: GetEventsStartedInRange :many
SELECT * FROM "Event"
WHERE "starts_at" BETWEEN $1 AND $2
  AND "deleted_at" IS NULL
ORDER BY "starts_at";

-- name: CreateEvent :one
INSERT INTO "Event" (
  "title",
  "description",
  "place",
  "coordinates",
  "starts_at",
  "ends_at",
  "capacity",
  "user_creator_id",
  "group_creator_id"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: ModifyEvent :one
UPDATE "Event"
SET 
    "title" = COALESCE($2, "title"),
    "description" = COALESCE($3, "description"),
    "place" = COALESCE($4, "place"),
    "coordinates" = COALESCE($5, "coordinates"),
    "starts_at" = COALESCE($6, "starts_at"),
    "ends_at" = COALESCE($7, "ends_at"),
    "capacity" = COALESCE($8, "capacity")
WHERE "id" = $1
RETURNING *;

-- name: SoftDeleteEvent :one
UPDATE "Event"
SET "deleted_at" = NOW()
WHERE "id" = $1
RETURNING *;


