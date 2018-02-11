package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bgpat/dhop"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dhop",
		Short: "Encode and decode DHCP options",
		Long: `dhop encodes and decodes DHCP option data.
With no options, dhop reads text data from stdin and writes available encoded data to stdout.`,
		RunE: execute,
	}
	isDecode    bool
	inputPath   string
	outputPath  string
	noTrimSpace bool
	codes       = codeRanges{
		codeRange{
			From: 0,
			To:   255,
		},
	}
	inputFormat  formatType
	outputFormat formatType
	separator    = " "
)

func init() {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "display this message")
	rootCmd.PersistentFlags().BoolVarP(&isDecode, "decode", "D", false, "decodes input")
	rootCmd.PersistentFlags().StringVarP(&inputPath, "input", "i", "-", "input file")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "-", "output file")
	rootCmd.PersistentFlags().BoolVarP(&noTrimSpace, "no-trim-space", "N", false, "do not trim spaces only binary format")
	rootCmd.PersistentFlags().VarP(&codes, "code", "c", "only output specified DHCP option code")
	rootCmd.PersistentFlags().VarP(&inputFormat, "input-format", "f", "input format")
	rootCmd.PersistentFlags().VarP(&outputFormat, "output-format", "t", "output format")
	rootCmd.PersistentFlags().StringVarP(&separator, "separator", "s", " ", "separator for hex format")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func readInput() (input []byte, err error) {
	if inputPath == "-" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else {
		input, err = ioutil.ReadFile(inputPath)
	}
	if err != nil {
		return nil, err
	}
	return inputFormat.Decode(input)
}

func writeOutput(output []byte, code byte) error {
	encoded, err := outputFormat.Encode(output)
	if err != nil {
		return err
	}
	fmt.Printf("%d: %s\n", code, encoded)
	return nil
}

func encode(input []byte, code byte) error {
	op, err := dhop.Unmarshal(code, input)
	if err != nil {
		return err
	}
	return writeOutput(op.Encode(), code)
}

func decode(input []byte, code byte) error {
	op, err := dhop.Decode(code, input)
	if err != nil {
		return err
	}
	return writeOutput(op.Marshal(), code)
}

func execute(cmd *cobra.Command, args []string) error {
	var convert func([]byte, byte) error
	if isDecode {
		convert = decode
		if inputFormat == FORMAT_TYPE_DEFAULT {
			inputFormat = FORMAT_TYPE_HEX
		}
		if outputFormat == FORMAT_TYPE_DEFAULT {
			outputFormat = FORMAT_TYPE_BINARY
		}
	} else {
		convert = encode
		if inputFormat == FORMAT_TYPE_DEFAULT {
			inputFormat = FORMAT_TYPE_BINARY
		}
		if outputFormat == FORMAT_TYPE_DEFAULT {
			outputFormat = FORMAT_TYPE_HEX
		}
	}
	input, err := readInput()
	if err != nil {
		return err
	}
	var success bool
	for _, code := range codes.Slice() {
		err := convert(input, code)
		if err != nil {
			continue
		}
		success = true
	}
	if !success {
		os.Exit(1)
	}
	return nil
}
