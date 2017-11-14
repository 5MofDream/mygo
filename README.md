# apollo
It is a  golang framework,Has rpc,http,tcp,queue...many widget!

## 文件命名规范

    全部小写

## 下载依赖

    安装包管理：go get -u github.com/kardianos/govendor

    查看依赖状态 govender status

    下载依赖 goverder sync

## 安装的第三方包如下

    gorm

    rpcx

    gin

    simpleyaml

## 开发约定

    逻辑转发和验证在app/controller实现

    主要业务代码写在app/service目录下

    表结构定义 写在 app/models中

    所有的模型操作写在 app/repository 中

### 从表结构生成model文件

    使用命令 go get github.com/jiazhoulvke/table2struct 获取工具

    生成models的命令为

```
    cd app

    table2struct -db_host 127.0.0.1 -db_name ginmon_db -db_port 3306 -db_user root -db_pwd 123456\
     -int64=true -output ./models -package_name models -tag_gorm=true -tag_json=true  users

```

### DI调用实现

    通过明文调用DI()实现 container的传递




