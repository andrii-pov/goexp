package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// encoding some basic types to json
	bol, _ := json.Marshal(true)
	fmt.Println(string(bol))

	intR, _ := json.Marshal(2)
	fmt.Println(string(intR))

	fltR, _ := json.Marshal(34.56)
	fmt.Println(string(fltR))

	strR, _ := json.Marshal("gopher")
	fmt.Println(string(strR))

	slc := []string{"some user", "another user", "yet another user"}
	slcR, _ := json.Marshal(slc)
	fmt.Println(string(slcR))

	mapStrInt := map[string]int{"apple": 5, "lettuce": 7}
	mapStrIntR, _ := json.Marshal(mapStrInt)
	fmt.Println(string(mapStrIntR))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", " peach", "pear"},
	}

	res1DR, _ := json.Marshal(res1D)
	fmt.Println(string(res1DR))

	res2D := &response2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2DR, _ := json.Marshal(res2D)
	fmt.Println(string(res2DR))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	fmt.Println(reflect.TypeOf(dat["num"]), dat["num"])

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)

	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	resR := response2{}
	json.Unmarshal([]byte(str), &resR)
	fmt.Println(resR)
	fmt.Println(resR.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
