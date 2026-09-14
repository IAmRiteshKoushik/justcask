# Exact-byte encoder passes

The learner implemented `Encode` using a fixed 20-byte header, big-endian integer fields, key and value copies, and CRC-32/IEEE over the protected suffix. `go test ./internal/record` passes the independent exact-byte vector. They also explained the difference between slice length and capacity when indexed writes and append are involved.
