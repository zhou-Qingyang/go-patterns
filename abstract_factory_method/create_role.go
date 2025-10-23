package abstract_factory_method

type FirstSkill interface {
	Shot() string
}
type SecondSkill interface {
	Jump() string
}

type JettFirstSkill struct {
}

func (j *JettFirstSkill) Shot() string {
	return "jett shoot"
}

type JettSecondSkill struct{}

func (j *JettSecondSkill) Jump() string {
	return "jett jump"
}

type NiHongFirstSkill struct{}

func (n *NiHongFirstSkill) Shot() string {
	return "nihong shoot"
}

type NiHongSecondSkill struct{}

func (j *NiHongSecondSkill) Jump() string {
	return "nihong jump"
}

type VolentRoleFactory interface {
	GetFirstSkill() FirstSkill
	GetSecondSkill() SecondSkill
}

type JettFactory struct {
}

func (j *JettFactory) GetFirstSkill() FirstSkill {
	return &JettFirstSkill{}
}
func (j *JettFactory) GetSecondSkill() SecondSkill {
	return &JettSecondSkill{}
}

type NiHongFactory struct{}

func (n *NiHongFactory) GetFirstSkill() FirstSkill {
	return &NiHongFirstSkill{}
}
func (n *NiHongFactory) GetSecondSkill() SecondSkill {
	return &NiHongSecondSkill{}
}
