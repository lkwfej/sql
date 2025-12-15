1. - # Go 用户认证系统（Gin + GORM + MySQL + JWT）

     一个基于 Go 的用户认证后端示例项目，包含注册/登录、JWT 无状态鉴权、中间件保护路由、用户信息查询等功能。

     ## 功能
     - 用户注册：bcrypt 加密存储密码
     - 用户登录：签发 JWT（HS256）实现无状态认证
     - JWT 鉴权中间件：保护 `/api/*` 路由
     - 用户信息：`GET /api/profile`（从 token 获取 user_id，再从 DB 查询用户信息）

     ## 技术栈
     Go / Gin / GORM / MySQL / JWT / Postman

     ## 接口列表
     | 方法 | 路径         | 说明                   | 是否需要登录 |
     | ---- | ------------ | ---------------------- | ------------ |
     | POST | /register    | 用户注册               | 否           |
     | POST | /login       | 用户登录（返回 token） | 否           |
     | GET  | /api/profile | 获取当前用户信息       | 是           |

     ## 快速启动

     ### 1. 配置 MySQL
     创建数据库：`goland_demo`

     ### 2. 配置环境变量（可选）
     > 不配置也可运行（使用默认值）。建议在真实使用时设置 JWT_SECRET。

     - `MYSQL_DSN`：MySQL 连接串（示例：`user:pass@tcp(127.0.0.1:3306)/goland_demo?charset=utf8mb4&parseTime=True&loc=Local`）
     - `JWT_SECRET`：JWT 签名密钥（示例：`your_secret_key`）
     - `ADDR`：服务监听地址（示例：`:9090`）

     ### 3. 启动
     ```bash
     go mod tidy
     go run main.go
     ```

## Postman 一键测试

1. 导入 `postman_collection.json`
   - Postman -> Import -> 选择文件
2. （推荐）创建环境变量
   - `baseUrl`: `http://localhost:9090`
   - `token`: （空）
3. 按顺序运行
   - `POST {{baseUrl}}/register`
   - `POST {{baseUrl}}/login`（拿到 token）
   - `GET {{baseUrl}}/api/profile`
     - Header：`Authorization: Bearer <token>`

## 返回结构（统一响应）

```
{
  "code": 0,
  "message": "ok",
  "data": {}
}
```

