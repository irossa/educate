CREATE TABLE "district" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "school" (
  "id" bigserial PRIMARY KEY,
  "district_id" bigint,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "student" (
  "id" bigserial PRIMARY KEY,
  "school_id" bigint,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "identifier" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "test" (
  "id" bigserial PRIMARY KEY,
  "district_id" bigint,
  "subject" varchar,
  "test_date" timestamptz NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "test_score" (
  "test_id" bigint,
  "student_id" bigint,
  "test_date" timestamptz NOT NULL,
  "score" decimal NOT NULL DEFAULT 0
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "user_student" (
  "user_id" bigint,
  "student_id" bigint
);

ALTER TABLE "school" ADD FOREIGN KEY ("district_id") REFERENCES "district" ("id");

ALTER TABLE "student" ADD FOREIGN KEY ("school_id") REFERENCES "school" ("id");

ALTER TABLE "test" ADD FOREIGN KEY ("district_id") REFERENCES "district" ("id");

ALTER TABLE "test_score" ADD FOREIGN KEY ("test_id") REFERENCES "test" ("id");

ALTER TABLE "test_score" ADD FOREIGN KEY ("student_id") REFERENCES "student" ("id");

ALTER TABLE "user_student" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_student" ADD FOREIGN KEY ("student_id") REFERENCES "student" ("id");

ALTER TABLE "users" ADD CONSTRAINT "username_key" UNIQUE ("username");
