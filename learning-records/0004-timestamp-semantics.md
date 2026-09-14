# Timestamp semantics

The learner chose a signed 8-byte `int64` Unix-nanosecond timestamp for records. The timestamp is record metadata; physical append order remains the authority for which record wins during recovery.
