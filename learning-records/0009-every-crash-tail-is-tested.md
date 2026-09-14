# Every crash tail is tested

The learner added a literal v1 test fixture and verified that all 24 truncated prefixes of the valid record return `ErrIncompleteRecord` and a zero `Record`. This turns the crash-tail rule into an exhaustive property for the first record format.
