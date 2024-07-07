package main

import (
	"compra-do-mes-curso-udemy-go/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua: "Alcides de Souza",
		Numero: 112,
		Cidade: "Valença",
	}

	mercado := model.Mercado{
		Endereco: endereco,
		Nome: "Guanabara",
	}

	item1 := model.Item{
		Nome: "Caixa de Leite",
		Quantidade: 1,
	}

	item2 := model.Item{
		Nome: "Pacote de Salaminho Fatiado",
		Quantidade: 2,
	}

	listaDeItens := []model.Item{
		item1, 
		item2,
	}

	listaJulho := model.ListaDeCompra{
		Data: time.Date(2024, 7, 10, 0, 0, 0, 0, time.Local),
		Lista: listaDeItens,
		Mercado: mercado,
	}

	fmt.Println(listaJulho)
}
