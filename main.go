package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Usuario struct {
	Usuarios []registro `json:"usuarios"`
}

type registro struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Tempo  string `json:"tempo"`
	Active bool   `json:"active"`
}

func main() {
	http.HandleFunc("/save", handlerSave)

	http.HandleFunc("/delete", handlerDelete)

	http.ListenAndServe(":1414", nil)
}

func handlerSave(w http.ResponseWriter, r *http.Request) {
	var (
		v        = r.URL.Query()
		nome     = v.Get("nome")
		tempo    = v.Get("tempo")
		registro = registro{
			Nome: nome,
			Tempo: tempo,
			Active: true,
		}
	)

	registros := contatenar(registro)
	fmt.Println(registros)
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
				registros[i].Active = false
				novosregistros = append(novosregistros, registros[i])
				found = true
				continue
			}
		}

		novosregistros = append(novosregistros, registros[i])
	}

	salvar(novosregistros)
}

func obterRegistros() []registro {
	file, _ := os.ReadFile("./dados.json")

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

	if tam == 0 {
		return append(novosregistros, r)
	}

	for i := 0; i < tam; i++ {
		if !found {
			if tempo(r.Tempo) < tempo(registros[i].Tempo) {
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

	fmt.Println(novosregistros)

	return novosregistros
}

func tempo(tempo string) int {
	tempos := strings.Split(tempo, ":")

	minuto, _ := strconv.Atoi(tempos[0])
	segundos, _ := strconv.Atoi(tempos[1])

	return segundos + (minuto * 60)
}

func salvar(rs []registro) {
	u := Usuario{Usuarios: rs}

	json, _ := json.Marshal(u)

	file, _ := os.Create("./dados.json")
	file.Write(json)
}
