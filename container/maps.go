package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "parker",
		"course":  "golang",
		"site":    "mok",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("...Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("...Getting values")
	couseName, ok := m["course"]
	fmt.Println(couseName, ok)
	if causeName, ok := m["case"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("Key does not exists")
	}
	fmt.Println("...Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name1, ok := m["name"]
	fmt.Println(name1, ok)

	testMap()

}

func testMap() {
	var maps map[string]int
	maps = make(map[string]int, 10)

	maps["admin"] = 10

	//fmt.Println(maps)

	var mapsmap = map[string](map[string]int)

	mapsmap = make(map[string]map[string]int)


	mapsmap["key1"] =maps["admin"]

	fmt.Println(mapsmap)
}
