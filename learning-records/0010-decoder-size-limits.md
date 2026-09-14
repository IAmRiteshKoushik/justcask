# Decoder size limits

The learner chose a 4 KiB maximum key size and a 16 MiB maximum value size. They explained that decode, not only the CLI, must enforce these limits because recovery accepts every record found in data files. The limits are implementation safety policy and do not alter the v1 `uint32` on-disk fields.
