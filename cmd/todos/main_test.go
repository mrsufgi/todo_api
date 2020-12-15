// +build integration unit

package main_test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("TESTING MAIN!")
	os.Exit(m.Run())
}
