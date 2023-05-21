CREATE SCHEMA IF NOT EXISTS AEC_EXECUTOR;

-- Professors table

CREATE TABLE Professors
(
    professor_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(255) UNIQUE
);

-- Alumni table

CREATE TABLE Alumni
(
    alumni_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(255) UNIQUE,
    graduation_year INT
);

-- Subjects table

CREATE TABLE Subjects
(
    subject_id SERIAL PRIMARY KEY,
    subject_name VARCHAR(100),
    semester INT,
    professor_id INT REFERENCES Professors(professor_id) ON DELETE CASCADE
);

-- Assignments table

CREATE TABLE Assignments
(
    assignment_id SERIAL PRIMARY KEY,
    assignment_title VARCHAR(100),
    assignment_description TEXT,
    subject_id INT REFERENCES Subjects(subject_id) ON DELETE CASCADE,
    professor_id INT REFERENCES Professors(professor_id) ON DELETE CASCADE
);

-- Alumni_Subjects table

CREATE TABLE Alumni_Subjects
(
    alumni_id INT REFERENCES Alumni(alumni_id) ON DELETE CASCADE,
    subject_id INT REFERENCES Subjects(subject_id) ON DELETE CASCADE,
    PRIMARY KEY (alumni_id, subject_id)
);

-- Alumni_Assignments table

CREATE TABLE Alumni_Assignments
(
    alumni_id INT REFERENCES Alumni(alumni_id) ON DELETE CASCADE,
    assignment_id INT REFERENCES Assignments(assignment_id) ON DELETE CASCADE,
    grade VARCHAR(2),
    PRIMARY KEY (alumni_id, assignment_id)
);

-- Executions table

CREATE TABLE Executions
(
    execution_id SERIAL PRIMARY KEY,
    language VARCHAR(50),
    workflow_id VARCHAR(50),
    run_id VARCHAR(50),
    code TEXT,
    assignment_id INT REFERENCES Assignments(assignment_id) ON DELETE CASCADE
);