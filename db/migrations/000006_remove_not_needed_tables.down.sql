CREATE TYPE "MemberRole" AS ENUM (
  'creator',
  'admin',
  'normal_user'
);

CREATE TYPE "JoinRequestStatus" AS ENUM (
  'pending',
  'accepted',
  'cancelled',
  'denied',
  'blocked'
);

CREATE TABLE "User" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "profile_pic" varchar
);

CREATE TABLE "UserInterests" (
  "user_id" uuid,
  "interest_id" uuid,
  PRIMARY KEY ("user_id", "interest_id")
);

CREATE TABLE "Event" (
  "id" uuid PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "place" varchar,
  "coordinates" point DEFAULT null,
  "starts_at" timestamp,
  "ends_at" timestamp,
  "capacity" int,
  "user_creator_id" uuid,
  "group_creator_id" uuid,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "GroupCategories" (
  "group_id" uuid,
  "interest_id" uuid,
  PRIMARY KEY ("group_id", "interest_id")
);

CREATE TABLE "Participant" (
  "user_id" uuid NOT NULL,
  "event_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "event_id")
);

CREATE TABLE "Member" (
  "id" uuid PRIMARY KEY,
  "role" "MemberRole" NOT NULL,
  "group_id" uuid NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "Category" (
  "id" uuid PRIMARY KEY,
  "category" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "JoinRequest" (
  "user_id" uuid NOT NULL,
  "group_id" uuid NOT NULL,
  "messsage_from_user" varchar,
  "messsage_from_admin" varchar,
  "status" "JoinRequestStatus",
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT null,
  PRIMARY KEY ("user_id", "group_id")
);

COMMENT ON COLUMN "Event"."place" IS 'Description about the place or reference to the place';

ALTER TABLE "UserInterests" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "UserInterests" ADD FOREIGN KEY ("interest_id") REFERENCES "Category" ("id");

ALTER TABLE "Event" ADD FOREIGN KEY ("user_creator_id") REFERENCES "User" ("id");

ALTER TABLE "Event" ADD FOREIGN KEY ("group_creator_id") REFERENCES "Group" ("id");

ALTER TABLE "GroupCategories" ADD FOREIGN KEY ("group_id") REFERENCES "Group" ("id");

ALTER TABLE "GroupCategories" ADD FOREIGN KEY ("interest_id") REFERENCES "Category" ("id");

ALTER TABLE "Participant" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "Participant" ADD FOREIGN KEY ("event_id") REFERENCES "Event" ("id");

ALTER TABLE "Member" ADD FOREIGN KEY ("group_id") REFERENCES "Group" ("id");

ALTER TABLE "Member" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "JoinRequest" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "JoinRequest" ADD FOREIGN KEY ("group_id") REFERENCES "Group" ("id");
