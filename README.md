# FreeMeal SDK

Search, browse, and look up meal recipes with ingredients, instructions, and images from TheMealDB

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Free Meal API

[TheMealDB](https://www.themealdb.com/) is a crowd-sourced, open recipe database that exposes its catalogue through a simple JSON HTTP API. This SDK targets the v1 endpoints, which are reachable using the public developer test key `1` at `https://www.themealdb.com/api/json/v1/1`.

What you get from the API:

- Search meals by name or by first letter (`/search.php?s=`, `/search.php?f=`)
- Look up a full meal record by id (`/lookup.php?i=`)
- Fetch a random meal (`/random.php`)
- Browse categories with images and descriptions (`/categories.php`)
- Filter meals by main ingredient, category, or area (`/filter.php?i=`, `?c=`, `?a=`)
- List the available categories, areas, and ingredients (`/list.php?c=list`, `?a=list`, `?i=list`)
- Meal records include title, category, area/cuisine, instructions, ingredient + measure pairs, image URL, tags, and optional YouTube / source links.

Operational notes: CORS is enabled, responses are JSON, and no auth header is needed for the test key (the key is part of the URL path). Some v2-only features advertised on the docs page — latest meals, a 10-random selection, and multi-ingredient filtering — are gated behind a supporter key and may return empty or unauthorised responses on the test key.

## Try it

**TypeScript**
```bash
npm install free-meal
```

**Python**
```bash
pip install free-meal-sdk
```

**PHP**
```bash
composer require voxgig/free-meal-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/free-meal-sdk/go
```

**Ruby**
```bash
gem install free-meal-sdk
```

**Lua**
```bash
luarocks install free-meal-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { FreeMealSDK } from 'free-meal'

const client = new FreeMealSDK({})

// List all categorys
const categorys = await client.Category().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o free-meal-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "free-meal": {
      "command": "/abs/path/to/free-meal-mcp"
    }
  }
}
```

## Entities

The API exposes 8 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Category** | Meal categories (e.g. Beef, Dessert) with thumbnail and description, served by `/categories.php` and usable as a filter via `/filter.php?c=`. | `/categories.php` |
| **Filter** | Filtered meal listings by main ingredient, category, or area: `/filter.php?i=`, `/filter.php?c=`, `/filter.php?a=`. | `/filter.php` |
| **Latest** | Most recently added meals — a v2/supporter-key endpoint that typically requires a paid key rather than the test key `1`. | `/latest.php` |
| **List** | Enumerations of the available categories, areas, and ingredients via `/list.php?c=list`, `?a=list`, `?i=list`. | `/list.php` |
| **Lookup** | Full meal detail records fetched by id at `/lookup.php?i=`, including instructions and ingredient/measure pairs. | `/lookup.php` |
| **Random** | A single randomly chosen meal from the database via `/random.php`. | `/random.php` |
| **Randomselection** | A batch of multiple random meals — a v2/supporter-key feature (commonly 10 meals) not available on the free test key. | `/randomselection.php` |
| **Search** | Meal search by name (`/search.php?s=`) or by first letter (`/search.php?f=`). | `/search.php` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from freemeal_sdk import FreeMealSDK

client = FreeMealSDK({})

# List all categorys
categorys, err = client.Category(None).list(None, None)
```

### PHP

```php
<?php
require_once 'freemeal_sdk.php';

$client = new FreeMealSDK([]);

// List all categorys
[$categorys, $err] = $client->Category(null)->list(null, null);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/free-meal-sdk/go"

client := sdk.NewFreeMealSDK(map[string]any{})

// List all categorys
categorys, err := client.Category(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "FreeMeal_sdk"

client = FreeMealSDK.new({})

# List all categorys
categorys, err = client.Category(nil).list(nil, nil)
```

### Lua

```lua
local sdk = require("free-meal_sdk")

local client = sdk.new({})

-- List all categorys
local categorys, err = client:Category(nil):list(nil, nil)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = FreeMealSDK.test()
const result = await client.Category().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = FreeMealSDK.test(None, None)
result, err = client.Category(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = FreeMealSDK::test(null, null);
[$result, $err] = $client->Category(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Category(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = FreeMealSDK.test(nil, nil)
result, err = client.Category(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Category(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Free Meal API

- Upstream: [https://www.themealdb.com/](https://www.themealdb.com/)
- API docs: [https://www.themealdb.com/api.php](https://www.themealdb.com/api.php)

- TheMealDB states the API and site will always remain free at point of access.
- The test/developer key `1` is intended for development and educational use.
- Production or commercial usage requires upgrading to a paid supporter key (via PayPal).
- No explicit attribution clause is published on the docs page; check `https://www.themealdb.com/api.php` for current terms.

---

Generated from the Free Meal API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
