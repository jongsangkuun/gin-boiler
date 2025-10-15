-- Drop index "idx_admins_admin_id" from table: "admins"
DROP INDEX "public"."idx_admins_admin_id";
-- Drop index "idx_admins_admin_name" from table: "admins"
DROP INDEX "public"."idx_admins_admin_name";
-- Rename a column from "admin_id" to "login_id"
ALTER TABLE "public"."admins" RENAME COLUMN "admin_id" TO "login_id";
-- Rename a column from "admin_name" to "nick_name"
ALTER TABLE "public"."admins" RENAME COLUMN "admin_name" TO "nick_name";
-- Create index "idx_admins_login_id" to table: "admins"
CREATE UNIQUE INDEX "idx_admins_login_id" ON "public"."admins" ("login_id");
-- Create index "idx_admins_nick_name" to table: "admins"
CREATE UNIQUE INDEX "idx_admins_nick_name" ON "public"."admins" ("nick_name");
-- Drop index "idx_users_user_id" from table: "users"
DROP INDEX "public"."idx_users_user_id";
-- Drop index "idx_users_username" from table: "users"
DROP INDEX "public"."idx_users_username";
-- Rename a column from "user_id" to "login_id"
ALTER TABLE "public"."users" RENAME COLUMN "user_id" TO "login_id";
-- Rename a column from "username" to "nick_name"
ALTER TABLE "public"."users" RENAME COLUMN "username" TO "nick_name";
-- Create index "idx_users_login_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_login_id" ON "public"."users" ("login_id");
-- Create index "idx_users_nick_name" to table: "users"
CREATE UNIQUE INDEX "idx_users_nick_name" ON "public"."users" ("nick_name");
-- Create "posts" table
CREATE TABLE "public"."posts" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "post_data" text NOT NULL DEFAULT '',
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_posts_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_posts_deleted_at" to table: "posts"
CREATE INDEX "idx_posts_deleted_at" ON "public"."posts" ("deleted_at");
-- Create index "idx_posts_user_id" to table: "posts"
CREATE INDEX "idx_posts_user_id" ON "public"."posts" ("user_id");
