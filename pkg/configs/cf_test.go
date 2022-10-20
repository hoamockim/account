package configs

import (
	"fmt"
	"testing"
)

func Test_config(t *testing.T) {
	fmt.Println(DBConnectionString())
}

func Test_Pin(t *testing.T) {
	fmt.Printf("Hello world")
}
