Excellent — let's break this down **cleanly and progressively**.  
You’ll get:

- A clear definition of each **normal form** (1NF, 2NF, 3NF).
- A **step-by-step example**, starting from a flawed table and applying each form **one at a time**.
- You’ll **see the problem**, then **fix it** by applying that normal form.

---

## 🔹 **What is Normalization?**

**Database normalization** is the process of organizing data to:
- Reduce **data redundancy** (repeating data).
- Eliminate **update anomalies**.
- Improve **data integrity**.

Normalization is applied through **normal forms**, each one fixing a specific kind of problem.

---

# ✅ Step-by-Step Normalization Example

### 🔴 Start: **Unnormalized Table**

| OrderID | CustomerName | CustomerPhone | ProductName | ProductPrice | Quantity |
|---------|--------------|----------------|-------------|--------------|----------|
| 101     | Alice         | 9876543210     | Pen         | 10           | 2        |
| 101     | Alice         | 9876543210     | Notebook     | 50           | 1        |
| 102     | Bob           | 9123456789     | Pencil       | 5            | 3        |

---

## 🔶 **1st Normal Form (1NF)**

### ✅ **Definition:**
- **No repeating groups or arrays.**
- Every field must be **atomic** (indivisible).

### ❌ **Problem in our table:**
Already appears atomic — every cell has a single value.  
✅ So, **already in 1NF**.

### 📌 Key takeaway:
1NF mainly checks for issues like:
```text
| Courses       |
|---------------|
| Math, Physics |
```
That’s **not 1NF**, but our table is OK here.

---

## 🔶 **2nd Normal Form (2NF)**

### ✅ **Definition:**
- Must be in **1NF**.
- **No partial dependency**: every non-key column must depend on the **whole** primary key.

### ❌ **Problem:**
Let’s assume `OrderID + ProductName` is the composite primary key.

Then:
- `CustomerName`, `CustomerPhone` depend **only on `OrderID`**, not on the full composite key.
- This is a **partial dependency** → breaks 2NF.

### ✅ **Fix: Split into two tables**

**Customers + Orders Table:**

| OrderID | CustomerName | CustomerPhone |
|---------|--------------|----------------|
| 101     | Alice         | 9876543210     |
| 102     | Bob           | 9123456789     |

**OrderItems Table:**

| OrderID | ProductName | ProductPrice | Quantity |
|---------|-------------|--------------|----------|
| 101     | Pen         | 10           | 2        |
| 101     | Notebook    | 50           | 1        |
| 102     | Pencil      | 5            | 3        |

---

## 🔶 **3rd Normal Form (3NF)**

### ✅ **Definition:**
- Must be in **2NF**.
- **No transitive dependencies**: non-key columns must depend **only on the key**, not on another non-key column.

### ❌ **Problem:**
In `OrderItems`, `ProductPrice` depends on `ProductName`, not on the key (`OrderID + ProductName`).

- If product price changes, we’ll have to update **every row** it appears in — bad!

### ✅ **Fix: Extract `Products` table**

**Products Table:**

| ProductName | ProductPrice |
|-------------|--------------|
| Pen         | 10           |
| Notebook    | 50           |
| Pencil      | 5            |

Now update `OrderItems`:

**OrderItems Table:**

| OrderID | ProductName | Quantity |
|---------|-------------|----------|
| 101     | Pen         | 2        |
| 101     | Notebook    | 1        |
| 102     | Pencil      | 3        |

---

# ✅ Final Schema After 3NF

### 1. `Customers` (or Orders if you keep 1:1)
| OrderID | CustomerName | CustomerPhone |
|---------|--------------|----------------|
| 101     | Alice         | 9876543210     |
| 102     | Bob           | 9123456789     |

### 2. `OrderItems`
| OrderID | ProductName | Quantity |
|---------|-------------|----------|
| 101     | Pen         | 2        |
| 101     | Notebook    | 1        |
| 102     | Pencil      | 3        |

### 3. `Products`
| ProductName | ProductPrice |
|-------------|--------------|
| Pen         | 10           |
| Notebook    | 50           |
| Pencil      | 5            |

---

## 🔁 Summary: Normalization Fixes

| Normal Form | Problem Fixed                                 | Fix Applied                         |
|-------------|------------------------------------------------|-------------------------------------|
| **1NF**     | Repeating values, arrays                       | Atomic columns                      |
| **2NF**     | Partial dependency on part of composite key    | Split to isolate full dependencies  |
| **3NF**     | Transitive dependency on non-key columns       | Move dependent data to new tables   |

---

Let me know if you want SQL `CREATE TABLE` code or want to go to **BCNF (Boyce-Codd Normal Form)** next.