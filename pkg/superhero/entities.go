package superhero

type SuperType string

const (
	HeroType    SuperType = "HERO"
	VillainType SuperType = "VILLAIN"
	NoType      SuperType = ""
)

type Super struct {
	ID           int64     `json:"id" gorm:"PRIMARY_KEY"`
	UUID         int64     `json:"uuid" gorm:"UNIQUE"`
	FullName     string    `json:"full_name" gorm:"column:full_name"`
	Name         string    `json:"name"`
	Intelligence int64     `json:"intelligence"`
	Power        int64     `json:"power"`
	Occupation   string    `json:"occupation"`
	Image        string    `json:"image"`
	Type         SuperType `json:"type"`
}
