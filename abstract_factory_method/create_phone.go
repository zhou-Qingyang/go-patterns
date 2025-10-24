package abstract_factory_method

import "fmt"

// 1. 抽象产品接口（先定义接口，再定义实现，结构更清晰）
// 手机产品接口（产品等级结构1）
type Phone interface {
	PowerOn() // 方法名首字母大写，支持包外调用
}

// 电脑产品接口（产品等级结构2）
type Computer interface {
	Boot() // 方法名首字母大写，支持包外调用
}

// 2. 抽象工厂接口（定义创建产品族的方法）
type DeviceFactory interface {
	CreatePhone() Phone       // 创建手机（属于产品族的一部分）
	CreateComputer() Computer // 创建电脑（属于产品族的另一部分）
}

// 3. 华为产品族（具体产品）
// 华为手机
type HuaweiPhone struct{}

// 实现Phone接口的PowerOn方法（职责：打印自身功能，而非工厂）
func (h *HuaweiPhone) PowerOn() {
	fmt.Println("Huawei Phone is powering on...") // 统一品牌名小写（Huawei规范拼写）
}

// 华为电脑
type HuaweiComputer struct{}

// 实现Computer接口的Boot方法
func (h *HuaweiComputer) Boot() {
	fmt.Println("Huawei Computer is booting...")
}

// 华为工厂（具体工厂：负责创建华为产品族）
type HuaweiFactory struct{}

// 创建华为手机（仅负责实例化，不包含产品功能逻辑）
func (h *HuaweiFactory) CreatePhone() Phone {
	return &HuaweiPhone{}
}

// 创建华为电脑
func (h *HuaweiFactory) CreateComputer() Computer {
	return &HuaweiComputer{}
}

// 4. 苹果产品族（具体产品）
// 苹果手机
type ApplePhone struct{}

func (a *ApplePhone) PowerOn() {
	fmt.Println("Apple Phone is powering on...") // 统一首字母大写，格式一致
}

// 苹果电脑
type AppleComputer struct{}

func (a *AppleComputer) Boot() {
	fmt.Println("Apple Computer is booting...")
}

// 苹果工厂（具体工厂：负责创建苹果产品族）
type AppleFactory struct{}

// 创建苹果手机（工厂仅创建对象，不执行产品方法）
func (a *AppleFactory) CreatePhone() Phone {
	return &ApplePhone{}
}

func (a *AppleFactory) CreateComputer() Computer {
	return &AppleComputer{}
}

// 5. 工厂生产者（简单工厂模式，用于获取具体工厂实例）
// 入参品牌字符串，返回对应的抽象工厂接口
func NewDeviceFactory(brand string) DeviceFactory {
	switch brand {
	case "huawei": // 统一小写入参，避免大小写敏感问题
		return &HuaweiFactory{}
	case "apple":
		return &AppleFactory{}
	default:
		return nil // 实际项目中可返回error，这里简化处理
	}
}
