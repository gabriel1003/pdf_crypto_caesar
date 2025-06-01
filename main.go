package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// constant with path of files to encrypt
const TARGET_DIRECTORY = "/home/gsmdev/Documentos"

const CAESAR_SHIFT = 3 // kay to encrypt

func main() {

	fmt.Printf("* Programa de Criptografia/Descriptografia de PDFs em '%s' (Recursivo) *\n", TARGET_DIRECTORY)
	fmt.Println("* ATENÇÃO: Este programa modifica arquivos.                       *")

	if len(os.Args) < 2 {
		fmt.Printf("Uso: ./seu_programa <encrypt|decrypt>\n")
		fmt.Printf("Exemplo para criptografar: ./seu_programa encrypt\n")

		fmt.Printf("Exemplo para descriptografar: ./seu_programa decrypt\n")
		os.Exit(1)
	}

	action := strings.ToLower(os.Args[1])
	var filesToProcess []string
	var err error
	var operation func(string, int) error
	var targetSuffix string
	var successMessageAction string
	switch action {
	case "encrypt":
		fmt.Printf("Modo: CRIPTOGRAFAR arquivos PDF em %s e seus subdiretórios.\n", TARGET_DIRECTORY)
		targetSuffix = ".pdf"
		filesToProcess, err = FindFilesRecursively(TARGET_DIRECTORY, targetSuffix, true)
		operation = EncryptFile
		successMessageAction = "Criptografia"
	case "decrypt":
		fmt.Printf("Modo: DESCRIPTOGRAFAR arquivos .pdf.encrypted em %s e seus subdiretórios.\n", TARGET_DIRECTORY)
		targetSuffix = ".pdf.encrypted"
		filesToProcess, err = FindFilesRecursively(TARGET_DIRECTORY, targetSuffix, false)
		operation = DecryptFile
		successMessageAction = "Descriptografia"
	default:
		fmt.Fprintf(os.Stderr, "Ação desconhecida: %s\n", action)
		fmt.Printf("Uso: ./seu_programa <encrypt|decrypt>\n")
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao buscar arquivos: %v\n", err)
		os.Exit(1)
	}

	if len(filesToProcess) == 0 {
		fmt.Printf("Nenhum arquivo '%s' encontrado em '%s' (incluindo subdiretórios) para processar.\n", targetSuffix, TARGET_DIRECTORY)
		return
	}

	fmt.Printf("\nArquivos a serem processados (de '%s' e subdiretórios):\n", TARGET_DIRECTORY)
	for _, fp := range filesToProcess {
		relativePath, relErr := filepath.Rel(TARGET_DIRECTORY, fp)
		if relErr == nil {
			fmt.Printf("- %s\n", relativePath)
		} else {
			fmt.Printf("- %s\n", fp)
		}
	}

	fmt.Printf("\nEncontrados %d arquivos para %s. Deseja continuar? (s/N): ", len(filesToProcess), strings.ToLower(successMessageAction))
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) == "n" {
		fmt.Println("Operação cancelada pelo usuário.")
		return
	}
	fmt.Printf("\nIniciando %s...\n", successMessageAction)
	var wg sync.WaitGroup
	processedCount := 0
	errorCount := 0
	for _, filePath := range filesToProcess {
		wg.Add(1)
		go func(fp string) {
			defer wg.Done()
			errOp := operation(fp, CAESAR_SHIFT)
			baseName := filepath.Base(fp)
			if errOp != nil {
				fmt.Fprintf(os.Stderr, "Falha ao %s %s: %v\n", strings.ToLower(successMessageAction), baseName, errOp)
				errorCount++
			} else {
				if action == "encrypt" {
					fmt.Printf("%s: %s -> %s.encrypted\n", successMessageAction, baseName, baseName)
				} else {
					fmt.Printf("%s: %s -> %s\n", successMessageAction, baseName, filepath.Base(strings.TrimSuffix(fp, ".encrypted")))
				}
				processedCount++
			}
		}(filePath)
	}
	wg.Wait()
	fmt.Printf("\nProcesso de %s concluído.\n", successMessageAction)
	fmt.Printf("Arquivos processados com sucesso: %d\n", processedCount)
	fmt.Printf("Arquivos com erro: %d\n", errorCount)
	if errorCount > 0 {
		fmt.Println("Alguns arquivos encontraram erros durante o processamento. Verifique as mensagens acima.")
	}

	if processedCount > 0 {
		if action == "encrypt" {
			fmt.Println("\nLEMBRE-SE: Os arquivos PDF originais foram removidos após a criptografia.")
		} else {
			fmt.Println("\nLEMBRE-SE: Os arquivos '.encrypted' NÃO foram removidos após a descriptografia por segurança.")
		}
	}
}
