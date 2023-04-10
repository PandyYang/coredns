## 加载插件生成二进制文件
go generate && go build

## 指定配置文件启动
./coredns -conf Corefile -dns.port 1053

## dig测试
dig @localhost -p 1053 a www.baidu.com