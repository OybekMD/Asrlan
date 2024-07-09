INSERT INTO languages (name, picture) VALUES ('English', 'app.asrlan.uz/media/images/flags/us.png');
INSERT INTO languages (name, picture) VALUES ('Russian', 'app.asrlan.uz/media/images/flags/ru.png');
INSERT INTO languages (name, picture) VALUES ('Arabic', 'app.asrlan.uz/media/images/flags/sa.png');

INSERT INTO levels (name, real_level, picture, language_id) VALUES ('Beginner', 1, 'app.asrlan.uz/media/images/english/levels/beginner.png', 1);
INSERT INTO levels (name, real_level, picture, language_id) VALUES ('Elementary', 2, 'app.asrlan.uz/media/images/english/levels/elementary.png', 1);
INSERT INTO levels (name, real_level, picture, language_id) VALUES ('Pre-interassetste', 3, 'app.asrlan.uz/media/images/english/levels/preintermidiet.png', 1);

INSERT INTO topics (name, level_id) VALUES
    ('Alphabet and Pronunciation|Alfavit va talaffuz', 1),
    ('Greetings and Introductions|Salomlashish va tanishtirish', 1);
    -- ('Numbers|Raqamlar', 1),
    -- ('Colors|Ranglar', 1),
    -- ('Days of the Week|Haftaning kunlari', 1),
    -- ('Months of the Year|Yil oylari', 1),
    -- ('Family Members|Oila a’zolari', 1),
    -- ('Personal Information|Shaxsiy ma’lumotlar', 1),
    -- ('Daily Routines|Kundalik odatlar', 1),
    -- ('Basic Questions|Oddiy savollar', 1);


-- V Vocabulary
-- G Grammar
-- E Exercise
-- C Video
-- L Listening 
-- S Speaking
-- R Reading

INSERT INTO lessons (lesson_type, topic_id) VALUES
    ('V', 1),
    ('C', 1),
    ('E', 1),

    ('V', 2), 
    ('C', 2), 
    ('E', 2);
    
-- BEGIN 1 Topic 
INSERT INTO contents (lesson_id, gentype, title, text_data) VALUES
    (1, 2, 'Xarf Aa: Apple | Olma', 'Aa [ eɪ ]'),
    (1, 2, 'Xarf Bb: Book | Kitob', 'Bb [ biː ]'),
    (1, 2, 'Xarf Cc: Car | Mashina', 'Cc [ siː ]'),
    (1, 2, 'Xarf Dd: Door | Eshik', 'Dd [ diː ]'),
    (1, 2, 'Xarf Ee: Egg | Tuxum', 'Ee [ iː ]'),
    (1, 2, 'Xarf Ff: Flower | Gul', 'Ff [ ɛf ]'),
    (1, 2, 'Xarf Gg: Gift | Sovg`a', 'Gg [ dʒiː ]'),
    (1, 2, 'Xarf Hh: House | Uy', 'Hh [ eɪtʃ ]'),
    (1, 2, 'Xarf Ii: Ice cream | Muzqaymoq', 'Ii [ aɪ ]'),
    (1, 2, 'Xarf Jj: Juice | Ichimlik', 'Jj [ dʒeɪ ]'),
    (1, 2, 'Xarf Kk: King | Podishox', 'Kk [ keɪ ]'),
    (1, 2, 'Xarf Ll: Lead | Barik yoki yaproq', 'Ll [ ɛl ]'),
    (1, 2, 'Xarf Mm: Milk | Sut', 'Mm [ ɛm ]'),
    (1, 2, 'Xarf Nn: Notebook | Daftar', 'Nn [ ɛn ]'),
    (1, 2, 'Xarf Oo: Ocean | Dengiz', 'Oo [ oʊ ]'),
    (1, 2, 'Xarf Pp: Pizza | Pitsa', 'Pp [ piː ]'),
    (1, 2, 'Xarf Qq: Question mark | So`roq belgisi', 'Qq [ kjuː ]'),
    (1, 2, 'Xarf Rr: Rain | Yomg`ir', 'Rr [ ɑr ]'),
    (1, 2, 'Xarf Ss: Sun | Quyosh', 'Ss [ ɛs ]'),
    (1, 2, 'Xarf Tt: Tree | Daraxt', 'Tt [ tiː ]'),
    (1, 2, 'Xarf Uu: Umbrella | Soyabon', 'Uu [ juː ]'),
    (1, 2, 'Xarf Vv: Voclano | Vulqon', 'Vv [ viː ]'),
    (1, 2, 'Xarf Ww: Window | Deraza', 'Ww [ ˈdʌbljuː ]'),
    (1, 2, 'Xarf Xx: X-ray | Rengen', 'Xx [ ɛks ]'),
    (1, 2, 'Xarf Yy: Yatch | Yaxta', 'Yy [ waɪ ]'),
    (1, 2, 'Xarf Zz: Zero | No`l', 'Zz [ zɛd ]');

INSERT INTO content_files (content_id, sound_data, image_data) VALUES 
    (1, 'app.asrlan.uz/media/audio/english/beginner/letters/a.mp3', 'app.asrlan.uz/media/images/objects/apple.png'),
    (2, 'app.asrlan.uz/media/audio/english/beginner/letters/b.mp3', 'app.asrlan.uz/media/images/objects/book.png'),
    (3, 'app.asrlan.uz/media/audio/english/beginner/letters/c.mp3', 'app.asrlan.uz/media/images/objects/car.png'),
    (4, 'app.asrlan.uz/media/audio/english/beginner/letters/d.mp3', 'app.asrlan.uz/media/images/objects/door.png'),
    (5, 'app.asrlan.uz/media/audio/english/beginner/letters/e.mp3', 'app.asrlan.uz/media/images/objects/egg.png'),
    (6, 'app.asrlan.uz/media/audio/english/beginner/letters/f.mp3', 'app.asrlan.uz/media/images/objects/flower.png'),
    (7, 'app.asrlan.uz/media/audio/english/beginner/letters/g.mp3', 'app.asrlan.uz/media/images/objects/gift.png'),
    (8, 'app.asrlan.uz/media/audio/english/beginner/letters/h.mp3', 'app.asrlan.uz/media/images/objects/house.png'),
    (9, 'app.asrlan.uz/media/audio/english/beginner/letters/i.mp3', 'app.asrlan.uz/media/images/objects/ice_cream.png'),
    (10, 'app.asrlan.uz/media/audio/english/beginner/letters/j.mp3', 'app.asrlan.uz/media/images/objects/juice.png'),
    (11, 'app.asrlan.uz/media/audio/english/beginner/letters/k.mp3', 'app.asrlan.uz/media/images/objects/king.png'),
    (12, 'app.asrlan.uz/media/audio/english/beginner/letters/l.mp3', 'app.asrlan.uz/media/images/objects/leaf.png'),
    (13, 'app.asrlan.uz/media/audio/english/beginner/letters/m.mp3', 'app.asrlan.uz/media/images/objects/milk.png'),
    (14, 'app.asrlan.uz/media/audio/english/beginner/letters/n.mp3', 'app.asrlan.uz/media/images/objects/notebook.png'),
    (15, 'app.asrlan.uz/media/audio/english/beginner/letters/o.mp3', 'app.asrlan.uz/media/images/objects/ocean.png'),
    (16, 'app.asrlan.uz/media/audio/english/beginner/letters/p.mp3', 'app.asrlan.uz/media/images/objects/pizza.png'),
    (17, 'app.asrlan.uz/media/audio/english/beginner/letters/q.mp3', 'app.asrlan.uz/media/images/objects/question_mark.png'),
    (18, 'app.asrlan.uz/media/audio/english/beginner/letters/r.mp3', 'app.asrlan.uz/media/images/objects/rainycloud.png'),
    (19, 'app.asrlan.uz/media/audio/english/beginner/letters/s.mp3', 'app.asrlan.uz/media/images/objects/sun.png'),
    (20, 'app.asrlan.uz/media/audio/english/beginner/letters/t.mp3', 'app.asrlan.uz/media/images/objects/tree.png'),
    (21, 'app.asrlan.uz/media/audio/english/beginner/letters/u.mp3', 'app.asrlan.uz/media/images/objects/umbrella.png'),
    (22, 'app.asrlan.uz/media/audio/english/beginner/letters/v.mp3', 'app.asrlan.uz/media/images/objects/volcano.png'),
    (23, 'app.asrlan.uz/media/audio/english/beginner/letters/w.mp3', 'app.asrlan.uz/media/images/objects/window.png'),
    (24, 'app.asrlan.uz/media/audio/english/beginner/letters/x.mp3', 'app.asrlan.uz/media/images/objects/x-ray.png'),
    (25, 'app.asrlan.uz/media/audio/english/beginner/letters/y.mp3', 'app.asrlan.uz/media/images/objects/yacht.png'),
    (26, 'app.asrlan.uz/media/audio/english/beginner/letters/z.mp3', 'app.asrlan.uz/media/images/objects/zero.png');

INSERT INTO contents (lesson_id, gentype, title) VALUES
    (1, 3, 'Alfavit va talaffuz');

INSERT INTO content_files (content_id, video_data) VALUES 
    (1, 'app.asrlan.uz/media/video/english/beginner/topic_1_v1');

INSERT INTO contents (lesson_id, gentype, title, correct_answer) VALUES 
    (3, 4, 'Qaysi rasmdagi narsa A harifidan boshlanadi?', 1),
    (3, 4, 'Qaysi rasmdagi narsa S harifidan boshlanadi?', 2),
    (3, 4, 'Qaysi rasmdagi narsa W harifidan boshlanadi?', 2);

INSERT INTO content_files (content_id, image_data) VALUES 
    (28, 'app.asrlan.uz/media/images/objects/apple.png'),
    (28, 'app.asrlan.uz/media/images/objects/sun.png'),
    (28, 'app.asrlan.uz/media/images/objects/car.png'),
    (28, 'app.asrlan.uz/media/images/objects/door.png'),

    (29, 'app.asrlan.uz/media/images/objects/book.png'),
    (29, 'app.asrlan.uz/media/images/objects/flower.png'),
    (29, 'app.asrlan.uz/media/images/objects/Good_morning.jpg'),
    (29, 'app.asrlan.uz/media/images/objects/Good_afternoon.jpg'),

    (30, 'app.asrlan.uz/media/images/objects/Good_evening.jpg'),
    (30, 'app.asrlan.uz/media/images/objects/window.png'),
    (30, 'app.asrlan.uz/media/images/objects/king.png'),
    (30, 'app.asrlan.uz/media/images/objects/ocean.png');


-- END 1 Topic 

-- ('V', 1), --vocab objs
INSERT INTO contents (lesson_id, gentype, title, text_data) VALUES
    (4, 1, 'Odamlar bilan salomlashish', 'Hello - Salom'),
    (4, 1, 'Odamlar bilan salomlashish', 'Hi - Salom'),
    (4, 1, 'Qanday qilib xayirlashish', 'Goodbye - Xayr'),
    (4, 2, 'Ertalab kimgadir salom berish uchun', 'Good morning - Hayrli tong'),
    (4, 2, 'Peshin vaqtida kimgadir salom berish uchun', 'Good afternoon - Hayrli kun'),
    (4, 2, 'Kechqurun odamlar bilan salomlashish', 'Good evening - Hayrli kech'),

    (4, 1, 'Birovning xolatini bilish uchun', 'How are you? - Qalaysiz?'),
    (4, 1, 'Biror kishini birinchi marta uchratganingizda', 'Nice to meet you - Tanishganimdan xursandman'),
    (4, 1, 'Minnatdorchilik bildirish', 'Thank you - Rahmat'),
    (4, 1, 'Muloyimlik bilan birovning etiborini jalb qilish', 'Excuse me - Kechirasiz');

INSERT INTO content_files (content_id, sound_data, image_data) VALUES 
    (34, 'app.asrlan.uz/media/audio/beginner/good_morning', 'app.asrlan.uz/media/images/objects/Good_morning.jpg'),
    (35, 'app.asrlan.uz/media/audio/beginner/good_afternoon', 'app.asrlan.uz/media/images/objects/Good_afternoon.jpg'),
    (36, 'app.asrlan.uz/media/audio/beginner/good_evening', 'app.asrlan.uz/media/images/objects/Good_evening.jpg');

INSERT INTO contents (lesson_id, gentype, title) VALUES
    (5, 3, 'Salomlashish va tanishtirish');

INSERT INTO content_files (content_id, video_data) VALUES 
    (40, 'app.asrlan.uz/media/video/english/beginner/topic_2_v1');

INSERT INTO contents (lesson_id, gentype, title, text_data) VALUES
    (5, 1, 'Ikki kishini suhbatini o`qing', 'Person A: Hello! Person B: Hi! How are you? Person A: I’m good, thank you. How about you? Person B: I’m fine, thanks. What’s your name? Person A: My name is John. What’s your name? Person B: I’m Sarah. Nice to meet you, John.');

INSERT INTO contents (lesson_id, gentype, title, correct_answer) VALUES 
    (6, 4, 'Qaysi rasm "Hayrli tong"ni ifodalaydi?', 2),
    (6, 4, 'Qaysi rasm "Hayrli kun"ni ifodalaydi?', 4),
    (6, 4, 'Qaysi rasm "Hayrli kech"ni ifodalaydi?', 1);

INSERT INTO content_files (content_id, image_data) VALUES 
    (41, 'app.asrlan.uz/media/images/objects/apple.png'),
    (41, 'app.asrlan.uz/media/images/objects/Good_morning.jpg'),
    (41, 'app.asrlan.uz/media/images/objects/car.png'),
    (41, 'app.asrlan.uz/media/images/objects/door.png'),

    (42, 'app.asrlan.uz/media/images/objects/book.png'),
    (42, 'app.asrlan.uz/media/images/objects/flower.png'),
    (42, 'app.asrlan.uz/media/images/objects/Good_morning.jpg'),
    (42, 'app.asrlan.uz/media/images/objects/Good_afternoon.jpg'),

    (43, 'app.asrlan.uz/media/images/objects/Good_evening.jpg'),
    (43, 'app.asrlan.uz/media/images/objects/door.png'),
    (43, 'app.asrlan.uz/media/images/objects/king.png'),
    (43, 'app.asrlan.uz/media/images/objects/ocean.png');