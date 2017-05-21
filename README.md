# nexmo
An idiomatic Go [nexmo](https://nexmo.com) library


[![GoDoc](https://godoc.org/github.com/ammario?status.svg)](https://godoc.org/github.com/ammario/nexmo)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Install](#install)
- [Status](#status)
- [Examples](#examples)
  - [Create a new client](#create-a-new-client)
  - [Get country number is in](#get-country-number-is-in)
  - [Send an SMS](#send-an-sms)
- [Testing](#testing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Install

```go get -u github.com/ammario/nexmo```

# Status 

Currently these APIs are supported:

- SMS
- Number Insight

TODO:

- Developer API
- Voice API

# Examples

## Create a new client

```go
client := nexmo.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
```

## Get country number is in

```go
resp, err := client.NumberInsightBasic("2562893333")
if err != nil {
    panic(err)
}
fmt.Printf("Country: %v\n", resp.CountryName)
```

## Send an SMS


```go
_, err := client.SMS("your number here", "2562893333")
if err != nil {
    panic(err)
}
```

# Testing

Nexmo doesn't provide a sandbox at the time of writing so tests only cover free endpoints.