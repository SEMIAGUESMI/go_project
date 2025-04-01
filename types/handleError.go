package types

func Calculate_revenu(quantita int, price_vente int, cost int) int {
	return cost - quantita*price_vente
}
