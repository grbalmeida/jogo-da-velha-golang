package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var jogadas map[string]string
var possibilidades map[int][]string

func main() {

	jogadas = map[string]string{
		"A1": " ",
		"A2": " ",
		"A3": " ",
		"B1": " ",
		"B2": " ",
		"B3": " ",
		"C1": " ",
		"C2": " ",
		"C3": " ",
	}

	possibilidades = make(map[int][]string)

	possibilidades[0] = []string{"A1", "A2", "B2"}
	possibilidades[1] = []string{"B1", "B2", "B3"}
	possibilidades[2] = []string{"C1", "C2", "C3"}
	possibilidades[3] = []string{"A1", "B1", "C1"}
	possibilidades[4] = []string{"A2", "B2", "C2"}
	possibilidades[5] = []string{"A3", "B3", "C3"}
	possibilidades[6] = []string{"A1", "B2", "C3"}
	possibilidades[7] = []string{"C1", "B2", "A3"}

	jogador := "O"
	ganhador := ""

	for {
		if numeroJogadas() != 9 {

		START:
			limparTela()

			fmt.Println(display())

			fmt.Println(fmt.Sprintf("Agora é a vez do jogador %s", jogador))
			var posicao string
			fmt.Scanf("%s", &posicao)
			posicao = strings.TrimSpace(posicao)

			jogada, ok := jogadas[posicao]

			if !ok {
				goto START
			}

			if strings.TrimSpace(jogada) != "" {
				fmt.Println("Jogada já executada")
				goto START
			}

			jogadas[posicao] = jogador

			if jogador == "O" {
				jogador = "X"
			} else {
				jogador = "O"
			}

			fmt.Println(display())

			ganhador = verificaSeTemVencedor()

			if ganhador != "" {
				fmt.Println(fmt.Sprintf("Parabéns %s. Você ganhou o jogo!", ganhador))
				break
			}
		} else {
			break
		}
	}

	if ganhador == "" {
		fmt.Println("O jogo empatou!")
	}
}

func display() string {
	var displayFormat string

	displayFormat = `
%s | %s | %s
---------
%s | %s | %s
---------
%s | %s | %s
`

	return fmt.Sprintf(displayFormat, jogadas["A1"], jogadas["A2"], jogadas["A3"], jogadas["B1"],
		jogadas["B2"], jogadas["B3"], jogadas["C1"], jogadas["C2"], jogadas["C3"])
}

func tem3Ocorrencias(coluna1, coluna2, coluna3, valor string) bool {
	return jogadas[coluna1] == valor && jogadas[coluna2] == valor && jogadas[coluna3] == valor
}

func verificaSeTemVencedor() string {
	// A1 == A2 == A3
	// B1 == B2 == B3
	// C1 == C2 == C3
	// A1 == B1 == C1
	// A2 == B2 == C2
	// A3 == B3 == C3
	// A1 == B2 == C3
	// C1 == B2 == A3

	ganhador := ""

	for _, possibilidade := range possibilidades {
		if tem3Ocorrencias(possibilidade[0], possibilidade[1], possibilidade[2], "O") {
			ganhador = "O"
		}

		if tem3Ocorrencias(possibilidade[0], possibilidade[1], possibilidade[2], "X") {
			ganhador = "X"
		}
	}

	return ganhador
}

func numeroJogadas() int {
	contador := 0

	for _, jogada := range jogadas {
		if strings.TrimSpace(jogada) != "" {
			contador++
		}
	}

	return contador
}

func limparTela() {
	time.Sleep(1 * time.Second)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
