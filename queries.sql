CREATE database local_db;

USE local_db;

CREATE TABLE IF NOT EXISTS users (
  id integer PRIMARY KEY AUTO_INCREMENT,
  user_id varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  userpassword varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  -- number of people the recipe will feed
  nickname varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  -- food items necessary to prepare the recipe
  comment varchar(300) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recipes (
  id integer PRIMARY KEY AUTO_INCREMENT,
  -- name of recipe
  title varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  -- time required to cook/bake the recipe
  making_time varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  -- number of people the recipe will feed
  serves varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  -- food items necessary to prepare the recipe
  ingredients varchar(300) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  -- price of recipe
  cost integer NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

SELECT * FROM recipes ;

SELECT COUNT(*) FROM recipes WHERE id = 1;

SELECT COUNT(*) as count FROM recipes WHERE id = 2;

INSERT INTO recipes (
  id,
  title,
  making_time,
  serves,
  ingredients,
  cost,
  created_at,
  updated_at
)
VALUES (
  1,
  'Chicken Curry',
  '45 min',
  '4 people',
  'onion, chicken, seasoning',
  1000,
  '2016-01-10 12:10:12',
  '2016-01-10 12:10:12'
);

INSERT INTO recipes (
  id,
  title,
  making_time,
  serves,
  ingredients,
  cost,
  created_at,
  updated_at
)
VALUES (
  2,
  'Rice Omelette',
  '30 min',
  '2 people',
  'onion, egg, seasoning, soy sauce',
  700,
  '2016-01-11 13:10:12',
  '2016-01-11 13:10:12'
); 

INSERT INTO recipes (title, making_time, serves, ingredients, cost)  VALUES ("Veg Biryani", "45 mins", "5 people", "onion, tomato, seasoning, water, Paneer, rice", "450");