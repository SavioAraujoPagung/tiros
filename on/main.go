package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const diretorio = "./dados.json"

type Usuario struct {
	Usuarios []registro `json:"usuarios"`
}

type registro struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Tempo    string `json:"tempo"`
	TempoStr string `json:"tempoStr"`
	Cidade   string `json:"cidade"`
	Telefone string `json:"telefone"`
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/salvar", handlerSave)
	http.HandleFunc("/deletar", handlerDelete)

	fmt.Println("Servico em execucao: http://localhost:8080")
	// Iniciar o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

func handlerSave(w http.ResponseWriter, r *http.Request) {
	var (
		v        = r.URL.Query()
		nome     = v.Get("nome")
		tempo    = v.Get("tempo")
		telefone = v.Get("telefone")
		cidade   = v.Get("cidade")
		registro = registro{
			Nome:     nome,
			Tempo:    tempo,
			Telefone: telefone,
			Cidade:   cidade,
		}
	)

	registros := contatenar(registro)
	salvar(registros)
}

func handlerDelete(w http.ResponseWriter, r *http.Request) {
	var (
		v     = r.URL.Query()
		id, _ = strconv.Atoi(v.Get("id"))
	)

	deletar(id)
}

func deletar(id int) {
	var (
		novosregistros = []registro{}
		registros      = obterRegistros()
		tam            = len(registros)
		found          = false
	)

	for i := 0; i < tam; i++ {
		if !found {
			if registros[i].Id == id {
				found = true
				continue
			}
		}

		novosregistros = append(novosregistros, registros[i])
	}

	salvar(novosregistros)
}

func obterRegistros() []registro {
	file, _ := os.ReadFile(diretorio)
	u := Usuario{}

	json.Unmarshal(file, &u)

	return u.Usuarios
}

func contatenar(r registro) []registro {
	var (
		registros      = obterRegistros()
		novosregistros = []registro{}
		found          = false
		tam            = len(registros)
	)

	r.Id = tam + 1
	rTempo, rTempoStr := tempo(r.Tempo)
	r.TempoStr = rTempoStr

	if tam == 0 {
		return append(novosregistros, r)
	}

	for i := 0; i < tam; i++ {
		if !found {
			auxTempo, _ := tempo(registros[i].Tempo)
			if rTempo < auxTempo {
				novosregistros = append(novosregistros, r)
				novosregistros = append(novosregistros, registros[i])
				found = true
				continue
			}
			fmt.Println(novosregistros)
		}

		novosregistros = append(novosregistros, registros[i])
	}

	if !found {
		novosregistros = append(novosregistros, r)
	}
	
	return novosregistros
}

func tempo(tempo string) (float64, string) {
	tempos := strings.Split(tempo, ":")
	tam := len(tempos)
	segundos := 0.0

	if tam == 1 {
		segundos, _ = strconv.ParseFloat(tempos[0], 64)
		return segundos, fmt.Sprintf("%dm%.2fs", 0, segundos)
	}

	minuto, _ := strconv.ParseFloat(tempos[0], 64)
	segundos, _ = strconv.ParseFloat(tempos[1], 64)

	return segundos + (minuto * 60), fmt.Sprintf("%.0fm%.2fs", minuto, segundos)
}

func salvar(rs []registro) {
	u := Usuario{Usuarios: rs}

	json, _ := json.Marshal(u)

	file, _ := os.Create(diretorio)
	file.Write(json)

	file.Close()
}
