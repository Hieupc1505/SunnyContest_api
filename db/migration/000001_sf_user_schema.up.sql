CREATE TABLE "sf_user" (
    "id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL,
    "role" int NOT NULL,
    "status" int NOT NULL,
    "token" varchar,
    "token_expried" bigint,
    "created_time" timestamptz NOT NULL DEFAULT (now()),
    "updated_time" timestamptz NOT NULL DEFAULT (now())
);