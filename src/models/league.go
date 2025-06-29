package models

type League struct {
    ID         uint   `gorm:"column:id" json:"id"`
    NameLeague string `gorm:"column:name_league" json:"name_league"`
    LogoLeague string `gorm:"column:logo_league" json:"logo_league"`
}

// TableName permite que gorm use tu tabla personalizada sin renombrarla a "leagues"
func (League) TableName() string {
    return "leagues"
}


