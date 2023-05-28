CREATE SCHEMA IF NOT EXISTS AEC_EXECUTOR;

-- Professors table

CREATE TABLE Professors
(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Alumni table

CREATE TABLE Alumni
(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(255) UNIQUE,
    graduation_year INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Subjects table

CREATE TABLE Subjects
(
    id SERIAL PRIMARY KEY,
    subject_name VARCHAR(100),
    semester INT,
    professor_id INT REFERENCES Professors(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Assignments table

CREATE TABLE Assignments
(
    id SERIAL PRIMARY KEY,
    assignment_title VARCHAR(100),
    assignment_description TEXT,
    subject_id INT REFERENCES Subjects(id) ON DELETE CASCADE,
    professor_id INT REFERENCES Professors(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Alumni_Subjects table

CREATE TABLE Alumni_Subjects
(
    alumni_id INT REFERENCES Alumni(id) ON DELETE CASCADE,
    subject_id INT REFERENCES Subjects(id) ON DELETE CASCADE,
    PRIMARY KEY (alumni_id, subject_id)
);

-- Alumni_Assignments table

CREATE TABLE Alumni_Assignments
(
    alumni_id INT REFERENCES Alumni(id) ON DELETE CASCADE,
    assignment_id INT REFERENCES Assignments(id) ON DELETE CASCADE,
    grade VARCHAR(2),
    PRIMARY KEY (alumni_id, assignment_id)
);

-- Executions table

CREATE TABLE Executions
(
    id SERIAL PRIMARY KEY,
    language VARCHAR(50),
    workflow_id VARCHAR(50),
    run_id VARCHAR(50),
    code TEXT,
    assignment_id INT REFERENCES Assignments(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);