package demo_04

import "fmt"

func TestMap() {

	countryCapitalMap := make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	/*查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(captial) */
	/*fmt.Println(ok) */
	if (ok) {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}
}

func TestTwoMap() {
	a := make(map[string]map[string]string)
	a["user1"] = make(map[string]string)
	a["user1"]["name"] = "dongdong"
	a["user1"]["language"] = "golang"
	a["user1"]["age"] = "26"
	fmt.Println("before: ", a)
	a["user1"]["age"] = "18"
	fmt.Println("after: ", a)
}

func TestMapping() {
	var a []map[int]int
	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 190
	fmt.Println(a)

}
