# ShortUrl

### 短链接生成工具

> 基本实现

使用gin作为web端，web端返回301永久重定向，location header为目标网址
使用MongoDB存储短链接与ID映射的bson数据,主要使用62进制转换ID与短地址


> 默认页面

![image](https://github.com/jiangyd/image/blob/master/shortUrl/%E9%BB%98%E8%AE%A4%E9%A1%B5%E9%9D%A2.jpg)

> 输入目标网址，点击生成短链接

![image](https://github.com/jiangyd/image/blob/master/shortUrl/%E7%94%9F%E6%88%90%E7%9F%AD%E9%93%BE%E6%8E%A5.jpg)

> 访问短链接，自动跳转目标网址

![image](https://github.com/jiangyd/image/blob/master/shortUrl/example.jpg)