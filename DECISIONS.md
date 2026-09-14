# Format decisions

This file is the source of truth for choices that change bytes written by Justcask. A later implementation must preserve existing decisions or introduce an explicit format migration.

## Record format v1

Each put record is encoded in this order:

```text
crc32 | timestamp | key length | value length | key | value
```

- All fixed-width integer fields use big-endian encoding.
- `crc32` occupies 4 bytes and uses CRC-32/IEEE.
- `timestamp` is an 8-byte signed `int64` Unix-nanosecond count.
- `key length` and `value length` are unsigned 4-byte fields.
- The checksum covers `timestamp | key length | value length | key | value`, never the checksum field itself.
- The fixed header occupies 20 bytes.
- The implementation accepts keys up to 4 KiB and values up to 16 MiB. These are decoder-enforced safety limits, not changes to the on-disk `uint32` field widths.

## Recovery v1

- A short read at the physical EOF is a crash tail. The partial record never enters the keydir.
- A checksum mismatch after a complete record read is corruption. Opening returns an error with the affected file and byte offset. It does not discard, truncate, compact, or otherwise modify data.
- A store publishes a keydir location only after the entire encoded record is successfully appended to the active data file and append returns that location. Power-loss durability is a separate sync-policy decision.
- A data-file location contains a `uint32` file ID, `int64` starting byte offset, and `uint32` full encoded record size. A successful location denotes a byte range that already exists in the file.
- After a data-file write fails, that `DataFile` is poisoned. Later append attempts return an error without writing. Restart recovery treats the partial EOF record as a crash tail.

## Deliberately deferred

- The explicit operation encoding for tombstones. Empty values will remain valid values.
