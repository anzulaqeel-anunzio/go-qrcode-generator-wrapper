# QR Code Generator Wrapper (Go)

A simple Go library to generate QR codes.

## Features

-   **Generate PNG**: Create QR code images as PNG files.
-   **Recovery Levels**: Support for Low, Medium, High, and Highest error recovery.
-   **Easy to Use**: Straightforward API.

## Installation

```bash
go get github.com/anunzio/qrcode-wrapper
```

## Usage

```go
package main

import (
	"log"

	"github.com/anunzio/qrcode-wrapper/qrcode"
)

func main() {
	// Generate a standard QR code
	err := qrcode.Generate("https://example.com", 256, "example.png")
	if err != nil {
		log.Fatal(err)
	}

    // Generate with custom recovery level
    err = qrcode.GenerateWithRecovery("https://example.com", 256, "example_high.png", qrcode.Highest)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Contact

Developed for Anunzio International by Anzul Aqeel.
Contact +971545822608 or +971585515742.

## License

MIT


---
### 🔗 Part of the "Ultimate Utility Toolkit"
This tool is part of the **[Anunzio International Utility Toolkit](https://github.com/anzulaqeel-anunzio/ultimate-utility-toolkit)**.
Check out the full collection of **180+ developer tools, scripts, and templates** in the master repository.

Developed for Anunzio International by Anzul Aqeel.
