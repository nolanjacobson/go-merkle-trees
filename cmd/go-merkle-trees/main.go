package main

import (
    "log",
    "go-merkle-trees"
    // Import other necessary packages
)

// func main() {
//     // Your application logic here
//     log.Println("Application started")
// }

func main() {
    dataBlocks := [][]byte{
        []byte("Block 1"),
        []byte("Block 2"),
        // ... more blocks
    }

    merkleTree := NewMerkleTree(dataBlocks)
    fmt.Println("Merkle Tree Root Hash:", merkleTree.Root.Data)
}