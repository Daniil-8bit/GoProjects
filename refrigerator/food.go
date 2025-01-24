package refrigerator

import "fmt"

type Refrigerator []string

func find(product string, amount []string) bool {
	for _, value := range amount {
		if value == product {
			return true
		}
	}
	return false
}

func (r Refrigerator) openRefrigerator() {
	fmt.Println("Refrigerator open!")
}

func (r Refrigerator) closeRefrigerator() {
	fmt.Println("Refrigerator closed!")
}

func (r Refrigerator) FindProduct(product string) error {
	r.openRefrigerator()
	defer r.closeRefrigerator()
	if find(product, r) {
		fmt.Println(product, "found!")
	} else {
		return fmt.Errorf(product, "wasn`t found!")
	}
	return nil
}
