package go-merkle-trees

// Define your types, functions, and methods

type MerkleNode struct {
    Left  *MerkleNode
    Right *MerkleNode
    Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
    mNode := MerkleNode{}

    if left == nil && right == nil {
        // If it's a leaf node, hash the data
        hash := sha256.Sum256(data)
        mNode.Data = hash[:]
    } else {
        // Concatenate the hashes of left and right and then hash
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

    // Create a leaf node for each data block
    for _, datum := range data {
        node := NewMerkleNode(nil, nil, datum)
        nodes = append(nodes, *node)
    }

    // Iterate over the nodes until the tree is complete
    for len(nodes) > 1 {
        var newLevel []MerkleNode

        for i := 0; i < len(nodes); i += 2 {
            // Handle odd number of nodes
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
