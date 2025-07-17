package util

// Constans for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSuppordtedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
