# Write errors have two meanings

The learner proved that the first failed append preserves its underlying write error, while later appends report the poisoned state and make no further write call. This separates the cause of a failed operation from the active file's later safety state.

## Evidence

The learner changed the regression test to inject a non-short-write sentinel error and verified the focused and full Go test suites.
