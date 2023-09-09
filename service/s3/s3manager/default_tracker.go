package s3manager

import (
	"log"

	"fmt"

	"github.com/schollz/progressbar/v3"
)

func DefaultTracker(ch <-chan ProgressInfo, size int64) {
	if ch == nil {
		log.Println("Progress tracking disabled, because it is not a multipart upload!")
		return
	}
	bar := progressbar.Default(size)

	// pList := []ProgressInfo{}

	for p := range ch {
		bar.Describe(fmt.Sprintf("Part %d uploading", p.PartNumber))
		// pList = append(pList, p)
		bar.Add(1)
	}

	bar.Describe("Upload completed")
	// bar.Finish()
}
