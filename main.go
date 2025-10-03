package main

import "fmt"

func main() {
	numeros := []int{}
	var err error

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1 - Adicionar número")
		fmt.Println("2 - Listar números")
		fmt.Println("3 - Remover por índice")
		fmt.Println("4 - Estatísticas: Mínimo, Máximo, Média")
		fmt.Println("5 - Divisão")
		fmt.Println("6 - Limpar lista")
		fmt.Println("0 - Sair")
		fmt.Println()

		var opcao int
		fmt.Print("Escolha uma opção: ")
		if _, scanErr := fmt.Scan(&opcao); scanErr != nil {
			fmt.Println("Erro ao ler a opção.")
			continue
		}
		switch opcao {

		case 1:
			fmt.Print("Digite um número para adicionar: ")
			var valor int
			if _, scanErr := fmt.Scan(&valor); scanErr != nil {
				fmt.Println("Erro ao ler o número.")
				continue
			}
			numeros, err = adicionar(numeros, valor)
			if err != nil {
				fmt.Println("Erro:", err)
			}else {
				fmt.Println("Número adicionado.")
			}
			
		
		case 2:
			listar(numeros)
		
		case 3:
			fmt.Print("Digite o índice do número para remover: ")
			var indice int
			if _, scanErr := fmt.Scan(&indice); scanErr != nil{
				fmt.Println("Digite um indice válido")
				continue
			}

			numeros, err = remover(numeros, indice)
			if err != nil {
				fmt.Println("Erro:", err)
			}else {
				fmt.Println("Número removido.")
			}

		case 4:
			min, max, media, err := estatisticas(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Resultado das estatísticas: Mínimo: %d, Máximo: %d, Média: %.2f\n", min, max, media)
			}

		case 5:
			var a float64
			var b float64
			fmt.Print("Digite o dividendo: ")
			if _, scanErr := fmt.Scan(&a); scanErr != nil {
				fmt.Println("Digite um número válido.")
				continue
			}
			fmt.Print("Digite o divisor: ")
			if _, scanErr := fmt.Scan(&b); scanErr != nil {
				fmt.Println("Digite um número válido.")
				continue
			}
			result, err := divisao(a, b)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
       		 	fmt.Printf("Resultado da divisão: %.2f\n", result) 
			}

		case 6:
			numeros = limpar(numeros)
			fmt.Println("Lista limpa.")

		case 0:
			fmt.Println("Saindo...")
			return
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

func estatisticas(numeros []int) (int, int, float64, error) {
	if len(numeros) == 0 {
		return 0, 0, 0, fmt.Errorf("A lista está vazia.")
	}

	min := numeros[0]
	max := numeros[0]
	count := 0
	sum := 0

	for i := 0; i < len(numeros); i++ {
		count++
		sum += numeros[i]
		if numeros[i] < min {
			min = numeros[i]
		}
		if numeros[i] > max {
			max = numeros[i]
		}
	}
	media := float64(sum) / float64(count)
	return min, max, media, nil
}

func divisao(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divisão por zero")
	}
	return float64(a) / float64(b), nil
}

func limpar(numeros []int) []int {
	numeros = []int{}
	return numeros
}