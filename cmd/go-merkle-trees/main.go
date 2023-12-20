package main

import (
    "fmt"
    "github.com/nolanjacobson/go-merkle-trees/internal/merkletree"
)

func main() {
    dataBlocks := [][]byte{
        []byte("Block 1"),
        []byte("Block 2"),
        // ... more blocks
    }

    merkleTree := merkletree.NewMerkleTree(dataBlocks)
    fmt.Printf("Merkle Tree Root Hash: %x\n", merkleTree.Root.Data)
}