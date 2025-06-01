package main

import (
	"fmt"
	"os"
	"strings"
)

func DecryptFile(filePath string, shift int) error {
	fmt.Printf("Descriptografando: %s\n", filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo %s: %w", filePath, err)
	}
	decryptedContent := make([]byte, len(content))
	for i, b := range content {
		decryptedContent[i] = CaesarCipher(b, -shift)
	}
	originalFilePath := strings.TrimSuffix(filePath, ".encrypted")
	err = os.WriteFile(originalFilePath, decryptedContent, 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever o arquivo descriptografado %s: %w", originalFilePath, err)
	}

	fmt.Printf("Descriptografado com sucesso: %s -> %s\n", filePath, originalFilePath)
	return nil
}
