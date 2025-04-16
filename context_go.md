In Go, the **`context`** package provides a way to manage and control the execution of processes, particularly in the context of **concurrency**. It allows you to manage **timeouts**, **cancellations**, and **deadlines** for processes that are running concurrently (e.g., goroutines). 

The `context` package is particularly useful in handling long-running operations, such as HTTP requests, database queries, or background tasks, that may need to be canceled, timed out, or have their scope limited.

### Key Concepts of `context`:

1. **Cancellation**: The ability to signal to a goroutine or process that it should stop its execution, typically when it's no longer needed (e.g., the user cancels an HTTP request).
2. **Timeouts**: Specify how long a process should be allowed to run before it times out.
3. **Deadlines**: Specify an absolute time after which a process should be aborted.
4. **Passing Data**: A context can carry values, which can be useful for passing request-scoped data (like authentication tokens) to different parts of a program.

### Why is `context` Used?

- **Managing Goroutines**: Goroutines are lightweight and can run independently, but sometimes you may want to cancel them or stop them after a certain period. The `context` package is helpful in controlling goroutines in these scenarios.
- **Request Lifetimes**: In HTTP servers or client applications, you often need to ensure that the lifetime of certain tasks (e.g., database queries, remote API calls) is tied to the lifetime of the request. If the request is canceled, all associated tasks should also be canceled.
- **Graceful Shutdown**: Context helps ensure that resources like database connections, file handlers, or network connections are properly released when a task is canceled or a timeout occurs.

### Common Use Cases:
- **HTTP requests**: Cancel an HTTP request after a timeout.
- **Database queries**: Cancel a query if it takes too long.
- **Worker tasks**: Stop worker goroutines after a signal (e.g., shutdown).
- **Distributed systems**: Propagate cancellation and deadlines across multiple services.

### Example:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure that the cancel function is called

	// Simulate a long-running task
	go longRunningTask(ctx)

	// Wait for the task to complete or the timeout to occur
	select {
	case <-ctx.Done():
		// If the context is done, print the error message
		fmt.Println("Operation timed out:", ctx.Err())
	}
}

// longRunningTask simulates a task that listens for context cancellation or timeout
func longRunningTask(ctx context.Context) {
	// Simulate a task that takes 3 seconds to complete
	select {
	case <-time.After(3 * time.Second): // Simulating a long task
		fmt.Println("Task completed")
	case <-ctx.Done(): // If context is canceled or times out
		fmt.Println("Task canceled:", ctx.Err())
	}
}
```

### Explanation:

1. **Creating a Context**: We use `context.WithTimeout` to create a context that automatically cancels after 2 seconds.
   - `context.Background()` creates a root context.
   - `WithTimeout` adds a timeout to that root context.
   - `cancel()` is used to manually cancel the context when it's no longer needed, although in this example, it is primarily used to handle cleanup.
   
2. **Goroutine**: A goroutine (`longRunningTask`) simulates a long-running task.
   - If the task completes before the timeout, it prints `Task completed`.
   - If the context is done (either canceled or timed out), the task prints `Task canceled` with the error.

3. **`select` Statement**: It waits for either the timeout or the completion of the task. If the context is done before the task completes, it handles the cancellation.

### Types of Contexts:

1. **`context.Background()`**:
   - The **empty context**. This is typically used as the starting point for requests, such as the root context for an HTTP request.
   
2. **`context.TODO()`**:
   - Similar to `context.Background()`, but used when you are not sure which context to use yet. It's a placeholder until the context is determined.
   
3. **`context.WithCancel(parent context.Context)`**:
   - Creates a new context derived from the parent context. It allows you to cancel the derived context, and when canceled, the cancellation is propagated to the parent context.
   
4. **`context.WithTimeout(parent context.Context, timeout time.Duration)`**:
   - Similar to `WithCancel`, but adds a timeout. The context will be automatically canceled after the specified duration.
   
5. **`context.WithDeadline(parent context.Context, deadline time.Time)`**:
   - Similar to `WithTimeout`, but allows you to specify an absolute deadline (a specific time after which the context will be canceled).

6. **`context.WithValue(parent context.Context, key, value interface{})`**:
   - Used to attach data to the context. The data can be retrieved later in the same context's lifecycle (e.g., to pass request-specific information like user authentication).

### Example with Context Value:

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// Create a context with a value (userID in this case)
	ctx := context.WithValue(context.Background(), "userID", 123)

	// Pass the context to a function
	printUserID(ctx)
}

func printUserID(ctx context.Context) {
	// Retrieve the value associated with the key "userID" from the context
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("No user ID found")
	}
}
```

### Explanation:
- We create a context with a value (`userID = 123`), and then pass it to the `printUserID` function.
- Inside the function, we retrieve the value using `ctx.Value("userID")`.

### Summary of `context`:

- **Purpose**: Manages cancellations, timeouts, and request-specific values in concurrent Go programs.
- **Key Functions**:
  - `context.WithCancel` — Create a cancellable context.
  - `context.WithTimeout` — Create a context with a timeout.
  - `context.WithDeadline` — Create a context with an absolute deadline.
  - `context.WithValue` — Attach data to the context.
- **Use Cases**:
  - **HTTP Request Handling**: Tie the lifetime of requests and their associated tasks (like database queries) to a context.
  - **Goroutine Management**: Control the lifetime of goroutines with cancellation signals.
  - **Distributed Systems**: Propagate cancellation signals across different parts of a system.

The `context` package provides powerful tools to manage goroutines, and its use is essential for handling timeouts, cancellations, and request lifetimes effectively in concurrent Go applications.