package decorator

type Coffee interface {
	GetDescription() string
}

type CommonCoffee struct {
}

func (c *CommonCoffee) GetDescription() string {
	return "Regular Coffee"
}

func NewCommonCoffee() Coffee {
	return &CommonCoffee{}
}

type DecoratorCoffee struct {
	coffee Coffee // 持有Coffee接口（关键：依赖抽象）
}

type MilkCoffee struct {
	DecoratorCoffee
}

func (c *MilkCoffee) GetDescription() string {
	return "Add Milk " + c.coffee.GetDescription()
}

type SugarCoffee struct {
	DecoratorCoffee
}

func (c *SugarCoffee) GetDescription() string {
	return "Add Sugar " + c.coffee.GetDescription()
}

func NewMilkCoffee(coffee Coffee) Coffee {
	return &MilkCoffee{
		DecoratorCoffee: DecoratorCoffee{
			coffee: coffee,
		},
	}
}

func NewSugarCoffee(coffee Coffee) Coffee {
	return &SugarCoffee{
		DecoratorCoffee: DecoratorCoffee{
			coffee: coffee,
		},
	}
}
