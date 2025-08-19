# Usamos Nginx como servidor web
FROM nginx:stable-alpine

# Establecemos el directorio de trabajo
WORKDIR /usr/share/nginx/html

# Eliminamos archivos por defecto
RUN rm -rf ./*

# Copiamos los archivos estáticos (html, css, js) al directorio de nginx
COPY . .

# Exponemos el puerto 80 (HTTP)
EXPOSE 80

# Nginx ya viene con el CMD por defecto, no necesitas definirlo