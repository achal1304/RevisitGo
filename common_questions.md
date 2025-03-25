### How does indexing improve DB performance? like back in database what happens on memory 
- **Answer**:Indexing is one of the most common and powerful techniques used in databases to improve query performance, especially for **search operations** like `SELECT`, `UPDATE`, and `DELETE`. Here's how it works and what happens in memory when you create an index:

### 1. **What is an Index in a Database?**
An **index** is a data structure that improves the speed of data retrieval operations on a database table at the cost of additional space and slower write operations (like `INSERT`, `UPDATE`, and `DELETE`). The most common types of indexes are:

- **B-tree Indexes**: Balanced tree structures that allow for efficient searching, insertion, and deletion.
- **Hash Indexes**: Efficient for equality comparisons but not for range queries.
- **Full-Text Indexes**: Specialized for text search in large data sets.

### 2. **How Does Indexing Improve Query Performance?**
Without an index, the database has to **scan** every row in a table (a full table scan) to find the data you are requesting. As tables grow in size, this operation becomes slower. **Indexes** speed up this process by allowing the database to locate data more efficiently, generally in logarithmic time (`O(log N)`), rather than linear time (`O(N)`).

- **Faster Searches**: An index provides a pointer to the data rather than scanning the entire table. For example, with a B-tree index, the database can traverse the index structure to locate the relevant data much faster.
- **Efficient Sorting**: Indexes also help in sorting operations. If a query requires ordering by a column that is indexed, the database can simply return the sorted index entries without scanning the whole table.
- **Faster Joins**: When joining tables on indexed columns, the database can use the index to locate matching rows quickly, reducing the need for full table scans.

### 3. **What Happens in Memory?**

When you create an index in a database, it is typically **stored in memory** and on disk in a manner that optimizes access. Here's a breakdown of what happens in memory:

#### **Memory Usage**:
- **In-memory Structures**: Indexes often reside in **RAM** to speed up query performance. Modern database management systems (DBMS) like MySQL, PostgreSQL, or SQL Server use in-memory data structures (such as B-trees or hash tables) to keep the index active and fast to query.
- **Buffer Pool**: Databases maintain a buffer pool (cache) to hold frequently accessed index data in memory. This cache speeds up operations by reducing disk I/O when the index is accessed repeatedly.
  
#### **How the Index Affects Memory**:
- **B-tree Index**: A B-tree is structured as a **balanced tree**, where each node contains a key and a pointer to the next node or data. The keys in the nodes are sorted, allowing for efficient searching. The tree structure is stored in memory as a series of linked nodes, and each key is associated with pointers to the actual data rows.
  
  - **Disk Storage**: The index itself is stored on disk in the same way as a regular table but structured to allow faster lookups. When a query is executed, the DBMS loads the index into memory, which improves performance as the structure is already prepared to quickly locate the desired data.
  
  - **In-memory Operations**: As queries use the index, operations like binary search are performed in memory. For example, when looking up a value, the system navigates down the B-tree structure to find the desired record more quickly.

- **Hash Index**: Hash indexes use a **hashing algorithm** to map values to a hash table. In memory, the DBMS stores the hash table in a way that makes searching for keys and values almost instantaneous.
  - **In-memory Search**: A hash lookup involves converting the search key into a hash value and directly accessing the value in the hash table, which is very fast.
  
- **Memory Usage Impact**: Creating an index consumes additional memory. If your table is large and the index is complex, it will increase memory usage significantly. However, the **trade-off is worth it** in terms of faster query performance, particularly for read-heavy operations.

#### **Disk and Memory Synchronization**:
- **Writing to Disk**: As data in the indexed column changes, the index is also updated. While reading is faster with indexes, writing can become slower because the index needs to be updated every time data is inserted, updated, or deleted in the table. To handle this, databases often rely on **transaction logs** to maintain consistency between memory and disk.
  
- **Data Caching**: If an index is frequently accessed, it's often cached in the DBMS’s memory cache to avoid having to perform I/O operations every time the index is queried. The effectiveness of this caching mechanism depends on the size of the index and the available system memory.

### 4. **Other Considerations**
- **Index Maintenance**: As data changes (insertions, updates, deletions), the index must be updated to reflect those changes. This **adds overhead** to write operations. While reads become faster, writes become slower.
- **Memory Consumption**: Large tables with many indexes can consume a significant amount of memory, and this should be considered when planning your database architecture.
- **Choosing the Right Index**: Indexes improve performance, but **too many indexes** can degrade write performance and take up significant memory. It’s important to balance the use of indexes based on query patterns.

### Summary:
- **Indexing** improves **query performance** by reducing the amount of data scanned during search operations, allowing the DBMS to use memory-efficient data structures (like B-trees or hash tables) to quickly access data.
- **In-memory** indexing structures are cached for faster access, reducing disk I/O, which is critical for performance, especially with large datasets.
- **Write operations** may suffer from overhead due to the maintenance of indexes, but the trade-off is that **read performance** is significantly improved.

By utilizing memory to store and manage indexes, databases can dramatically improve the efficiency of data retrieval operations. However, the benefits should always be weighed against the overhead involved in maintaining and updating indexes.

---

### 1. **What is the correct way to implement authorization? What factors do we need to consider?**

**Answer**:
Authorization is the process of granting or denying access to resources based on the permissions assigned to a user or entity. The correct way to implement authorization involves:
- **Role-Based Access Control (RBAC)**: Define roles and permissions for different users (e.g., Admin, User, Guest). Each role should have clearly defined access levels.
- **Attribute-Based Access Control (ABAC)**: Allow access based on attributes (e.g., department, job title).
- **Least Privilege Principle**: Users and processes should have the minimum level of access needed for their function.
- **Authorization Tokens**: Use tokens (e.g., JWT) to validate permissions for accessing protected resources.
- **Access Control Lists (ACLs)**: Define specific permissions on resources for particular users or roles.

**Factors to Consider**:
- **Granularity of Permissions**: Decide how detailed the permissions are for each resource.
- **Security**: Ensure tokens and credentials are stored and transmitted securely.
- **Scalability**: The authorization system should support changes in roles and permissions easily.
- **Auditing**: Track who accessed what resources and when to provide an audit trail.

---

### 2. **What is the difference between authorization and authentication?**

**Answer**:
- **Authentication**: The process of verifying the identity of a user or system (e.g., via username and password). It answers the question: "Who are you?"
- **Authorization**: The process of granting or denying access to resources based on the authenticated user's permissions. It answers the question: "What can you do?"

In simple terms, **authentication** is about identifying the user, while **authorization** is about granting access to resources based on that identity.

---


### 1. **What is JWT Token and Its Structure?**

**JWT (JSON Web Token)** is a widely used standard for transmitting information between parties as a compact, URL-safe string. JWTs are commonly used in stateless authentication and authorization mechanisms, especially in modern web applications and microservices.

#### **JWT Structure:**

A **JWT** consists of **three parts**, separated by dots (`.`):
1. **Header**
2. **Payload**
3. **Signature**

Each part is base64-encoded.

```
<Header>.<Payload>.<Signature>
```

#### **1. Header**:
The header typically consists of two parts:
- **alg** (algorithm): The algorithm used to sign the token (e.g., `HS256` for HMAC SHA-256, `RS256` for RSA).
- **typ** (type): The type of token, usually `JWT`.

**Example:**

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

This header indicates that the JWT is using HMAC SHA-256 for signing.

#### **2. Payload**:
The payload contains **claims**—statements about an entity (usually the user) and additional data. There are three types of claims:
- **Registered claims**: Predefined claims such as `sub` (subject), `exp` (expiration time), `iat` (issued at), etc.
- **Public claims**: Claims that can be freely defined (e.g., username, role).
- **Private claims**: Claims that are agreed upon by both parties (e.g., user-specific data).

**Example**:
```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022,
  "role": "admin"
}
```

- **sub**: The subject of the token (usually the user ID).
- **name**: Custom claim (user’s name).
- **iat**: The time the token was issued (Unix timestamp).
- **role**: Custom claim indicating the user’s role (e.g., `admin`).

#### **3. Signature**:
The signature is used to verify that the sender of the JWT is who it says it is and to ensure that the message hasn’t been tampered with. It’s created by combining the encoded header, the encoded payload, and a secret key (for symmetric algorithms like `HS256`) or a private key (for asymmetric algorithms like `RS256`).

**For example**, using the **HMAC SHA256 algorithm**, the signature is created as follows:

```
HMACSHA256(
  base64UrlEncode(header) + "." + base64UrlEncode(payload),
  secretKey
)
```

**Example of a complete JWT**:
```
eyJhbGciOiAiSFMyNTYiLCJ0eXAiOiAiSldUIn0.eyJzdWIiOiAiMTIzNDU2Nzg5MCIsICJuYW1lIjogIkpvaG4gRG9lIiwgImlhdCI6IDE1MTYyMzkwMjIsICJyZWQiOiAiYWRtaW4ifQ.I4l2E_RrH6FgLfKyl1mrwWzGFzFmbC5JnlThfVhNl7lPYmO2I4t2ZYrB7e38Ktft
```

#### **JWT Validation in Backend**:
When the backend receives the JWT, it performs the following steps:
1. **Verify the Signature**: The backend re-generates the signature from the received token's header and payload using the secret key or public key and compares it with the signature part of the JWT.
2. **Check Expiration**: If the `exp` claim is present, the backend checks if the token has expired.
3. **Extract Claims**: After verification, the backend can extract and use claims like `sub`, `role`, and others to authorize the user.

---

### 2. **How Are Sessions Used?**

**Sessions** are a server-side method to maintain state between HTTP requests. Since HTTP is a stateless protocol, the session is used to remember the user between requests by storing information on the server (e.g., user ID, roles, preferences) and linking it to a session identifier.

#### **How Does the Server Manage Sessions?**

1. **Session Creation**:
   - When a user logs in successfully, the server creates a session and stores user-related information (like user ID and roles) on the server. The server generates a **unique session ID**.
   - This session ID is sent to the client as a **cookie** (via `Set-Cookie` header in HTTP response).

2. **Session Storage**:
   - **In-memory**: The session data can be stored in server memory (e.g., using `Redis` or an in-memory store).
   - **Database**: In larger systems, session data can be stored in a database, typically in a **sessions** table.

3. **Session ID**:
   - The session ID is sent back to the client in the form of a cookie. Every time the client sends a request, the cookie is included automatically (if the `HttpOnly` flag is set).
   - The server uses the session ID to fetch the corresponding session data.

4. **Session Expiry**:
   - Sessions often have an **expiration time** (e.g., after 30 minutes of inactivity) and are automatically destroyed once the time has passed.

#### **Client’s Involvement with the Server**:
- The **client** stores the session ID (usually in a **cookie**) and sends it with each subsequent request.
- The **server** is responsible for storing the session data and validating the session ID received from the client.
- The session ID links the request to stored user information on the server, allowing the server to "remember" the user’s state (authenticated user, roles, etc.).

#### **Common Issues**:
- **Session Hijacking**: If an attacker can steal the session ID (via XSS or man-in-the-middle attack), they could impersonate the user. To prevent this, use **secure cookies**, **HttpOnly**, and **SameSite** attributes.
- **Scalability**: For larger applications, managing sessions across multiple servers can be challenging. Solutions like **Sticky Sessions** or external session stores (e.g., Redis) are commonly used.

---

### 3. **What Are Cookies?**

**Cookies** are small pieces of data sent from the server and stored in the client's browser. Cookies are used for maintaining state, tracking user sessions, and storing small amounts of data between requests.

#### **Cookie Attributes**:
1. **`name`**: The name of the cookie.
2. **`value`**: The value of the cookie (often a session ID or JWT).
3. **`expires`**: The expiration date of the cookie (optional).
4. **`path`**: The URL path for which the cookie is valid.
5. **`domain`**: The domain for which the cookie is valid.
6. **`secure`**: The cookie will only be sent over HTTPS.
7. **`HttpOnly`**: The cookie can only be accessed by the server, not by JavaScript (prevents XSS attacks).
8. **`SameSite`**: Controls whether the cookie is sent with cross-site requests (helps prevent CSRF attacks).

### **Example**:
```http
Set-Cookie: session_id=abcd1234; Path=/; HttpOnly; Secure; SameSite=Strict;
```

#### **Cookies for Sessions**:
- **Session Cookies**: Store session IDs on the client-side. These cookies are **temporary** and deleted when the browser is closed.
- **Persistent Cookies**: Used to store data for longer periods, like user preferences or login credentials.

---

### Summary of Key Concepts:

| Concept               | Description                                                                 |
|-----------------------|-----------------------------------------------------------------------------|
| **JWT**               | A compact, URL-safe token format used for stateless authentication and authorization. |
| **Session**           | Server-side data store that holds user-specific information (e.g., user ID, roles). |
| **Cookie**            | A small piece of data sent from the server and stored on the client-side, often used for session management. |
| **Authentication**    | Verifying a user’s identity (e.g., login credentials).                       |
| **Authorization**     | Granting or denying access to resources based on user permissions.         |

### Additional Points:
- **JWT vs Session**: JWT is stateless, and the server doesn’t need to store user data. Sessions are stateful, and the server stores the user data.
- **Session Hijacking**: Protect cookies with **Secure** and **HttpOnly** flags and ensure sensitive data is encrypted.
