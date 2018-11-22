CREATE TABLE "greeting" (
  "start_time" TIMESTAMP WITH TIME ZONE NOT NULL,
  "end_time" TIMESTAMP WITH TIME ZONE NOT NULL,
  "raw_venue_id" BIGINT NOT NULL,
  "raw_character_id" BIGINT NOT NULL,
  PRIMARY KEY ("start_time", "end_time", "raw_venue_id", "raw_character_id")
);
CREATE INDEX "index_raw_venue_id_start_time" ON "greeting" ("raw_venue_id", "start_time");
CREATE INDEX "index_raw_character_id_start_time" ON "greeting" ("raw_character_id", "start_time");

CREATE TABLE "raw_venue" (
  "id" BIGSERIAL NOT NULL,
  "name" VARCHAR(128) NOT NULL,
  PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "raw_venue_name" ON "raw_venue" ("name");

CREATE TABLE "raw_character" (
  "id" BIGSERIAL NOT NULL,
  "name" VARCHAR(128) NOT NULL,
  PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "raw_character_name" ON "raw_character" ("name");

CREATE TABLE "venue" (
  "id" BIGSERIAL NOT NULL,
  "name" VARCHAR(128) NOT NULL,
  PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "venue_name" ON "venue" ("name");

CREATE TABLE "character" (
  "id" BIGSERIAL NOT NULL,
  "name" VARCHAR(128) NOT NULL,
  PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "character_name" ON "character" ("name");

CREATE TABLE "character_style" (
  "id" BIGSERIAL NOT NULL,
  "character_id" BIGINT NOT NULL,
  "name" VARCHAR(128) NOT NULL,
  "is_default" BOOLEAN NOT NULL,
  PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "character_style_character_id_name" ON "character_style" ("character_id", "name");

CREATE TABLE "venue_canonicalization" (
  "venue_id" BIGINT NOT NULL,
  "raw_venue_id" BIGINT NOT NULL,
  PRIMARY KEY ("venue_id", "raw_venue_id")
);

CREATE TABLE "character_canonicalization" (
  "character_id" BIGINT NOT NULL,
  "character_style_id" BIGINT NOT NULL,
  "raw_character_id" BIGINT NOT NULL,
  PRIMARY KEY ("character_id", "character_style_id", "raw_character_id")
);
