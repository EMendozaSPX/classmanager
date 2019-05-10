CREATE SCHEMA classmanager;

CREATE TABLE classmanager.admins (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    role VARCHAR(100) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE classmanager.classes (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id VARCHAR(20) UNIQUE NOT NULL,
    teacher_id INTEGER NOT NULL
);

CREATE TABLE classmanager.class_student (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id INTEGER NOT NULL,
    student_id INTEGER NOT NULL
);

CREATE TABLE classmanager.year (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year INTEGER NOT NULL,
    year_group INTEGER NOT NULL
);

CREATE TABLE classmanager.periods (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);
