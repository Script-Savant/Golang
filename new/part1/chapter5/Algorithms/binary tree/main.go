package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// insert a value into the binary search tree
func insert(root *Node, value int) *Node {
	// create a node if there is none
	if root == nil {
		return &Node{Data: value}
	}

	// if the value is less that the current node data insert it into the left subtree
	if value < root.Data {
		root.Left = insert(root.Left, value)
	} else {
		// otherwise, insert into the right subtree
		root.Right = insert(root.Right, value)
	}

	// return root of the modified tree
	return root
}

// In-order Traversal: Left -> Root -> Right
func inOrder(root *Node) {
	if root != nil {
		inOrder(root.Left)        // traverse the left subtree
		fmt.Print(root.Data, " ") // visit the root node
		inOrder(root.Right)       // traverse the right subtree
	}
}

// Root -> left -> Right
func preOrder(root *Node) {
	if root != nil {
		fmt.Print(root.Data, " ")
		preOrder(root.Left)
		preOrder(root.Right)
	}
}

// Left -> right -> Root
func postOrder(root *Node) {
	if root != nil {
		postOrder(root.Left)
		postOrder(root.Right)
		fmt.Print(root.Data, " ")
	}
}

// search for a value in the BST
func Search(root *Node, value int) bool {
	// if current node is nil, return value not found
	if root == nil {
		return false
	}

	// if the current node's value == target valu return true
	if root.Data == value {
		return true
	}

	// if target value < current node's value, search the left subtree
	if value < root.Data {
		return Search(root.Left, value)
	}

	// otherwise, search the right subtree
	return Search(root.Right, value)
}

// Find the minimum value in a tree
func findMin(root *Node) *Node {
	// traverse to the left most node to find the minimum value
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// Delete a node with the given value from the BST
func delete(root *Node, value int) *Node {
	// if current node is nil return nil -> base case for recursion
	if root == nil {
		return nil
	}

	// if the value is less than the current node's value. delete from the left subtree
	if value < root.Data {
		root.Left = delete(root.Left, value)
	} else if value > root.Data {
		// if the value is greater than the current node's value. delete from the right subtree
		root.Right = delete(root.Right, value)
	} else {
		// if the value matches the current node's value, delete this node

		// case 1: Node has no left child(only right child or no children
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			// case 2: Node has no right child only left child
			return root.Left
		}

		// Case 3: Node has two children
		// Find the in-order successor (smallest value in the right subtree)
		temp := findMin(root.Right)

		// Replace the current node's value with the successor's value
		root.Data = temp.Data

		// Delete the successor node from the right subtree
		root.Right = delete(root.Right, temp.Data)
	}

	return root
}

func main() {
	var root *Node

	// Insert values into the BST
	root = insert(root, 50)
	root = insert(root, 30)
	root = insert(root, 20)
	root = insert(root, 40)
	root = insert(root, 70)
	root = insert(root, 60)
	root = insert(root, 80)

	// Perform traversals
	fmt.Println("In-order Traversal:")
	inOrder(root) // Output: 20 30 40 50 60 70 80
	fmt.Println("\nPre-order Traversal:")
	preOrder(root) // Output: 50 30 20 40 70 60 80
	fmt.Println("\nPost-order Traversal:")
	postOrder(root) // Output: 20 40 30 60 80 70 50

	// Search for values
	fmt.Println("\nSearch for 40:", Search(root, 40)) // Output: true
	fmt.Println("Search for 90:", Search(root, 90))   // Output: false

	// Delete nodes and print the tree after each deletion
	fmt.Println("\nDeleting 20...")
	root = delete(root, 20)
	fmt.Println("In-order Traversal after deletion:")
	inOrder(root) // Output: 30 40 50 60 70 80

	fmt.Println("\nDeleting 50...")
	root = delete(root, 50)
	fmt.Println("In-order Traversal after deletion:")
	inOrder(root) // Output: 30 40 60 70 80
}
