package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const UPLOADPATH = "./uploads" //Será utilizado como volúmen en Docker

func SaveProfileImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	//Crear la carpeta si no existe
	if _, err := os.Stat(UPLOADPATH); os.IsNotExist(err) {
		os.MkdirAll(UPLOADPATH, os.ModePerm)
	}

	//Generación nombre único para el archivo
	extension := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	filePath := filepath.Join(UPLOADPATH, fileName)

	//Crear archivo en el servidor
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	//Copiar contenido del archivo cargado
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	//Retornar ruta relativa (Se almacena en BD)
	return fileName, nil
}
