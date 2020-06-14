package Account

type Details struct {
	InitialSum   float64
	Transactions []float64
	Owner        string
}

type Debit struct {
	Details   Details
	Overdraft float64
}

func (d Debit) YearLeftover() (result float64) {
	result = d.Details.InitialSum
	for _, transaction := range d.Details.Transactions {
		result = result + transaction
	}

	return result
}

type Credit struct {
	Details            Details
	YearlyInterestRate float64
}

func (c Credit) YearLeftover() (result float64) {
	result = c.Details.InitialSum
	for _, transaction := range c.Details.Transactions {
		result = result + transaction
	}

	interest := result / 100 * c.YearlyInterestRate
	return result + interest
}
