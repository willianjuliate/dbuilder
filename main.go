package main

import (
	"bufio"
	"dbuilder/newf"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	new := newf.New()

	fmt.Println("====================== BEM-VINDO AO DIR BUILDER =======================")
	// Define o nome do projeto
	nomeProjeto := input("INFORME O NOME DO PROJETO > ")
	fmt.Println(new.NameProject(nomeProjeto))
	// Define quantidade módulos
	fmt.Println("=============================== MÓDULOS ===============================")
	quantidade := input("INFORME A QUANTIDADE DE MÓDULOS > ")
	fmt.Println(new.QuantityModules(quantidade))
	// Define Dir
	for {
		fmt.Println("============================== DIRETÓRIO ==============================")
		opcoes := input("[A] - ADICIONAR NOVA PASTA\n[D] - DELETAR PASTA PADRÃO\n[ ] - VAZIO PARA SAIR\n> ")

		switch opcoes {
		case "A", "C", "ADD", "CRIAR", "ADICIONAR":
			new.ShowDirs()
			opcao_dir := input(fmt.Sprintf("[%s]: INFORME O NOME DA(s) NOVA(s) PASTA(s) SEPARADAS POR VIRGULAS (,)\n> ", opcoes))
			new.AddNewDirs(opcao_dir)
			new.ShowDirs()
		case "D", "R", "E", "EXCLUIR", "DELETAR", "REMOVER":
			new.ShowDirs()
			opcao_dir := input(fmt.Sprintf("[%s]: INFORME O NOME DA(s) PASTAS A SER DELETADAS SEPARADAS POR VIRGULA (,)\n>", opcoes))
			new.RemoveDirs(opcao_dir)
			new.ShowDirs()
		case "":
			fmt.Println(new.CreateDirFiles())
			if runtime.GOOS != "linux" {
				time.Sleep(time.Second * 3)
			}
			return
		default:
			fmt.Println("# OPÇÃO INVÁLIDA #")
			continue
		}
	}

}
// Entrada padrão do teclado
func input(txt string) string {
	ler := bufio.NewReader(os.Stdin)
	fmt.Print(txt)
	text, _ := ler.ReadString('\n')
	return strings.TrimSpace(strings.ToUpper(text))
}
