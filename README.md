# text-summary
A lightweight go library for summarizing text into n lines.

# Installation
```
go get github.com/arjunmahishi/text-summary
```

# Usage
```go
package main

import (
	"fmt"

	textsummary "github.com/arjunmahishi/text-summary"
)

func main() {
	para := `So, keep working. Keep striving. Never give up. Fall down seven times, get up eight. Ease is a greater threat to progress than hardship. Ease is a greater threat to progress than hardship. So, keep moving, keep growing, keep learning. See you at work.`

	summarizer := textsummary.CreateSummarizer()

	summary := summarizer.Summary(para, 1) // Gives a 1 line summary
	fmt.Println(summary)

    // Output: Ease is a greater threat to progress than hardship.
}

```