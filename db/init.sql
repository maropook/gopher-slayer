CREATE DATABASE IF NOT EXISTS gopher_slayer;
USE gopher_slayer;

-- -----------------------------------------------
-- heroes
-- -----------------------------------------------
CREATE TABLE IF NOT EXISTS heroes (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(50)  NOT NULL DEFAULT 'Gopher',
    hp          INT          NOT NULL DEFAULT 100,
    max_hp      INT          NOT NULL DEFAULT 100,
    attack      INT          NOT NULL DEFAULT 15,
    level       INT          NOT NULL DEFAULT 1,
    experience  INT          NOT NULL DEFAULT 0
);

-- -----------------------------------------------
-- stages
-- -----------------------------------------------
CREATE TABLE IF NOT EXISTS stages (
    id                  INT AUTO_INCREMENT PRIMARY KEY,
    name                VARCHAR(100) NOT NULL,
    description         TEXT,
    required_experience INT NOT NULL DEFAULT 0,
    order_num           INT NOT NULL
);

-- -----------------------------------------------
-- enemies
-- -----------------------------------------------
CREATE TABLE IF NOT EXISTS enemies (
    id                INT AUTO_INCREMENT PRIMARY KEY,
    stage_id          INT NOT NULL,
    name              VARCHAR(100) NOT NULL,
    hp                INT NOT NULL,
    max_hp            INT NOT NULL,
    attack            INT NOT NULL,
    experience_reward INT NOT NULL DEFAULT 0,
    FOREIGN KEY (stage_id) REFERENCES stages(id)
);

-- -----------------------------------------------
-- Seed: Hero (single hero, id=1)
-- -----------------------------------------------
INSERT INTO heroes (name, hp, max_hp, attack, level, experience)
VALUES ('Gopher', 100, 100, 15, 1, 0);

-- -----------------------------------------------
-- Seed: Stages
-- -----------------------------------------------
INSERT INTO stages (name, description, required_experience, order_num) VALUES
('Forest',         'A quiet forest, but beware of slimes!',          0,   1),
('Cave',           'Dark caves filled with bats and rock monsters.',  40,  2),
('Castle',         'An abandoned castle haunted by the undead.',      100, 3),
('Hell Gate',      'The entrance to the underworld. Beware!',        180, 4),
('Dragon''s Lair', 'Face the legendary dragon and save the world!',  300, 5);

-- -----------------------------------------------
-- Seed: Enemies (1 enemy per stage)
-- XP rewards are tuned to match stage unlock thresholds:
--   Stage 2 requires 40 XP → Stage 1 reward = 40
--   Stage 3 requires 100 XP → Stage 2 reward = 60 (cumulative 100)
--   Stage 4 requires 180 XP → Stage 3 reward = 80 (cumulative 180)
--   Stage 5 requires 300 XP → Stage 4 reward = 120 (cumulative 300)
-- -----------------------------------------------
-- Stage 1: Forest
INSERT INTO enemies (stage_id, name, hp, max_hp, attack, experience_reward) VALUES
(1, 'Goblin', 40, 40, 8, 40);

-- Stage 2: Cave
INSERT INTO enemies (stage_id, name, hp, max_hp, attack, experience_reward) VALUES
(2, 'Rock Monster', 70, 70, 12, 60);

-- Stage 3: Castle
INSERT INTO enemies (stage_id, name, hp, max_hp, attack, experience_reward) VALUES
(3, 'Dark Knight', 100, 100, 18, 80);

-- Stage 4: Hell Gate
-- NOTE: This enemy has a bug in battle_service.go for the Lv4 workshop task.
INSERT INTO enemies (stage_id, name, hp, max_hp, attack, experience_reward) VALUES
(4, 'Demon', 150, 150, 22, 120);

-- Stage 5: Dragon's Lair
-- NOTE: Boss Dragon has attack=50. This is intentionally high for the Lv3 workshop task.
-- Students must create PUT /api/hero/hp to edit HP before this fight.
INSERT INTO enemies (stage_id, name, hp, max_hp, attack, experience_reward) VALUES
(5, 'Boss Dragon', 300, 300, 50, 200);
