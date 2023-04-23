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
	case somaHistorico > float64(1000):
		disconto = somaHistorico * 0.1
	case somaHistorico <= float64(1000) && somaHistorico >= 501:
		disconto = somaHistorico * 0.05
	case somaHistorico <= 500:
		disconto = somaHistorico * 0.02
	case somaHistorico == float64(0):
		disconto = somaHistorico * 0.2
	case somaHistorico/float64(len(purchaseHistory)) > float64(1000):
		disconto = somaHistorico * 0.2
	}
	return disconto, nil
}
