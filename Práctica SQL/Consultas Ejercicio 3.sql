-- Consulta 1
SELECT * FROM movies;

-- Consulta 2
SELECT first_name, last_name, rating FROM actors;

-- Consulta 3
SELECT title AS titulo FROM series AS series_en_espanol;

-- Consulta 4
SELECT first_name, last_name FROM actors WHERE rating > 7.5;

-- Consulta 5
SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;

-- Consulta 6
SELECT title, rating FROM movies ORDER BY rating ASC;

-- Consulta 7
SELECT title FROM movies LIMIT 3;

-- Consulta 8
SELECT title, rating FROM movies ORDER BY rating DESC LIMIT 5;

-- Consulta 9
SELECT * FROM actors LIMIT 10;

-- Consulta 10
SELECT title, rating FROM movies WHERE title LIKE '%Toy Story%';

-- Consulta 11
SELECT * FROM actors WHERE first_name LIKE 'Sam%';

-- Consulta 12
SELECT title FROM movies WHERE release_date BETWEEN '2004-01-01 00:00:00' AND '2008-12-31 23:59:59';

-- Consulta 13
SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND release_date BETWEEN '1988-01-01 00:00:00' AND '2009-12-31 23:59:59' ORDER BY rating DESC;