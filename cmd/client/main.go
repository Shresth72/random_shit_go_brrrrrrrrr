package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Shresth72/tor_client/pkg/decode"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "bencode",
		Short: "A tool to decode bencoded values",
	}

	decodeCmd := &cobra.Command{
		Use:   "decode [encoded_value]",
		Short: "Decode a bencoded string",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			encodedValue := args[0]
			decoded, _, err := decode.DecodeBencodedValue(encodedValue)

			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}

			fmt.Println(decoded)
			fmt.Println()
			fmt.Println(prettyPrint(decoded))
		},
	}

	rootCmd.AddCommand(decodeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prettyPrint(val interface{}) string {
	switch v := val.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case string:
		return fmt.Sprintf("%q", v) // quoted string
	case []interface{}:
		var items []string
		for _, item := range v {
			items = append(items, prettyPrint(item))
		}
		return "[" + strings.Join(items, ",") + "]"
	case map[string]interface{}:
		var items []string
		for k, v := range v {
			items = append(items, fmt.Sprintf("%q:%s", k, prettyPrint(v)))
		}
		return "{" + strings.Join(items, ",") + "}"
	default:
		return fmt.Sprintf("%v", v)
	}
}
