package main

import (
	"bufio"
	"dbuilder/newf"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var opcoes, addDir string

	new := newf.New()

	fmt.Println("====================== BEM-VINDO AO DIR BUILDER =======================")
	// Define name Project
	nomeProjeto := input("INFORME O NOME DO PROJETO > ")
	fmt.Println(new.NameProject(nomeProjeto))
	// Define Quantity modules
	fmt.Println("=============================== MÓDULOS ===============================")
	quantidade := input("INFORME A QUANTIDADE DE MÓDULOS > ")
	fmt.Println(new.QuantityModules(quantidade))
	// Define Dir
	for {
		fmt.Println("============================== DIRETÓRIO ==============================")
		opcoes := input("[A] - ADICIONAR NOVA PASTA\n[D] - DELETAR PASTA PADRÃO\n[ ] - VAZIO PARA SAIR\n> ")

		switch opcoes {
		case "A", "C", "ADD", "CRIAR", "ADCIONAR":
			new.ShowDirs()
			opcao_dir := input("ADD: INFORME O NOME DA(s) NOVA(s) PASTA(s) SEPARADAS POR VIRGULAS (,)\n> ")
			new.AddNewDirs(opcao_dir)
			new.ShowDirs()
		case "D", "R", "E", "EXCLUIR", "DELETAR", "REMOVER":
			new.ShowDirs()
			opcao_dir := input("DEL: INFORME O NOME DA(s) PASTAS A SER DELETADAS SEPARADAS POR VIRGULA (,)\n>")
			new.RemoveDirs(opcao_dir)
			new.ShowDirs()
		case "":
			fmt.Println(new.CreateDirFiles())
			return
		default:
			fmt.Println("# OPÇÃO INVÁLIDA #")
			continue
		}
	}

}
func input(txt string) string {
	ler := bufio.NewReader(os.Stdin)
	fmt.Print(txt)
	text, _ := ler.ReadString('\n')
	return strings.TrimSpace(strings.ToUpper(text))
}