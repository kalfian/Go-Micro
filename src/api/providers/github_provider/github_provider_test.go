package githubprovider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("tester123")
	assert.EqualValues(t, "token tester123", header)
}

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")

	fmt.Println("function Body")
}
