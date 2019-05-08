$(document).ready(function() {
    $('#example').DataTable( {
        "bSort": false,
        "bInfo":false,
        "bLengthChange":false,
        "bAutoWidth":false,
        "bFilter": false,
        "autoWidth": false,
        "processing": true,
        "serverSide": true,

        "sAjaxSource": "/vehicle/pagedatalist",
        columns: [
            { data: "Id"            },
            { data: 'Sn'            },
            { data: 'CarNo'         },
            { data: "Time"          },
            { data: 'Longitude'      },
            { data: 'Latitude'       } ,
            { data: 'Height'            },
            { data: "Gear"           },
            { data: 'FootBrakeAngle'  },
            { data: 'HandBrake'         },
            { data: "Safetybelt"     },
            { data: 'Trumpet'           },
            { data: 'HeadLight'         } ,
            { data: 'DistanceLight'     },//远光灯
            { data: 'ShutDownSignal'  },
            { data: 'FrontFogLamp'      },
            { data: "RearFogLamp"       },
            { data: 'EmergencyLight'  },
            { data: 'StartSignal'       } ,
            { data: 'MotoSpeed'         },
            { data: "WidthLamp"         },
            { data: 'LeftTurnSignal'  },
            { data: 'RightTurnSignal' },
            { data: "WindscreenWiper" },
            { data: 'CarSpeed'          },
            { data: 'ACC'               } ,
            { data: 'ElectricQuantity' }
        ],
        fnServerData:function (sSource, aoData, fnCallback) {
            var globalSearch = {};
            var searchConditions = {};
            searchConditions.limit = 100;//页面显示记录条数，在页面显示每页显示多少项的时候
            searchConditions.start = 10;//开始的记录序号
            globalSearch.name= 'globalSearch';
            globalSearch.value= JSON.stringify(searchConditions);
            console.log("********globalSearch*******"+JSON.stringify(globalSearch));
            console.log("********searchConditions*******"+JSON.stringify(searchConditions));
            aoData.push(globalSearch);

            console.log("********sSource*******"+JSON.stringify(sSource));
            console.log("********aoData******"+JSON.stringify(aoData));
            var param = {};
            param.limit = aoData.iDisplayLength;//页面显示记录条数，在页面显示每页显示多少项的时候
            param.start = aoData.iDisplayStart;//开始的记录序号
            param.page = (aoData.start / aoData.length)+1;//当前页码
            console.log("===================="+param);
            $.ajax({
                type: "GET",
                url: sSource,
                dataType: "json",
                contentType: "application/json;charset=UTF-8",
                data: aoData, // 以json格式传递
                success: function(result) {
                    if(result.flag){
                        var obj =result;
                        var data = obj.results;
                        var aaData = [];
                        var dataResult = {};
                        dataResult.iTotalRecords = obj.totalcount;
                        dataResult.iTotalDisplayRecords= obj.totalcount;
                        dataResult.iDisplayStart = obj.offset;
                        dataResult.iDisplayLength= obj.limit;
                        dataResult.aaData = data;
                        fnCallback(dataResult);
                    }
                }
            });
        }

    } );
} );