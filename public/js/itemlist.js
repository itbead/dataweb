$(document).ready(function() {
    $('#example').DataTable( {
        "processing": true,
        "serverSide": true,
        "ajax": "/iteminfolist",
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
        dataSrc: function(data){
            var obj =result;
            var data = obj.results;
            var aaData = [];
            var dataResult = {};
            dataResult.iTotalRecords = obj.totalcount;
            dataResult.iTotalDisplayRecords= obj.totalcount;
            dataResult.iDisplayStart = obj.offset;
            dataResult.iDisplayLength= obj.limit;
            dataResult.aaData = aaData;

            var datasource={}
            if (data.code == 0) {
                if(data.Results==null)
                    datasource.length=0;
                else
                    datasource=data.Results;

            }else{
                datasource.length=0;
            }

            return datasource
        }
    } );
} );