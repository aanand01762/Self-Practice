// Practice program to create linked list with insertion at beginning.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	data int
	next *node
}

func printll(ll_head *node) {
	fmt.Println("Printing the linked list:")
	for ll_head != nil {
		fmt.Println(ll_head.data)
		ll_head = ll_head.next
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Start entering numbers and to stop just press enter key:")
	var ll_head *node
	flag := 1

	// Get the numbers until just "Enter key" is pressed
	for flag == 1 {
		str, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(strings.TrimSpace(str))

		// Check if it is end of linked list elements then set flag to
		// zero else print error and return
		if err != nil {
			if err.Error() != "strconv.Atoi: parsing \"\": invalid syntax" {
				fmt.Println(err.Error())
				return
			}
			flag = 0
		} else {

			// Create new node with its next pointer pointing to current
			// head and point the ll_head pointer to this new node
			new_node := node{num, ll_head}
			ll_head = &new_node
		}
	}
	// Print the linked list
	printll(ll_head)
}
