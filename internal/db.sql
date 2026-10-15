CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "email" VARCHAR(50) UNIQUE,
    "password" VARCHAR(100),
    "updated_at" TIMESTAMP DEFAULT NOW(),
    "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "user_profile" (
    "id" UUID PRIMARY KEY,
    "id_user" UUID REFERENCES "users"("id"),
    "name" VARCHAR(50),
    "profile_pic" VARCHAR(100),
    "phone" VARCHAR(15),
    "birth_date" TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT NOW(),
    "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "conversations" (
    "id" UUID PRIMARY KEY,
    "updated_at" TIMESTAMP DEFAULT NOW(),
    "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "conversations_member" (
    "id" UUID PRIMARY KEY,
    "conversations_id" UUID REFERENCES "conversations"("id"),
    "users_id" UUID REFERENCES "users"("id"),
    "updated_at" TIMESTAMP DEFAULT NOW(),
    "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "messages" (
    "id" UUID PRIMARY KEY,
    "conversations_id" UUID REFERENCES "conversations"("id"),
    "sender_id" UUID REFERENCES "users"("id"),
    "content" TEXT,
    "updated_at" TIMESTAMP DEFAULT NOW(),
    "created_at" TIMESTAMP DEFAULT NOW()
);