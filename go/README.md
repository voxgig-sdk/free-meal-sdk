# FreeMeal Golang SDK



The Golang SDK for the FreeMeal API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/free-meal-sdk/go
```

If the module is not yet published to a registry, use a `replace` directive
in your `go.mod` to point to a local checkout:

```bash
go mod edit -replace github.com/voxgig-sdk/free-meal-sdk/go=../path/to/github.com/voxgig-sdk/free-meal-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```go
package main

import (
    "fmt"
    "os"

    sdk "github.com/voxgig-sdk/free-meal-sdk/go"
    "github.com/voxgig-sdk/free-meal-sdk/go/core"
)

func main() {
    client := sdk.NewFreeMealSDK(map[string]any{
        "apikey": os.Getenv("FREE-MEAL_APIKEY"),
    })
```

### 2. List categorys

```go
    result, err := client.Category(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }

    rm := core.ToMapAny(result)
    if rm["ok"] == true {
        for _, item := range rm["data"].([]any) {
            p := core.ToMapAny(item)
            fmt.Println(p["id"], p["name"])
        }
    }
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

result, err := client.Planet(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
// result contains mock response data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewFreeMealSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
FREE-MEAL_TEST_LIVE=TRUE
FREE-MEAL_APIKEY=<your-key>
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewFreeMealSDK

```go
func NewFreeMealSDK(options map[string]any) *FreeMealSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *FreeMealSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### FreeMealSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Category` | `(data map[string]any) FreeMealEntity` | Create a Category entity instance. |
| `Filter` | `(data map[string]any) FreeMealEntity` | Create a Filter entity instance. |
| `Latest` | `(data map[string]any) FreeMealEntity` | Create a Latest entity instance. |
| `List` | `(data map[string]any) FreeMealEntity` | Create a List entity instance. |
| `Lookup` | `(data map[string]any) FreeMealEntity` | Create a Lookup entity instance. |
| `Random` | `(data map[string]any) FreeMealEntity` | Create a Random entity instance. |
| `Randomselection` | `(data map[string]any) FreeMealEntity` | Create a Randomselection entity instance. |
| `Search` | `(data map[string]any) FreeMealEntity` | Create a Search entity instance. |

### Entity interface (FreeMealEntity)

All entities implement the `FreeMealEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(any, error)`. The `any` value is a
`map[string]any` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `"ok"` | `bool` | `true` if the HTTP status is 2xx. |
| `"status"` | `int` | HTTP status code. |
| `"headers"` | `map[string]any` | Response headers. |
| `"data"` | `any` | Parsed JSON response body. |

On error, `"ok"` is `false` and `"err"` contains the error value.

### Entities

#### Category

| Field | Description |
| --- | --- |
| `"id_category"` |  |
| `"str_category"` |  |
| `"str_category_description"` |  |
| `"str_category_thumb"` |  |

Operations: List.

API path: `/categories.php`

#### Filter

| Field | Description |
| --- | --- |
| `"id_meal"` |  |
| `"str_meal"` |  |
| `"str_meal_thumb"` |  |

Operations: List.

API path: `/filter.php`

#### Latest

| Field | Description |
| --- | --- |
| `"date_modified"` |  |
| `"id_meal"` |  |
| `"str_area"` |  |
| `"str_category"` |  |
| `"str_creative_commons_confirmed"` |  |
| `"str_drink_alternate"` |  |
| `"str_image_source"` |  |
| `"str_ingredient1"` |  |
| `"str_ingredient10"` |  |
| `"str_ingredient11"` |  |
| `"str_ingredient12"` |  |
| `"str_ingredient13"` |  |
| `"str_ingredient14"` |  |
| `"str_ingredient15"` |  |
| `"str_ingredient16"` |  |
| `"str_ingredient17"` |  |
| `"str_ingredient18"` |  |
| `"str_ingredient19"` |  |
| `"str_ingredient2"` |  |
| `"str_ingredient20"` |  |
| `"str_ingredient3"` |  |
| `"str_ingredient4"` |  |
| `"str_ingredient5"` |  |
| `"str_ingredient6"` |  |
| `"str_ingredient7"` |  |
| `"str_ingredient8"` |  |
| `"str_ingredient9"` |  |
| `"str_instruction"` |  |
| `"str_meal"` |  |
| `"str_meal_thumb"` |  |
| `"str_measure1"` |  |
| `"str_measure10"` |  |
| `"str_measure11"` |  |
| `"str_measure12"` |  |
| `"str_measure13"` |  |
| `"str_measure14"` |  |
| `"str_measure15"` |  |
| `"str_measure16"` |  |
| `"str_measure17"` |  |
| `"str_measure18"` |  |
| `"str_measure19"` |  |
| `"str_measure2"` |  |
| `"str_measure20"` |  |
| `"str_measure3"` |  |
| `"str_measure4"` |  |
| `"str_measure5"` |  |
| `"str_measure6"` |  |
| `"str_measure7"` |  |
| `"str_measure8"` |  |
| `"str_measure9"` |  |
| `"str_source"` |  |
| `"str_tag"` |  |
| `"str_youtube"` |  |

Operations: List.

API path: `/latest.php`

#### List

| Field | Description |
| --- | --- |
| `"str_area"` |  |
| `"str_category"` |  |
| `"str_ingredient"` |  |

Operations: List.

API path: `/list.php`

#### Lookup

| Field | Description |
| --- | --- |
| `"date_modified"` |  |
| `"id_meal"` |  |
| `"str_area"` |  |
| `"str_category"` |  |
| `"str_creative_commons_confirmed"` |  |
| `"str_drink_alternate"` |  |
| `"str_image_source"` |  |
| `"str_ingredient1"` |  |
| `"str_ingredient10"` |  |
| `"str_ingredient11"` |  |
| `"str_ingredient12"` |  |
| `"str_ingredient13"` |  |
| `"str_ingredient14"` |  |
| `"str_ingredient15"` |  |
| `"str_ingredient16"` |  |
| `"str_ingredient17"` |  |
| `"str_ingredient18"` |  |
| `"str_ingredient19"` |  |
| `"str_ingredient2"` |  |
| `"str_ingredient20"` |  |
| `"str_ingredient3"` |  |
| `"str_ingredient4"` |  |
| `"str_ingredient5"` |  |
| `"str_ingredient6"` |  |
| `"str_ingredient7"` |  |
| `"str_ingredient8"` |  |
| `"str_ingredient9"` |  |
| `"str_instruction"` |  |
| `"str_meal"` |  |
| `"str_meal_thumb"` |  |
| `"str_measure1"` |  |
| `"str_measure10"` |  |
| `"str_measure11"` |  |
| `"str_measure12"` |  |
| `"str_measure13"` |  |
| `"str_measure14"` |  |
| `"str_measure15"` |  |
| `"str_measure16"` |  |
| `"str_measure17"` |  |
| `"str_measure18"` |  |
| `"str_measure19"` |  |
| `"str_measure2"` |  |
| `"str_measure20"` |  |
| `"str_measure3"` |  |
| `"str_measure4"` |  |
| `"str_measure5"` |  |
| `"str_measure6"` |  |
| `"str_measure7"` |  |
| `"str_measure8"` |  |
| `"str_measure9"` |  |
| `"str_source"` |  |
| `"str_tag"` |  |
| `"str_youtube"` |  |

Operations: List.

API path: `/lookup.php`

#### Random

| Field | Description |
| --- | --- |
| `"date_modified"` |  |
| `"id_meal"` |  |
| `"str_area"` |  |
| `"str_category"` |  |
| `"str_creative_commons_confirmed"` |  |
| `"str_drink_alternate"` |  |
| `"str_image_source"` |  |
| `"str_ingredient1"` |  |
| `"str_ingredient10"` |  |
| `"str_ingredient11"` |  |
| `"str_ingredient12"` |  |
| `"str_ingredient13"` |  |
| `"str_ingredient14"` |  |
| `"str_ingredient15"` |  |
| `"str_ingredient16"` |  |
| `"str_ingredient17"` |  |
| `"str_ingredient18"` |  |
| `"str_ingredient19"` |  |
| `"str_ingredient2"` |  |
| `"str_ingredient20"` |  |
| `"str_ingredient3"` |  |
| `"str_ingredient4"` |  |
| `"str_ingredient5"` |  |
| `"str_ingredient6"` |  |
| `"str_ingredient7"` |  |
| `"str_ingredient8"` |  |
| `"str_ingredient9"` |  |
| `"str_instruction"` |  |
| `"str_meal"` |  |
| `"str_meal_thumb"` |  |
| `"str_measure1"` |  |
| `"str_measure10"` |  |
| `"str_measure11"` |  |
| `"str_measure12"` |  |
| `"str_measure13"` |  |
| `"str_measure14"` |  |
| `"str_measure15"` |  |
| `"str_measure16"` |  |
| `"str_measure17"` |  |
| `"str_measure18"` |  |
| `"str_measure19"` |  |
| `"str_measure2"` |  |
| `"str_measure20"` |  |
| `"str_measure3"` |  |
| `"str_measure4"` |  |
| `"str_measure5"` |  |
| `"str_measure6"` |  |
| `"str_measure7"` |  |
| `"str_measure8"` |  |
| `"str_measure9"` |  |
| `"str_source"` |  |
| `"str_tag"` |  |
| `"str_youtube"` |  |

Operations: List.

API path: `/random.php`

#### Randomselection

| Field | Description |
| --- | --- |
| `"date_modified"` |  |
| `"id_meal"` |  |
| `"str_area"` |  |
| `"str_category"` |  |
| `"str_creative_commons_confirmed"` |  |
| `"str_drink_alternate"` |  |
| `"str_image_source"` |  |
| `"str_ingredient1"` |  |
| `"str_ingredient10"` |  |
| `"str_ingredient11"` |  |
| `"str_ingredient12"` |  |
| `"str_ingredient13"` |  |
| `"str_ingredient14"` |  |
| `"str_ingredient15"` |  |
| `"str_ingredient16"` |  |
| `"str_ingredient17"` |  |
| `"str_ingredient18"` |  |
| `"str_ingredient19"` |  |
| `"str_ingredient2"` |  |
| `"str_ingredient20"` |  |
| `"str_ingredient3"` |  |
| `"str_ingredient4"` |  |
| `"str_ingredient5"` |  |
| `"str_ingredient6"` |  |
| `"str_ingredient7"` |  |
| `"str_ingredient8"` |  |
| `"str_ingredient9"` |  |
| `"str_instruction"` |  |
| `"str_meal"` |  |
| `"str_meal_thumb"` |  |
| `"str_measure1"` |  |
| `"str_measure10"` |  |
| `"str_measure11"` |  |
| `"str_measure12"` |  |
| `"str_measure13"` |  |
| `"str_measure14"` |  |
| `"str_measure15"` |  |
| `"str_measure16"` |  |
| `"str_measure17"` |  |
| `"str_measure18"` |  |
| `"str_measure19"` |  |
| `"str_measure2"` |  |
| `"str_measure20"` |  |
| `"str_measure3"` |  |
| `"str_measure4"` |  |
| `"str_measure5"` |  |
| `"str_measure6"` |  |
| `"str_measure7"` |  |
| `"str_measure8"` |  |
| `"str_measure9"` |  |
| `"str_source"` |  |
| `"str_tag"` |  |
| `"str_youtube"` |  |

Operations: List.

API path: `/randomselection.php`

#### Search

| Field | Description |
| --- | --- |
| `"date_modified"` |  |
| `"id_meal"` |  |
| `"str_area"` |  |
| `"str_category"` |  |
| `"str_creative_commons_confirmed"` |  |
| `"str_drink_alternate"` |  |
| `"str_image_source"` |  |
| `"str_ingredient1"` |  |
| `"str_ingredient10"` |  |
| `"str_ingredient11"` |  |
| `"str_ingredient12"` |  |
| `"str_ingredient13"` |  |
| `"str_ingredient14"` |  |
| `"str_ingredient15"` |  |
| `"str_ingredient16"` |  |
| `"str_ingredient17"` |  |
| `"str_ingredient18"` |  |
| `"str_ingredient19"` |  |
| `"str_ingredient2"` |  |
| `"str_ingredient20"` |  |
| `"str_ingredient3"` |  |
| `"str_ingredient4"` |  |
| `"str_ingredient5"` |  |
| `"str_ingredient6"` |  |
| `"str_ingredient7"` |  |
| `"str_ingredient8"` |  |
| `"str_ingredient9"` |  |
| `"str_instruction"` |  |
| `"str_meal"` |  |
| `"str_meal_thumb"` |  |
| `"str_measure1"` |  |
| `"str_measure10"` |  |
| `"str_measure11"` |  |
| `"str_measure12"` |  |
| `"str_measure13"` |  |
| `"str_measure14"` |  |
| `"str_measure15"` |  |
| `"str_measure16"` |  |
| `"str_measure17"` |  |
| `"str_measure18"` |  |
| `"str_measure19"` |  |
| `"str_measure2"` |  |
| `"str_measure20"` |  |
| `"str_measure3"` |  |
| `"str_measure4"` |  |
| `"str_measure5"` |  |
| `"str_measure6"` |  |
| `"str_measure7"` |  |
| `"str_measure8"` |  |
| `"str_measure9"` |  |
| `"str_source"` |  |
| `"str_tag"` |  |
| `"str_youtube"` |  |

Operations: List.

API path: `/search.php`



## Entities


### Category

Create an instance: `category := client.Category(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_category` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_category_description` | ``$STRING`` |  |
| `str_category_thumb` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.Category(nil).List(nil, nil)
```


### Filter

Create an instance: `filter := client.Filter(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_meal` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.Filter(nil).List(nil, nil)
```


### Latest

Create an instance: `latest := client.Latest(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.Latest(nil).List(nil, nil)
```


### List

Create an instance: `list := client.List(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_ingredient` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.List(nil).List(nil, nil)
```


### Lookup

Create an instance: `lookup := client.Lookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.Lookup(nil).List(nil, nil)
```


### Random

Create an instance: `random := client.Random(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.Random(nil).List(nil, nil)
```


### Randomselection

Create an instance: `randomselection := client.Randomselection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.Randomselection(nil).List(nil, nil)
```


### Search

Create an instance: `search := client.Search(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

#### Example: List

```go
results, err := client.Search(nil).List(nil, nil)
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller. An unexpected panic triggers the
`PreUnexpected` hook.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/free-meal-sdk/go/
├── free-meal.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/free-meal-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
moon := client.Moon(nil)
moon.Load(map[string]any{"planet_id": "earth", "id": "luna"}, nil)

// moon.Data() now returns the loaded moon data
// moon.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
