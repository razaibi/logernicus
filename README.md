# **logernicus - Automatic Log Format Detection and Parsing in Golang**

## **Overview**
`logernicus` is a Golang package designed to **automatically detect and parse log files** in various formats. Instead of requiring manual format specification, this package identifies the log structure and extracts relevant information, making log processing seamless and efficient.

## **Project Objective**
- **Automatic log format detection** based on file content.
- **Support multiple log formats** such as JSON, Common Log Format (CLF), Syslog, Key-Value, and more.
- **Provide a unified structure** (`LogEntry`) for parsed logs.
- **Enable easy integration** with Golang applications.

## **Features**
✅ **Automatic log format detection**  
✅ **Support for multiple log types**  
✅ **Structured log parsing into a `LogEntry` format**  
✅ **Lightweight and efficient**  
✅ **Easy integration into any Golang project**  

---

## **Supported Log Formats & Examples**

### **1. JSON Log Format**
```json
{
  "timestamp": "2025-03-10T14:32:14Z",
  "level": "INFO",
  "message": "User logged in",
  "ip": "127.0.0.1",
  "user": "frank"
}
```

### **2. Common Log Format (CLF)**
```log
127.0.0.1 - frank [10/Mar/2025:14:32:14 +0000] "GET /index.html HTTP/1.1" 200 1234
```

### **3. Extended Log Format (ELF)**
```log
127.0.0.1 - frank [10/Mar/2025:14:32:14 +0000] "GET /index.html HTTP/1.1" 200 1234 "https://example.com" "Mozilla/5.0"
```

### **4. Key-Value (KV) Log Format**
```log
timestamp=2025-03-10T14:32:14Z level=INFO user=frank event="User logged in" ip=127.0.0.1
```

### **5. Syslog Format**
```log
<34>1 2025-03-10T14:32:14Z myserver app - ID47 [exampleSDID@32473 event="User logged in"]
```

### **6. Apache/Nginx Access Log**
```log
192.168.1.1 - - [10/Mar/2025:14:32:14 +0000] "GET /home HTTP/1.1" 200 1024 "-" "Mozilla/5.0"
```

---

## **Installation**
To use `logernicus`, install it via `go get`:

```sh
go get github.com/razaibi/logernicus
```

Then, import the package in your project:

```go
import "github.com/razaibi/logernicus"
```

---

## **Usage**
### **Basic Example**
```go
package main

import (
	"fmt"
	"log"

	"github.com/razaibi/logernicus"
)

func main() {
	// Read and parse logs
	logs, err := logernicus.ReadLogFile("logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Print parsed log entries
	for _, log := range logs {
		fmt.Printf("%+v\n", log)
	}
}
```

### **Expected Output**
```
{Timestamp:2025-03-10T14:32:14Z Level:INFO Message:User logged in IP:127.0.0.1 UserAgent: Request: StatusCode:200}
```

---

## **How It Works**
### **1. Automatic Log Format Detection**
- The package **reads the first few lines** of the log file and applies **regular expressions** to detect the format.
- If it matches **JSON**, **Common Log Format**, **Syslog**, **Key-Value**, or **Apache/Nginx**, it selects the appropriate parser.

### **2. Parsing into a `LogEntry` Struct**
- Regardless of the log format, the package **normalizes** the parsed data into a structured **LogEntry**.

#### **LogEntry Struct**
```go
type LogEntry struct {
    Timestamp  string
    Level      string
    Message    string
    IP         string
    UserAgent  string
    Request    string
    StatusCode int
}
```

---

## **Advanced Usage**
### **Manually Specify Log Format**
If you know the log format beforehand, you can call the parsers directly:
```go
package main

import (
	"fmt"
	"github.com/razaibi/logernicus/parsers"
)

func main() {
	log := `{"timestamp":"2025-03-10T14:32:14Z", "level":"INFO", "message":"User logged in"}`
	entry := parsers.ParseJSON(log)

	fmt.Printf("%+v\n", entry)
}
```

---

## **Contributing**
1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch`
3. Commit your changes: `git commit -m "Added new log format"`
4. Push the branch: `git push origin feature-branch`
5. Open a Pull Request.

---

## **License**
MIT License. See `LICENSE` for details.

---
