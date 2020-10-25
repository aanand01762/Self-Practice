// Solution of problem "7. Reverse Integer" at leetcode
// link: https://leetcode.com/problems/reverse-integer/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter the num:")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reverse(num))
	}

}

func reverse(x int) int {
	var result int
	for x != 0 {
		remainder := x % 10
		x /= 10
		result = result*10 + remainder
	}
	if result > int(math.Pow(2, 31)-1) || result < -int(math.Pow(2, 31)) {
		return 0
	}
	return result
}
