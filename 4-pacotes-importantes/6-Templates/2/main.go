package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n"))
	
	// tmp := template.New("CursoTemplate")

	// tmp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n")
	// if err != nil {
	// 	panic(err)
	// }

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}