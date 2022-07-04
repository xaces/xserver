# 说明

## Step1

* 1、中心服务管理车辆信息

* 2、添加车辆时指定工作站，默认所属当前用户所在组织

* 3、工作站启动时从中心服务获取所属车辆

* 4、中心服务对车辆信息增删改查通过mq(topic=stationGuid)通知给所属工作节点（并实现内部管理）

## Step2

* 1、工作站mq推送信息给中心服务

* 2、中心服务websocket推送用户对应设备信息

* 3、中心服务不存储设备status信息，因为不支持组合设备查询

## Step3

* 1、中心服务存储设备alarm、online、event信息

* 2、由于alarm信息量巨大，并且需要合并。合并策略

    * 1）、alarm等信息直接入库 t_devalarm表，并合并alarm信息

    * 2）、t_devlarm表信息超过两天的数据搬运到t_devalarm_history表

    * 3）、查询近两天的数据从t_devalarm表获取

    * 4）、查询超过两天从t_devalarm_history和t_devalarm分别获取
