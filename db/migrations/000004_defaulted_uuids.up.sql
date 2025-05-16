ALTER TABLE "Event" ALTER COLUMN "id" SET DEFAULT uuid_generate_v4();
ALTER TABLE "Group" ALTER COLUMN "id" SET DEFAULT uuid_generate_v4();
ALTER TABLE "Category" ALTER COLUMN "id" SET DEFAULT uuid_generate_v4();
ALTER TABLE "Member" ALTER COLUMN "id" SET DEFAULT uuid_generate_v4();
