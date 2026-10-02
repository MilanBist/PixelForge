-- +goose Up

CREATE TABLE "users" (
  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "userName" VARCHAR(50) UNIQUE,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "passwordHash" TEXT NOT NULL,
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "uploadedFiles" (
  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "userId" BIGINT NOT NULL,
  "fileName" VARCHAR(255) NOT NULL,
  "storageKey" TEXT NOT NULL,
  "fileType" VARCHAR(50) NOT NULL,
  "mimeType" VARCHAR(100) NOT NULL,
  "fileSize" BIGINT,
  "width" INTEGER DEFAULT -1,
  "height" INTEGER DEFAULT -1,
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY ("userId")
    REFERENCES "users" ("id")
    DEFERRABLE INITIALLY IMMEDIATE
);

CREATE TABLE "images" (
  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "userId" BIGINT NOT NULL,
  "sourceFileId" BIGINT NOT NULL,
  "storageKey" TEXT NOT NULL,
  "fileName" VARCHAR(255),
  "mimeType" VARCHAR(100),
  "width" INTEGER,
  "height" INTEGER,
  "fileSize" BIGINT,
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),

  FOREIGN KEY ("userId")
    REFERENCES "users" ("id")
    DEFERRABLE INITIALLY IMMEDIATE,

  FOREIGN KEY ("sourceFileId")
    REFERENCES "uploadedFiles" ("id")
    DEFERRABLE INITIALLY IMMEDIATE
);

CREATE TABLE "history" (
  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "userId" BIGINT NOT NULL,
  "sourceFileId" BIGINT NOT NULL,
  "outputImageId" BIGINT NOT NULL,
  "operationType" VARCHAR(50) NOT NULL,
  "prompt" TEXT,
  "parameters" JSONB,
  "status" VARCHAR(30) NOT NULL DEFAULT 'pending',
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
  "completedAt" TIMESTAMP,

  FOREIGN KEY ("userId")
    REFERENCES "users" ("id")
    DEFERRABLE INITIALLY IMMEDIATE,

  FOREIGN KEY ("sourceFileId")
    REFERENCES "uploadedFiles" ("id")
    DEFERRABLE INITIALLY IMMEDIATE,

  FOREIGN KEY ("outputImageId")
    REFERENCES "images" ("id")
    DEFERRABLE INITIALLY IMMEDIATE
);


-- Indexes

CREATE INDEX ON "uploadedFiles" ("userId");
CREATE INDEX ON "uploadedFiles" ("createdAt");
CREATE INDEX ON "uploadedFiles" ("userId", "createdAt");

CREATE INDEX ON "images" ("userId");
CREATE INDEX ON "images" ("sourceFileId");
CREATE INDEX ON "images" ("createdAt");
CREATE INDEX ON "images" ("userId", "createdAt");

CREATE INDEX ON "history" ("userId");
CREATE INDEX ON "history" ("sourceFileId");
CREATE INDEX ON "history" ("outputImageId");
CREATE INDEX ON "history" ("operationType");
CREATE INDEX ON "history" ("status");
CREATE INDEX ON "history" ("createdAt");
CREATE INDEX ON "history" ("userId", "createdAt");
CREATE INDEX ON "history" ("userId", "status");


-- +goose Down

-- Drop tables in reverse dependency order.
-- Foreign keys are automatically removed when their tables are dropped.

DROP TABLE IF EXISTS "history";
DROP TABLE IF EXISTS "images";
DROP TABLE IF EXISTS "uploadedFiles";
DROP TABLE IF EXISTS "users";