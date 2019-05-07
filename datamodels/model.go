package datamodels

import "time"

type DatatablesDemo struct {
	First_name string    `json:"first_name" xorm:"varchar(32) 'first_name'"`
	Last_name  string    `json:"last_name" xorm:"int 'last_name'"`
	Position   string    `json:"position" xorm:"varchar(32) 'position'"`
	Office     string    `json:"office" xorm:"varchar(32) 'office'"`
	Start_date time.Time `json:"start_date" xorm:" 'start_date'"`
	Salary     int       `json:"salary" xorm:"int 'salary'"`
}


type HttpRequestResult struct {
	Time            int64       `json:"tm"`              //时间戳
	Code            int         `json:"draw"`            //返回码
	RecordsTotal    int         `json:"recordsTotal"`    //时间戳
	RecordsFiltered int         `json:"recordsFiltered"` //时间戳
	Message         string      `json:"msg"`             //消息内容
	Data            interface{} `json:"data"`
}

type ReturnDataInfos struct {
	Code       int         `json:"code"`       //返回码
	Flag       bool        `json:"flag"`       //标志
	Time       int64       `json:"tm"`         //时间戳
	Offset     int         `json:"offset"`     //偏移
	Limit      int         `json:"limit"`      //Limit
	TotalCount int         `json:"totalcount"` //总记录数
	Message    string      `json:"msg"`        //消息内容
	Results    interface{} `json:"results"`    //多个消息对象
}


//VehicleItem
//设备上行
type DbVehicleItem struct {
	Id               int64   `xorm:"not null pk autoincr int64 'id'"` //id
	Cmd              int     `xorm:"int 'cmd'"`                       //命令
	Sn               string  `xorm:"varchar(32) 'sn'"`                //sn
	CarNo            string  `xorm:"varchar(12) 'carno'"`             //车牌号
	Version          string  `xorm:"varchar(12) 'version'"`           //版本号
	Time             int64   `xorm:"BigInt 'time'"`                   //时间戳
	Gear             int     `xorm:"int 'gear'"`                      //档位
	FootBrakeAngle   float32 `xorm:"float32 'footbrakeangle'"`        //脚刹
	HandBrake        int     `xorm:"int 'handbrake'"`                 //手刹
	Safetybelt       int     `xorm:"int 'safetybelt'"`                //安全带
	Trumpet          int     `xorm:"int 'trumpet'"`                   //喇叭
	HeadLight        int     `xorm:"int 'headlight'"`                 //近光灯
	DistanceLight    int     `xorm:"int 'distancelight'"`             //远光灯
	ShutDownSignal   int     `xorm:"int 'shutdownsignal'"`            //熄火信号
	FrontFogLamp     int     `xorm:"int 'frontfoglamp'"`              //前雾灯
	RearFogLamp      int     `xorm:"int 'rearfoglamp'"`               //后雾灯
	EmergencyLight   int     `xorm:"int 'emergencylight'"`            //应急灯
	StartSignal      int     `xorm:"int 'startsignal'"`               //点火start信号
	MotoSpeed        int     `xorm:"int 'motospeed'"`                 //电机转速
	WidthLamp        int     `xorm:"int 'widthlamp'"`                 //示宽灯
	LeftTurnSignal   int     `xorm:"int 'leftturnsignal'"`            //左转向灯
	RightTurnSignal  int     `xorm:"int 'rightturnsignal'"`           //右转向灯
	WindscreenWiper  int     `xorm:"int 'windscreenwiper'"`           //雨刷
	CarSpeed         int     `xorm:"int 'carspeed'"`                  //车速
	ACC              int     `xorm:"int 'acc'"`                       //ACC
	ElectricQuantity float32 `xorm:"float32 'electricquantity'"`      //电量
	SteeringAngle    float32 `xorm:"float32 'steeringangle'"`         //方向盘角度
	ThrottleAngle    float32 `xorm:"float32 'throttleangle'"`         //油门角度
	CarDoor          int     `xorm:"int 'cardoor'"`                   //车门
	BeforeTheHatch   int     `xorm:"int 'beforethehatch'"`            //前舱盖
	AfterTheHatch    int     `xorm:"int 'afterthehatch'"`             //后舱盖
	IgnitionSignal   int     `xorm:"int 'ignitionsignal'"`            //点火ON档信号
	CarWindow        int     `xorm:"int 'carwindow'"`                 //车窗
	Longitude        float64 `xorm:"float64 'longitude'"`             //经度
	Latitude         float64 `xorm:"float64 'latitude'"`              //纬度
	Height           float64 `xorm:"float64 'height'"`                //高度
	Created          int64   `xorm:"created"`
	Updated          int64   `xorm:"updated"`
}