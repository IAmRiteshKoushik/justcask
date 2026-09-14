# Decoder rejects oversized headers

The learner added 4 KiB key and 16 MiB value limits to `Decode` immediately after parsing the length fields. Header-only oversized-key and oversized-value tests now return `ErrRecordTooLarge` before body-presence and checksum checks. The focused codec suite passes.
