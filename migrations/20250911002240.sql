-- Create "base_models" table
CREATE TABLE "public"."base_models" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_base_models_deleted_at" to table: "base_models"
CREATE INDEX "idx_base_models_deleted_at" ON "public"."base_models" ("deleted_at");
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "username" character varying(255) NOT NULL,
  "password" character varying(255) NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create index "idx_users_user_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_user_id" ON "public"."users" ("user_id");
-- Create index "idx_users_username" to table: "users"
CREATE UNIQUE INDEX "idx_users_username" ON "public"."users" ("username");
