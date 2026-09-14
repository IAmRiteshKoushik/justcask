# Data-file location types

The learner chose `uint32` for file ID, `int64` for byte offset, and `uint32` for encoded record size. They understand that a successful append location must denote bytes that physically exist, otherwise a later keydir entry could point at missing data.
