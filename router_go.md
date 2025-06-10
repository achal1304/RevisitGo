A **router** in Golang (or in web development generally) is a **component that matches incoming HTTP requests to specific handler functions based on the request’s URL path and method (GET, POST, etc.)**. It acts as the “traffic controller” for your web server, determining how each request should be processed.

### Uses of a Router in Golang

1. **Route Matching**:
   Maps different URL paths (like `/home`, `/api/users`, etc.) to corresponding handler functions.

2. **Parameter Extraction**:
   Extracts variables from URLs (e.g., `/users/{id}` can capture the `id` from the path).

3. **Middleware Support**:
   Allows you to add middleware functions (like authentication, logging, CORS, etc.) that run before or after your main handler.

4. **Cleaner Code Structure**:
   Keeps the code organized, especially for APIs with many endpoints.

### Example: Using the Standard Library

Golang’s standard library has a very basic router via [`http.ServeMux`](https://pkg.go.dev/net/http#ServeMux):

```go
http.HandleFunc("/hello", helloHandler)
http.HandleFunc("/users", usersHandler)
http.ListenAndServe(":8080", nil)
```

But this is limited (no path parameters, weak middleware support), so for bigger projects, you usually use a third-party router.

### Popular Routers in Go

* **[Gorilla Mux](https://github.com/gorilla/mux)** (archived, but still popular)
* **[chi](https://github.com/go-chi/chi)** (very lightweight and modern)
* **[Gin](https://github.com/gin-gonic/gin)** (feature-rich, fast, commonly used for APIs)

### Example with `chi`

```go
import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()
    r.Get("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
        userID := chi.URLParam(r, "userID")
        w.Write([]byte("User ID: " + userID))
    })
    http.ListenAndServe(":8080", r)
}
```

---

## **Summary Table**

| Purpose                | Example                          |
| ---------------------- | -------------------------------- |
| Match route to handler | `/users` → `usersHandler`        |
| Extract parameters     | `/users/123` → `id=123`          |
| Add middleware         | Logging, Auth, CORS              |
| Clean API structure    | Organize all routes in one place |

---

### In Short

A **router in Go** is what maps HTTP requests to functions.
**Use it** to organize your API, handle different endpoints, extract parameters from URLs, and attach middleware.

Let me know if you want a more in-depth code sample or explanation of how routing works under the hood!
