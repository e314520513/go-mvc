package connection

import ( 
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


func ConnectDatabase(){
	//組合sql連線字串
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",UserName,Password,Addr,Port,Database)
	//連接mysql
	database, err := gorm.Open(mysql.Open(addr), &gorm.Config{})

	if err != nil {
		fmt.Println("get db failed:",err)
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)



	
}