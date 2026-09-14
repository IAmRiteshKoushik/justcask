# Independent v1 decode passes

The learner implemented `Decode` for a complete v1 record. It reads the big-endian header fields at the correct widths, verifies the declared record size and CRC-32/IEEE, and returns correctly bounded key and value slices. The independent literal-vector decode test and encoder test both pass.
