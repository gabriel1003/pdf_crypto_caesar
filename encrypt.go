package main

import (
	"fmt"
	"os"
)

func CaesarCipher(data byte, shift int) byte {
	return data + byte(shift)
}

func EncryptFile(filePath string, shift int) error {
	fmt.Printf("Criptografando: %s\n", filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo %s: %w", filePath, err)
	}

	encryptedContent := make([]byte, len(content))
	for i, b := range content {
		encryptedContent[i] = CaesarCipher(b, shift)
	}

	encryptedFilePath := filePath + ".encrypted"
	err = os.WriteFile(encryptedFilePath, encryptedContent, 0644)

	if err != nil {
		return fmt.Errorf("erro ao escrever o arquivo criptografado %s: %w", encryptedFilePath, err)
	}

	if errRemove := os.Remove(filePath); errRemove != nil {
		return fmt.Errorf("arquivo criptografado com sucesso, mas falha ao remover original %s: %w", filePath, errRemove)
	}
	return nil
}
