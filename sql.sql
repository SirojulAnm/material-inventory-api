/*
 Navicat Premium Data Transfer

 Source Server         : postgres_docker
 Source Server Type    : PostgreSQL
 Source Server Version : 140000
 Source Host           : localhost:5432
 Source Catalog        : tripatra
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140000
 File Encoding         : 65001

 Date: 11/05/2023 23:21:37
*/


-- ----------------------------
-- Sequence structure for materials_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."materials_id_seq";
CREATE SEQUENCE "public"."materials_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."materials_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for my_table_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."my_table_id_seq";
CREATE SEQUENCE "public"."my_table_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."my_table_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for my_table_id_seq1
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."my_table_id_seq1";
CREATE SEQUENCE "public"."my_table_id_seq1" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."my_table_id_seq1" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for my_table_id_seq2
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."my_table_id_seq2";
CREATE SEQUENCE "public"."my_table_id_seq2" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."my_table_id_seq2" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for my_table_id_seq3
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."my_table_id_seq3";
CREATE SEQUENCE "public"."my_table_id_seq3" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."my_table_id_seq3" OWNER TO "postgres";

-- ----------------------------
-- Table structure for materials
-- ----------------------------
DROP TABLE IF EXISTS "public"."materials";
CREATE TABLE "public"."materials" (
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "type" varchar(255) COLLATE "pg_catalog"."default",
  "quantity" int4,
  "updated_at" timestamp(6),
  "created_at" timestamp(6),
  "id" int4 NOT NULL DEFAULT nextval('materials_id_seq'::regclass)
)
;
ALTER TABLE "public"."materials" OWNER TO "postgres";

-- ----------------------------
-- Table structure for notifications
-- ----------------------------
DROP TABLE IF EXISTS "public"."notifications";
CREATE TABLE "public"."notifications" (
  "id" int4 NOT NULL DEFAULT nextval('my_table_id_seq3'::regclass),
  "material_id" int4,
  "message" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(6),
  "created_at" timestamp(6),
  "sender_id" int4,
  "receiver_id" int4
)
;
ALTER TABLE "public"."notifications" OWNER TO "postgres";

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS "public"."transactions";
CREATE TABLE "public"."transactions" (
  "id" int4 NOT NULL DEFAULT nextval('my_table_id_seq2'::regclass),
  "material_id" int4,
  "quantity" int4,
  "updated_at" timestamp(6),
  "created_at" timestamp(6),
  "status" varchar(255) COLLATE "pg_catalog"."default",
  "reason" varchar(255) COLLATE "pg_catalog"."default",
  "sender_id" int4,
  "receiver_id" int4,
  "warehouse_category" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."transactions" OWNER TO "postgres";

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('my_table_id_seq'::regclass),
  "email" varchar(255) COLLATE "pg_catalog"."default",
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "password_hash" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(0),
  "created_at" timestamp(6),
  "role" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."users" OWNER TO "postgres";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."materials_id_seq"
OWNED BY "public"."materials"."id";
SELECT setval('"public"."materials_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."my_table_id_seq"', 8, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."my_table_id_seq1"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."my_table_id_seq2"', 80, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."my_table_id_seq3"', 1, false);

-- ----------------------------
-- Primary Key structure for table materials
-- ----------------------------
ALTER TABLE "public"."materials" ADD CONSTRAINT "materials_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table notifications
-- ----------------------------
ALTER TABLE "public"."notifications" ADD CONSTRAINT "my_table_pkey3" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table transactions
-- ----------------------------
ALTER TABLE "public"."transactions" ADD CONSTRAINT "my_table_pkey2" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "my_table_pkey" PRIMARY KEY ("id");
