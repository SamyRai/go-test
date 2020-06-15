package Account

// General structure to store money related data
type Details struct {
	InitialSum   float64
	Transactions []float64
	Owner        string
}

// Debit account with overdraft
type Debit struct {
	Details   Details
	Overdraft float64
}

// Implementing a method for the Money interface
func (d Debit) YearLeftover() (result float64) {
	result = d.Details.InitialSum
	for _, transaction := range d.Details.Transactions {
		result = result + transaction
	}

	return result
}

// Credit account with interest
type Credit struct {
	Details            Details
	YearlyInterestRate float64
}

// Implementing a method for the Money interface
func (c Credit) YearLeftover() (result float64) {
	result = c.Details.InitialSum
	for _, transaction := range c.Details.Transactions {
		result = result + transaction
	}

	interest := result / 100 * c.YearlyInterestRate
	return result + interest
}
