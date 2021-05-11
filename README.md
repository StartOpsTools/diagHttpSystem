diagHttpSystem
===============

diagHttpSystem 是一个HTTP系统诊断工具, 具备以下功能:

    1. 解析url
        1.1 解析url中域名
    2. 获取域名解析结果
    3. 获取LocalDNS
    4. 获取HTTP访问结果
        4.1 HTTP CODE
        4.2 HTTP Response
    5. 用户端IP
    6. TCP 包体过程 (暂时未实现)
	
start
------

./dist/linux/digHttpSystem  -url http://www.baidu.com  -reportUrl http://hook.startops.com.cn/v1/hook/diag/http

./dist/linux/digHttpSystem

    -url         探测url
    -reportUrl   上报消息url
    
json格式上报数据
--------------

```
error_message      错误消息
url                传入的url
domain             url中解析出的域名
domain_ip_addr     域名解析的IP地址
local_dns          本地DNS
request_body       请求url返回内容
tcp_body           TCP内容
```



