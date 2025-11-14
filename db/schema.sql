CREATE TYPE "currency" AS ENUM (
  'rupee',
  'dollar'
);

CREATE TYPE "status" AS ENUM (
  'active',
  'inactive'
);

CREATE TABLE "entries" (
  "id" uuid PRIMARY KEY,
  "account_id" uuid,
  "created_at" timestamp,
  "amount" int
);

CREATE TABLE "accounts" (
  "id" uuid PRIMARY KEY,
  "first_name" varchar(50),
  "last_name" varchar(50),
  "created_at" timestamp,
  "balance" int,
  "currency" currency,
  "status" status
);

CREATE TABLE "transactions" (
  "id" uuid PRIMARY KEY,
  "from_account" uuid,
  "to_account" uuid,
  "created_at" timestamp,
  "amount" int
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account") REFERENCES "accounts" ("id");
