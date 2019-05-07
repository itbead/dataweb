package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/itbread/dataweb/datamodels"
	"github.com/itbread/dataweb/inits"
	"github.com/itbread/dataweb/services"
	"github.com/kataras/iris"
	"strconv"
	"time"
)


//HomeController
//HomeController
type HomeController struct {
	Ctx     iris.Context
	Service services.InfoService
}

func (c *HomeController) Get()  {
	c.Ctx.ViewData("message", "HomeController page!")
	c.Ctx.View("index.html")
}

func (c *HomeController) GetItemlist()  {
	c.Ctx.ViewData("message", "HomeController page!")
	c.Ctx.View("itemlist.html")
}

func (c *HomeController) GetPeopleinfolistBy(sn string) {
	var err error
	var result datamodels.ReturnDataInfos
	//var para datamodels.HttpRequestPara
	fmt.Printf("FormValues====%v \n",c.Ctx.FormValues())
	fmt.Println("GetPeopleinfolistBy==================sn%v",c.Ctx.Params())
	c.Service.Engine, err = inits.NewEngine()
		if err == nil {
		fmt.Println("inits.NewEngine() err=", err)
	}
	begin := time.Now().Add(-time.Duration(24) * time.Hour).Unix()
	end := time.Now().Unix()
	fmt.Printf("GetHistorylist============begin=%v end=%v",begin,end)
	vehicles, err2 := c.Service.SelectMany(0,100)
	if err2 != nil {
		result.Code = 1002
		result.Message = fmt.Sprintf("error=%v \n", err2)
		c.Ctx.JSON(result)
		return
	} else {
		result.Results = vehicles
		json.Marshal(result)
	}
	result.Code = 0
	result.Time = time.Now().Unix()
	var tmp string
	tmp="success"
	if err!=nil {
		tmp=fmt.Sprintf("%v",err)
	}

	result.Message = fmt.Sprintf("PostHistorylist==================%v",tmp)
	c.Ctx.JSON(result)
}

func (c *HomeController) GetPeopleinfolist() {
	var err error
	var result datamodels.HttpRequestResult
	//var para datamodels.HttpRequestPara
	fmt.Printf("FormValues====%v \n",c.Ctx.FormValues())
	tmp_start :=c.Ctx.FormValue("start")
	tmp_length :=c.Ctx.FormValue("length")
	search :=c.Ctx.FormValue("search")

	start,err:=strconv.Atoi(tmp_start)
	length,err:=strconv.Atoi(tmp_length)
	fmt.Printf("start====%v \n",start)
	fmt.Printf("length====%v \n",length)
	fmt.Printf("search====%v \n",search)
	fmt.Println("GetPeopleinfolist0==================sn%v",c.Ctx.Params())
	c.Service.Engine, err = inits.NewEngine()
		if err == nil {
		fmt.Println("inits.NewEngine() err=", err)
	}
	begin := time.Now().Add(-time.Duration(24) * time.Hour).Unix()
	end := time.Now().Unix()
	fmt.Printf("GetPeopleinfolist1============begin=%v end=%v \n",begin,end)
	vehicles, err2 := c.Service.SelectMany(start,length)
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

	result.Message = fmt.Sprintf("GetPeopleinfolist2==================%v",tmp)
	c.Ctx.JSON(result)
}


func (c *HomeController) GetIteminfolist() {
	var err error
	var result datamodels.HttpRequestResult
	//var para datamodels.HttpRequestPara
	fmt.Printf("GetIteminfolist FormValues====%v \n",c.Ctx.FormValues())
	tmp_start :=c.Ctx.FormValue("start")
	tmp_length :=c.Ctx.FormValue("length")
	search :=c.Ctx.FormValue("search")

	start,err:=strconv.Atoi(tmp_start)
	length,err:=strconv.Atoi(tmp_length)
	fmt.Printf("start====%v \n",start)
	fmt.Printf("length====%v \n",length)
	fmt.Printf("search====%v \n",search)
	fmt.Println("GetIteminfolist 0==================sn%v",c.Ctx.Params())
	c.Service.Engine, err = inits.NewEngine()
	if err == nil {
		fmt.Println("inits.NewEngine() err=", err)
	}
	begin := time.Now().Add(-time.Duration(24) * time.Hour).Unix()
	end := time.Now().Unix()
	fmt.Printf("GetIteminfolist 1============begin=%v end=%v \n",begin,end)
	sn:="sn123"
	vehicles, err2 := c.Service.SelectManyItems(sn,start,length)
	totals:=c.Service.CountItems()
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

	result.Message = fmt.Sprintf("GetIteminfolist 2==================%v",tmp)
	c.Ctx.JSON(result)
}


