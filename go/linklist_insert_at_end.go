// Practice program to create linked list with insertion at end.

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
	var new_ptr *node
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

			// Check if it is the first element then create a node
			// and intialize the ll_head pointer to it
			if ll_head == nil {
				new_node := node{num, ll_head}
				new_ptr = &new_node
				ll_head = &new_node
			} else {
				// if first element already exist, keep on
				// adding nodes and iterate new_ptr
				new_node := node{num, nil}
				new_ptr.next = &new_node
				new_ptr = new_ptr.next
			}
		}
	}
	// Print the linked list
	printll(ll_head)
}
