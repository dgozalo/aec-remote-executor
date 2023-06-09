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

CREATE TABLE Assignment_Code_Templates
(
    id SERIAL PRIMARY KEY,
    assignment_id INT REFERENCES Assignments(id) ON DELETE CASCADE,
    language VARCHAR(50),
    code TEXT,
    test_runner_code TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE Assignment_Examples
(
    id SERIAL PRIMARY KEY,
    assignment_id INT REFERENCES Assignments(id) ON DELETE CASCADE,
    example_title VARCHAR(100),
    example_description TEXT,
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

-- inserts

INSERT INTO alumni (first_name, last_name, email, graduation_year) VALUES ('John', 'Doe', 'john@edu.com', null);
INSERT INTO Professors (first_name, last_name, email) VALUES ('William', 'Doe', 'will@edu.com');
INSERT INTO Subjects (subject_name, semester, professor_id) VALUES ('Programming', 1, 1);
INSERT INTO Subjects (subject_name, semester, professor_id) VALUES ('Data Structures and Algorithms', 1, 1);
INSERT INTO Assignments (assignment_title, assignment_description, subject_id, professor_id) VALUES ('Two Sum', 'Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.', 1, 1);
INSERT INTO Assignment_Examples (assignment_id, example_title, example_description) VALUES (1, 'Example 1', 'Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].');
INSERT INTO Assignment_Examples (assignment_id, example_title, example_description) VALUES (1, 'Example 2', 'Input: nums = [3,2,4], target = 6
Output: [1,2]');
INSERT INTO Assignment_Examples (assignment_id, example_title, example_description) VALUES (1, 'Example 3', 'Input: nums = [3,3], target = 6
Output: [0,1]');
INSERT INTO Assignment_Code_Templates (assignment_id, language, code) VALUES (1, 'Python', 'class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        pass');
INSERT INTO Assignment_Code_Templates (assignment_id, language, code, test_runner_code) VALUES (1, 'java', '
    public int[] twoSum(int[] nums, int target) {

    }', '   public static void runTestCases(Solution solution, TestsRunner runner, String testsFilePath) {
        java.util.List<Test<int[], Integer, int[]>> tests = java.util.List.of(
                new Test<>(new int[]{1, 2, 3}, 1, new int[]{1, 2}),
                new Test<>(new int[]{1, 2, 3}, 2, new int[]{1, 2}),
                new Test<>(new int[]{1, 2, 3}, 3, new int[]{1, 2}),
                new Test<>(new int[]{1, 2, 3}, 4, new int[]{1, 2}));
        for (int i = 0; i < tests.size(); i++) {
            Test<int[], Integer, int[]> test = tests.get(i);
            int[] sol = solution.twoSum(test.getInput(), test.getTarget());
            runner.writeResultToTestsFile(testsFilePath, String.format("TestCase#%1d::%2s::%3s::%4b", i,
                    java.util.Arrays.toString(sol),
                    java.util.Arrays.toString(test.getExpectedOutput()),
                    java.util.Arrays.equals(sol, test.getExpectedOutput())));
        }
    }');
INSERT INTO Assignment_Code_Templates (assignment_id, language, code) VALUES (1, 'javascript', 'var twoSum = function(nums, target) {

};');

INSERT INTO Alumni_Subjects (alumni_id, subject_id) VALUES (1, 1);
INSERT INTO Alumni_Subjects (alumni_id, subject_id) VALUES (1, 2);
INSERT INTO Alumni_Assignments (alumni_id, assignment_id) VALUES (1, 1);