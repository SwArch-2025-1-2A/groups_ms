-- This guarantees that any event created in the past has a user creator
ALTER TABLE "Event" ALTER COLUMN user_creator_id SET NOT NULL;
