package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var inputWavFile string
var outputMp3File string

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputWavFile, "input", "i", "", "Input WAV file path")
	rootCmd.PersistentFlags().StringVarP(&outputMp3File, "output", "o", "output.mp3", "Output MP3 file path")
}

var rootCmd = &cobra.Command{
	Use:   "wav2mp3",
	Short: "Convert WAV to MP3",
	Run: func(cmd *cobra.Command, args []string) {
		if inputWavFile == "" {
			fmt.Println("Input file is required")
			return
		}
		if err := convertWavToMp3(inputWavFile, outputMp3File); err != nil {
			fmt.Println("Conversion failed:", err)
			return
		}

		fmt.Println("Conversion successful:", inputWavFile, "->", outputMp3File)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convertWavToMp3(inputWav, outputMp3 string) error {
	cmd := exec.Command("ffmpeg", "-i", inputWav, "-vn", "-ar", "44100", "-ac", "2", "-ab", "192k", "-f", "mp3", outputMp3)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
