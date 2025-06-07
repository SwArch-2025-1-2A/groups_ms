ALTER TABLE "Group" ALTER COLUMN "name" SET NOT NULL;
ALTER TABLE "Group" ALTER COLUMN "isOpen" SET NOT NULL;
ALTER TABLE "Group" ALTER COLUMN "isVerified" SET NOT NULL;


ALTER TABLE "Group" RENAME COLUMN "profile_pic" TO "profilePic";
ALTER TABLE "Group" ALTER COLUMN "profilePic" TYPE BYTEA USING decode("profilePic", 'escape');