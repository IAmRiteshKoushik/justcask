# Building My Own Bitcask

Here's the [paper](https://riak.com/assets/bitcask-intro.pdf) behind this KV store.

Bitcask is a log-structured hash table for fast KV data storage and retrieval.
It writes data to an append-only log file and uses an in-memory hash table to 
keep track of the positions of it's log entries. They chose this approach as it 
provides high read/write throughput and is relatively easy to implement and 
reason about.

Technical goals of this implementation:
- Low latency per item read and write
- High throughput, especially when writing an incoming stream or random items.
- The ability to handle datasets much larger than RAM without performance 
degradation.
- Crash "friendly", meaning it provides fast recovery and zero to minimal 
data loss.
- Easy to backup and restore.
- A relatively simple, understandable and thus supportable codebase and format.
- Predictable behaviour under heavy access load or large data volume.
