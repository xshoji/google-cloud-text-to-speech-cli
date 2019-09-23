## google-cloud-text-to-speech-cli

Command line interface for google-cloud-text-to-speech-api.

> Cloud Text-to-Speech Client Libraries  ｜  Cloud Text-to-Speech API  ｜  Google Cloud  
> https://cloud.google.com/text-to-speech/docs/reference/libraries?hl=ja

## How to use

#### Export environment value

```
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/credential.json
```

#### Build and run

```
go build .
```

## Development

#### `go get` google libraries

```
go get -u cloud.google.com/go/texttospeech/apiv1
go get -u github.com/googleapis/google-cloud-go
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