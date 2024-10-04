CREATE TABLE IF NOT EXISTS workout_records (
    id SERIAL PRIMARY KEY,
    exercise VARCHAR(255) NOT NULL,
    reps INT NOT NULL,
    sets INT NOT NULL,
    date DATE NOT NULL
);

INSERT INTO workout_records (exercise, reps, sets, date) VALUES
('Push-ups', 20, 3, '2024-10-01'),
('Squats', 15, 4, '2024-10-02'),
('Pull-ups', 10, 3, '2024-10-03');