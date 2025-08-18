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
INSERT INTO usuarios (username, password, avatar_path, created_at, updated_at, deleted_at) VALUES
('Benjamin', '$2a$12$YW1ikzXms3OV8/wNfF4avOxi67/1MJfBHseqKCgbcusZTF.WztTbu', '/static/default.jpg', '2025-07-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
('CARLOS98', '$2a$12$akd3aYgtsx2jrRpYZ9SvUeaU9isSc./tXVl36j8f7kxzw4Rc473VK', '/static/default.jpg', '2025-07-19T05:01:22.410608857Z', '2025-07-19T05:01:22.410608857Z', NULL),
('Carlos18', '$2a$12$MqVjwxXzp1W2ZbyfW8EqyO94FYp10pddqVpdOJZrkrth98PANBm.u', '/static/default.jpg', '2025-07-23T06:06:22.410608857Z', '2025-07-23T06:06:22.410608857Z', NULL),
('Andres', '$2a$12$DdFD7OlqqrBtAFk3kKOTV.Px9O9idbRq1VkRw5jXEB2lFds3LUfTC', '/static/default.jpg', '2025-07-27T08:43:22.410608857Z', '2025-07-27T08:43:22.410608857Z', NULL),
('CamilaSuarez', '$2a$12$IUJIB1hR9s/j9Qh2x8oivePlMk5jCY.DAXTBry.Jo01X1/RqV1DTG', '/static/default.jpg', '2025-07-30T02:01:22.410608857Z', '2025-07-30T02:01:22.410608857Z', NULL),
('Maicol Lopez', '$2a$12$tc5Zw/v3ES9sRk93..SZeeMZQawywb6JvnB2RNoONrMxzJgGtbwVO', '/static/default.jpg', '2025-08-01T05:01:22.410608857Z', '2025-08-01T05:01:22.410608857Z', NULL),
('FranciscaE', '$2a$12$46yuHKn5uiZixS7nuh6gxutAcFG6iZQkq8MDcK9nbtSbOELm7BRcy', '/static/default.jpg', '2025-08-04T15:01:22.410608857Z', '2025-08-04T15:01:22.410608857Z', NULL),
('Correo', '$2a$12$m7WrXGkuOsE9spvVOOYf7OCsnBdgg/deFr5dh7HTVSP6SrwMFONce', '/static/default.jpg', '2025-08-7T00:01:22.410608857Z', '2025-08-7T00:01:22.410608857Z', NULL),
('485787', '$2a$12$5yt9NzJu3Y565LyYiemM5O8LQnXmVYExrLO.OcEW9rpWPDiX8f1L2', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
('Gaston Fe', '$2a$12$ZjV0woMT/R/I8mTWdQ7v/uhEN/ZJCAKPGjniYs2feyqyPqvEGpDDi', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
('Palmera', '$2a$12$AZ9pjThmv8a4NRE5ZGop7eT7MXZRfW2VrYQGRcfT.snDRb86AwNUS', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
('Deadpool', '$2a$12$dwE8HXNZcmYmfQiCNQOxBuujUr2UdSEeQ9HCTNG4zs3p1JIAAKhR2', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
('Correo PRUEBA', '$2a$12$29ajc1GO3G0m7MAOL0RSIO.XET.NXfXqGTdKuU4ZDnec2l9BUgepi', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL),
('Carlos98', '$2a$12$sToX6YvezfEN0zaQEegKVOXi19f0V96F.kQBRmWXE7BRKGFOJ4LF2', '/static/default.jpg', '2025-08-17T05:01:22.410608857Z', '2025-08-17T05:01:22.410608857Z', NULL)
ON CONFLICT DO NOTHING;

-- Definición tabla de categoría
CREATE TABLE IF NOT EXISTS categorias(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);


-- Insertar datos en la tabla de categorias
INSERT INTO categorias (name, description, created_at) VALUES
('Trabajos', 'Trabajos universidad', '2025-07-17 12:00:00'),
('Trabajos', 'Trabajos universidad', '2025-07-19 12:00:00'),
('Proyectos', 'Proyectos migración Apps Empresa', '2025-07-20 12:35:00'),
('Hogar', '', '2025-07-17 12:01:22'),
('Estudio', 'Actividades de estudio', '2025-07-24 05:02:48'),
('Salud', 'Actividades para mejorar hábitos', '2025-07-28 01:48:12'),
('Finanzas', 'Validaciones de mis finanzas', '2025-07-30 20:08:01'),
('Social/Ocio', 'Actividades de descanso', '2025-08-08 08:08:08'),
('Estudio', 'Actividades de estudio matemáticas', '2025-08-10 10:12:00'),
('Finanzas', 'Validaciones de mis finanzas', '2025-08-13 17:53:25'),
('Hogar', 'Actividades para el apartamento', '2025-07-17 12:00:00'),
('Proyectos/creatividad', 'Creación de contenido', '2025-08-15 04:05:06'),
('Estudio', '', '2025-07-17 12:00:00'),
('Finanzas', '', '2025-08-17 17:17:17')
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
