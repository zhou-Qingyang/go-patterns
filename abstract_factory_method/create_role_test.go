package abstract_factory_method

import (
	"fmt"
	"testing"
)

func TestCreateRole(t *testing.T) {
	var jett VolentRoleFactory = &JettFactory{}
	firstSkill := jett.GetFirstSkill()
	secondSkill := jett.GetSecondSkill()

	fmt.Println(firstSkill.Shot())
	fmt.Println(secondSkill.Jump())

	var nihong VolentRoleFactory = &NiHongFactory{}
	firstSkill = nihong.GetFirstSkill()
	secondSkill = nihong.GetSecondSkill()
	fmt.Println(firstSkill.Shot())
	fmt.Println(secondSkill.Jump())
}
