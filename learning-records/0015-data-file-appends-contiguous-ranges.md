# Data-file appends contiguous ranges

The learner implemented a single-writer data file that opens with append mode and owner-only permissions, captures the pre-write file size, checks for a complete write, and returns a location with file ID, offset, and full written size. The focused append test passes and confirms two records occupy contiguous ranges with matching file bytes.
