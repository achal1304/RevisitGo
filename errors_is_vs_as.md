In Go, `errors.Is` and `errors.As` are two functions from the `errors` package that help with error handling, particularly when dealing with wrapped errors. They allow you to inspect and interact with errors in different ways.

Here’s a quick breakdown of the differences:

### **1. `errors.Is`**
- **Purpose**: It checks whether a given error is **equal** to a specific target error or if it's wrapped within an error that matches the target error.
- **Use Case**: Useful when you want to check if an error is the same as a particular error or if it has been wrapped with that error.

- **How it works**: 
  - It will compare the given error to the target error, checking if the error is either **directly the target error** or **wrapped within** the target error (using `fmt.Errorf` with `%w` for wrapping).
  - It is commonly used to check for **specific error types** or for **known errors**.

#### **Example**:

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func checkError(err error) {
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Error is ErrNotFound!")
	} else {
		fmt.Println("Error is not ErrNotFound.")
	}
}

func main() {
	// Example with ErrNotFound
	err := fmt.Errorf("wrapped error: %w", ErrNotFound)
	checkError(err)
}
```

**Output**:
```
Error is ErrNotFound!
```

#### **Key Points**:
- `errors.Is` checks if the error is **equal** to or **wrapped by** the target error.
- It is useful for checking **specific error matches**.

---

### **2. `errors.As`**
- **Purpose**: It checks whether an error can be **cast** to a specific type. It is used to **extract a specific error type** from a wrapped error and assign it to a variable of that type.
- **Use Case**: Useful when you want to work with the specific fields or methods of a custom error type or interface that a wrapped error implements.

- **How it works**: 
  - `errors.As` tries to match the **error type** with the target type and, if successful, assigns the underlying error to a variable of the target type.
  - This is useful for extracting **detailed information** from errors when the error type is known or when you need access to custom methods on an error type.

#### **Example**:

```go
package main

import (
	"errors"
	"fmt"
)

// Custom error type
type MyError struct {
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError with code: %d", e.Code)
}

func checkError(err error) {
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("Extracted MyError with code: %d\n", myErr.Code)
	} else {
		fmt.Println("Error is not of type MyError.")
	}
}

func main() {
	// Example with MyError wrapped
	err := fmt.Errorf("wrapped error: %w", &MyError{Code: 404})
	checkError(err)
}
```

**Output**:
```
Extracted MyError with code: 404
```

#### **Key Points**:
- `errors.As` is used to **extract the underlying error type** if it's wrapped in another error.
- It is useful for working with **custom error types** or extracting additional data from a specific error type.

---

### **Summary of Differences**:

| Feature              | `errors.Is`                                             | `errors.As`                                                |
|----------------------|---------------------------------------------------------|------------------------------------------------------------|
| **Purpose**           | Checks if an error is **equal to** or **wrapped by** another error. | Extracts the **underlying error** of a specific type. |
| **Use Case**          | To check if an error matches a specific target error, including wrapped errors. | To **extract** a specific error type from a wrapped error. |
| **How it Works**      | Returns `true` if the error is **equal** to the target error or is wrapped in it. | Returns `true` if the error can be **cast** to the target type and assigns the underlying error. |
| **Typical Use**       | Checking if an error matches a **known error** type (e.g., `ErrNotFound`). | Extracting **custom error types** or working with specific error details. |

Both functions are useful for **error inspection** and **handling**, but they serve slightly different purposes:
- Use **`errors.Is`** for checking error equality or wrapping.
- Use **`errors.As`** when you need to **extract** and work with specific error types or details.