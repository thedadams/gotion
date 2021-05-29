# gotion
⛔ WARNING: This is a work in progress! There are no gaurantees until the package reaches v1.0.0 ⛔

[![Go Reference](https://pkg.go.dev/badge/github.com/thedadams/gotion.svg)](https://pkg.go.dev/github.com/thedadams/gotion) [![Go Report Card](https://goreportcard.com/badge/github.com/thedadams/gotion)](https://goreportcard.com/report/github.com/thedadams/gotion) [![Maintainability](https://api.codeclimate.com/v1/badges/8acee9c1a9139a82b98b/maintainability)](https://codeclimate.com/github/thedadams/gotion/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/8acee9c1a9139a82b98b/test_coverage)](https://codeclimate.com/github/thedadams/gotion/test_coverage)

Go client for the Official Notion API. The goal of this project is to provide type-safety to the Notion API. Instead of page/database properties being `interface{}` in Go, this project tries to provide safety by having custom (un)marshal methods for the types to transition from Go-like code to the Javascript-type code that is compatible with the Notion API.

There are also some helper functions in the client that will get a page and its children from the Notion API in one call. Also, since Notion has a maximum page size for paginated requests, `gotion` can make multiple requests to get all results from a paginated request. Never fear: `gotion` is rate-limited so you won't overwhelm the Notion API.

## Installation

```bash
go get github.com/thedadams/gotion
```

## Getting started

To obtain an API key, follow Notion’s [getting started guide](https://developers.notion.com/docs/getting-started).

### Code example


```go
import (
    "context"

    "github.com/thedadams/gotion"
)

func main() {
    client := gotion.NewClient("api-key")
    
    page, err := client.GetPage(context.Background(), "page-id")
    if err != nil {
        // Handler error
    }

    for _, prop := range page.Properties {
        fmt.Printf("Prop name %s with type %s\n", prop.Name, prop.Type)
    }

    page, err = client.GetPageAndChildren(context.Background(), "page-id")
    if err != nil {
        // Handle error
    }
    // Use page
}
```

## Status

All the basic methods (and some helpers) are implemented to enble communciation with the Notion API with one excpetion (see next section).

- GetPage
- GetPageAndChildren
- UpdatePageProperties
- CreatePage
- GetDatabae
- GetDatabases
- GetDatabaseAndChildren
- GetBlockChildren
- AppendBlockChildren
- QueryDatabase
- Search

### TODO
- [ ] Add basic examples
- [ ] Nested Compound filters
- [ ] Github Actions
- [ ] Unit tests
- [ ] Integration tests
