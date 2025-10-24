package decorator

import (
	"fmt"
	"testing"
)

func TestCreateCoffee(t *testing.T) {
	commonCoffee := NewCommonCoffee()
	fmt.Println(commonCoffee.GetDescription())
	fmt.Println("----------")

	milkCoffee := NewMilkCoffee(commonCoffee)
	milkCoffee.GetDescription()
	fmt.Println(milkCoffee.GetDescription())
	fmt.Println("----------")

	sugarCoffee := NewSugarCoffee(commonCoffee)
	sugarCoffee.GetDescription()
	fmt.Println(sugarCoffee.GetDescription())
}
