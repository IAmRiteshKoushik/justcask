# Self-delimiting records and crash tails

The learner correctly proposed a record layout with a checksum, timestamp, key length, value length, key, and value. They explained how the fixed header and the two length fields determine a record boundary, and that an incomplete final record must not enter the recovered keydir. This is the foundation for the codec and recovery tests that follow.
