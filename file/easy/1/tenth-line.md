# 195. Tenth Line

**Difficulty:** Easy

## Problem Description

Given a text file `file.txt`, print just the **10th line** of the file.

If the file has **fewer than 10 lines**, print **nothing**.

---

## Example

**Input (file.txt):**
```
Line 1
Line 2
Line 3
Line 4
Line 5
Line 6
Line 7
Line 8
Line 9
Line 10
Line 11
```
**Expected output:**
```
Line 10
```

---

## Bash Solutions

### 1) `sed` (simple and robust)
```bash
# Prints only the 10th line (does nothing if fewer than 10 lines)
sed -n '10p' file.txt
```

### 2) `awk`
```bash
# Print the record (line) whose number is 10
awk 'NR==10{print; exit}' file.txt
```

### 3) `tail` + `head`
```bash
# Get last 10 lines, then take the first of those (fails quietly if <10 lines)
tail -n +10 file.txt | head -n 1
```

> Note: On LeetCode, your submission should be a **single bash file** named `solution.sh` that reads `file.txt` from the current directory.

---

## Edge Cases & Notes
- If the file has fewer than 10 lines, the commands above **produce no output** (as required).
- `sed -n '10p'` and `awk 'NR==10{print; exit}'` stop scanning early and are efficient for large files.
- Ensure the script has executable permission: `chmod +x solution.sh`.

---

## Example `solution.sh` (using `sed`)
```bash
#!/usr/bin/env bash
sed -n '10p' file.txt
```
