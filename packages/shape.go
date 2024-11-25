package packages

import "fmt"

type Shape struct {
	Type   string
	volume int64
}

func (shape *Shape) GetVolume() int64 {
	if shape.Type == "" {
		fmt.Println("Please, add type of shape")
	}
	return shape.volume * 4
}

func init() {
	circle := Shape{
		Type:   "Circle",
		volume: 4,
	}

	block := Shape{
		Type:   "Block",
		volume: 8,
	}

	fmt.Println(circle.GetVolume())
	fmt.Println(block.GetVolume())
}
