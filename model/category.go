package model

type Category struct {
	Model
	Name string
}

func (category *Category) GetList() map[string]Category {
	return map[string]Category{
		"1":  {Model: Model{Id: 1}, Name: "php"},
		"2":  {Model: Model{Id: 2}, Name: "laravel"},
		"3":  {Model: Model{Id: 3}, Name: "mysql"},
		"4":  {Model: Model{Id: 4}, Name: "docker"},
		"5":  {Model: Model{Id: 5}, Name: "redis"},
		"6":  {Model: Model{Id: 6}, Name: "rabbitmq"},
		"7":  {Model: Model{Id: 7}, Name: "go"},
		"8":  {Model: Model{Id: 8}, Name: "js"},
		"99": {Model: Model{Id: 99}, Name: "其他"},
	}
}
