CREATE TABLE users (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    role VARCHAR(100) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE classes (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id VARCHAR(20) UNIQUE NOT NULL,
    teacher_id INTEGER NOT NULL REFERENCES users(id)
);

CREATE TABLE class_student (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES classes(id),
    student_id INTEGER NOT NULL REFERENCES users(id)
);

CREATE TABLE year (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year INTEGER,
    year_group INTEGER
);

CREATE TABLE public_holidays (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES year(id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES year(id),
    class_id INTEGER NOT NULL REFERENCES classes(id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE timetable (
    id          SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id    INTEGER NOT NULL REFERENCES classes (id),
    period_name VARCHAR(50) NOT NULL,
    week_day    INTEGER NOT NULL
);

CREATE TABLE behaviour_notes (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    student_class_id INTEGER NOT NULL REFERENCES class_student(id),
    time_stamp TIMESTAMP NOT NULL
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES classes(id),
    task_name VARCHAR(100) NOT NULL,
    total_mark INTEGER NOT NULL,
    task_description TEXT NOT NULL,
    due_time TIMESTAMP NOT NULL
);

CREATE TABLE task_marks (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    task_id INTEGER NOT NULL REFERENCES tasks(id),
    student_class_id INTEGER NOT NULL REFERENCES class_student(id),
    task_mark INTEGER NOT NULL,
    feedback TEXT,
    time_stamp TIMESTAMP NOT NULL
);
