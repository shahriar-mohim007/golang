package main
import (
	"fmt",
	"sort"
)
func main(){
	//different way to declare map
	ages := make(map[string]int)
	ages := map[string]int{
		"alice":31,
		"charlie": 34,
	}
	ages := make(map[string]int)
    ages["alice"] = 31
    ages["charlie"] = 34
	delete(ages, "alice")
	ages["bob"] += 1
	ages["bob"]++
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
		}
	// To enumerate the key/value pairs in order, we must sort
    //the keys explicitly, for instance, using the Strings function from the sort package if the keys
    //are strings.
	var names []string
	for name := range ages {
	names = append(names, name)
	}
    sort.Strings(names)
	for _, name := range names {
	fmt.Printf("%s\t%d\n", name, ages[name])
	}
	//Using Maps as Sets
	intSet := map[int]bool{}
    vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
	intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
	fmt.Println("100 is in the set")
	}
}