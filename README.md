# AT-5000 Auto-Tweeter [![GoDoc](https://godoc.org/github.com/spotlightpa/at-5000?status.svg)](https://godoc.org/github.com/spotlightpa/at-5000) [![Go Report Card](https://goreportcard.com/badge/github.com/spotlightpa/at-5000)](https://goreportcard.com/report/github.com/spotlightpa/at-5000)

An autotweeter that sends a randomly selected Tweet from a JSON array of choices.

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN=$(pwd) go install github.com/spotlightpa/at-5000@latest
```

## Usage

```
$ at-5000 -h
at-5000 - sends a randomly selected Tweet from a JSON array of choices

Usage:

        at-5000 [options]

Options can also be specified as environment variables prefixed with AUTOTWEETER_.
  -blob-url URL
        URL for S3 blob store (mock if not set)
  -mock
        mock calls rather than use real thing
  -silent
        don't log debug output
  -src file or URL
        file or URL source for Tweets (default stdin)
  -template template
        Go-style template for Tweet text
  -twitter-access-token string
    
  -twitter-access-token-secret string
    
  -twitter-consumer-key string
    
  -twitter-consumer-secret string
```

- - - -

*Designed to alert schoolchildren about snow days and such*

![AT-5000 Autodialer](images/at-5000.png)
