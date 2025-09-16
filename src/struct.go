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
	// d := &dict{}
	// d.add(1, "one")
	// d.add(2, "two")
	// fmt.Println(d.get(1))
	// d.del(1)
	// fmt.Println(d.get(1))

	d := newDict()
	d.add(1, "one")
	d.add(2, "two")
	fmt.Println(d.get(1))
	d.del(1)
	fmt.Println(d.get(1))
}