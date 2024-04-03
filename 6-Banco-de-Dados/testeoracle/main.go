package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

type Banco struct {
	Numero int64
	Nome   string
}

func main() {

	fmt.Println("=> Conectando ao Oracle")

	db, err := sql.Open("godror", `user="CRMCARONE" password="sql123" connectString="ORALINUXCRM"`)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("=> Pesquisando Banco 70.")

	var numero int64 = 70

	banco, err := selectBanco(db, numero)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Banco: %d - %s\n", banco.Numero, banco.Nome)

	fmt.Println("=> Listar Bancos.")

	bancos, err := selectAllBancos(db)
	if err != nil {
		panic(err)
	}
	for _, b := range bancos {
		fmt.Printf("Banco: %d - %s\n", b.Numero, b.Nome)
	}

	fmt.Println("=> Fim.")

}

func selectBanco(db *sql.DB, num int64) (*Banco, error) {

	fmt.Println("=> Preparando SQL.")
	stmt, err := db.Prepare("select bannumero, bannome from bancos where bannumero = :1 ")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	fmt.Println("=> Pesquisando SQL. ", num)

	var b Banco

	err = stmt.QueryRow(num).Scan(&b.Numero, &b.Nome)

	if err != nil {
		return nil, err
	}

	fmt.Println("=> Retornando Banco. ", b)

	return &b, nil
}

func selectAllBancos(db *sql.DB) ([]Banco, error) {

	rows, err := db.Query("select bannumero, bannome from bancos order by bannumero")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bancos []Banco
	for rows.Next() {
		var b Banco
		err = rows.Scan(&b.Numero, &b.Nome)
		if err != nil {
			return nil, err
		}

		bancos = append(bancos, b)
	}

	return bancos, nil
}
