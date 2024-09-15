package main

func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

func NewBill() map[string]int {
	return make(map[string]int)
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	currentUnit, ok := units[unit]
	if !ok {
		return false
	}

	currentItemBill, ok := bill[item]
	if !ok {
		bill[item] = currentUnit
		return true
	}
	bill[item] = currentItemBill + currentUnit
	return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentUnit, existsUnit := units[unit]
	currentItemBill, existsBill := bill[item]
	if !existsUnit || !existsBill {
		return false
	}

	newQuantity := currentItemBill - currentUnit
	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = newQuantity
	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	currentItemBill, ok := bill[item]
	if !ok {
		return 0, false
	}
	return currentItemBill, true
}
