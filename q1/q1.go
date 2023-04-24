package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	var somaHistorico float64
	for i := 0; i < len(purchaseHistory); i++ {
		somaHistorico += purchaseHistory[i]
	}

	var disconto float64
	disconto = 0
	switch {
	case currentPurchase <= float64(0):
		return 0, fmt.Errorf("Valor da compra inválido.")
	case somaHistorico/float64(len(purchaseHistory)) > float64(1000):
		disconto = currentPurchase * 0.2
	case somaHistorico > float64(1000):
		disconto = currentPurchase * 0.1
	case somaHistorico <= float64(1000) && somaHistorico >= 501:
		disconto = currentPurchase * 0.05
	case somaHistorico <= 500:
		disconto = currentPurchase * 0.02
	case somaHistorico == float64(0):
		disconto = currentPurchase * 0.2
	}
	return disconto, nil
}
//func main() {
//	numero := 1500.0
//	slice := []float64{800, 1200, 700, 1500}
//	resultado, err := CalculateDiscount(numero, slice)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(resultado)
//	}
//}
