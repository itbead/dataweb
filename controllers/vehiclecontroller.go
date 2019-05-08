package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/itbread/dataweb/datamodels"
	"github.com/itbread/dataweb/inits"
	"github.com/itbread/dataweb/services"
	"github.com/kataras/iris"
	"reflect"
	"strconv"
	"time"
)


//HomeController
//HomeController
type VehicleController struct {
	Ctx     iris.Context
	Service services.VehicleItemService
}

func (c *VehicleController) Get()  {
	c.Ctx.ViewData("message", "HomeController page!")
	c.Ctx.View("vehicleitemlist.html")
}

func (c *VehicleController) GetItemlist()  {
	c.Ctx.ViewData("message", "HomeController page!")
	c.Ctx.View("itemlist.html")
}

func (c *VehicleController) GetPagelist()  {
	c.Ctx.ViewData("message", "HomeController page!")
	c.Ctx.View("pagelist.html")
}


func (c *VehicleController) GetIteminfolist() {

	typeName := reflect.TypeOf(c)
	var err error
	var result datamodels.HttpRequestResult
	//var para datamodels.HttpRequestPara
	fmt.Printf("%v FormValues====%v \n",typeName,c.Ctx.FormValues())
	tmp_start :=c.Ctx.FormValue("start")
	tmp_length :=c.Ctx.FormValue("length")
	search :=c.Ctx.FormValue("search")

	start,err:=strconv.Atoi(tmp_start)
	length,err:=strconv.Atoi(tmp_length)
	fmt.Printf("start  ====%v =%v  \n",tmp_start,start)
	fmt.Printf("length ====%v =%v  \n",tmp_length,length)
	fmt.Printf("search ====%v      \n",search)
	c.Service.Engine, err = inits.NewEngine()
	if err != nil {
		fmt.Println("inits.NewEngine() err=", err)
	}
	sn:="sn123"
	vehicles, err2 := c.Service.SelectMany(sn,start,length)
	totals:=c.Service.Count()
	if err2 != nil {
		result.Code = 1002
		result.Message = fmt.Sprintf("error=%v \n", err2)
		result.Data = vehicles
		result.RecordsTotal=len(vehicles)
		result.RecordsFiltered=len(vehicles)
		c.Ctx.JSON(result)
		return
	} else {
		result.Data = vehicles
		result.RecordsTotal=int(totals)
		result.RecordsFiltered=int(totals)
		json.Marshal(result)
	}
	result.Code = 0
	result.Time = time.Now().Unix()
	var tmp string
	tmp="success"
	if err!=nil {
		tmp=fmt.Sprintf("%v",err)
	}

	result.Message = fmt.Sprintf("%v 2====%v",typeName,tmp)
	c.Ctx.JSON(result)
}

func (c *VehicleController) GetPagedatalist() {

	typeName := reflect.TypeOf(c)
	var err error
	var result datamodels.ReturnDataInfos
	fmt.Printf("%v FormValues====%v \n",typeName,c.Ctx.FormValues())
	tmp_start :=c.Ctx.FormValue("iDisplayStart")
	tmp_length :=c.Ctx.FormValue("iDisplayLength")
	globalSearch :=c.Ctx.FormValue("globalSearch")

	start,err:=strconv.Atoi(tmp_start)
	length,err:=strconv.Atoi(tmp_length)
	fmt.Printf("start====%v \n",start)
	fmt.Printf("length====%v \n",length)
	fmt.Printf("search====%v \n",globalSearch)
	fmt.Println("v% 0==================sn%v",typeName,c.Ctx.Params())
	c.Service.Engine, err = inits.NewEngine()
	if err == nil {
		fmt.Println("inits.NewEngine() err=", err)
	}
	begin := time.Now().Add(-time.Duration(24) * time.Hour).Unix()
	end := time.Now().Unix()
	fmt.Printf("%v 1============begin=%v end=%v \n",typeName,begin,end)
	sn:="sn123"

	vehicles, err2 := c.Service.SelectMany(sn,start,length)
	totals:=c.Service.Count()
	result.Offset=start
	result.Limit=length
	result.TotalCount=int(totals)
	if err2 != nil {
		result.Code = 1002
		result.Message = fmt.Sprintf("error=%v \n", err2)
		result.Results = vehicles
		c.Ctx.JSON(result)
		return
	} else {
		result.Flag=true
		result.Results = vehicles
		json.Marshal(result)
	}
	result.Code = 0
	result.Time = time.Now().Unix()
	tmp :="success"
	if err!=nil {
		tmp=fmt.Sprintf("%v",err)
	}

	result.Message = fmt.Sprintf("%v 2==================%v",typeName,tmp)
	c.Ctx.JSON(result)
}


