-- Drop index "idx_posts_user_id" from table: "posts"
DROP INDEX "public"."idx_posts_user_id";
-- Modify "posts" table
ALTER TABLE "public"."posts" DROP CONSTRAINT "fk_posts_user", DROP COLUMN "post_data", ADD COLUMN "title" character varying(255) NOT NULL, ADD COLUMN "content" character varying(255) NOT NULL, ADD CONSTRAINT "fk_users_posts" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
