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