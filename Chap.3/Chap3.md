# 3. Basic Data Types

## 3.5 Strings

- A string is an immutable sequence of bytes.
   - Strings may contain arbitrary data, including bytes with value 0.
- Text strings are conventionally interpreted as UTF-8 encoded sequences of Unicoe code points (runes).
- The built-in function len() returns the number of bytes (not runes) in a string.
   - The index operation s[i] retrieves the i-th byte of string s, where <= i <len(s)
- The + operator makes a new string by concatenating two strings
- Strings may be compared with comparison operators like == and <; 
   - The comparison is done bye by byte, so the result is the natural lexicographic ordering.
- Since strings are immutable, constructions that try to modify a string's data in place are not allowed:
```go
s := "left foot"
s[0] = 'L' // compile error; cannot assign to s[0]
```

### 3.5.1 String Literals
- Go source files are always encoded in UTF-8
- Go text strings are conventionally interpreted as UTF-8.
- We can include Unicode code points in string literals
- Within a double quoted string literal, escapte sequences that begin with a backslash \ can be used to insert arbitrary byte values into the string.
   - One set of escapes handles ASCII control codes like newline, carriage return, and tab:
```
\a	"alert" or bell
\b	backspace
\f	form feed
```
