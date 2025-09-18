-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "account_status" TYPE character varying(20), ALTER COLUMN "account_status" SET DEFAULT 'pending';
-- Create "admins" table
CREATE TABLE "public"."admins" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "admin_id" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "admin_name" character varying(255) NOT NULL,
  "password" character varying(255) NOT NULL,
  "role" character varying(20) NOT NULL DEFAULT 'normal_admin',
  PRIMARY KEY ("id")
);
-- Create index "idx_admins_admin_id" to table: "admins"
CREATE UNIQUE INDEX "idx_admins_admin_id" ON "public"."admins" ("admin_id");
-- Create index "idx_admins_admin_name" to table: "admins"
CREATE UNIQUE INDEX "idx_admins_admin_name" ON "public"."admins" ("admin_name");
-- Create index "idx_admins_deleted_at" to table: "admins"
CREATE INDEX "idx_admins_deleted_at" ON "public"."admins" ("deleted_at");
-- Create index "idx_admins_email" to table: "admins"
CREATE UNIQUE INDEX "idx_admins_email" ON "public"."admins" ("email");
