package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFilesRecursively(targetDir string, suffix string, skipAlreadyProcessed bool) ([]string, error) {
	var files []string
	dirInfo, err := os.Stat(targetDir)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("o diretório alvo '%s' não existe", targetDir)
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar o diretório alvo '%s': %w", targetDir, err)
	}
	if !dirInfo.IsDir() {
		return nil, fmt.Errorf("o caminho alvo '%s' não é um diretório", targetDir)
	}

	fmt.Printf("Buscando arquivos recursivamente em: %s (sufixo: %s)\n", targetDir, suffix)
	err = filepath.Walk(targetDir, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			fmt.Fprintf(os.Stderr, "Aviso: Erro ao acessar %s: %v (ignorado)\n", path, walkErr)
			if info != nil && info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if info.IsDir() {
			return nil
		}
		lowerName := strings.ToLower(info.Name())
		lowerSuffix := strings.ToLower(suffix)
		if strings.HasSuffix(lowerName, lowerSuffix) {
			if skipAlreadyProcessed && lowerSuffix == ".pdf" && strings.HasSuffix(lowerName, ".pdf.encrypted") {
				fmt.Printf("Ignorando arquivo já criptografado: %s\n", path)
				return nil
			}

			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao caminhar pelo diretório '%s': %w", targetDir, err)
	}

	return files, nil
}
