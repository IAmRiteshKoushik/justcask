# Encoder and decoder limits are symmetric

The learner made `Encode` reject the same oversized keys and values that `Decode` rejects. The focused codec suite passes, establishing that a successful encoder call cannot create a record outside the decoder's accepted size policy.
