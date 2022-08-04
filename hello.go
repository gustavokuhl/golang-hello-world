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

const monitoramentos = 3
const delay = 5 * time.Second

func main() {
	showIntro()
	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			initMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Gustavo"
	var age = 21
	var version float32 = 1.1

	fmt.Println("Time to Hack The Planet mr", name, " you already are", age, "years old!")
	fmt.Println("Program version:", version)
	fmt.Println()
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var comando int
	fmt.Print("-> ")
	fmt.Scan(&comando)
	fmt.Println()
	return comando
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	fmt.Println()

	sites := getSitesFromFite()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay)
		fmt.Println()
	}
}

func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problema. Status code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func getSitesFromFite() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)
	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	fmt.Println("Exibindo logs...")
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(file))
}
