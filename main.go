// Command quickstart generates an audio file with the content "Hello, World!".
package main

import (
	"context"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

const AppName = "google-cloud-text-to-speech-cli"

type options struct {
	Text           string  `short:"t" long:"text" description:"[required] Text content." default:""`
	LanguageCode   string  `short:"l" long:"language" description:"LanguageCode." default:"en"`
	Gender         string  `short:"g" long:"gender" description:"SsmlGender." default:"FEMALE"`
	Voice          string  `short:"v" long:"voice" description:"Voice type. [ see --listvoicetype, --gender is ignored. ]"`
	SpeakingRate   float64 `short:"s" long:"rate" description:"SpeakingRate. [ 0.25 <= rate <= 4.0 ]" default:"1.0"`
	Pitch          float64 `short:"p" long:"pitch" description:"Pitch. [ -20.0 <= pitch <= 20.0 ] " default:"0.0"`
	OutputFilePath string  `short:"o" long:"output" description:"Output file path." default:"out/output.mp3"`
	ListVoiceType  bool    `long:"listvoicetype" description:"Display voice types."`
	FilterByLang   bool    `long:"filterbylang" description:"Filter voice types by language."`
}

func main() {

	// [ --help and error help ]
	opts := *new(options)
	parser := flags.NewParser(&opts, flags.Default)
	// set name
	parser.Name = AppName
	if _, err := parser.Parse(); err != nil {
		flagsError, _ := err.(*flags.Error)
		// help時は何もしない
		if flagsError.Type == flags.ErrHelp {
			return
		}
		fmt.Println()
		parser.WriteHelp(os.Stdout)
		fmt.Println()
		return
	}

	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	fmt.Println("text: ", opts.Text)
	fmt.Println("language: ", opts.LanguageCode)
	fmt.Println("gender: ", opts.Gender)
	fmt.Println("speakingRate: ", opts.SpeakingRate)
	fmt.Println("pitch: ", opts.Pitch)
	fmt.Println("output: ", opts.OutputFilePath)
	fmt.Println()

	// list mode
	if opts.ListVoiceType {
		ListVoices(os.Stdout, opts.LanguageCode, opts.FilterByLang)
		os.Exit(0)
	}

	if opts.Text == "" {
		log.Fatal("text is empty.")
	}
	ssmlVoiceGender, ok := texttospeechpb.SsmlVoiceGender_value[opts.Gender]
	if !ok {
		log.Fatal("Undefined voice type.")
	}

	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: opts.Text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice genderFlag ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: opts.LanguageCode,
			Name:         opts.Voice,
			SsmlGender:   texttospeechpb.SsmlVoiceGender(ssmlVoiceGender),
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			SpeakingRate:  opts.SpeakingRate,
			Pitch:         opts.Pitch,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	err = ioutil.WriteFile(opts.OutputFilePath, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", opts.OutputFilePath)
}

// > サポートされているすべての音声の一覧表示  ｜  Cloud Text-to-Speech Documentation  ｜  Google Cloud
// > https://cloud.google.com/text-to-speech/docs/list-voices?hl=ja
// ListVoices lists the available text to speech voices.
func ListVoices(w io.Writer, language string, enableFilter bool) error {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return err
	}

	// Performs the list voices request.
	resp, err := client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{})
	if err != nil {
		return err
	}

	for _, voice := range resp.Voices {

		// Display the supported language codes for this voice. Example: "en-US"
		isSupported := false
		for _, languageCode := range voice.LanguageCodes {
			if strings.HasPrefix(languageCode, language) {
				isSupported = true
			}
		}
		if enableFilter && !isSupported {
			continue
		}

		// Display the voice's name. Example: tpc-vocoded
		fmt.Fprintf(w, "Name: %v\n", voice.Name)

		// Display the supported language codes for this voice. Example: "en-US"
		for _, languageCode := range voice.LanguageCodes {
			fmt.Fprintf(w, "  Supported language: %v\n", languageCode)
		}

		// Display the SSML Voice Gender.
		fmt.Fprintf(w, "  SSML Voice Gender: %v\n", voice.SsmlGender.String())

		// Display the natural sample rate hertz for this voice. Example: 24000
		fmt.Fprintf(w, "  Natural Sample Rate Hertz: %v\n",
			voice.NaturalSampleRateHertz)
	}

	return nil
}
