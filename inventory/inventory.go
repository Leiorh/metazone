package inventory

var Stock = make(map[int]int)

func AddStock(productID int, quantity int) {
    Stock[productID] += quantity
}

func ReduceStock(productID int, quantity int) bool {
    if Stock[productID] >= quantity {
        Stock[productID] -= quantity
        return true
    }
    return false
}