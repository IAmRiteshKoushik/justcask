# Exact-byte codec test is trustworthy

The learner created the `internal/record` package and a compiling exact-byte encoding test. They corrected the CRC literal and ensured the input record contains the timestamp, key, and value represented by the expected bytes. The remaining failure is only the intentional `Encode` TODO, so the test can now drive implementation.
