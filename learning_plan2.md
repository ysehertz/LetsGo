# Go Web Development Learning Plan (Native Path)

## Current Status
- **Stage**: 2.2 Middleware Pattern
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
    - **Task**: Refactor code into `cmd/web/`, `internal/handler/`, `internal/model/`. Keep business logic separate from handler.
    - **Key Concept**: Go Modules, Internal packages, Visibility, Package organization principles (e.g., separating handler from model/service).
- [ ] **2.2 Middleware Pattern**
    - **Task**: Write a Logging Middleware that wraps a handler.
    - **Key Concept**: Functions as First-class Citizens, Decorator Pattern (AOP in Go), comparing to Servlet Filters or Spring Interceptors.

### Phase 3: Transition to Frameworks (Gin)
> 目标：理解框架到底帮我们做了什么。
- [ ] (Future) Switch to Gin
- [ ] (Future) Integrate GORM

## Immediate Task Instructions
请引导用户编写一个日志中间件 (Logging Middleware)。
1.  创建一个新目录 `internal/middleware/`。
2.  在该目录中创建一个 `logger.go` 文件。
3.  在 `logger.go` 中，创建一个 `Logging` 函数，它的签名应该是 `func(next http.Handler) http.Handler`。
4.  这个函数返回一个新的 `http.Handler`，该处理器会：
    *   记录下收到的请求的 `Method` 和 `URL`。
    *   记录请求处理所花费的时间。
5.  在 `cmd/web/main.go` 中，使用这个中间件来“包裹”我们现有的 `mux`。
6.  **重点讲解**:
    *   **`http.Handler` 接口**：它只有一个方法 `ServeHTTP(ResponseWriter, *Request)`。`http.ServeMux` 和我们接下来要创建的中间件都实现了这个接口。
    *   **Go 中间件模式**：解释 `func(next http.Handler) http.Handler` 这个函数签名是如何像洋葱一样层层包裹核心业务逻辑的。
    *   **与 Java 的对比**：将它与 Java Servlet 的 `Filter` 或 Spring 的 `HandlerInterceptor` 进行类比。