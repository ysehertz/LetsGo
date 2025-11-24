 # Go Web Development Learning Plan (Native Path)

 ## Current Status
 - **Stage**: 1.2 JSON Handling
 - **Status**: 🟢 In Progress

 ## Roadmap

 ### Phase 1: The Standard Library (No Frameworks)
 > 目标：理解 Go 处理 HTTP 的底层逻辑，类似学习 Servlet。

 - [x] **1.1 The Basics of net/http**
     - **Task**: Create a `main.go`. Use `http.HandleFunc` to handle root `/` request responding "Hello Native Go". Listen on port 8080.
     - **Key Concept**: `ResponseWriter` vs `HttpServletResponse`, `Request` vs `HttpServletRequest`.
 - [ ] **1.2 JSON Handling**
     - **Task**: Return a User struct as JSON. Handle POST request to parse JSON body.
     - **Key Concept**: `encoding/json`, Struct Tags (similar to Jackson annotations).
 - [ ] **1.3 Routing & Mux**
     - **Task**: Implement separate handlers for `GET /users` and `POST /users` without a router framework (using switch case or ServeMux).
     - **Key Concept**: `http.ServeMux`, Method checking.

 ### Phase 2: Project Structure & Middleware
 > 目标：学习如何不依赖框架组织代码结构。

 - [ ] **2.1 Standard Project Layout**
     - **Task**: Refactor code into `cmd/`, `internal/handler`, `internal/model`.
     - **Key Concept**: Go Modules, Internal packages, Visibility.
 - [ ] **2.2 Middleware Pattern**
     - **Task**: Write a Logging Middleware that wraps the handler.
     - **Key Concept**: Functions as First-class Citizens, Decorator Pattern (AOP in Go).

 ### Phase 3: Transition to Frameworks (Gin)
 > 目标：理解框架到底帮我们做了什么。
 - [ ] (Future) Switch to Gin
 - [ ] (Future) Integrate GORM

 ## Immediate Task Instructions
 请引导用户修改 `main.go`。
 1. 创建一个 `User` struct。
 2. 创建一个新的 handler `/user`，当收到 GET 请求时，返回一个 `User` 实例的 JSON 字符串。
 3. **重点讲解**: `encoding/json` 包的 `Marshal` 函数，以及如何设置 `Content-Type` 为 `application/json`。讲解 struct tag (如 `` `json:"name"` ``) 与 Java Jackson 的 `@JsonProperty` 注解的类比。
