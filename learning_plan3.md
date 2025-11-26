# Go Web Development Learning Plan (Native Path)

## Current Status
- **Stage**: 3.2 Gin Middleware
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
    - **Key Concept**: Gin's `*gin.Context` vs `http.ResponseWriter`/`*http.Request`, Gin's simplified routing, Gin's JSON rendering (`c.JSON`), Gin's request binding (`c.dJSON`).
- [ ] **3.2 Gin Middleware**
    - **Task**: Implement the logging middleware using Gin's middleware system.
    - **Key Concept**: Gin's middleware chaining, comparing to `net/http` middleware and native Go implementation.
- [ ] **3.3 Integrate GORM**
    - **Task**: Integrate GORM for database interaction (e.g., SQLite), create a `repository` layer, and update handlers/services to use GORM.
    - **Key Concept**: ORM vs raw SQL, GORM's models, migrations, CRUD operations.

## Immediate Task Instructions
请引导用户重新实现我们之前的日志中间件，但这次要使用 Gin 的方式。
1.  **修改 `internal/middleware/logger.go`**：
    *   将 `Logging` 函数的签名改为 `gin.HandlerFunc` (即 `func(c *gin.Context)`)。
    *   在函数内部，仍然记录请求的 `Method` 和 `URL`。
    *   使用 `c.Next()` 将控制权传递给链中的下一个中间件或最终的处理函数。
    *   在 `c.Next()` 调用之后，记录请求处理所花费的时间。
2.  **修改 `cmd/web/main.go`**：
    *   导入 `native-web-demo/internal/middleware` 包。
    *   使用 `router.Use(middleware.Logging())` 将自定义的日志中间件添加到 Gin 路由器中。
3.  **重点讲解**：
    *   Gin 中间件的函数签名 (`gin.HandlerFunc`)。
    *   `c.Next()` 的作用。
    *   Gin 如何通过 `router.Use()` 应用中间件。
    *   将 Gin 的中间件与 `net/http` 原生中间件进行对比。