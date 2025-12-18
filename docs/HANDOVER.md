# 项目交接总结（Go + Gin + GORM + MySQL + JWT）

## 1. 项目概览
- **技术栈**：Go 1.20+、Gin、GORM、MySQL、JWT。
- **目标**：提供标准用户体系（注册 / 登录 / JWT 认证 / 受保护接口）。
- **当前状态**：功能可运行，已完成注册、登录、JWT 鉴权与受保护接口。项目入口在 `main.go`，默认监听 `:9090`。

## 2. 环境与依赖
- 数据库连接改为读取环境变量（见 `config/config.go`），默认值与仓库内的 `docker-compose.yml` 保持一致：
  - `DB_HOST=127.0.0.1`
  - `DB_PORT=3306`
  - `DB_USER=admin_sql`
  - `DB_PASSWORD=admin_sql`
  - `DB_NAME=goland_demo`
- 启动时 `InitDB` 会带重试逻辑并 AutoMigrate `models.User`（`InitTables` 中的 `Role` 仍为占位）。
- JWT 密钥当前为 `change_this_secret`（见 `utils/jwt.go`）；上线前需替换为安全的值并考虑使用环境变量管理。

## 3. 已有功能
- **注册 `POST /register`**：创建用户并对密码做 bcrypt 加密，避免用户名重复。
- **登录 `POST /login`**：校验凭证、生成 JWT（HS256，24 小时有效，含 `user_id`、`username`、`iat`、`exp`）。
- **鉴权中间件**：从 `Authorization: Bearer <token>` 读取 Token，调用 `utils.ParseToken`，将 `user_id`/`username` 注入 Gin Context，否则返回 401。
- **受保护接口 `GET /api/profile`**：在通过中间件后返回当前用户身份信息。

## 4. 启动与联调
### 使用 docker-compose（一键启动 MySQL）
1. 复制环境变量：`cp .env.example .env`（可按需调整端口/密码）。
2. 启动数据库：`docker compose up -d db`（MySQL 会自动初始化库和账号，无需手动打开 phpstudy）。
3. 启动服务：`go run main.go`（默认端口 9090，重试机制可等待数据库就绪）。

### 直接使用本地 MySQL
1. 确保本地 MySQL 正在运行，并创建数据库 `goland_demo` 与账号 `admin_sql/admin_sql`。
2. 如有不同配置，设置对应的 `DB_*` 环境变量后执行 `go run main.go`。
3. 接口验证示例：
   - 注册：`POST /register`，Body `{ "username":"test1", "password":"123456" }`。
   - 登录：`POST /login` → 复制返回的 Token。
   - 访问受保护接口：`GET /api/profile`，Header `Authorization: Bearer <token>`。

## 5. 已知经验与坑点
- 确认连接的 MySQL 实例与端口（3306 vs 33060）一致，避免连错库。
- `Authorization` 头必须以 `Bearer ` 前缀，且避免多余空格/换行，否则解析失败。
- 数据库重建后需要重新迁移以生成 `users` 表。

## 6. 推荐的后续迭代
1. **完善 Profile**：在 `/api/profile` 中直接查询数据库返回完整用户信息（如创建时间、角色）。
2. **统一返回格式与错误码**：定义标准响应结构，例如 `{ "code": 401, "message": "Token 无效" }`。
3. **角色/权限体系**：设计 `roles`/`user_roles`，在路由层做基于角色的访问控制。
4. **刷新 Token**：引入 `refresh_token` + `access_token` 流程，缩短 Access Token 有效期并支持刷新。
5. **接口文档**：使用 Swagger/OpenAPI 自动生成接口说明，方便联调。
6. **配置化密钥/数据库**：将 JWT 密钥等敏感信息通过环境变量或配置文件管理，避免硬编码。

## 7. 交接要点
- 核心逻辑集中于 `controllers`（请求处理）、`services`（业务逻辑）、`routers`（路由）、`middleware`（JWT 认证）、`utils`（JWT 生成/解析）。
- 如需扩展数据模型，先调整 `models` 并通过 `InitTables`/`AutoMigrate` 迁移。
- 目前缺少自动化测试，建议补充针对注册、登录、鉴权流程的集成测试后再调整代码。
