-- consulta 1
SELECT * FROM autor;

-- consulta 2
SELECT nombre, edad FROM estudiante;

-- consulta 3
SELECT nombre FROM estudiante WHERE carrera = 'Informatica';

-- consulta 4
SELECT * FROM autor WHERE nacionalidad = "francesa" OR nacionalidad = "italiana";

-- consulta 5
SELECT * FROM libro WHERE area != "internet";

-- consulta 6
SELECT * FROM libro WHERE editorial = "Salamandra";

-- consulta 7
SELECT * FROM estudiante WHERE edad > (SELECT AVG(edad) FROM estudiante);

-- consulta 8 
SELECT nombre FROM estudiante WHERE nombre LIKE 'G%';

-- consulta 9
SELECT autor.nombre FROM autor
INNER JOIN libroautor ON libroautor.id_autor = autor.id_autor
INNER JOIN libro ON libroautor.id_libro = libro.id_libro 
AND libro.titulo = 'El Universo: Guía de viaje';

-- consulta 10
SELECT libro.titulo FROM libro
INNER JOIN prestamo ON prestamo.id_libro = libro.id_libro
INNER JOIN estudiante ON prestamo.id_lector = estudiante.id_estudiante
AND estudiante.nombre = "Filippo" AND estudiante.apellido = "Galli";

-- consulta 11
SELECT * FROM estudiante
WHERE edad=(SELECT MIN(edad) FROM estudiante);

-- consulta 12
SELECT DISTINCT estudiante.nombre FROM estudiante
INNER JOIN prestamo ON prestamo.id_lector = estudiante.id_estudiante
INNER JOIN libro ON prestamo.id_libro = libro.id_libro
AND libro.area = "Base de Datos";

-- Consulta 13
SELECT libro.* FROM libro
INNER JOIN libroautor ON libroautor.id_libro = libro.id_libro
INNER JOIN autor ON libroautor.id_autor = autor.id_autor
AND autor.nombre = "J.K. Rowling";

-- consulta 14
SELECT * FROM libro
INNER JOIN prestamo ON prestamo.id_libro = libro.id_libro
AND prestamo.fecha_devolucion = '2021-07-16';