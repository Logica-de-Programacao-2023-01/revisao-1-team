package q3

import (
	"fmt"
	"math"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	maximo := math.Inf(-1) // Define o máximo como o menor número possível
	minimo := math.Inf(1)  // Define o mínimo como o maior número possível
	var soma, media, denominador float64
	if len(numbers) <= 0 {
		return 0, 0, 0, fmt.Errorf("Lista vazia.") // erro caso a lista seja igual a 0
	}
	for i := 0; i < len(numbers); i++ {
		soma += float64(numbers[i]) //soma de todos os numeros
		denominador++               //soma dos denominadores ate o len(numbers)
		valorMax := numbers[i]
		valorMin := numbers[i]
		if float64(valorMax) > maximo {
			maximo = float64(valorMax)
		}
		if float64(valorMin) < minimo {
			minimo = float64(valorMin)
		}
	}
	media = (soma - float64(minimo) - float64(maximo)/(denominador-2))
	return int(maximo), int(minimo), media, nil
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print(FindMinMaxAverage(slice))

}

//ouuuu
//package q3
//
//import "fmt"
//
//func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
//
//	maior := -9999999999
//	menor := 99999999999
//	var soma float64
//	var contagem float64
//	var media float64
//	if len(numbers) == 0 {
//		return 0, 0, 0, fmt.Errorf("Lista vazia")
//	}
//	for i := 0; i < len(numbers); i++ {
//		soma += float64(numbers[i])
//		contagem++
//		if numbers[i] > maior {
//			maior = numbers[i]
//		}
//		if numbers[i] < menor {
//			menor = numbers[i]
//		}
//	}
//	media = (soma - float64(menor) - float64(maior)) / (contagem - 2)
//	return maior, menor, media, nil
//}
