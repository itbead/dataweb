package services

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/itbread/dataweb/datamodels"
)

//DBService
//数据库接口
type BaseService interface {
	SelectOne(code string) (item interface{}, found bool)
	SelectMany(code string, max int64, small int64) (results []interface{},err error)
	Insert(item interface{}) (success bool, err error)
	Update(item interface{}) (success bool, err error)
	Delete(code string, limit int) (success bool, err error)
}

type InfoService struct {
	Engine *xorm.Engine
}

func (c *InfoService) SelectMany(start int,limit int ) (results []datamodels.DatatablesDemo,err error) {
	c.Engine.ShowSQL(true)
	err =c.Engine.Limit(limit, start).Find(&results)
	fmt.Printf("InfoService SelectMany==%v \n err=%v\n", results,err)
	return results,err
}

func (c *InfoService) Count( ) (counts int64 ) {
	//c.Engine.ShowSQL(true)
	var tmp datamodels.DatatablesDemo
	var err error
	counts, err = c.Engine.Count(&tmp)
	fmt.Printf("InfoService DatatablesDemo Count==%v \n err=%v\n", counts,err)
	return counts
}

func (c *InfoService) SelectManyItems(code string, start int, limit int) (results []datamodels.DbVehicleItem,err error) {
	//c.Engine.ShowSQL(true)
	err =c.Engine.Where("sn=?", code).Limit(limit, start).Find(&results)
	if err!=nil{
		fmt.Println("InfoService SelectManyItems err=",err,"results=",results)
	}

	return results,err
}
func (c *InfoService) CountItems( ) (counts int64 ) {
	//c.Engine.ShowSQL(true)
	var tmp datamodels.DbVehicleItem
	var err error
	counts, err = c.Engine.Count(&tmp)
	fmt.Printf("InfoService DbVehicleItem Count==%v \n err=%v\n", counts,err)
	return counts
}