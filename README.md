# gpi

&nbsp;&nbsp;&nbsp;&nbsp;Gpi是一个基于gin的api微服务封装
## Information

#### Framework

Gin

#### Folder

> gpi
> * config &nbsp;&nbsp;&nbsp;&nbsp;//配置文件目录 
>  - config.yml &nbsp;&nbsp;&nbsp;&nbsp;//yaml格式配置文件
> * controllers &nbsp;&nbsp;&nbsp;&nbsp;//控制器 
> * entities &nbsp;&nbsp;&nbsp;&nbsp;//数据库实例 
> * libriries &nbsp;&nbsp;&nbsp;&nbsp;//应用库，包括mysql，redis，email等功能的基类封装
>   - config &nbsp;&nbsp;&nbsp;&nbsp;//config获取配置功能库
>   - database &nbsp;&nbsp;&nbsp;&nbsp;//获取数据库功能库
>   - elog &nbsp;&nbsp;&nbsp;&nbsp;//log处理log库
>   - redis &nbsp;&nbsp;&nbsp;&nbsp;//redis链接库
>   - verify &nbsp;&nbsp;&nbsp;&nbsp;//验证功能库
>   - wmail &nbsp;&nbsp;&nbsp;&nbsp;//邮件功能库
>   - apolloCli &nbsp;&nbsp;&nbsp;&nbsp;//连接apollo库
>   - httpReq &nbsp;&nbsp;&nbsp;&nbsp;//发送http请求
> * middlewares &nbsp;&nbsp;&nbsp;&nbsp;//中间件
>   - authentication &nbsp;&nbsp;&nbsp;&nbsp;//route验证中间件
>   - exception &nbsp;&nbsp;&nbsp;&nbsp;//统一的异常捕获中间件
>   - limiter &nbsp;&nbsp;&nbsp;&nbsp;//限流中间件（padding）
>   - timeout &nbsp;&nbsp;&nbsp;&nbsp;//超时返回504中间件（developing）
> * models &nbsp;&nbsp;&nbsp;&nbsp;//数据模型 
> * router &nbsp;&nbsp;&nbsp;&nbsp;//路由 
> * tmp &nbsp;&nbsp;&nbsp;&nbsp;//tmp 
> * vendor &nbsp;&nbsp;&nbsp;&nbsp;//框架及组件目录 
> * main.go &nbsp;&nbsp;&nbsp;&nbsp;//启动文件
## Install
注意在环境变量里加入GPIACTIVE(fat-测试, pro-正式)
## Milestone
* milestone1 (2019-10-26)
  - 由GOROOT模式改为module模式，并针对1.13进行优化处理 


