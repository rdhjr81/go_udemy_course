package arrays

import "fmt"

func HandsOnExercise40() {
	iceCreamFlavors := [...]string{
		"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	iceCreamFlavorsLen := len(iceCreamFlavors)
	fmt.Println("number of ice cream flavors:", iceCreamFlavorsLen)
	fmt.Println("ice cream flovors", iceCreamFlavors)
}
func HandsOnExercise41() {
	iceCreamFlavors := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println("number of ice cream flavors:", len(iceCreamFlavors))
	// fmt.Println("ice cream flovors", iceCreamFlavors)
	for index, flavor := range iceCreamFlavors {
		fmt.Println("flavor number: ", index, "flavor name: ", flavor)
	}
}

func Slicing() {
	iceCreamFlavors := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	printSlice(iceCreamFlavors)
	iceCreamFlavors2 := append(iceCreamFlavors, "Alfalfa Sprouts", "Old Bagel")
	printSlice(iceCreamFlavors2)
}

func SlicingASlice() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(si)
	//inclusive:exclusive
	fmt.Println("si[0:1]", si[0:1])
	fmt.Println("si[1:4]", si[1:4])
	//exclusive
	fmt.Println("si[:4]", si[:4])

	//inclusive
	fmt.Println("si[5:]", si[5:])

}

func DeletingASlice() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(si)
	si = append(si[:1], si[3:]...)
	fmt.Println(si)
}

func MakeASlice() {
	// ai := [3]int{1, 2, 3}
	si2 := make([]int, 0, 10)
	fmt.Println(si2)
	fmt.Println(len(si2))
	fmt.Println(cap(si2))

	si2 = append(si2, 1, 2, 3, 4)

	fmt.Println(si2)
	fmt.Println(len(si2))
	fmt.Println(cap(si2))

	fmt.Println(si2)
	fmt.Println(len(si2))
	fmt.Println(cap(si2))
}

func printSlice(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println("position: ", i, "value: ", s[i])
	}
	fmt.Println("--------------------")
}
