# 1. 安装
安装 `swag` : 
`go install github.com/swaggo/swag/cmd/swag@latest`

验证: `swag --version`

# 2. 注解语法
注解: API的基本信息
// @title API的标题
// @version 1.0 [API的版本]
// @description API的描述
// @host localhost:8080 [API的主机地址]
// @BasePath /api/v1 [API的基础路径]
// @securityDefinitions.apikey ApiKeyAuth [定义 API Key 认证]
// @in header [API key 放在请求头中]
// @name Authorization [API Key 的名称，常用 Authorization]

注解: 路由和处理函数
// 函数标题
// @Summary 简单描述
// @Description 详细描述
// @Tags hello [路由的标签，用于分组]
// @Accept json [接受的请求格式]
// @Produce json [返回的响应格式]
// @Security ApiKeyAuth  [安全定义]
// @Success 200 {string} string "Hello, world!" [响应的格式和描述]
// @Router /hello [get] [路由的路径和方法]

注解: 请求参数
// @Param name query string true "用户名称"
// @Param age query int false "用户年龄"

注解: 响应模型
// @Success 200 {object} User

注解: 定义模型配合响应
type User struct {
    ID int `json:"id" example:"1"`
    Name string `json:"name" example:"John Doe"`
}

# 3. 生成文档
生成 Swagger 文档 `swag init`
可以使用 `swag init -o ./docs/swagger` 来指定生成的位置
json、yaml文件可以传到其他API文档管理工具中使用部署
 docs.go 
 swagger.json
 swagger.yaml
 
# 4. 在项目中集成 Swagger UI
1. 导入包
`swaggerFiles "github.com/swaggo/files"`
`ginSwagger "github.com/swaggo/gin-swagger"`
`_ "github.com/username/projectname/docs"`
2. 配置路由
`g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))`
3. 访问界面
`http://localhost:8080/swagger/index.html`