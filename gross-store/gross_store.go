package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMap := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	valueUnit, existsUnit := units[unit]
	valueItem, existsItem := bill[item]
	newQuantities := valueItem + valueUnit
	if !existsUnit {
		return false
	} else if !existsItem {
		bill[item] = valueUnit
		return true
	} else {
		bill[item] = newQuantities
		return true
	}

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	valueBill, existBill := bill[item]
	valueUnit, existUnit := units[unit]
	newQuantities := valueBill - valueUnit
	if !existBill || !existUnit || newQuantities < 0 {
		return false
	} else if newQuantities == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = newQuantities
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exist := bill[item]

	if !exist {
		return 0, false
	} else {
		return value, true
	}
}
