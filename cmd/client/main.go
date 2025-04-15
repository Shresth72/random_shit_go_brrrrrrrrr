package main

import (
	"fmt"
	"os"

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
		},
	}

	rootCmd.AddCommand(decodeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
