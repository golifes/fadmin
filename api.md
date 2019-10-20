host:127.0.0.1

```
一:域操作(增删改查):
    1:添加接口:
        uri:/admin/domain
        method:post
        params:
                {
                    "name":"xxxx"// 域名称(必传)
                 }
        返回值:
                {
                    "code": 200, 自定义状态码 200代表正常成功，其他的代表异常，对应异常信息msg会显示出来
                    "data": "",//数据部分
                    "msg": "成功"
                }
    2:删除接口:
        uri:/admin/del
                method:post
        params:
                {
                    "id":1,//域id (必传)
                    "name":"xxxx"//域名称 (必传)
                 }
        返回值:
                {
                    "code": 200,
                    "data": "",
                    "msg": "成功"
                }
       3:列表接口:
             uri:/admin/find
             method:get
             params:
                    {
                        "id":1,//域id 
                        "name":"xxxx",//域名称
                        "ps":1,//每页显示多少(默认是10，最多展示10条每页)
                        "pn":1//页码(默认从1开始，传0后台会处理成1)
                    }
             返回值:
                   {
                       "code": 200,
                       "data": {
                           "count": 1,//总条数
                           "data": [
                               {
                                   "id": 1918826438459392,
                                   "name": "狗蛋",
                                   "status": 1,
                                   "ctime": 1571566729,
                                   "mtime": 1571566729
                               }
                           ]
                       },
                       "msg": "成功"
                   }
       4:删除接口:
           uri:/admin/update
                   method:post
           params:
                   {
                       "id":1,//域id (必传)
                       "name":"xxxx"//域名称 (必传)传name表示修改名称(后端统一一个model，隐藏这里name不修改的时候默认也传)
                        "status": （1或者-1 -1表示禁用）
                    }
           返回值:
                   {
                       "code": 200,
                       "data": "",
                       "msg": "成功"
                   }       