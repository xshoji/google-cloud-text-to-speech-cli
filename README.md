# google-cloud-text-to-speech-cli

Command line interface for google-cloud-text-to-speech-api.

> Cloud Text-to-Speech Client Libraries  ｜  Cloud Text-to-Speech API  ｜  Google Cloud  
> https://cloud.google.com/text-to-speech/docs/reference/libraries?hl=ja

# How to use

## Export environment value

```
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/credential.json
```

#### Google Service Account

Credential file of Google Service Account ( /path/to/credential.json ) is required.
If you don't have account, you need create it.


## Execute

Download latest application based on your platform.

> Releases · xshoji/google-cloud-text-to-speech-cli  
> https://github.com/xshoji/google-cloud-text-to-speech-cli/releases

```
// Show help
$ ./google-cloud-text-to-speech-cli --help
Usage:
  google-cloud-text-to-speech-cli [OPTIONS]

Application Options:
  -t, --text=          [required] Text content.
  -l, --language=      LanguageCode. (default: en)
  -g, --gender=        SsmlGender. (default: FEMALE)
  -v, --voice=         Voice type. [ see --listvoicetype, --gender is ignored. ]
  -s, --rate=          SpeakingRate. [ 0.25 <= rate <= 4.0 ] (default: 1.0)
  -p, --pitch=         Pitch. [ -20.0 <= pitch <= 20.0 ]  (default: 0.0)
  -o, --output=        Output file path. (default: out/output.mp3)
      --listvoicetype  Display voice types.
      --filterbylang   Filter voice types by language.

Help Options:
  -h, --help           Show this help message  



// Run
$ ./google-cloud-text-to-speech-cli --text="Thank you download my apps. This is Command line interface for google-cloud-text-to-speech-api." --rate=1.5 --pitch=-5.0
language:  en
gender:  FEMALE
speakingRate:  1.5
pitch:  -5
output:  out/output.mp3

Audio content written to file: out/output.mp3
```

## Development

- dep
- go 1.12 ( or later )

#### `dep ensure`

```
dep ensure
```

> googleapis/google-cloud-go： Google Cloud Client Libraries for Go.  
> https://github.com/googleapis/google-cloud-go

#### Create `Service Account Key`

> Cloud Text-to-Speech Client Libraries  ｜  Cloud Text-to-Speech API  ｜  Google Cloud  
> https://cloud.google.com/text-to-speech/docs/reference/libraries?hl=ja

Then, download credential json.

```
2019/09/23 20:00:08 rpc error: code = PermissionDenied desc = Cloud Text-to-Speech API has not been used in project xxxxxx before or it is disabled. Enable it by visiting https://console.cloud.google.com/apis/api/texttospeech.googleapis.com/overview?project=xxxxx then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
```

And enable API account ( register your credit card info ).

## References

> サポートされているすべての音声の一覧表示  ｜  Cloud Text-to-Speech Documentation  ｜  Google Cloud  
> https://cloud.google.com/text-to-speech/docs/list-voices?hl=ja