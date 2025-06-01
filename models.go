package hibiscus

type Config struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Key      string `json:"key" gorm:"size:128;uniqueIndex"`
	Desc     string `json:"desc" gorm:"size:200"`
	Autoload bool   `json:"autoload" gorm:"index"`
	Public   bool   `json:"public" gorm:"index" default:"false"`
	Format   string `json:"format" gorm:"size:20" default:"text" comment:"json,yaml,int,float,bool,text"`
	Value    string
}
