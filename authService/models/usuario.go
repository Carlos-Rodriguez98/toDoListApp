type Usuario struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	name         string `json:"nombreUsuario"`
	password     string `json:"-"` //This information isn't returned.
	imagenPerfil string `json:"imagenPerfil"`
}