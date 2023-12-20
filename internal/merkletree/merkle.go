package merkletree

import "crypto/sha256"

type MerkleNode struct {
    Left  *MerkleNode
    Right *MerkleNode
    Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
    mNode := MerkleNode{}

    if left == nil && right == nil {
        hash := sha256.Sum256(data)
        mNode.Data = hash[:]
    } else {
        prevHashes := append(left.Data, right.Data...)
        newHash := sha256.Sum256(prevHashes)
        mNode.Data = newHash[:]
    }

    mNode.Left = left
    mNode.Right = right

    return &mNode
}

type MerkleTree struct {
    Root *MerkleNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
    var nodes []MerkleNode

    for _, datum := range data {
        node := NewMerkleNode(nil, nil, datum)
        nodes = append(nodes, *node)
    }

    for len(nodes) > 1 {
        var newLevel []MerkleNode

        for i := 0; i < len(nodes); i += 2 {
            if i+1 == len(nodes) {
                newLevel = append(newLevel, nodes[i])
                break
            }

            node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
            newLevel = append(newLevel, *node)
        }

        nodes = newLevel
    }

    mTree := MerkleTree{&nodes[0]}
    return &mTree
}