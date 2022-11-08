-- Consulta 1
SELECT m.title, g.name 'genre' 
FROM movies m 
INNER JOIN genres g ON m.genre_id = g.id;

-- Consulta 2
SELECT e.title, a.first_name, a.last_name FROM episodes e
INNER JOIN actor_episode ae ON ae.episode_id = e.id
INNER JOIN actors a ON ae.actor_id = a.id;

-- Consulta 3
SELECT s.title, COUNT(se.id) count_seasons FROM series s
INNER JOIN seasons se ON se.serie_id = s.id
GROUP BY s.id;

-- Consulta 4
SELECT g.name, COUNT(m.id) count_movies FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.id
HAVING count_movies >= 3;

-- Consulta 5
SELECT DISTINCT a.first_name, a.last_name FROM actors a
INNER JOIN actor_movie am ON am.actor_id = a.id
INNER JOIN movies m ON am.movie_id = m.id AND m.title LIKE '%La Guerra de las galaxias%';