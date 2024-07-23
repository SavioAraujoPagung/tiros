package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

type logger string

func (r registro) toString() string {
	return fmt.Sprintf("[Id: %d, Nome: %s, Tempo: %s, Cidade: %s, Telefone: %s]", r.Id, r.Nome, r.Tempo, r.Cidade, r.Telefone)
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
			Nome:     strings.ToUpper(nome),
			Tempo:    tempo,
			Telefone: telefone,
			Cidade:   strings.ToUpper(cidade),
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
				log(fmt.Sprintf("- deletado: %s, horario: %v", registros[i].toString(), time.Now()))
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

func getID(registros []registro) int {
	id := 0

	for i, r := range registros {
		if r.Id > id {
			id = registros[i].Id
		}
	}

	return id + 1
}

func contatenar(r registro) []registro {
	var (
		registros      = obterRegistros()
		novosregistros = []registro{}
		found          = false
		tam            = len(registros)
	)

	r.Id = getID(registros)
	rTempo, rTempoStr := tempo(r.Tempo)
	r.TempoStr = rTempoStr

	if tam <= 10 {
		novosregistros = append(novosregistros, r)
		return append(novosregistros, registros...)
	}

	for i := 0; i < tam; i++ {
		if !found {
			auxTempo, _ := tempo(registros[i].Tempo)
			if rTempo < auxTempo || auxTempo == -1 {
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
	log(fmt.Sprintf("+ salvo: %s, horario: %v", r.toString(), time.Now()))
	return novosregistros
}

func tempo(tempo string) (float64, string) {
	if tempo == "" {
		return -1, ""
	}

	tempos := strings.Split(tempo, ":")
	tam := len(tempos)
	segundos := 0.0

	if tam == 1 {
		segundos, _ = strconv.ParseFloat(tempos[0], 64)
		if segundos < 60 {
			return segundos, fmt.Sprintf("%.2fs", segundos)
		}

		min := segundos / 60

		resto := int(segundos) % 60
		parteDecimal := segundos - math.Floor(segundos)

		seg := float64(resto) + parteDecimal

		return segundos, fmt.Sprintf("%.0fm%.2fs", min, seg)
	}

	minuto, _ := strconv.ParseFloat(tempos[0], 64)
	segundos, _ = strconv.ParseFloat(tempos[1], 64)

	return segundos + (minuto * 60), fmt.Sprintf("%.0fm%.2fs", minuto, segundos)
}

func salvar(rs []registro) {
	u := Usuario{Usuarios: rs}

	json, _ := json.Marshal(u)

	file, _ := os.Create(diretorio)
	defer file.Close()

	file.Write(json)
}

func log(msg string) {
	fmt.Println(msg)
	file, _ := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	ll := ""

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Imprimir cada linha do arquivo
		ll = fmt.Sprintf("%s", scanner.Text())
	}

	// Verificar se houve algum erro durante a leitura
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura do arquivo:", err)
	}

	b := []byte(fmt.Sprintf("%s\n%s", ll, msg))
	_, err := file.Write(b) 
	if err != nil {
		fmt.Println(err)
	}

	file.Close()
}
