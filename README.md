# K-th Permutation Sequence

> Striver SDE Sheet 054 · Day 9 Recursion — the factorial number system: divide k by (slots−1)! to pick each digit directly, no enumeration. Go, stdlib only, tested.

## Problem

Among the n! lexicographically ordered permutations of 1..n, return the k-th — without generating the k−1 before it (k can be enormous).

## Approach

Fix the first digit: all permutations starting with a given digit form a block of (n−1)! consecutive ranks. So `(k−1) / (n−1)!` indexes which remaining digit leads, and the remainder re-poses the same question one position shorter. This is writing k−1 in the *factorial number system*, one "digit" per output position. The off-by-one discipline — convert k to 0-based once, up front — is where most implementations break; everything after is clean division and a shrinking candidate list.

**Complexity:** O(n²) time (the list removal), O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
