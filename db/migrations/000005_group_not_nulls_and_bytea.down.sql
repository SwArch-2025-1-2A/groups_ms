ALTER TABLE "Group" ALTER COLUMN "profilePic" TYPE VARCHAR USING encode("profilePic", 'escape');
ALTER TABLE "Group" RENAME COLUMN "profilePic" TO "profile_pic";

ALTER TABLE "Group" ALTER COLUMN "isVerified" DROP NOT NULL;
ALTER TABLE "Group" ALTER COLUMN "isOpen" DROP NOT NULL;
ALTER TABLE "Group" ALTER COLUMN "name" DROP NOT NULL;