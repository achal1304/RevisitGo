# Go Garbage Collection: Heap vs Stack

## Introduction to Go Garbage Collection
Go uses a garbage collector (GC) to manage memory automatically. The GC frees up memory that is no longer in use, reducing the need for manual memory management. Go's GC operates on two main areas of memory: **heap** and **stack**.

## Stack vs Heap in Go

- **Stack**: The stack is a region of memory that stores local variables for functions and is managed in a Last In, First Out (LIFO) manner. Variables that are declared inside a function or method live on the stack.
- **Heap**: The heap is used for dynamically allocated memory. Variables that are allocated using `new()` or `make()` and pointers to large data structures are stored on the heap.

### Garbage Collection on the Stack

- **Function Variables**: Local variables within functions are typically allocated on the stack. When a function exits, the memory for these variables is automatically reclaimed (no garbage collection needed). The stack frame is discarded when the function returns, so there’s no need for the GC to manage stack memory.
- **Global Variables**: Global variables are not directly managed by the GC because they live for the entire lifetime of the program. However, the GC will track whether they are referenced or not. If a global variable is no longer referenced, it can be cleaned up manually or via a tool like `go tool gc`.

### Garbage Collection on the Heap

- **Dynamic Variables**: Variables that are created dynamically (using `new()` or `make()`) are allocated on the heap. The garbage collector tracks these variables and frees them when they are no longer reachable (i.e., no references to them remain). For example, when a struct or slice is created dynamically, the memory is managed by the garbage collector.
- **Pointers**: A pointer itself is usually allocated on the stack, but it may point to memory on the heap (e.g., if it points to dynamically allocated data). The garbage collector tracks these heap-allocated objects, ensuring that memory is freed when the pointer is no longer pointing to valid memory or when the object it points to becomes unreachable.

## Garbage Collection Process for Different Variables

### 1. Global Variables
- Global variables are not automatically garbage collected. They persist throughout the program's life. The GC does not manage them because they are considered to be "reachable" for the duration of the program. However, if global variables are pointers to heap-allocated objects, those objects will be managed by the GC.

### 2. Function Variables
- Function variables (local variables) are stored on the stack. Once the function returns, these variables are no longer needed, and the stack frame is discarded. No garbage collection is needed for these variables as they are automatically cleaned up when the function exits.

### 3. Dynamic Variables (using `new()` or `make()`)
- Memory allocated for dynamic variables (such as slices, maps, and structs created with `new()` or `make()`) resides on the heap. The GC keeps track of these objects and reclaims the memory when they are no longer reachable.

### 4. Pointers
- Pointers themselves are stored on the stack. However, if a pointer points to memory on the heap (e.g., a pointer to a dynamically allocated struct), the GC will track that heap memory. If no other references exist to the memory being pointed to, it will eventually be marked for garbage collection.
  
  Example:
  ```go
  type MyStruct struct {
      value int
  }

  func example() {
      p := &MyStruct{value: 10}  // 'p' is a pointer stored on the stack, MyStruct is on the heap.
      // The MyStruct instance will be garbage collected once 'p' goes out of scope or is no longer referenced.
  }
