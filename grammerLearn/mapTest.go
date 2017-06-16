package grammerLearn

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	/* 创建集合*/
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"

	for country:= range countryCapitalMap {
		fmt.Println(countryCapitalMap[country])
	}

	capitial, ok := countryCapitalMap["China"]
	if(ok) {
		fmt.Println("Capitial of China is %s", capitial)
	}else {
		fmt.Println("Capitial of China is no present")
	}

	delete(countryCapitalMap, "France")
	for country:= range countryCapitalMap {
		fmt.Println(countryCapitalMap[country])
	}
}
