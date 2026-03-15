# ASCII Art Color (Go)

A command-line program written in **Go** that converts text into **ASCII art** and allows you to **color specific parts of the text** in the terminal.

The program reads characters from a banner file and prints stylized ASCII text using **ANSI color codes**.

---

# Features

* Convert text into ASCII art
* Apply colors to the entire text
* Apply colors to specific substrings
* Supports multi-line input (`\n`)
* Uses ANSI terminal colors
* Lightweight CLI application

---

# Project Structure

```text
.
├── main.go
├── standard.txt
└── README.md
```

### standard.txt

Contains ASCII representations of printable characters used to render the output.

Each character in the file occupies **8 lines** of ASCII art.

---

# Requirements

* Go **1.18 or later**
* A terminal that supports **ANSI color codes**

Check Go installation:

```bash
go version
```

---

# Usage

```bash
go run . [OPTION] [STRING]
```

or

```bash
go run . --color=<color> [SUBSTRING] [STRING]
```

---

# Parameters

| Parameter         | Description                  |
| ----------------- | ---------------------------- |
| `--color=<color>` | Color to apply               |
| `SUBSTRING`       | Optional substring to color  |
| `STRING`          | Text to convert to ASCII art |

---

# Supported Colors

The program supports the following colors:

* black
* red
* green
* yellow
* blue
* magenta
* cyan

Example:

```bash
--color=red
```

---

# Examples

### Basic ASCII Art

```bash
go run . "Hello"
```

---

### Color the Entire Text

```bash
go run . --color=green Hello
```

---

### Color a Specific Substring

```bash
go run . --color=red lo Hello
```

This colors only the substring **"lo"** in the ASCII output.

---

### Multi-line Input

```bash
go run . "Hello\nWorld"
```

---

# Example Output

```
 _   _      _ _       
| | | | ___| | | ___  
| |_| |/ _ \ | |/ _ \ 
|  _  |  __/ | | (_) |
|_| |_|\___|_|_|\___/
```

(The colored characters will appear in your terminal.)

---

# Error Handling

The program validates:

* Invalid number of arguments
* Incorrect option format
* Missing banner file
* Empty input

Example error message:

```
Invalid number of arguments
Usage: go run . [OPTION] [STRING]
```

---

# How It Works

1. Reads command-line arguments.
2. Parses the `--color` option.
3. Loads ASCII character definitions from `standard.txt`.
4. Converts each character into ASCII art.
5. Applies ANSI color codes to selected characters.
6. Prints the result to the terminal.

---

# License

Open-source project intended for **learning and experimentation with Go**.

---

# Author
>Dillon Ofili

Built using **Go** to practice:

* CLI tools
* File handling
* String processing
* Terminal color formatting
* ASCII rendering
