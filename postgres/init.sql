CREATE TABLE "users" (
  user_id     int,
  lastname    varchar(255),
  firstname   varchar(255),
  state       varchar(2),
  city        varchar(255)
);

INSERT INTO users (user_id, lastname, firstname, state, city)
VALUES
  (0001, 'Regenold', 'Kali',  'MN', 'St. Paul'),
  (0002, 'Seaborn',  'Sam',   'CA', 'Laguna Beach'),
  (0003, 'Cregg',    'C. J.', 'OH', 'Dayton'),
  (0004, 'McGarry',  'Leo',   'IL', 'Chicago'),
  (0005, 'Lyman',    'Josh',  'CT', 'Westport'),
  (0006, 'Bartlet',  'Jed',   'NH', 'Manchester'),
  (0007, 'Ziegler',  'Toby',  'NY', 'New York City');
