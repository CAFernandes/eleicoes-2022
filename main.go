package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Candidato struct {
	Seq string `json:"seq"`
	Sqcand string `json:"sqcand"`
	N string `json:"n"`
	Nm string `json:"nm"`
	Cc string `json:"cc"`
	Nv string `json:"nv"`
	E string `json:"e"`
	St string `json:"st"`
	Dvt string `json:"dvt"`
	Vap string `json:"vap"`
	Pvap string `json:"pvap"`
}

type Eleicoes struct {
	Ele string `json:"ele"`
	Tpabr string `json:"tpabr"`
	Cdabr string `json:"cdabr"`
	Carper string `json:"carper"`
	Md string `json:"md"`
	T string `json:"t"`
	F string `json:"f"`
	Dg string `json:"dg"`
	Hg string `json:"hg"`
	Dt string `json:"dt"`
	Ht string `json:"ht"`
	Dv string `json:"dv"`
	Tf string `json:"tf"`
	V string `json:"v"`
	Esae string `json:"esae"`
	Mnae string `json:"mnae"`
	S string `json:"s"`
	St string `json:"st"`
	Pst string `json:"pst"`
	Snt string `json:"snt"`
	Psnt string `json:"psnt"`
	Si string `json:"si"`
	Psi string `json:"psi"`
	Sni string `json:"sni"`
	Psni string `json:"psni"`
	Sa string `json:"sa"`
	Psa string `json:"psa"`
	Sna string `json:"sna"`
	Psna string `json:"psna"`
	E string `json:"e"`
	Ea string `json:"ea"`
	Pea string `json:"pea"`
	Ena string `json:"ena"`
	Pena string `json:"pena"`
	Esi string `json:"esi"`
	Pesi string `json:"pesi"`
	Esni string `json:"esni"`
	Pesni string `json:"pesni"`
	C string `json:"c"`
	Pc string `json:"pc"`
	A string `json:"a"`
	Pa string `json:"pa"`
	Vscv string `json:"vscv"`
	Vnom string `json:"vnom"`
	Pvnom string `json:"pvnom"`
	Vvc string `json:"vvc"`
	Pvvc string `json:"pvvc"`
	Vb string `json:"vb"`
	Pvb string `json:"pvb"`
	Tvn string `json:"tvn"`
	Ptvn string `json:"ptvn"`
	Vn string `json:"vn"`
	Pvn string `json:"pvn"`
	Vnt string `json:"vnt"`
	Pvnt string `json:"pvnt"`
	Vp string `json:"vp"`
	Pvp string `json:"pvp"`
	Vv string `json:"vv"`
	Pvv string `json:"pvv"`
	Van string `json:"van"`
	Pvan string `json:"pvan"`
	Vansj string `json:"vansj"`
	Pvansj string `json:"pvansj"`
	Tv string `json:"tv"`
	Cand []Candidato `json:"cand"`
}

func main() {
	url := "https://resultados.tse.jus.br/oficial/ele2022/544/dados-simplificados/br/br-c0001-e000544-r.json"
	body := getBody(url)
	printEleicao(body, 1)

	url = "https://resultados.tse.jus.br/oficial/ele2022/545/dados-simplificados/br/br-c0001-e000545-r.json"
	body = getBody(url)
	printEleicao(body, 2)
}

func getBody(url string) []byte {
	// fmt.Println("URL:>", url)

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func printEleicao(body []byte, turno int) {
	eleicao := Eleicoes{}
	jsonErr := json.Unmarshal(body, &eleicao)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println()
	fmt.Printf("\t\t\t\t%dº Turno", turno)
	fmt.Println()
	fmt.Printf("\t%-20s %-14s %-6s\n", "Nº Votos Válidos", "Nº Votos Nulos", "Nº Votos Brancos")
	fmt.Printf("\t%-20s %14s %16s \n", eleicao.Vnom, eleicao.Vn, eleicao.Vb)
	fmt.Println()

	fmt.Printf("\t%-20s %-10s %-10s\n", "Candidato", "Nº Votos", "Porcentagem")
	for _, cand := range eleicao.Cand {
		fmt.Printf("\t%-20s %8s %12s%s\n", cand.Nm, cand.Vap, cand.Pvap, "%")
	}
	fmt.Println()
}
