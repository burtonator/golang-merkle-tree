# Overview

Simple merkle tree implementation implemented in Go.

# Usage

```go

if slabs, err := SplitFile(path, 1024); err == nil {
    tree := BuildTree(slabs)
}

```