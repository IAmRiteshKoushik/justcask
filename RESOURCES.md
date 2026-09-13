# Bitcask and durable storage resources

## Knowledge

- [Paper: "Bitcask: A Log-Structured Hash Table for Fast Key/Value Data" by Sheehy and Smith](https://riak.com/assets/bitcask-intro.pdf)
  The primary design source. Use for the keydir, append-only files, merges, hint files, and the original tradeoffs.
- [Riak Bitcask documentation](https://docs.riak.com/riak/kv/2.2.3/setup/planning/backend/bitcask/index.html)
  A concise explanation of the engine's intended operating model. Use to cross-check terminology from the paper.
- [Go `os.File` documentation](https://pkg.go.dev/os#File)
  The contract for opening, seeking, reading at offsets, writing, truncating, and syncing data files.
- [Go `encoding/binary` documentation](https://pkg.go.dev/encoding/binary)
  The standard-library tools for a stable byte order and fixed-width record fields.
- [Go `hash/crc32` documentation](https://pkg.go.dev/hash/crc32)
  The standard-library checksum package. Use when designing and checking record integrity.
- [Go testing package documentation](https://pkg.go.dev/testing)
  The base reference for table tests, benchmarks, fuzz targets, and test fixtures.

## Wisdom (communities)

- [Database Internals Discord](https://discord.gg/ezK7dYv)
  A practitioner community for asking focused questions about storage-engine tradeoffs. Bring a small reproducer or design sketch, not a broad "how do I build a DB" question.
- [r/databasedevelopment](https://www.reddit.com/r/databasedevelopment/)
  A place to compare design choices and learn from other implementers. Treat replies as prompts to verify against primary sources and experiments.

## Gaps

- We will choose and document a concrete durability promise after the first record codec exists. The Bitcask paper describes goals and mechanisms, but it does not substitute for a precise contract for this CLI.
