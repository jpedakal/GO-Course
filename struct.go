package main

type mesure struct {
	height, width int
}

func (r *mesure) area() int {
	return r.height * r.width
}

func main() {
	r := mesure{height: 25, width: 25}
	fmt.Println("area", r.area)
}
