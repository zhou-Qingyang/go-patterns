package abstract_factory_method

import "testing"

func TestCreatePhone(t *testing.T) {
	huaweiDevice := NewDeviceFactory("huawei")
	huaweiPhone := huaweiDevice.CreatePhone()
	huaweiPhone.PowerOn()
	huaWeiComputer := huaweiDevice.CreateComputer()
	huaWeiComputer.Boot()
}
