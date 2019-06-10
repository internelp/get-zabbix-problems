# get-zabbix-problems

获取 zabbix 当前问题数量,用于 zcate 发送消息时设定 app 图标角标为 zabbix 当前问题数量.

https://www.qiansw.com/how-to-use-zcate-to-receive-zabbix-alarm-messages.html

![image](4120784843.jpg)

## 使用方法

### 1 自行编译

修改 `main.go` 第 18 行开始为你的 zabbix 信息,然后编译.
```
var (
	// Your zabbix info
	zUser   = "read"
	zPasswd = "read"
	zURL    = "http://www.qiansw.com/api_jsonrpc.php"
)
```

编译完成后执行 `./get-zabbix-problems`,成功返回数字,失败返回错误.

### 2 使用已编译完成的

下载适合您操作系统的二进制文件,然后根据您自己的 zabbix 信息执行下面指令,必须使用命令行参数指定 zabbix 信息:

```
./get-zabbix-problems -u username -p password -url http://your.zabbix.url/api_jsonrpc.php
```

成功返回数字,失败返回错误.
