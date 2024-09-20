package types

type TClient struct {
	Id           int
	Email        string
	Name         string
	Transactions []TMovements
	GroupedBy    []TSummary
	TotalBalance float64
	AvgDebit     float64
	AvgCredit    float64
}
