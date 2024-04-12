package handles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type RetCep struct {
	Origem      string `json:"origem"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilApiCep struct {
	Cep        string `json:"cep"`
	Uf         string `json:"state"`
	Localidade string `json:"city"`
	Bairro     string `json:"neighborhood"`
	Logradouro string `json:"street"`
	Servico    string `json:"service"`
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c1 := make(chan ViaCEP)
	c2 := make(chan BrasilApiCep)

	go func() {
		viacep, err := BuscarViaCep(cep)
		if err != nil {
			fmt.Println("Viacep - Cep Não Encontrado")
			return
		}
		c1 <- *viacep
	}()

	go func() {
		brasilcep, err := BuscarBrasilApiCep(cep)
		if err != nil || brasilcep.Cep == "" {
			fmt.Println("BrasilApi - Cep Não Encontrado")
			return
		}
		c2 <- *brasilcep
	}()

	var rCep RetCep

	sair := false

	for {
		select {
		case cepret := <-c1:
			fmt.Printf("Received from ViaCep - Cep: %s\n", cepret.Cep)
			rCep.Origem = "ViaCep"
			rCep.Cep = cepret.Cep
			rCep.Logradouro = cepret.Logradouro
			rCep.Complemento = cepret.Complemento
			rCep.Bairro = cepret.Bairro
			rCep.Localidade = cepret.Localidade

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(rCep)

			sair = true

		case brcepret := <-c2:
			fmt.Printf("Received from BrasilAPI - Cep: %s\n", brcepret.Cep)
			rCep.Origem = "BrasilApi"
			rCep.Cep = brcepret.Cep
			rCep.Logradouro = brcepret.Logradouro
			rCep.Complemento = ""
			rCep.Bairro = brcepret.Bairro
			rCep.Localidade = brcepret.Localidade

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(rCep)

			sair = true

		case <-time.After(time.Second * 1):
			println("Timeout")
			w.WriteHeader(http.StatusInternalServerError)
			sair = true
		}

		if sair {
			break
		}
	}

}

func BuscarViaCep(cep string) (*ViaCEP, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func BuscarBrasilApiCep(cep string) (*BrasilApiCep, error) {
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep + " + cep")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c BrasilApiCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
