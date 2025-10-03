package main

import "fmt"

func main() {
	numeros := []int{1,2,3,4,5}
	var err error

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1 - Adicionar número")
		fmt.Println("2 - Listar números")
		fmt.Println("3 - Remover por índice")
		fmt.Println("4 - Estatísticas: Mínimo, Máximo, Média")
		fmt.Println("5 - Divisão por índice")
		fmt.Println("6 - Limpar lista")
		fmt.Println("0 - Sair")
		fmt.Println()

		var opcao int
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {

		case 1:
			fmt.Print("Digite um número para adicionar: ")
			var valor int
			fmt.Scan(&valor)
			numeros, err = adicionar(numeros, valor)
			if err != nil {
				fmt.Println("Erro:", err)
			}
			
		
		case 2:
			listar(numeros)
		
		case 3:
			fmt.Print("Digite o índice do número para remover: ")
			var indice int
			fmt.Scan(&indice)
			numeros, err = remover(numeros, indice)
			if err != nil {
				fmt.Println("Erro:", err)
			}

		case 4:
			estatisticas(numeros)

		case 5:
			var a int
			var b int
			fmt.Print("Digite o dividendo: ")
			fmt.Scan(&a)
			fmt.Print("Digite o divisor: ")
			fmt.Scan(&b)
			err = divisao(a, b)
			if err != nil {
				fmt.Println("Erro:", err)
			}

		case 6:
			numeros = limpar(numeros)
			fmt.Println("Lista limpa.")

		case 0:
			fmt.Println("Saindo...")
			return
		
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func adicionar(numeros []int, valor int) ([]int, error) {
	if valor < 0 {
		return numeros, fmt.Errorf("digite numeros maiores que zero")
	}
    numeros = append(numeros, valor)
	return numeros, nil
}

func listar(numeros []int) {
	fmt.Println("Números na lista:", numeros)
}

func remover(numeros []int, indice int) ([]int, error) {
	if indice < 0 || indice >= len(numeros){
		return numeros, fmt.Errorf("índice inválido")
	}
	numeros = append(numeros[:indice], numeros[indice+1:]...)
	return numeros, nil
}

func estatisticas(numeros []int) {
	if len(numeros) == 0 {
		fmt.Println("A lista está vazia.")
		return
	}

	min := numeros[0]
	max := numeros[len(numeros)-1]
	count := 0
	sum := 0

	for i := 0; i < len(numeros); i++ {
		count++
		sum += numeros[i]
	}
	media := float64(sum) / float64(count)
	fmt.Print("Estatísticas: ")
	fmt.Printf("Mínimo: %d, Máximo: %d, Média: %.2f\n", min, max, media)
}

func divisao(a int, b int) (error) {
	if b == 0 {
		return fmt.Errorf("divisão por zero")
	}
	divisao := a / b
	fmt.Print("Resultado da divisão: ", divisao, "\n")
	return nil
}

func limpar(numeros []int) []int {
	numeros = []int{}
	return numeros
}