package models

type Tobuy struct {
	// tobuy data
	// Name string `gorm:"type:varchar(64) NOT NULL;" json:"name"`
	ID   string `json:"tobuy_id"`
	Name string `json:"name"`
}

// [Cmd] [Action] [Item] [other]
// ex) @tobuy add sugar other
type TobuyArgs struct {
	Cmd      string // tobuy | use CmdConversionMap
	Action   string // add, list, delete | use ActionConversionMap
	Item     string // [any] | no conversionMap
	Category string // [any of categories] | CategoryConversionMap
}

type TobuyRepository interface {
	Add(args *TobuyArgs) (string, error)
	GetOne(tobuyID string) (*Tobuy, error)
	GetList() (*[]Tobuy, error)
	Delete(tobuyID string) (string, error)
}
