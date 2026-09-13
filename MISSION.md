# Mission: build a crash-recoverable Bitcask CLI in Go

## Why
Build a small, understandable database engine rather than treating persistence as a black box. The finished command-line store should make its on-disk format and recovery behavior visible enough to explain and demonstrate to another engineer.

## Success looks like
- `justcask` can put, get, delete, list, and inspect values in a chosen data directory.
- A deliberate unclean process stop leaves the store able to reopen, preserve committed records, and report how it handled a damaged tail.
- The repository includes a documented file format, tests for recovery and corruption, benchmarks, and a short reproducible demo.

## Constraints
- Learn by making the design decisions and writing the implementation in small stages. Use reviews, hints, and adversarial tests rather than generated whole features.
- Start from the current minimal Go program. Prefer the Go standard library until a dependency has a clear job.
- The learner is an experienced Go backend developer, somewhat familiar with file I/O and concurrency, and new to binary formats, checksums, and crash consistency.

## Out of scope
- Distributed replication, transactions, range scans, and an LSM tree.
- Matching every behavior or implementation detail of Riak's production Bitcask.
