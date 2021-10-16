package connection

import ( 
	"fmt"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const(
	UserName      string = "golang_user"
	Password      string = "golang_passw0rd"
	Addr          string = "127.0.0.1"
	Port      	  int    = 4306	   
	Database      string = "godb"
	MaxLifetime   int    = 10
	MaxOpenConns  int    = 10
	MaxIdleConns  int    = 10
)
type User struct {
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string    `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Status    int32     `gorm:"type:int(5);" json:"status,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func ConnectDatabase(){
	//組合sql連線字串
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",UserName,Password,Addr,Port,Database)
	//連接mysql
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})

	if err != nil {
		fmt.Println("get db failed:",err)
	}

	
	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	
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