package superhero

type SuperType string

const (
	HeroType    SuperType = "HERO"
	VillainType SuperType = "VILLAIN"
)

type Super struct {
	ID           int64  `gorm:"PRIMARY_KEY"`
	UUID         int64  `gorm:"UNIQUE"`
	FullName     string `gorm:"column:full_name"`
	Name         string
	Intelligence int64
	Power        int64
	Occupation   string
	Image        string
	Type         SuperType
}
