# Cifra de César em Go

---
Este projeto é uma aplicação Go simples que demonstra a Cifra de César para criptografar e descriptografar arquivos de texto. Ele foi construído com uma clara separação de responsabilidades, tornando as operações criptográficas e o tratamento de arquivos distintos.
---

## Sobre a Cifra de César

A Cifra de César é um dos métodos de criptografia mais antigos e simples. É um tipo de cifra de substituição em que cada letra no texto original é substituída por uma letra que está um certo número de posições abaixo no alfabeto. Por exemplo, com um deslocamento (shift) de 3, 'A' se tornaria 'D', 'B' se tornaria 'E', e assim por diante.
---

## Como Usar

Para usar este projeto, você precisará ter o Go instalado em sua máquina.

### Clonar o Repositório
Primeiro, clone o repositório para o seu ambiente local:
```Bash
git clone (https://github.com/gabriel1003/pdf_crypto_caesar.git
cd pdf_crypto_caesar
```
---

### Construir o Executável (Build)
Depois de clonar, você pode construir o executável do seu projeto Go. Isso criará um arquivo binário que você pode executar diretamente.
```Bash
go build -o <nome_do_executavel> main.go
```
* `go build`: O comando para compilar seu código Go.
* `-o <nome_do_executavel>`: Define o nome do arquivo executável que será gerado. Substitua <nome_do_executavel> pelo nome que você preferir (por exemplo, caesar_cipher).
* `main.go`: O arquivo principal do seu projeto Go.

Após este comando, você terá um arquivo executável com o nome que você escolheu no diretório do seu projeto.
---

### Executando o Programa (Linux)
Uma vez que você construiu o executável, executar o programa é bem direto. Navegue até o diretório onde o executável foi salvo no seu terminal.

Observação Importante sobre a Configuração:
O valor do deslocamento (shift) para a Cifra de César e os caminhos dos arquivos de entrada/saída estão atualmente definidos diretamente no código-fonte. Para alterá-los, você precisará modificar o arquivo main.go e, em seguida, reconstruir o executável.
---

#### Criptografar um Arquivo
Para criptografar um arquivo, use a seguinte sintaxe:

```Bash
./<nome_do_executavel> encrypt
```
* `./<nome_do_executavel> encrypt`: Inicia o programa no modo de criptografia. Lembre-se de substituir <nome_do_executavel> pelo nome que você deu ao seu arquivo executável.

O programa usará o caminho do arquivo de entrada definido no código como fonte e o caminho do arquivo de saída também definido no código para o conteúdo criptografado. O valor de deslocamento definido no código será aplicado durante a criptografia.
---

#### Descriptografar um Arquivo
Para descriptografar um arquivo, use a seguinte sintaxe:

```Bash
./<nome_do_executavel> decrypt
```
* `./<nome_do_executavel> decrypt`: Inicia o programa no modo de descriptografia.

O programa usará o caminho do arquivo de entrada definido no código para a fonte criptografada e o caminho do arquivo de saída também definido no código para o conteúdo descriptografado. Ele usará o valor de deslocamento definido no código para a descriptografia. É crucial que este valor corresponda ao usado para a criptografia!
---

## Estrutura do Projeto
* **`main.go`**: Contém a lógica principal para orquestrar as operações, incluindo a chamada das funções de criptografia, descriptografia e busca de arquivos.
* **`encrypt.go`**: Abriga a lógica central para o processo de criptografia da Cifra de César.
* **`decrypt.go`**: Contém a lógica central para o processo de descriptografia da Cifra de César.
* **`filefinder.go`**: Inclui a lógica responsável por localizar e manipular os caminhos dos arquivos.

---
## Conciderações finais
Espero que gostem do projeto, é o primeiro que fiz em Go.
