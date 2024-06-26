DELETE FROM hello_worlds;
DELETE FROM sketches;
DELETE FROM posts;
DELETE FROM users;

INSERT INTO hello_worlds (lang, message) VALUES ('en', 'Hello World'), ('ja', 'こんにちは 世界');

INSERT INTO users (name, password) VALUES ('taro', 'password'),('hanako', 'PASSWORD');

INSERT INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行'),(1, 'test2', '質問2\n改行');

INSERT INTO sketches (image_name, user_id) VALUES
('happy.png', 1),
('quickMTG.png', 1),
('animal_mark04_neko.png', 1),
('animal_mark07_lion.png', 1),
('animal_mark09_tora.png', 1),
('animal_mark10_usagi.png', 1),
('animal_mark11_panda.png', 1),
('animal_mark15_koala.png', 1);
