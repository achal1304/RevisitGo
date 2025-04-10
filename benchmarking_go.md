In Go, benchmarking is done using the **testing** package, which includes built-in support for benchmarks. This allows you to measure the performance of your code and identify bottlenecks. Benchmarking helps determine how long a particular piece of code takes to execute under different conditions.

### **Steps to Perform Benchmarking in Go**

1. **Import the Testing Package**:
   The `testing` package provides the `Benchmark` function, which is used to measure the performance of code.

2. **Write a Benchmark Function**:
   Benchmark functions should follow the naming convention `BenchmarkXxx`, where `Xxx` is the function being tested. A benchmark function accepts a `testing.B` type as an argument.

3. **Run the Benchmark**:
   You run benchmarks using the `go test` command, but with the `-bench` flag to only execute benchmarks.

### **Example of Benchmarking in Go**

#### 1. **Create a Benchmark Function**

Here's how you can write a benchmark for a simple function like calculating the factorial of a number.

```go
package main

import (
	"fmt"
	"testing"
)

// Function to benchmark (Factorial)
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Benchmark for the Factorial function
func BenchmarkFactorial(b *testing.B) {
	// Run the factorial function b.N times for benchmarking
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

func main() {
	// Running a single test, normally you'd use "go test"
	fmt.Println("Benchmarking completed.")
}
```

#### 2. **Explanation of Code**:
- `Factorial`: A simple recursive function that calculates the factorial of a number.
- `BenchmarkFactorial`: The benchmark function uses the `testing.B` object and performs the factorial computation `b.N` times, where `b.N` is the number of iterations Go automatically chooses to run the benchmark based on the execution time.

#### 3. **Running the Benchmark**:
To run the benchmark, you would use the following command in the terminal:

```bash
go test -bench .
```

- The `-bench` flag runs all benchmarks in the current package (denoted by `.`).
- By default, `go test -bench` runs each benchmark multiple times and reports the result in terms of time taken for each iteration.

**Output**:
```bash
goos: linux
goarch: amd64
pkg: your-package-name
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkFactorial-12     5000000    238 ns/op
PASS
ok      your-package-name  1.252s
```

The output shows:
- **BenchmarkFactorial-12**: This indicates that the benchmark was run on a 12-core CPU (this will vary depending on your machine).
- **5000000**: The number of iterations that were run.
- **238 ns/op**: The time taken for each operation (in nanoseconds per operation).

#### 4. **Understanding Benchmark Output**:
- **ns/op**: This stands for "nanoseconds per operation," representing how long each individual operation took, on average, during the benchmark.
- The `go test` output gives you a statistical summary after running the benchmark.

---

### **Benchmarking with Setup and Teardown**

If your benchmark involves any setup or teardown steps (such as preparing data or resources before benchmarking and cleaning them up afterward), you can use the `testing.B` methods `StartTimer`, `StopTimer`, `ResetTimer`, and `SetBytes` to manage the time measurement more precisely.

#### **Example**: Benchmarking with Setup/Teardown

```go
package main

import (
	"fmt"
	"testing"
)

var data []int

// Setup function for preparing data
func setup() {
	data = make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
}

// Function to benchmark (Summing elements in a slice)
func Sum(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

// Benchmark with Setup
func BenchmarkSum(b *testing.B) {
	// Setup: prepare data
	setup()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		Sum(data)
	}
}

func main() {
	// Running the benchmark test
	fmt.Println("Benchmarking completed.")
}
```

### **Key Methods**:
- `b.StartTimer()`: Starts the timer for the benchmark.
- `b.StopTimer()`: Stops the timer.
- `b.ResetTimer()`: Resets the timer (useful for separating setup and actual benchmark time).
- `b.SetBytes(n int)`: Records the number of bytes processed during the benchmark, which is useful for benchmarks that involve processing large data sets.

---

### **Advanced Benchmarking with Parallel Tests**

Go also allows you to run benchmarks in parallel using `b.RunParallel` for testing how code performs when run concurrently. Here’s an example of a parallel benchmark:

```go
package main

import (
	"fmt"
	"testing"
)

func BenchmarkParallelExample(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Code to benchmark (sum elements in a slice)
			_ = Sum([]int{1, 2, 3, 4, 5})
		}
	})
}

func Sum(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func main() {
	// Running the benchmark test
	fmt.Println("Benchmarking completed.")
}
```

This allows the benchmark to run multiple iterations in parallel, simulating more real-world conditions where multiple workers may be operating concurrently.

---

### **Benchmarking Tips**:
- **Avoid measuring setup time**: Ensure that setup time (data preparation) is excluded from the benchmarking process.
- **Use realistic workloads**: Make sure your benchmark reflects real-world use cases for accurate results.
- **Avoid small iterations**: Running very small benchmarks might lead to inaccurate results due to micro-optimizations. Ensure enough iterations are run to get stable results.

### **Conclusion**:
Benchmarking in Go is a powerful tool to measure the performance of your code. The `testing` package provides an easy way to run and analyze benchmarks. Use this in combination with proper setup, teardown, and parallelism to get accurate performance data. This will allow you to identify bottlenecks and make your code more efficient.



### **Running Benchmarks in Go**

To run the benchmarks in Go, you use the `go test` command with the `-bench` flag. Here's how you can run benchmarks in Go and inspect performance.

#### **Running Benchmarks**
1. **Basic Benchmark Command**:
   You can run the benchmarks in your Go tests by executing the following command:
   
   ```bash
   go test -bench .
   ```

   This command will run all the benchmark functions in the current package. The `.` means "current directory", and you can use the package name as well to run benchmarks for a specific package.

2. **Run Specific Benchmark Function**:
   If you have multiple benchmark functions and want to run a specific one, you can do so by specifying the benchmark name:
   
   ```bash
   go test -bench BenchmarkSum
   ```

   This will only run the `BenchmarkSum` function.

3. **Benchmark with Memory Allocation Details**:
   To check the memory allocation in addition to the benchmark results, use the `-benchmem` flag:
   
   ```bash
   go test -bench . -benchmem
   ```

   This will output memory allocations (in bytes) and allocations per operation (allocs/op).

#### **Useful Benchmarking Commands**

- **Running Benchmarks with Multiple Iterations**:
   By default, Go runs the benchmark multiple times (based on the performance of the test). If you want to control the number of iterations, you can use the `-count` flag:
   
   ```bash
   go test -bench . -count 10
   ```

   This will run the benchmark 10 times.

- **Run Benchmarks for a Specific Subset of Tests**:
   If you want to run only specific benchmarks, you can use a regular expression with the `-bench` flag:
   
   ```bash
   go test -bench "Benchmark.*Sort"
   ```

   This will run all benchmarks whose name contains `Benchmark` and `Sort`.

---

### **CPU Profiling in Go**

To get a **CPU profile**, Go provides a way to record CPU usage and analyze how much time your program spends in each part of the code. The CPU profile helps in identifying bottlenecks.

1. **Enabling CPU Profiling**:
   You can start CPU profiling by importing the `pprof` package and using the `StartCPUProfile` and `StopCPUProfile` methods in your test code.

   ```go
   package main

   import (
       "fmt"
       "os"
       "runtime/pprof"
       "testing"
   )

   // Function to start CPU profiling
   func BenchmarkWithCPUProfile(b *testing.B) {
       file, err := os.Create("cpu_profile.prof")
       if err != nil {
           b.Fatal("Could not create CPU profile: ", err)
       }
       defer file.Close()

       pprof.StartCPUProfile(file)  // Start CPU profiling
       defer pprof.StopCPUProfile() // Stop CPU profiling

       // Run the benchmark logic
       for i := 0; i < b.N; i++ {
           // Some function to benchmark
           _ = Factorial(10)
       }
   }

   func Factorial(n int) int {
       if n == 0 {
           return 1
       }
       return n * Factorial(n-1)
   }
   ```

2. **Running the Benchmark with Profiling**:
   Run the benchmark as usual:
   
   ```bash
   go test -bench BenchmarkWithCPUProfile
   ```

   This will generate a CPU profile in the file `cpu_profile.prof`.

3. **Analyzing the CPU Profile**:
   After running the benchmark with the CPU profile, you can analyze the profile using the `go tool pprof` command:

   ```bash
   go tool pprof cpu_profile.prof
   ```

   This will start an interactive shell that allows you to examine the profile and identify bottlenecks in the code.

---

### **Memory Profiling in Go**

Memory profiling is another useful way to understand how your program is using memory.

1. **Enabling Memory Profiling**:
   You can capture a memory profile in a similar way to the CPU profile. Use `pprof` to capture memory allocations during the benchmark.

   ```go
   package main

   import (
       "fmt"
       "os"
       "runtime/pprof"
       "testing"
   )

   // Function to start Memory Profiling
   func BenchmarkWithMemoryProfile(b *testing.B) {
       file, err := os.Create("mem_profile.prof")
       if err != nil {
           b.Fatal("Could not create memory profile: ", err)
       }
       defer file.Close()

       // Start memory profiling
       pprof.WriteHeapProfile(file)

       // Run the benchmark logic
       for i := 0; i < b.N; i++ {
           _ = Factorial(10)
       }
   }
   ```

2. **Running the Benchmark**:
   You can run the benchmark and generate a memory profile using:

   ```bash
   go test -bench BenchmarkWithMemoryProfile
   ```

   This will generate a memory profile in the `mem_profile.prof` file.

3. **Analyzing the Memory Profile**:
   After running the benchmark with the memory profile, you can use `go tool pprof` to analyze memory usage:

   ```bash
   go tool pprof mem_profile.prof
   ```

---

### **Other Useful Profiling Tools**

- **Block Profile**: This profile helps identify goroutine blocking times.
- **Goroutine Profile**: Helps identify issues with goroutine lifecycles.
- **Threadcreate Profile**: Identifies thread creation activities.

To enable block profiling, for example, you can use:

```go
pprof.Lookup("block").WriteTo(file, 0)
```

### **Combining All Profiles**

You can also combine CPU, memory, and goroutine profiles in a single test:

```go
pprof.StartCPUProfile(cpuFile)
pprof.WriteHeapProfile(memFile)
```

Then, analyze them with respective `go tool pprof` commands for each profile type.

---

### **Summary of Go Benchmarking and Profiling Commands:**

| Command                                    | Description                                                           |
|--------------------------------------------|-----------------------------------------------------------------------|
| `go test -bench .`                          | Run all benchmarks in the current package.                            |
| `go test -bench BenchmarkName`             | Run a specific benchmark.                                             |
| `go test -bench . -benchmem`               | Run benchmarks and also measure memory allocation per operation.      |
| `go test -bench . -count 10`               | Run benchmarks 10 times.                                              |
| `go test -cpuprofile=cpu.prof`             | Enable CPU profiling and write to `cpu.prof`.                         |
| `go test -memprofile=mem.prof`             | Enable memory profiling and write to `mem.prof`.                      |
| `go tool pprof cpu.prof`                   | Analyze the CPU profile using the `pprof` tool.                       |
| `go tool pprof mem.prof`                   | Analyze the memory profile using the `pprof` tool.                    |
| `go test -bench . -v`                      | Run benchmarks with verbose output.                                   |

---

### Conclusion:
- **Benchmarking** in Go is a powerful way to measure the performance of your functions. The `testing` package provides built-in support for running benchmarks.
- Use **CPU profiling** to analyze performance bottlenecks and **memory profiling** to identify excessive memory usage.
- Profiling tools like **pprof** are very helpful in identifying areas where your code might need optimization, allowing you to focus on performance improvements.

By using these tools, you can continuously monitor and improve the performance of your Go applications.
---

To know the number of goroutines launched during benchmarking or profiling in Go, you can use the `runtime.NumGoroutine()` function, which returns the number of goroutines that currently exist. This function can be helpful in understanding how many goroutines are active at a specific point in your program.

Here is an example of how to use `runtime.NumGoroutine()` during a benchmark or profiling:

### Example:

```go
package main

import (
	"fmt"
	"runtime"
	"testing"
)

func ExampleGoroutineCount() {
	// Start by checking the initial number of goroutines
	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("Initial number of goroutines: %d\n", initialGoroutines)

	// Example goroutines (simulated work)
	go func() {
		// Simulate work
	}()
	go func() {
		// Simulate work
	}()

	// Check number of goroutines after launching new ones
	newGoroutines := runtime.NumGoroutine()
	fmt.Printf("New number of goroutines: %d\n", newGoroutines)
}

func BenchmarkGoroutineCount(b *testing.B) {
	// Record initial goroutine count
	initialGoroutines := runtime.NumGoroutine()

	b.ResetTimer() // Reset timer so setup code isn't included in the benchmark

	// Example benchmark code that launches goroutines
	for i := 0; i < b.N; i++ {
		go func() {
			// Simulate work
		}()
	}

	// Measure goroutine count after the benchmark execution
	finalGoroutines := runtime.NumGoroutine()
	fmt.Printf("Initial Goroutines: %d, Final Goroutines: %d\n", initialGoroutines, finalGoroutines)
}
```

### Steps:
1. Use `runtime.NumGoroutine()` to get the initial number of goroutines.
2. Launch new goroutines (e.g., using `go` keyword).
3. Use `runtime.NumGoroutine()` again to check how many goroutines are active after launching them.
4. In a benchmark function (`BenchmarkGoroutineCount`), compare the number of goroutines before and after the benchmark.

When running the benchmarks, Go will print the results, including the number of active goroutines, as shown in the `fmt.Printf` statements.

This approach is useful during benchmarking or profiling to help you track the goroutine count and understand the goroutine behavior during your test or workload execution.

**GOOD READ** - https://k21academy.com/docker-kubernetes/docker-vs-virtual-machine/