package models

import(
	"time"
	"fmt"
	"gomvc/connection"
)
type User struct {
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string    `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Status    int32     `gorm:"type:int(5);" json:"status,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func Users(){
	conn :=connection.ConnectDatabase()
	//產生table
	conn.Debug().AutoMigrate(&User{})
	//判斷有沒有table存在
	migrator := conn.Migrator()
	has := migrator.HasTable(&User{})
	//has := migrator.HasTable("GG")
	if !has {
		fmt.Println("table not exist")
	}
}