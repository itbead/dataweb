package services

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/itbread/dataweb/datamodels"
	"reflect"
)


type VehicleItemService struct {
	Engine *xorm.Engine
}

func (c *VehicleItemService) SelectMany(code string, start int, limit int) (results []datamodels.DbVehicleItem,err error) {
	//c.Engine.ShowSQL(true)
	typeName := reflect.TypeOf(c)
	err =c.Engine.Where("sn=?", code).Limit(limit, start).Find(&results)
	if err!=nil{
		fmt.Println("%v err=",err,"results=",typeName,results)
	}

	return results,err
}
func (c *VehicleItemService) Count( ) (counts int64 ) {
	//c.Engine.ShowSQL(true)
	typeName := reflect.TypeOf(c)
	var tmp datamodels.DbVehicleItem
	var err error
	counts, err = c.Engine.Count(&tmp)
	fmt.Printf("%v Count==%v \n err=%v \n",typeName, counts,err)
	return counts
}