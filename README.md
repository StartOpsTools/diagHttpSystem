diagHttpSystem
===============

diagHttpSystem 是一个HTTP系统诊断工具, 具备以下功能:

    1. 解析url
        1.1 解析url中域名
    2. 
    
    1. 获取解析信息      --    解析结果
	2. 获取访问结果      --    直接访问测试
	3. 获取绑定访问结果   --    绑定访问测试
	4. 抓包测试TCP过程   --     是否存在某个包被丢，或者被reject (未实现)
	

./dist/linux/digHttpSystem  -url http://www.baidu.com  -reportUrl http://hook.startops.com.cn/v1/hook/diag/http

./dist/linux/digHttpSystem

    -url         探测url
    -reportUrl   上报消息url
    
json 格式上报数据

```
error_message      错误消息
url                传入的url
domain             url中解析出的域名
domain_ip_addr     域名解析的IP地址
local_dns          本地DNS
request_body       请求url返回内容
tcp_body           TCP内容
```



