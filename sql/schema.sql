/*
 Navicat Premium Data Transfer

 Source Server         : Docker
 Source Server Type    : PostgreSQL
 Source Server Version : 110005
 Source Host           : localhost:5432
 Source Catalog        : mini_project
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110005
 File Encoding         : 65001

 Date: 01/01/2021 20:58:35
*/

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- ----------------------------
-- Table structure for hub
-- ----------------------------

DROP TABLE IF EXISTS "public"."hub";
CREATE TABLE "public"."hub" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "location_lat" float4 NOT NULL DEFAULT 0.0,
  "location_long" float4 NOT NULL DEFAULT 0.0
)
;

-- ----------------------------
-- Table structure for team
-- ----------------------------
DROP TABLE IF EXISTS "public"."team";
CREATE TABLE "public"."team" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "team_type" varchar(255) COLLATE "pg_catalog"."default",
  "hub_id" uuid NOT NULL
)
;

-- ----------------------------
-- Table structure for team_type
-- ----------------------------
DROP TABLE IF EXISTS "public"."team_type";
CREATE TABLE "public"."team_type" (
  "type_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "role" varchar(255) COLLATE "pg_catalog"."default",
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "team_id" uuid
)
;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_role";
CREATE TABLE "public"."user_role" (
  "role_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Function structure for uuid_generate_v1
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v1mc
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1mc"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1mc"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1mc'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v3
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v3"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v3"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v3'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v4
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v4"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v4"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v4'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v5
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v5"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v5"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v5'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_nil
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_nil"();
CREATE OR REPLACE FUNCTION "public"."uuid_nil"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_nil'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_dns
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_dns"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_dns"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_dns'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_oid
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_oid"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_oid"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_oid'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_url
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_url"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_url"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_url'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_x500
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_x500"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_x500"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_x500'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Uniques structure for table hub
-- ----------------------------
ALTER TABLE "public"."hub" ADD CONSTRAINT "hub_name_unique" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table hub
-- ----------------------------
ALTER TABLE "public"."hub" ADD CONSTRAINT "hub_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table team
-- ----------------------------
ALTER TABLE "public"."team" ADD CONSTRAINT "team_name_unique" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table team
-- ----------------------------
ALTER TABLE "public"."team" ADD CONSTRAINT "team_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table team_type
-- ----------------------------
ALTER TABLE "public"."team_type" ADD CONSTRAINT "team_type_pkey" PRIMARY KEY ("type_name");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_role
-- ----------------------------
ALTER TABLE "public"."user_role" ADD CONSTRAINT "user_role_pkey" PRIMARY KEY ("role_name");

-- ----------------------------
-- Foreign Keys structure for table team
-- ----------------------------
ALTER TABLE "public"."team" ADD CONSTRAINT "hub_id_fk" FOREIGN KEY ("hub_id") REFERENCES "public"."hub" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."team" ADD CONSTRAINT "team_type_fk" FOREIGN KEY ("team_type") REFERENCES "public"."team_type" ("type_name") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "role_fk" FOREIGN KEY ("role") REFERENCES "public"."user_role" ("role_name") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."user" ADD CONSTRAINT "team_id_fk" FOREIGN KEY ("team_id") REFERENCES "public"."team" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
