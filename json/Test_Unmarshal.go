package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

var data = `{"employees":{"first":{"firstName":"Bill","lastName":"Gates"},"second":{"firstName":"George","lastName":"Bush"},"third":{"firstName":"Thomas","lastName":"Carter"}}}`

func Parser(data []byte) interface{} {
	var i interface{}
	json.Unmarshal(data, &i)
	return i
}


func Test_Unmarshal() {
	da := Parser([]byte(data))
	fmt.Printf("%v\n", da)
}

func Test_UnmarshalV(t *testing.T) {
	da := Parser([]byte(data))
	fmt.Printf("%v\n", da)
}
// go test -v -timeout 30s . -run ^Test_Unmarshal$
// map[employees:map[first:map[firstName:Bill lastName:Gates] second:map[firstName:George lastName:Bush] third:map[firstName:Thomas lastName:Carter]]]

func Benchmark_Unmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parser([]byte(data))
	}
}
// go test -benchmem -run=^$ . -bench ^Benchmark_Unmarshal$
// Benchmark_Unmarshal-4   300000   4652 ns/op   2320 B/op   38 allocs/op

func Test_Marshal(t *testing.T) {
	da := Parser([]byte(data))
	d, _ := json.Marshal(da)
	fmt.Println(string(d))
}
// go test -v -timeout 30s . -run ^Test_Unmarshal$
// map[employees:map[first:map[firstName:Bill lastName:Gates] second:map[firstName:George lastName:Bush] third:map[firstName:Thomas lastName:Carter]]]

func Benchmark_Marshal(b *testing.B) {
	da := Parser([]byte(data))
	for i := 0; i < b.N; i++ { //use b.N for looping
		json.Marshal(da)
	}
}
// go test -benchmem -run=^$ . -bench ^Benchmark_Marshal$
// Benchmark_Marshal-4   200000   6455 ns/op   2192 B/op   49 allocs/op

func Test_MarshalLevel1(t *testing.T) {
	da := Parser([]byte(data))
	daJSON := da.(map[string]interface{})
	for k, v := range daJSON {
		d, _ := json.Marshal(v)
		daJSON[k] = string(d)
	}
	s, _ := json.Marshal(daJSON)
	fmt.Printf("Test_MarshalLevel1 %v\n", string(s))
}
// go test -v -timeout 30s . -run ^Test_MarshalLevel1$
// Test_MarshalLevel1 {"employees":"{\"first\":{\"firstName\":\"Bill\",\"lastName\":\"Gates\"},\"second\":{\"firstName\":\"George\",\"lastName\":\"Bush\"},\"third\":{\"firstName\":\"Thomas\",\"lastName\":\"Carter\"}}"}

func Benchmark_MarshalLevel1(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		da := Parser([]byte(data))
		daJSON := da.(map[string]interface{})
		for k, v := range daJSON {
			d, _ := json.Marshal(v)
			daJSON[k] = string(d)
		}
		json.Marshal(daJSON)
	}
}
// go test -benchmem -run=^$ . -bench ^Benchmark_MarshalLevel1$
// Benchmark_MarshalLevel1-4   100000   12399 ns/op   4880 B/op   90 allocs/op

func Test_MarshalLevel2(t *testing.T) {
	da := Parser([]byte(data))
	daJSON := da.(map[string]interface{})
	for k1, v1 := range daJSON {
		v1Map := v1.(map[string]interface{})
		for k2, v2 := range v1Map {
			d, _ := json.Marshal(v2)
			v1Map[k2] = string(d)
		}
		daJSON[k1] = v1Map
	}
	s, _ := json.Marshal(daJSON)
	fmt.Printf("Test_MarshalLevel2 %v\n", string(s))
}
// go test -v -timeout 30s . -run ^Test_MarshalLevel2$
// Test_MarshalLevel2 {"employees":{"first":"{\"firstName\":\"Bill\",\"lastName\":\"Gates\"}","second":"{\"firstName\":\"George\",\"lastName\":\"Bush\"}","third":"{\"firstName\":\"Thomas\",\"lastName\":\"Carter\"}"}}

func Benchmark_MarshalLevel2(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		da := Parser([]byte(data))
		daJSON := da.(map[string]interface{})
		for _, v1 := range daJSON {
			v1Map := v1.(map[string]interface{})
			for k2, v2 := range v1Map {
				d2, _ := json.Marshal(v2)
				v1Map[k2] = string(d2)
			}
		}
		json.Marshal(daJSON)
	}
}
// go test -benchmem -run=^$ . -bench ^Benchmark_MarshalLevel2$
// Benchmark_MarshalLevel2-4   100000   13297 ns/op   4881 B/op   96 allocs/op

func Test_MarshalLevel3(t *testing.T) {
	da := Parser([]byte(data))
	daJSON := da.(map[string]interface{})
	for k1, v1 := range daJSON {
		v1Map := v1.(map[string]interface{})
		for k2, v2 := range v1Map {
			v2Map := v2.(map[string]interface{})
			for k3, v3 := range v2Map {
				d, _ := json.Marshal(v3)
				v2Map[k3] = string(d)
			}
			v1Map[k2] = v2Map
		}
		daJSON[k1] = v1Map
	}
	s, _ := json.Marshal(daJSON)
	fmt.Printf("Test_MarshalLevel2 %v\n", string(s))
}
// go test -v -timeout 30s . -run ^Test_MarshalLevel3$
// Test_MarshalLevel2 {"employees":{"first":{"firstName":"\"Bill\"","lastName":"\"Gates\""},"second":{"firstName":"\"George\"","lastName":"\"Bush\""},"third":{"firstName":"\"Thomas\"","lastName":"\"Carter\""}}}

func Benchmark_MarshalLevel3(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		da := Parser([]byte(data))
		daJSON := da.(map[string]interface{})
		for k1, v1 := range daJSON {
			v1Map := v1.(map[string]interface{})
			for k2, v2 := range v1Map {
				v2Map := v2.(map[string]interface{})
				for k3, v3 := range v2Map {
					d, _ := json.Marshal(v3)
					v2Map[k3] = string(d)
				}
				v1Map[k2] = v2Map
			}
			daJSON[k1] = v1Map
		}
		json.Marshal(daJSON)
	}
}

func main() {
	Test_Unmarshal()
}
