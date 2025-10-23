package factory_method

import "errors"

// 定义枪的类型常量，避免字符串硬编码
type GunType string

const (
	GunTypeAK47 GunType = "ak47"
	GunTypeM4A1 GunType = "m4a1"
	// 新增类型只需添加常量，无需修改工厂核心逻辑
	// GunTypeAWM   GunType = "awm"
)

// IGun 接口规范（方法名改为驼峰式，符合Go规范）
type IGun interface {
	GetPower() int
	GetGunName() string
}

// Ak47 具体实现
type Ak47 struct{}

func (a *Ak47) GetPower() int {
	return 2
}

func (a *Ak47) GetGunName() string {
	return string(GunTypeAK47)
}

// M4A1 具体实现
type M4A1 struct{}

func (m *M4A1) GetPower() int {
	return 1
}

func (m *M4A1) GetGunName() string {
	return string(GunTypeM4A1)
}

// ------------------------------
// 注册式工厂核心逻辑
// ------------------------------

// 定义创建函数类型，每个具体枪类型需提供对应的创建函数
type gunCreator func() IGun

// 工厂注册表，存储类型与创建函数的映射
var gunRegistry = make(map[GunType]gunCreator)

// 注册函数：将枪类型与对应的创建函数关联，供工厂调用
func RegisterGun(gunType GunType, creator gunCreator) {
	if creator == nil {
		panic("factory_method: 注册的创建函数不能为nil")
	}
	if _, exists := gunRegistry[gunType]; exists {
		panic("factory_method: 该枪类型已注册")
	}
	gunRegistry[gunType] = creator
}

// NewGun 工厂方法：根据类型创建对应的枪实例（返回错误而非nil）
func NewGun(gunType GunType) (IGun, error) {
	creator, exists := gunRegistry[gunType]
	if !exists {
		return nil, errors.New("factory_method: 不支持的枪类型")
	}
	return creator(), nil
}

// 初始化时注册所有支持的枪类型（类似“自注册”）
func init() {
	RegisterGun(GunTypeAK47, func() IGun { return &Ak47{} })
	RegisterGun(GunTypeM4A1, func() IGun { return &M4A1{} })
	// 新增类型时，只需在这里添加注册即可，无需修改NewGun和RegisterGun
	// RegisterGun(GunTypeAWM, func() IGun { return &AWM{} })
}
