# Go Web Development Learning Plan (Native Path)

## Current Status
- **Stage**: 3.3 Integrate GORM
- **Status**: 🟢 In Progress

## Roadmap

### Phase 1: The Standard Library (No Frameworks)
> 目标：理解 Go 处理 HTTP 的底层逻辑，类似学习 Servlet。

- [x] **1.1 The Basics of net/http**
    - **Task**: Create a `main.go`. Use `http.HandleFunc` to handle root `/` request responding "Hello Native Go". Listen on port 8080.
    - **Key Concept**: `ResponseWriter` vs `HttpServletResponse`, `Request` vs `HttpServletRequest`.
- [x] **1.2 JSON Handling**
    - **Task**: Return a User struct as JSON. Handle POST request to parse JSON body.
    - **Key Concept**: `encoding/json`, Struct Tags (similar to Jackson annotations).
- [x] **1.3 Routing & Mux**
    - **Task**: Implement separate handlers for `GET /users` and `POST /users` without a router framework (using switch case or ServeMux).
    - **Key Concept**: `http.ServeMux`, Method checking.

### Phase 2: Project Structure & Middleware
> 目标：学习如何不依赖框架组织代码结构。

- [x] **2.1 Standard Project Layout**
    - **Task**: Refactor code into `cmd/web/`, `internal/handler`, `internal/model`.
    - **Key Concept**: Go Modules, Internal packages, Visibility, Package organization principles (e.g., separating handler from model/service).
- [x] **2.2 Middleware Pattern**
    - **Task**: Write a Logging Middleware that wraps the handler.
    - **Key Concept**: Functions as First-class Citizens, Decorator Pattern (AOP in Go), comparing to Servlet Filters or Spring Interceptors.

### Phase 3: Transition to Frameworks (Gin)
> 目标：理解框架到底帮我们做了什么。

- [x] **3.1 Introduction to Gin**
    - **Task**: Introduce Gin, initialize a Gin router, and adapt existing handlers (`Home`, `UsersHandler`) to Gin's context. Replace `net/http` server with Gin's run method.
    - **Key Concept**: Gin's `*gin.Context` vs `http.ResponseWriter`/`*http.Request`, Gin's simplified routing, Gin's JSON rendering (`c.JSON`), Gin's request binding (`c.BindJSON`).
- [x] **3.2 Gin Middleware**
    - **Task**: Implement the logging middleware using Gin's middleware system.
    - **Key Concept**: Gin's middleware chaining, comparing to `net/http` middleware and native Go implementation.
- [ ] **3.3 Integrate GORM**
    - **Task**: Integrate GORM for database interaction (e.g., SQLite), create a `repository` layer, and update handlers/services to use GORM.
    - **Key Concept**: ORM vs raw SQL, GORM's models, migrations, CRUD operations.

## Immediate Task Instructions
请引导用户集成 GORM，使用 SQLite 数据库。
1.  **添加 GORM 和 SQLite 驱动依赖**：指导用户运行 `go get gorm.io/gorm gorm.io/driver/sqlite`。
2.  **创建 `internal/database/` 目录**：用于存放数据库连接和初始化逻辑。
3.  **在 `internal/database/database.go` 中**：
    *   定义 `ConnectDatabase()` 函数，初始化 SQLite 数据库连接。
    *   使用 GORM 的 `AutoMigrate()` 自动创建 `User` 表。
    *   返回 `*gorm.DB` 实例。
4.  **修改 `internal/models/user.go`**：
    *   导入 `gorm.io/gorm`。
    *   `User` struct 嵌入 `gorm.Model`，用于获得 GORM 提供的 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` 字段。
5.  **修改 `internal/handlers/handlers.go`**：
    *   修改 `GetUsers` 和 `CreateUser` 函数签名，接收一个 `*gorm.DB` 实例。
    *   在 `GetUsers` 中，使用 `db.Find(&users)` 查询所有用户。
    *   在 `CreateUser` 中，使用 `db.Create(&newUser)` 创建新用户。
6.  **修改 `cmd/web/main.go`**：
    *   导入 `native-web-demo/internal/database`。
    *   在 `main` 函数开头调用 `database.ConnectDatabase()` 获取 `*gorm.DB` 实例。
    *   将 `*gorm.DB` 实例通过闭包或依赖注入的方式传递给 `handlers`。
7.  **重点讲解**：
    *   **ORM 概念**：GORM 作为 Go 语言的 ORM，与 Java Hibernate/JPA 的类比。
    *   **GORM `gorm.Model`**：自动提供的字段。
    *   **数据库连接和初始化**：如何建立连接和自动迁移。
    *   **基本 CRUD 操作**：`db.Find()`, `db.Create()`。
    *   **依赖注入**：如何将 `*gorm.DB` 实例传递给需要它的处理函数。