package config

import (
	"fmt"
	"testing"
)

func TestParserConfig(t *testing.T) {
	ParserConfig()
	fmt.Printf("%##v\n",Config)
	fmt.Println("serviceport",ServerInfo)
}