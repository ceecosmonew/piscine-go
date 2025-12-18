package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	// Menu using struct food
	menu := map[string]food{
		"burger":  {15},
		"chips":   {10},
		"nuggets": {12},
	}

	// Check if item exists
	if item, ok := menu[order]; ok {
		return item.preptime
	}

	// If item doesn't exist → return 404
	return 404
}
