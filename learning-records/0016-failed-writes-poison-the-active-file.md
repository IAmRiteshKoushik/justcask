# Failed writes poison the active file

The learner chose to poison a data file after a write error rather than repair it in place. They explained that later appends must return a zero location and avoid another write, which keeps a partial record at physical EOF for restart recovery.

## Evidence

The learner added a failing regression test that simulates a short write and proves the current implementation incorrectly attempts a second write.
