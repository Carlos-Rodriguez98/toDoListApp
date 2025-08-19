# To Do List App

## Integrantes
- Carlos Stiven Rodriguez Aroca (cs.rodrigueza1@uniandes.edu.co)
- Jefferson Alberto Hernandez Garcia (ja.hernandezg1@uniandes.edu.co)
- Valentin Charles Michel Pecqueux (v.pecqueux@uniandes.edu.co)
- David Andres Paredes Bravo (da.paredes2@uniandes.edu.co)
- Harold Nicolas Coca Peña (h.coca@uniandes.edu.co)

## Introducción
To Do List App es una aplicación de gestión de tareas moderna y distribuida, construida con una arquitectura de microservicios. La aplicación permite a los usuarios gestionar sus tareas diarias de manera eficiente, organizándolas por categorías y manteniendo un seguimiento de su progreso.

## Características Principales
- Sistema de autenticación y registro de usuarios
- Perfiles de usuario con avatares personalizables
- Organización de tareas por categorías
- Gestión completa de tareas (crear, editar, eliminar, marcar como completadas)
- Interfaz de usuario moderna y responsiva
- Arquitectura de microservicios
- Seguridad mediante tokens JWT
- Containerización con Docker para fácil despliegue

## Arquitectura

### Vista de Componentes
![Vista de Componentes](docs/artifacts/components-vw.png "Vista de Componentes")

**Componentes y responsabilidades**:

* **Frontend (HTML/CSS/JavaScript)**

  * Renderiza la UI (listas de tareas/categorías, formularios de login/registro).
  * Llama a los microservicios vía **HTTP/JSON**.
  * Gestiona el token JWT en el navegador (localStorage o cookies seguras).
  * Aplica filtros de tareas (por estado/categoría) a partir de datos de la API.

* **Auth Service (Go)**

  * **Controllers**: exponen endpoints `/usuarios` (registro), `/usuarios/iniciar-sesion` (login).
  * **Logic/Utils**: hashing de contraseñas (bcrypt), emisión/verificación de **JWT**, validaciones, manejo de avatar (subida y asignación de avatar por defecto).
  * Persiste y consulta usuarios en la base de datos.

* **Category Service (Go)**

  * **Controllers**: CRUD de categorías (`/categorias`).
  * Valida datos (nombre requerido, unicidad si se define), y verifica JWT si el acceso es autenticado.
  * Lee/escribe categorías en la base de datos.

* **Task Service (Go)**

  * **Controllers**: CRUD de tareas (`/tareas`) y lecturas filtradas (`/tareas/usuario?categoria=...&status=...`).
  * **Logic**: reglas de negocio (solo ver/editar tareas propias, transición de estados `NOT_STARTED → STARTED → COMPLETED`, validación de fechas).
  * Persiste y consulta tareas en la base de datos.

* **PostgreSQL**

  * Almacena **usuarios, categorías y tareas**.
  * Los servicios acceden mediante consultas parametrizadas (evitando inyección SQL).
  * Índices para acelerar búsquedas por `user_id` y `category_id`.

**Flujo de trabajo**

1. **Registro**: Web App → Auth Service (`POST /usuarios`, opcional imagen). Auth Service guarda usuario y avatar (o asigna uno por defecto).
2. **Login**: Web App → Auth Service (`POST /usuarios/iniciar-sesion`). Devuelve **JWT**.
3. **Navegación**:

   * Categorías: Web App → Category Service (`GET /categorias`).
   * Tareas del usuario: Web App → Task Service (`GET /tareas/usuario`) con `Authorization: Bearer <JWT>`.
4. **Edición**: Crear/actualizar/borrar tareas y categorías mediante endpoints REST; Task Service valida la autoría con el **JWT**.

* Cada servicio tiene su “capa de controladores” (HTTP) y “lógica” (reglas/validaciones).
* Se comparte una única BD (patrón **DB-shared** entre microservicios); simple en local.
* Contratos REST simples, formato JSON, y autenticación **Bearer JWT**.

### Vista de Despliegue
![Vista de Despliegue](docs/artifacts/deploy-vw.png "Vista de Despliegue")

* **Host Local (localhost)** con **Docker Engine**.
* Contenedores separados:

  * `frontend-service` (Nginx sirviendo estáticos) **localhost:8083**.
  * `auth-service` (Go) — p.ej. **localhost:8080**.
  * `category-service` (Go) — p.ej. **localhost:8081**.
  * `task-service` (Go) — p.ej. **localhost:8082**.
  * `todolistapp-database` — **localhost:5432** (con volumen persistente).
* **Red de Docker** compartida para que los servicios se resuelvan por nombre (ej.: `postgres:5432`).

**Conexiones**

* Web App → cada microservicio vía **HTTP/JSON** (puertos publicados al host).
* Microservicios → **PostgreSQL** vía **SQL/TCP** (con `DATABASE_URL`).
* Variables de entorno típicas:

  * `DATABASE_URL=postgres://user:pass@postgres:5432/tododb?sslmode=disable`
  * `JWT_SECRET=...`
  * `MAX_UPLOAD_SIZE`, `DEFAULT_AVATAR_URL`, etc.

### Diagrama Entidad-Relación
![Diagrama Entidad-Relación](docs/artifacts/er-vw.png "Diagrama Entidad Relación")

**Entidades**

* **USUARIO**

  * `id` (PK), `username` (único), `password_hash`, `avatar_path` (nullable), `created_at`.
  * Reglas: `username` único; contraseña almacenada **hash** (bcrypt); si no hay `avatar_path`, se usa un avatar por defecto.

* **CATEGORIA**

  * `id` (PK), `name`, `description` (nullable), `created_at`.
  * Reglas: nombre requerido; puedes añadir unicidad si lo deseas por usuario (si las categorías son personales) o global.

* **TAREA**

  * `id` (PK), `user_id` (FK→USUARIO), `category_id` (FK→CATEGORIA, **nullable** con `ON DELETE SET NULL`), `text`, `created_at`, `due_date` (nullable), `status`.
  * `status` ∈ {`NOT_STARTED`, `STARTED`, `COMPLETED`} (puedes implementar como `CHECK` o enumeración de aplicación).
  * Índices por `user_id` y `category_id` para filtros frecuentes.

**Cardinalidades y comportamiento**

* **USUARIO 1—N TAREA**: un usuario puede tener muchas tareas; `ON DELETE CASCADE` en `tasks.user_id` elimina sus tareas al borrar el usuario.
* **CATEGORIA 1—N TAREA**: una categoría puede agrupar muchas tareas; `ON DELETE SET NULL` evita borrar tareas si se elimina la categoría (quedan sin categoría).

### Flujo de Trabajo
La aplicación está construida siguiendo una arquitectura de microservicios, donde cada servicio es responsable de una funcionalidad específica:

1. **Auth Service (Puerto 8080)**: 
   - Gestiona la autenticación y registro de usuarios
   - Maneja tokens JWT para sesiones seguras
   - Almacena y sirve avatares de usuario

2. **Category Service (Puerto 8081)**:
   - Gestiona las categorías de tareas
   - Permite organizar tareas en grupos lógicos

3. **Task Service (Puerto 8082)**:
   - Maneja el CRUD completo de tareas
   - Asocia tareas con usuarios y categorías

4. **Frontend Service (Puerto 8083)**:
   - Sirve la interfaz de usuario
   - Interactúa con los demás servicios
   - Proporciona una experiencia fluida al usuario

## Estructura del Repositorio
```
├── authService/           # Servicio de autenticación
│   ├── config/           # Configuración de BD y variables de entorno
│   ├── controllers/      # Controladores de login y registro
│   ├── models/          # Modelo de usuario
│   ├── services/        # Lógica de negocio de autenticación
│   └── utils/           # Utilidades (hash, JWT, almacenamiento)
├── categoryService/      # Servicio de gestión de categorías
├── taskService/         # Servicio de gestión de tareas
│   ├── controllers/     # Controladores de tareas
│   ├── models/         # Modelo de tarea
│   └── services/       # Lógica de negocio de tareas
├── frontend/           # Interfaz de usuario
│   ├── categories/     # Componentes de categorías
│   └── tasks/         # Componentes de tareas
└── database/          # Scripts de inicialización de BD
```

## Uso

### Requisitos Previos
- Docker y Docker Compose
- Espacio en disco para imágenes Docker
- Puertos 8080-8083 disponibles

### Instrucciones de Ejecución

1. Clonar el repositorio:
```bash
git clone git@github.com:Carlos-Rodriguez98/toDoListApp.git
cd toDoListApp
```

2. Iniciar los servicios con Docker Compose:
```bash
docker-compose up -d
```

3. Acceder a la aplicación:
   - Frontend: http://localhost:8083
   - Servicios API:
     - Auth Service: http://localhost:8080
     - Category Service: http://localhost:8081
     - Task Service: http://localhost:8082

### Variables de Entorno
Cada servicio utiliza sus propias variables de entorno para configuración:

- Base de datos:
  - POSTGRES_USER=Admin
  - POSTGRES_PASSWORD=Admin
  - POSTGRES_DB=toDoListApp

Los servicios se conectarán automáticamente a la base de datos usando las credenciales configuradas en el docker-compose.yml.
