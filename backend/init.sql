-- init.sql

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    surname TEXT NOT NULL,
    gender CHAR(1) DEFAULT 'M',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS fortunes (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    category TEXT DEFAULT 'random',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    added_by INTEGER REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS subjects (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "name_en" TEXT NOT NULL,
    "credits" REAL NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "class" (
    "start_year" INTEGER NOT NULL,
    "id" TEXT NOT NULL PRIMARY KEY,
    "subject_id" TEXT,
    "name" TEXT,
    "start_semester_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Check if any users exist before inserting
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM users) THEN
        INSERT INTO users (surname) VALUES 
        ('Myra'), 
        ('Hady'), 
        ('Baber'), 
        ('Sidi'), 
        ('Aguida');
    END IF;
END
$$;

-- Check if any fortunes exist before inserting
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM fortunes) THEN
        INSERT INTO fortunes (title, added_by) VALUES
        ('You will become a millionaire by %d 🚀!', 1),
        ('A pigeon is plotting against you 🐦!', 2),
        ('Someone will compliment your hair soon 💇‍♂️!', 3),
        ('Beware of tacos 🌮 today!', 4),
        ('You will discover a hidden talent: %s 🎨!', 5),
        ('An unexpected journey awaits you ✈️!', 1),
        ('Your code will compile perfectly on the first try 🧑‍💻!', 2),
        ('You will find money on the ground: %d yen 💵!', 3),
        ('Beware of sneaky cats 🐱!', 4),
        ('Great news will come in your email 📧!', 5),
        ('You will become invisible for 5 minutes 🫥!', 1),
        ('Your favorite food will betray you 🍕!', 2),
        ('A bird will deliver an important message 🦅!', 3),
        ('You will wake up with super strength 💪!', 4),
        ('Be cautious of mysterious elevators 🛗!', 5),
        ('Your singing voice will charm someone 🎤!', 1),
        ('You will receive a mysterious gift 🎁!', 2),
        ('Unexpected rain will bless your day 🌧️!', 3),
        ('You will find a secret passage in your city 🛤️!', 4),
        ('You will meet a talking cat 🐈‍⬛!', 5);
    END IF;
END
$$;
