package models

// aca defino el modelo que devera tener la peticion que ire a hacer
// es decir el formato que sera en este caso Json y como debera de ser lo que esperaé resibir

// estrucuta de lo que espero recibir, tanto en el tipo de informacion como el orden
type User struct {
	Id       int    `Json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}
