package user

type AgeCategory int

const (
	Adult AgeCategory = iota
	Retiree
)

func (u *User) GetAgeCategory() AgeCategory {
	switch {
	case u.Age > 60:
		return Retiree
	default:
		return Adult
	}
}
