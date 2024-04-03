package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Moeda struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Cotacao struct {
	Moeda Moeda `json:"USDBRL"`
}

type DadosCotacao struct {
	Id         int
	Code       string
	Codein     string
	Name       string
	Valor      float64
	CreateDate string
}


var (
	db *sql.DB
)

func main() {

	var err error

	fmt.Println("===> Iniciando - Conexão com Banco de Dados")
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		fmt.Println("#> Falha na Conexão Banco de Dados.")
		panic(err)
	} else {
		fmt.Println("===> Iniciando - Web Service")
		http.HandleFunc("/cotacao", BuscarCotacaoHandler)
		http.ListenAndServe(":8081", nil)
	}
	defer db.Close()
	fmt.Println("===> Encerrando o Serviço.")
}

func BuscarCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	cotacao, err := BuscaCotacaoDolar()
	if err != nil {
		fmt.Println("Falha na Busca da Cotação.")
		w.Write([]byte(`{"erro": "Falha na Busca da Cotação."}`))
		//w.WriteHeader(http.StatusInternalServerError)
	} else {

		d := NewCotacao(cotacao.Moeda)

		fmt.Println(d)

		json.NewEncoder(w).Encode(cotacao.Moeda)

		err = InsertCotacao(d)
		if err != nil {
			fmt.Println(err)
			//w.WriteHeader(http.StatusInternalServerError)
		}

	}

}

func BuscaCotacaoDolar() (*Cotacao, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Falha de Conexão: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c Cotacao
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func NewCotacao(c Moeda) *DadosCotacao {

	fmt.Println(c)

	valor, err := strconv.ParseFloat(c.Bid, 64)
	if err != nil {
		valor = 0
	}

	return &DadosCotacao{
		Id:         0,
		Code:       c.Code,
		Codein:     c.Codein,
		Name:       c.Name,
		Valor:      valor,
		CreateDate: c.CreateDate,
	}

}

func InsertCotacao(dados *DadosCotacao) error {

	ctxDB, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	var dtcotacao = time.Now()
	dt := dtcotacao.Format(time.RFC3339)
	dtcotacao, err := time.Parse("2006-01-02 15:04:05", dados.CreateDate)
	if err != nil {
		errRetorno := fmt.Sprintf("Falha na Conversão da Data: %s", err.Error())
		fmt.Println(errRetorno)
		return errors.New(errRetorno)
	}

	fmt.Println("Insert Cotacao: ", dados, dt)

	stmt, err := db.Prepare("insert into cotacoes (code, codein, name, valor, create_time) values (?, ?, ?, ?, ?)")
	if err != nil {
		errRetorno := fmt.Sprintf("Falha na preparação do Insert: %s", err.Error())
		fmt.Println(errRetorno)
		return errors.New(errRetorno)
	}
	defer stmt.Close()

	//_, err = stmt.Exec(dados.Code, dados.Codein, dados.Name, dados.Valor, dtcotacao)
	_, err = stmt.ExecContext(ctxDB, dados.Code, dados.Codein, dados.Name, dados.Valor, dtcotacao)
	if err != nil {
		errRetorno := fmt.Sprintf("Falha no Insert: %s", err.Error())
		fmt.Println(errRetorno)
		return errors.New(errRetorno)
	}

	fmt.Println("Insert com Sucesso.")

	return nil
}
