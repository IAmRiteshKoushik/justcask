# Stable decode errors pass

The learner made incomplete input and checksum corruption distinct, stable decoder outcomes. Tests use `errors.Is` and prove failed decoding returns a zero `Record`. The codec suite now passes exact encoding, independent decoding, incomplete-record handling, and checksum-mismatch handling.
