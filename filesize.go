package wtype

import "fmt"

// FileSize represents file size units in bytes.
type FileSize int64

func (s FileSize) String() string {
	const unit = 1024
	if s < unit {
		return fmt.Sprintf("%dB", s)
	}

	div, exp := int64(unit), 0
	for n := s / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	value := float64(s) / float64(div)
	units := []string{"KB", "MB", "GB", "TB"}
	if exp >= len(units) {
		return fmt.Sprintf("%.1fPB", value/float64(unit)) // 避免超出範圍
	}
	return fmt.Sprintf("%.1f%s", value, units[exp])
}

const (
	B  FileSize = 1 << (10 * iota) // 1 B
	KB                             // 1 KB = 1024 B
	MB                             // 1 MB = 1024 KB
	GB                             // 1 GB = 1024 MB
	TB                             // 1 TB = 1024 GB
)
