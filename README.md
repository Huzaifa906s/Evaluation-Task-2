# Wave Pattern Array Rearrangement in Go

## Problem Statement
Given an array of integers and an integer `x` (where `x ≥ 1`), rearrange the array into a **wave pattern** as follows:

### **Rules for Rearrangement:**
1. **Partitioning**: Divide the array into consecutive blocks of `2x+1` elements. (Assume that the array's length is a multiple of `2x+1`)
2. **Wave Pattern in Each Block**:
   - The element at index `x` in each block must be the **maximum**.
   - The elements to the **left** of index `x` must be arranged in **non-decreasing** order.
   - The elements to the **right** of index `x` must be arranged in **non-increasing** order.

## **Example**
### **Input:**
```plaintext
Array: [3, 6, 5, 10, 7, 1, 8, 4, 9, 2, 11, 12, 15, 14, 13]
x = 2
```
- Block size = `2(2) + 1 = 5`
- Array length = `15`, which is a multiple of `5` → ✅ Valid.

### **Step-by-Step Rearrangement:**
1. **Block 1:** `[3, 6, 5, 10, 7]` → `[3, 6, 10, 7, 5]`
2. **Block 2:** `[1, 8, 4, 9, 2]` → `[1, 2, 9, 8, 4]`
3. **Block 3:** `[11, 12, 15, 14, 13]` → `[11, 12, 15, 14, 13]`

### **Output:**
```plaintext
[3, 6, 10, 7, 5, 1, 2, 9, 8, 4, 11, 12, 15, 14, 13]
```

## **Solution Approach**
1. **Partition the array** into blocks of size `2x+1`.
2. **Sort each block** in ascending order.
3. **Move the maximum element** to the middle index `x`.
4. **Ensure the left side** is in **non-decreasing** order.
5. **Ensure the right side** is in **non-increasing** order.

## **Installation and Usage**
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/wave-pattern-go.git
   cd wave-pattern-go
   ```

2. Run the Go program:
   ```sh
   go run main.go
   ```


## **Edge Cases Considered**
✅ If the array length is not a multiple of `2x+1`, an error message is displayed.  
✅ Works for **large arrays** and different values of `x`.  
✅ Handles cases where elements are **already sorted or random**.

## **Contributing**
Feel free to submit pull requests or open issues for improvements.

## **Author**
Muhammad Huzaifa

**Email**: huzaifashamayl90@gmail.com


