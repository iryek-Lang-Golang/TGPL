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
- A string value can be written as a string literal.
   - "string가나다'
- Go source files are always encoded in UTF-8
- Go text strings are conventionally interpreted as UTF-8.
- We can include Unicode code points in string literals
- Within a double quoted string literal, escapte sequences that begin with a backslash \ can be used to insert arbitrary byte values into the string.
   - One set of escapes handles ASCII control codes like newline, carriage return, and tab:
```
\a	"alert" or bell
\b	backspace
\f	form feed
\n	newline
\r	carriage return
\t	tab
\v 	vertical tab
\'	single quote (only in the rune literal '\'')
\"	double quote (only within "..." literals)
\\	backslash
```

- Arbitrary bytes can be included in literal strings using hexadecimal or octal escapes.
   - \xhh
   - \ooo
- A raw string literal is written ``...`` using backquotes instead of double quotes.
   - A raw string literal may spread over several lines in the program source.
   - The only processing is that carriage returns are deleted so that the value of the string is the same on all platforms.

### 3.5.2 Unicode

### 3.5.3 UTF-8
- The lexographic byte order equals the Unicode code point order.
  - Sorting UTF-8 works naturally.
- There are no embedded NUL(zero) bytes.
- UTF-8 is the preferred encoding for text strings manipulated by Go programs.
- unicode/utf8 package
- Unicode escape
```
"世界"
UTF-8: "\xe4\xb8\x96\xe7\x95\x8c
16bit: "\u4e16\u754c
32bit: "\U00004e16\U0000754c
```
  - Unicode escapes may also be used in rune literals.
  ```
  '世' '\u4e16' '\U00004e16'
  ```
  - A rune whose value is less than 256 may be written with a single hexadecimal escape, such as 'x41' for 'a'
    - For higher values, a \u or \U escape must be used.
      - UTF-8 encoding '\xe4\xb8\x96' is not a legal rune literal.

- Thanks to the nice properties of UTF-8, many string operations don't require decoding.
  - Test whether one string contains another as a prefix, a suffix, or a substring
      ```go
      // prefix
      func HasPrefix(s, prefix string) bool {

      }

      // suffix
      func HasSuffix(s, suffix string) bool {

      }

      // substring
      func Contains(s, substr string) bool {

      }

      ```
### 3.5.5 Conversions between Strings and Numbers
- 'strconv' package
  - Conversion between numberic values and their string representations
  - strconv.Iota
  ```go
  x := 123
  y := fmt.Sprintf("%d", x)
  fmt.Println(y, strconv.Iota(x))
  ```
  - FomatInt or FormatUint to formant numbers in a different base
  ```go
  fmt.Println(strconv.FormtInt(Int64(x), 2))   // "1111011
  ```
  - The fmt.Printf verbs %d, %d, %o, and %x are often more convenient than format functions with additional information
  ```go
  s := fmt.Sprintf("x=%b", x)   // "x=1111011
  ```
  - Atoi, ParseInt, or ParseUint for unsigned integers
  ```go
  x, err := strconv.Atoi("123")               // x is an int
  x, err := strconv.ParseInt("123", 10, 64)   // base 10, up to 64bit
  ```
    - The third argument of ParseInt gives the size of the integer type that the result must fit into
      - 16 implies int16
      - 0 implies int
        - In any case, the type of the result y is always int64, which you can then convert to a smaller type.
- fmt.Scanf
