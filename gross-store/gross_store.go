package gross

const QrtDozen = "quarter_of_a_dozen"
const HlfDozen = "half_of_a_dozen"
const Dozen = "dozen"
const SmGross = "small_gross"
const Gross = "gross"
const GrGross = "great_gross"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		QrtDozen: 3,
		HlfDozen: 6,
		Dozen:    12,
		SmGross:  120,
		Gross:    144,
		GrGross:  1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if u, ok := units[unit]; ok {
		bill[item] += u
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := bill[item]; !ok {
		return false
	}
	if u, ok := units[unit]; ok {
		if bill[item]-u < 0 {
			return false
		}
		if bill[item]-u == 0 {
			delete(bill, item)
			return true
		}

		bill[item] -= u
		return true
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if i, ok := bill[item]; ok {
		return i, true
	}
	return 0, false
}
