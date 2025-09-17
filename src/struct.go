package main

import "fmt"

type dict struct {
	data map[int]string
}
func newDict() *dict {
	return &dict{data: make(map[int]string)}
}

func (d *dict) add(key int, value string) {
	if d.data == nil {
		d.data = make(map[int]string)
	}
	d.data[key] = value
}

func (d *dict) get(key int) string {
	return d.data[key]
}

func (d *dict) del(key int) {
	delete(d.data, key)
}

func main() {
	fmt.Println("=== 기본 방법: 구조체 리터럴 사용 ===")
	d1 := &dict{}
	d1.add(1, "one")
	d1.add(2, "two")
	fmt.Println("d1.get(1):", d1.get(1))

	fmt.Println("\n=== new(dict) 사용 ===")
	d2 := new(dict)
	d2.data = make(map[int]string) // new로 생성하면 zero value이므로 map 초기화 필요
	d2.add(3, "three")
	d2.add(4, "four")
	fmt.Println("d2.get(3):", d2.get(3))

	fmt.Println("\n=== newDict 함수 사용 ===")
	d3 := newDict()
	d3.add(5, "five")
	d3.add(6, "six")
	fmt.Println("d3.get(5):", d3.get(5))
}