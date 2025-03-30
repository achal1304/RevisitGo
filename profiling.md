### 🔍 What is **Profiling**?

**Profiling** is the process of **measuring the performance** of a program to identify:
- **CPU usage**: Which functions consume the most processing time.
- **Memory usage**: Where memory is being allocated (heap, stack).
- **Goroutine usage**: How many goroutines are running, if any are stuck or leaking.
- **Blocking operations**: Where goroutines are blocked (channels, mutexes, etc).

> 📌 Profiling helps **optimize performance**, debug **bottlenecks**, and fix **resource leaks**.

---

### 🛠️ How to Do Profiling in **Go**

Go provides built-in support for profiling via the [`runtime/pprof`](https://pkg.go.dev/runtime/pprof) and [`net/http/pprof`](https://pkg.go.dev/net/http/pprof) packages.

---

## ✅ 1. **CPU Profiling**

### 📄 Example: Manual CPU Profiling

```go
package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Your performance-heavy function
	work()
}

func work() {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i
	}
}
```

### 🧪 Run the app:

```bash
go run main.go
```

This generates `cpu.prof`.

### 🔍 Analyze with:

```bash
go tool pprof cpu.prof
(pprof) top         # shows most time-consuming functions
(pprof) list work   # shows annotated source
```

---

## ✅ 2. **Memory (Heap) Profiling**

```go
package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Allocate memory
	_ = make([]int, 1e6)

	pprof.WriteHeapProfile(f) // Write current memory usage
}
```

### Analyze with:

```bash
go tool pprof mem.prof
```

---

## ✅ 3. **Live Profiling with `net/http/pprof`**

Best for **long-running services or servers** (like web APIs).

### Example:

```go
package main

import (
	"net/http"
	_ "net/http/pprof" // Import for side effect
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	select {} // Keep program alive
}
```

Now visit:
```
http://localhost:6060/debug/pprof/
```

You can:
- View profiles like `/heap`, `/goroutine`, `/cpu`
- Download CPU profiles for `pprof`

### Run & Record CPU profile for 30 seconds:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

---

## ✅ 4. **Other Profiles**

| Profile        | Description                                  |
|----------------|----------------------------------------------|
| `/debug/pprof/goroutine` | Shows all goroutines and stack traces |
| `/debug/pprof/heap`      | Shows heap allocations               |
| `/debug/pprof/threadcreate` | Thread creation count             |
| `/debug/pprof/block`     | Blocking on channels, mutex, etc.    |

You can also import and visualize profiles in [**https://github.com/google/pprof**](https://github.com/google/pprof) or in the **GoLand** IDE.

---

## ✅ 5. **Useful Flags for Testing**

You can also profile during testing:

```bash
go test -bench . -cpuprofile=cpu.out -memprofile=mem.out
go tool pprof cpu.out
```

---

### ✅ Summary

| Profiling Type | Tool/API               | Use Case                             |
|----------------|------------------------|--------------------------------------|
| CPU            | `pprof.StartCPUProfile`| Detect CPU-heavy code paths          |
| Memory         | `pprof.WriteHeapProfile`| Analyze heap allocations             |
| Live Server    | `net/http/pprof`       | Web-accessible profiling             |
| Goroutines     | `/debug/pprof/goroutine` | Debug leaks or stuck goroutines     |
| Benchmarking   | `go test -bench`       | Measure performance & allocations    |

---

### ✅ Best Practices

- **Profile in realistic environments** (not synthetic or dev-only data).
- Use `pprof` to find hot paths.
- Combine profiling with **benchmarking** to confirm improvements.
- **Be cautious with premature optimization** – profile first, optimize next.

---

Let me know if you want to visualize profile output in HTML or generate flame graphs! 🔥