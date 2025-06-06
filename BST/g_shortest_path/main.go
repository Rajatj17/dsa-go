package shortestpath

// Shortest Tree Path
// Given a Binary Search Tree and values of two nodes that lie inside the tree, find the Shortest Path Length between the two nodes.

//          10
//        /     \
//       4      15
//     /   \   /   \
//    2     5 13   22
//  /           \
// 1            14
// Examples

// Shortest Distance between (1,4) is 2.
// Shortest Distance between (2,13) is 4.
// Shortest Distance between (5,14) is 5.

// Hint:
// First find LCA, then note down the distance of two nodes from LCA.

