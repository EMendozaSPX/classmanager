CREATE SCHEMA classmanager;

CREATE TABLE classmanager.users (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    role VARCHAR(100) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE classmanager.classes (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id VARCHAR(20) UNIQUE NOT NULL,
    teacher_id INTEGER NOT NULL REFERENCES classmanager.users(id)
);

CREATE TABLE classmanager.class_student (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES classmanager.classes(id),
    student_id INTEGER NOT NULL REFERENCES classmanager.users(id)
);

CREATE TABLE classmanager.year (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year INTEGER NOT NULL,
    year_group INTEGER NOT NULL
);

CREATE TABLE classmanager.terms (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES classmanager.year(id),
    term INTEGER NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE classmanager.public_holidays (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES classmanager.year(id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE classmanager.events (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES classmanager.year(id),
    class_id INTEGER NOT NULL REFERENCES classmanager.classes(id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE classmanager.periods (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES classmanager.year(id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);
