# JSON-translator

A tool that translates json files via Google Translate


![Screenshot](docs/images/screenshot.png)


## Installation Requirements

This application makes use of the Google Translate API, this means you will have
to provide your own API token, generated in the Google cloud console.

Once you have it, set it to the following Environment variable so that JSON-translator can find it.

`GOOGLE_API_KEY`

## Running For Development


Windows

```sh
go build ; .\JSON-translator.exe
```


Other

```sh
go build && ./JSON-translator
```


### License
Apache 2