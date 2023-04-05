package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibndo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheco este comando")
			os.Exit(-1)
		}

	}

	//if comando == 1 {
	//	fmt.Println("Monitorando")
	//} else if comando == 2 {
	//	fmt.Println("Exibndo Logs...")
	//} else if comando == 0 {
	//	fmt.Println("Saindo do programa")
	//} else {
	//	fmt.Println("Não conheco este comando")
	//}

}
func devolveNome() string {
	nome := "Italo"
	return nome
}
func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do progrmama")
}

func leComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}
func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://wwww.alura.com.br", "https://work-media.vercel.app/"}
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")

}
func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um ero", err)
	}
	if resp.StatusCode == 200 {
		registraLog(site, true)
		fmt.Println("O site ", site, "foi carregado com sucesso")
	} else {
		registraLog(site, false)
		fmt.Println("Site", site, "esta com problemas. StatusCode:", resp.StatusCode)
	}

}
func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um ero", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
		fmt.Println(linha)
	}
	arquivo.Close()

	return sites

}
func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "= online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}
func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))

}
