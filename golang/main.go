package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Registrar o tempo inicial
	inicio := time.Now()

	// Estabelecer conexão com o banco de dados MySQL
	db, err := sql.Open("mysql", "root:usbw@tcp(localhost:3307)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela se não existir
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cnpjs(id INT AUTO_INCREMENT PRIMARY KEY, cnpj VARCHAR(20)) ENGINE=MyISAM")
	if err != nil {
		log.Fatal(err)
	}

	// Definir o caminho da pasta que deseja percorrer
	pasta := "C:\\Users\\jphal\\Desktop\\Teste\\folder"

	// Listar os arquivos na pasta especificada
	arquivosNaPasta, err := os.ReadDir(pasta)
	if err != nil {
		log.Fatal(err)
	}

	valor := 0
	// Iterar sobre cada arquivo na lista
	for _, arquivo := range arquivosNaPasta {
		// Fazer o que desejar com cada arquivo
		caminhoCompleto := filepath.Join(pasta, arquivo.Name())
		fmt.Println(caminhoCompleto)

		// Abrir arquivo CSV
		arquivoCSV, err := os.Open(caminhoCompleto)
		if err != nil {
			log.Fatal(err)
		}
		defer arquivoCSV.Close()

		// Criar um objeto leitor de CSV com o delimitador correto
		leitorCSV := csv.NewReader(arquivoCSV)
		leitorCSV.Comma = ';'

		// Iterar sobre as linhas do arquivo
		for {
			linha, err := leitorCSV.Read()
			if err != nil {
				break // Chegou ao final do arquivo
			}

			// Acessar os valores das colunas desejadas
			// As colunas 0, 2 e 5 correspondem aos índices 0, 2 e 5 na lista
			coluna := linha[0] + linha[2] + linha[5]

			// Inserir dados no MySQL
			_, err = db.Exec("INSERT INTO cnpjs (cnpj) VALUES (?)", coluna)
			if err != nil {
				log.Fatal(err)
			}

			//fmt.Printf("\rEstamos na linha: %d de 4,494,861", valor)
			os.Stdout.Sync()

			valor++ // valor que irá parar em: 4,494,861
		}
	}

	// Registrar o tempo final
	fim := time.Now()
	fmt.Printf("\n Tempo total para processar os arquivos: %.2f minutos.\n", fim.Sub(inicio).Minutes())
	fmt.Println("Concluído!!!")
}
