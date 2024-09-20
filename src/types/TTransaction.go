package types

type TTransaction struct {
	Id          int    `csv:"ID"`
	Date        string `csv:"DATE"`
	Transaction string `csv:"TRANSACTION"`
}
