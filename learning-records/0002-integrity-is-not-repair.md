# Integrity is not repair

The learner distinguished an incomplete EOF tail from a complete record with a checksum mismatch. They correctly chose to keep a short EOF read out of the keydir and to treat a checksum mismatch as corruption rather than silently dropping it. Future recovery work must fail without mutating data when it sees a complete corrupt record.
