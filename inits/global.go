package inits

import (
	"carauto/datamodels"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
)

const (
	host = "localhost"
	//host = "192.168.0.11"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "cardb"
)
func NewEngine() (*xorm.Engine, error) {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	//fmt.Printf("NewEngine======%v \n", connStr)
	engine, err := xorm.NewEngine("postgres", connStr)
	return engine, err
}

func init2()  {
	engine, err := NewEngine()
	defer engine.Close()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
		fmt.Println("数据库连接失败")
	}

	if err := engine.Sync(new(datamodels.DbVehicleItem)); err != nil {
		log.Fatal("数据表同步失败:", err)
		fmt.Println("数据表同步失败")
	}
	
}
