# 🧑‍💻 Guía de colaboración para el proyecto de la aplicación Web To Do List

Este documento define las reglas y buenas prácticas para colaborar en este repositorio, garantizando una estructura organizada y un flujo de trabajo eficiente.

---

## 🌱 Estructura de ramas

- `main`: Rama principal (producción). Contiene código aprobado y estable.
- `dev`: Rama de integración. Aquí se integran todos los desarrollos antes de pasar a `main`.
- `feature/[nombre]`: Ramas individuales por desarrollador para trabajar de forma aislada.

---

## 🔁 Flujo de trabajo recomendado

1. **Crear tu rama de desarrollo**
   - Desde `dev`, crea una rama con tu nombre o tarea específica:
     git checkout dev
     git pull origin dev
     git checkout -b feature/tu-nombre
     git push -u origin feature/tu-nombre

2. **Desarrollar en tu rama**
   - Realiza cambios y haz commits con mensajes claros:
     git add .
     git commit -m "feat: agregar formulario de tareas"
     git push

3. **Hacer Pull Request a 'dev'**
   3.1. Tener instalada la extensión GitHub Pull Request.
   3.2. Abre la extensión desde el panel lateral de Visual Studio Code.
   3.3. Seleccionar el icono de Create Pull Request
   3.4. Seleccionar tu rama de trabajo (Ej: `feature/carlos`) como origen y `dev` como destino.
   3.5. Asigna un titulo y una breve descripción.
   3.6. Haz Clic en Create. 

4. **Actualizar tu rama con cambios de `dev` (cuando sea necesario):**
   git checkout feature/tu-nombre
   git pull origin dev


