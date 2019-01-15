CREATE USER postgres_user;
CREATE DATABASE mydatabase;
GRANT ALL PRIVILEGES ON DATABASE mydatabase TO postgres_user;

CREATE TABLE users (
    user_id     int,
    lastname    varchar(255),
    firstname   varchar(255),
    state       varchar(2),
    city        varchar(255)
);

INSERT INTO users (user_id, lastname, firstname, state, city) VALUES (0001, 'Regenold', 'Kali',  'MN', 'St. Paul');
INSERT INTO users (user_id, lastname, firstname, state, city) VALUES (0002, 'Seaborn',  'Sam',   'CA', 'Laguna Beach');
INSERT INTO users (user_id, lastname, firstname, state, city) VALUES (0003, 'Cregg',    'C. J.', 'OH', 'Dayton');
INSERT INTO users (user_id, lastname, firstname, state)       VALUES (0004, 'Ziegler',  'Toby',  'NY');
INSERT INTO users (user_id, lastname, firstname, state, city) VALUES (0005, 'McGarry',  'Leo',   'IL', 'Chicago');
INSERT INTO users (user_id, lastname, firstname, state, city) VALUES (0006, 'McGarry',  'Leo',   'IL', 'Chicago');
INSERT INTO users (user_id, lastname, firstname, state, city) VALUES (0007, 'Bartlet',  'Jed',   'NH', 'Manchester');
