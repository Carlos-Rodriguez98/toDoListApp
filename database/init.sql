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

