-- Definición tabla de usuarios
CREATE TABLE IF NOT EXISTS usuarios(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    avatar_path VARCHAR(255) DEFAULT '/static/default.jpg',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Insertar datos en la tabla de usuarios
INSERT INTO usuarios (id, username, password, avatar_path, created_at, updated_at, deleted_at) VALUES
(1,'Benjamin', '$2a$12$abc123usuario01hash', '/static/default.jpg', '2025-07-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
(2, 'CARLOS98', '$2a$12$abc123usuario02hash', '/static/default.jpg', '2025-07-19T05:01:22.410608857Z', '2025-07-19T05:01:22.410608857Z', NULL),
(3, 'Carlos18', '$2a$12$abc123usuario03hash', '/static/default.jpg', '2025-07-23T06:06:22.410608857Z', '2025-07-23T06:06:22.410608857Z', NULL),
(4, 'Andres', '$2a$12$abc123usuario04hash', '/static/default.jpg', '2025-07-27T08:43:22.410608857Z', '2025-07-27T08:43:22.410608857Z', NULL),
(5, 'CamilaSuarez', '$2a$12$abc123usuario05hash', '/static/default.jpg', '2025-07-30T02:01:22.410608857Z', '2025-07-30T02:01:22.410608857Z', NULL),
(6, 'Maicol Lopez', '$2a$12$abc123usuario06hash', '/static/default.jpg', '2025-08-01T05:01:22.410608857Z', '2025-08-01T05:01:22.410608857Z', NULL),
(7, 'FranciscaE', '$2a$12$abc123usuario07hash', '/static/default.jpg', '2025-08-04T15:01:22.410608857Z', '2025-08-04T15:01:22.410608857Z', NULL),
(8, 'Correo', '$2a$12$abc123usuario08hash', '/static/default.jpg', '2025-08-7T00:01:22.410608857Z', '2025-08-7T00:01:22.410608857Z', NULL),
(9, '485787', '$2a$12$abc123usuario09hash', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
(10, 'Gaston Fe', '$2a$12$abc123usuario10hash', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
(11, 'Palmera', '$2a$12$abc123usuario11hash', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
(12, 'Deadpool', '$2a$12$abc123usuario12hash', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
(13, 'Correo PRUEBA', '$2a$12$abc123usuario13hash', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
(14, 'Carlos98', '$2a$12$abc123usuario14hash', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL)
ON CONFLICT DO NOTHING;

-- Definición tabla de categoría
CREATE TABLE IF NOT EXISTS categorias(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);


-- Insertar datos en la tabla de categorias
INSERT INTO categorias (id, name, description, created_at) VALUES
(1, 'Trabajos', 'Trabajos universidad', '2025-07-17 12:00:00'),
(2, 'Trabajos', 'Trabajos universidad', '2025-07-19 12:00:00'),
(3, 'Proyectos', 'Proyectos migración Apps Empresa', '2025-07-20 12:35:00'),
(4, 'Hogar', '', '2025-07-17 12:01:22'),
(5, 'Estudio', 'Actividades de estudio', '2025-07-24 05:02:48'),
(6, 'Salud', 'Actividades para mejorar hábitos', '2025-07-28 01:48:12'),
(7, 'Finanzas', 'Validaciones de mis finanzas', '2025-07-30 20:08:01'),
(8, 'Social/Ocio', 'Actividades de descanso', '2025-08-08 08:08:08'),
(9, 'Estudio', 'Actividades de estudio matemáticas', '2025-08-10 10:12:00'),
(10, 'Finanzas', 'Validaciones de mis finanzas', '2025-08-13 17:53:25'),
(11, 'Hogar', 'Actividades para el apartamento', '2025-07-17 12:00:00'),
(12, 'Proyectos/creatividad', 'Creación de contenido', '2025-08-15 04:05:06'),
(13, 'Estudio', '', '2025-07-17 12:00:00'),
(14, 'Finanzas', '', '2025-08-17 17:17:17')
ON CONFLICT DO NOTHING;

-- Definición tabla de de tareas
CREATE TABLE IF NOT EXISTS tareas(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    category_id BIGINT,
    text TEXT NOT NULL,
    due_date TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    status VARCHAR(50) DEFAULT 'Pendiente',

    -- Relaciones
    CONSTRAINT fk_tares_usuarios FOREIGN KEY (user_id)
        REFERENCES usuarios(id) ON DELETE CASCADE,

    CONSTRAINT fk_tareas_categorias FOREIGN KEY (category_id)
        REFERENCES categorias (id) ON DELETE CASCADE
);

-- Insertar datos en la tabla de tareas
INSERT INTO tareas (user_id, category_id, text, due_date, status)
VALUES
(1, 1, 'Entregar informe final de la universidad', '2025-08-20 23:59:59', 'Pendiente'),
(2, 3, 'Subir avances de migración de apps', '2025-08-25 18:00:00', 'En Progreso'),
(3, 5, 'Estudiar para examen de bases de datos', '2025-08-19 08:00:00', 'Pendiente'),
(4, 6, 'Ir al médico general', '2025-08-22 10:00:00', 'Pendiente'),
(5, 7, 'Revisar gastos del mes', '2025-08-30 20:00:00', 'Pendiente'),
(6, 8, 'Reunión con amigos', '2025-08-21 19:30:00', 'Pendiente'),
(7, 9, 'Resolver ejercicios de matemáticas', '2025-08-23 15:00:00', 'Pendiente'),
(8, 10, 'Validar presupuesto familiar', '2025-08-28 21:00:00', 'Pendiente'),
(9, 11, 'Limpieza general del apartamento', '2025-08-19 09:00:00', 'Pendiente'),
(10, 12, 'Crear contenido para redes sociales', '2025-08-24 14:00:00', 'Pendiente'),
(1, 1, 'Preparar presentación para reunión del lunes', '2025-08-24 14:00:00', 'En progreso'),
(1, 4, 'Comprar víveres', '2025-08-19 09:00:00', 'Pendiente'),
(1, 4, 'Limpiar cocina y baño', '2025-08-19 09:00:00', 'Pendiente'),
(1, 4, 'Lavar ropa', '2025-08-22 09:00:00', 'Pendiente'),
(1, 4, 'Pagar servicios públicos', '2025-08-25 09:00:00', 'Pendiente');
