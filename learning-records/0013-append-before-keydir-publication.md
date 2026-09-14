# Append before keydir publication

The learner separated the record codec from the data-file layer: data files append already encoded bytes and return a location containing file ID, starting byte offset, and encoded record size. They established that a future keydir may update only after append successfully returns that location.
