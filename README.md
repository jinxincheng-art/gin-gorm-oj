# gin-gorm-oj

语言: GoLang 框架:Gin,Gorm

GORM中文文档:https://learnku.com/docs/gorm/v2

GIN中文文文档:https://www.kancloud.cn/shuangdeyu/gin_book/949411

导入gin,gorm

$ go get -u github.com/gin-gonic/gin

$ go get -u gorm.io/gorm

## 整合Swagger
参考文档: https://github.com/swaggo/gin-swagger

接口访问地址: http://localhost:8080/swagger/index.html#/

写在对应的函数上

```go
// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
```

## 整合JWT

参考文档：https://github.com/dgrijalva/jwt-go

导入依赖

$ go get github.com/dgrijalva/jwt-go

## 邮件发送验证码

参考文档：https://github.com/jordan-wright/email

导入依赖

$ go get github.com/jordan-wright/email

## go-redis
参考文档：https://github.com/go-redis/redis

导入依赖

$ go get github.com/go-redis/redis

## go.uuid

参考文档：https://github.com/satori/go.uuid

导入依赖

$go get github.com/satori/go.uuid


