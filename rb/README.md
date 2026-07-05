# FreeMeal Ruby SDK



The Ruby SDK for the FreeMeal API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Category` — with named operations (`list`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/free-meal-sdk/releases](https://github.com/voxgig-sdk/free-meal-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "FreeMeal_sdk"

client = FreeMealSDK.new({
  "apikey" => ENV["FREE_MEAL_APIKEY"],
})
```

### 2. List category records

```ruby
begin
  # list returns an Array of Category records — iterate directly.
  categorys = client.Category.list
  categorys.each do |item|
    puts "#{item["id_category"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  categorys = client.Category.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = FreeMealSDK.test

# Entity ops return the bare mock record (raises on error).
category = client.Category.list()
puts category
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = FreeMealSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
FREE_MEAL_TEST_LIVE=TRUE
FREE_MEAL_APIKEY=<your-key>
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### FreeMealSDK

```ruby
require_relative "FreeMeal_sdk"
client = FreeMealSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = FreeMealSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### FreeMealSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `Category` | `(data) -> CategoryEntity` | Create a Category entity instance. |
| `Filter` | `(data) -> FilterEntity` | Create a Filter entity instance. |
| `Latest` | `(data) -> LatestEntity` | Create a Latest entity instance. |
| `List` | `(data) -> ListEntity` | Create a List entity instance. |
| `Lookup` | `(data) -> LookupEntity` | Create a Lookup entity instance. |
| `Random` | `(data) -> RandomEntity` | Create a Random entity instance. |
| `Randomselection` | `(data) -> RandomselectionEntity` | Create a Randomselection entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `FreeMealError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### Category

| Field | Description |
| --- | --- |
| `id_category` |  |
| `str_category` |  |
| `str_category_description` |  |
| `str_category_thumb` |  |

Operations: List.

API path: `/categories.php`

#### Filter

| Field | Description |
| --- | --- |
| `id_meal` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |

Operations: List.

API path: `/filter.php`

#### Latest

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/latest.php`

#### List

| Field | Description |
| --- | --- |
| `str_area` |  |
| `str_category` |  |
| `str_ingredient` |  |

Operations: List.

API path: `/list.php`

#### Lookup

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/lookup.php`

#### Random

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/random.php`

#### Randomselection

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/randomselection.php`

#### Search

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/search.php`



## Entities


### Category

Create an instance: `category = client.Category`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_category` | `String` |  |
| `str_category` | `String` |  |
| `str_category_description` | `String` |  |
| `str_category_thumb` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Category records (raises on error).
categorys = client.Category.list
```


### Filter

Create an instance: `filter = client.Filter`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_meal` | `String` |  |
| `str_meal` | `String` |  |
| `str_meal_thumb` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Filter records (raises on error).
filters = client.Filter.list
```


### Latest

Create an instance: `latest = client.Latest`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `String` |  |
| `id_meal` | `String` |  |
| `str_area` | `String` |  |
| `str_category` | `String` |  |
| `str_creative_commons_confirmed` | `String` |  |
| `str_drink_alternate` | `String` |  |
| `str_image_source` | `String` |  |
| `str_ingredient1` | `String` |  |
| `str_ingredient10` | `String` |  |
| `str_ingredient11` | `String` |  |
| `str_ingredient12` | `String` |  |
| `str_ingredient13` | `String` |  |
| `str_ingredient14` | `String` |  |
| `str_ingredient15` | `String` |  |
| `str_ingredient16` | `String` |  |
| `str_ingredient17` | `String` |  |
| `str_ingredient18` | `String` |  |
| `str_ingredient19` | `String` |  |
| `str_ingredient2` | `String` |  |
| `str_ingredient20` | `String` |  |
| `str_ingredient3` | `String` |  |
| `str_ingredient4` | `String` |  |
| `str_ingredient5` | `String` |  |
| `str_ingredient6` | `String` |  |
| `str_ingredient7` | `String` |  |
| `str_ingredient8` | `String` |  |
| `str_ingredient9` | `String` |  |
| `str_instruction` | `String` |  |
| `str_meal` | `String` |  |
| `str_meal_thumb` | `String` |  |
| `str_measure1` | `String` |  |
| `str_measure10` | `String` |  |
| `str_measure11` | `String` |  |
| `str_measure12` | `String` |  |
| `str_measure13` | `String` |  |
| `str_measure14` | `String` |  |
| `str_measure15` | `String` |  |
| `str_measure16` | `String` |  |
| `str_measure17` | `String` |  |
| `str_measure18` | `String` |  |
| `str_measure19` | `String` |  |
| `str_measure2` | `String` |  |
| `str_measure20` | `String` |  |
| `str_measure3` | `String` |  |
| `str_measure4` | `String` |  |
| `str_measure5` | `String` |  |
| `str_measure6` | `String` |  |
| `str_measure7` | `String` |  |
| `str_measure8` | `String` |  |
| `str_measure9` | `String` |  |
| `str_source` | `String` |  |
| `str_tag` | `String` |  |
| `str_youtube` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Latest records (raises on error).
latests = client.Latest.list
```


### List

Create an instance: `list = client.List`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `str_area` | `String` |  |
| `str_category` | `String` |  |
| `str_ingredient` | `String` |  |

#### Example: List

```ruby
# list returns an Array of List records (raises on error).
lists = client.List.list
```


### Lookup

Create an instance: `lookup = client.Lookup`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `String` |  |
| `id_meal` | `String` |  |
| `str_area` | `String` |  |
| `str_category` | `String` |  |
| `str_creative_commons_confirmed` | `String` |  |
| `str_drink_alternate` | `String` |  |
| `str_image_source` | `String` |  |
| `str_ingredient1` | `String` |  |
| `str_ingredient10` | `String` |  |
| `str_ingredient11` | `String` |  |
| `str_ingredient12` | `String` |  |
| `str_ingredient13` | `String` |  |
| `str_ingredient14` | `String` |  |
| `str_ingredient15` | `String` |  |
| `str_ingredient16` | `String` |  |
| `str_ingredient17` | `String` |  |
| `str_ingredient18` | `String` |  |
| `str_ingredient19` | `String` |  |
| `str_ingredient2` | `String` |  |
| `str_ingredient20` | `String` |  |
| `str_ingredient3` | `String` |  |
| `str_ingredient4` | `String` |  |
| `str_ingredient5` | `String` |  |
| `str_ingredient6` | `String` |  |
| `str_ingredient7` | `String` |  |
| `str_ingredient8` | `String` |  |
| `str_ingredient9` | `String` |  |
| `str_instruction` | `String` |  |
| `str_meal` | `String` |  |
| `str_meal_thumb` | `String` |  |
| `str_measure1` | `String` |  |
| `str_measure10` | `String` |  |
| `str_measure11` | `String` |  |
| `str_measure12` | `String` |  |
| `str_measure13` | `String` |  |
| `str_measure14` | `String` |  |
| `str_measure15` | `String` |  |
| `str_measure16` | `String` |  |
| `str_measure17` | `String` |  |
| `str_measure18` | `String` |  |
| `str_measure19` | `String` |  |
| `str_measure2` | `String` |  |
| `str_measure20` | `String` |  |
| `str_measure3` | `String` |  |
| `str_measure4` | `String` |  |
| `str_measure5` | `String` |  |
| `str_measure6` | `String` |  |
| `str_measure7` | `String` |  |
| `str_measure8` | `String` |  |
| `str_measure9` | `String` |  |
| `str_source` | `String` |  |
| `str_tag` | `String` |  |
| `str_youtube` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Lookup records (raises on error).
lookups = client.Lookup.list
```


### Random

Create an instance: `random = client.Random`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `String` |  |
| `id_meal` | `String` |  |
| `str_area` | `String` |  |
| `str_category` | `String` |  |
| `str_creative_commons_confirmed` | `String` |  |
| `str_drink_alternate` | `String` |  |
| `str_image_source` | `String` |  |
| `str_ingredient1` | `String` |  |
| `str_ingredient10` | `String` |  |
| `str_ingredient11` | `String` |  |
| `str_ingredient12` | `String` |  |
| `str_ingredient13` | `String` |  |
| `str_ingredient14` | `String` |  |
| `str_ingredient15` | `String` |  |
| `str_ingredient16` | `String` |  |
| `str_ingredient17` | `String` |  |
| `str_ingredient18` | `String` |  |
| `str_ingredient19` | `String` |  |
| `str_ingredient2` | `String` |  |
| `str_ingredient20` | `String` |  |
| `str_ingredient3` | `String` |  |
| `str_ingredient4` | `String` |  |
| `str_ingredient5` | `String` |  |
| `str_ingredient6` | `String` |  |
| `str_ingredient7` | `String` |  |
| `str_ingredient8` | `String` |  |
| `str_ingredient9` | `String` |  |
| `str_instruction` | `String` |  |
| `str_meal` | `String` |  |
| `str_meal_thumb` | `String` |  |
| `str_measure1` | `String` |  |
| `str_measure10` | `String` |  |
| `str_measure11` | `String` |  |
| `str_measure12` | `String` |  |
| `str_measure13` | `String` |  |
| `str_measure14` | `String` |  |
| `str_measure15` | `String` |  |
| `str_measure16` | `String` |  |
| `str_measure17` | `String` |  |
| `str_measure18` | `String` |  |
| `str_measure19` | `String` |  |
| `str_measure2` | `String` |  |
| `str_measure20` | `String` |  |
| `str_measure3` | `String` |  |
| `str_measure4` | `String` |  |
| `str_measure5` | `String` |  |
| `str_measure6` | `String` |  |
| `str_measure7` | `String` |  |
| `str_measure8` | `String` |  |
| `str_measure9` | `String` |  |
| `str_source` | `String` |  |
| `str_tag` | `String` |  |
| `str_youtube` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Random records (raises on error).
randoms = client.Random.list
```


### Randomselection

Create an instance: `randomselection = client.Randomselection`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `String` |  |
| `id_meal` | `String` |  |
| `str_area` | `String` |  |
| `str_category` | `String` |  |
| `str_creative_commons_confirmed` | `String` |  |
| `str_drink_alternate` | `String` |  |
| `str_image_source` | `String` |  |
| `str_ingredient1` | `String` |  |
| `str_ingredient10` | `String` |  |
| `str_ingredient11` | `String` |  |
| `str_ingredient12` | `String` |  |
| `str_ingredient13` | `String` |  |
| `str_ingredient14` | `String` |  |
| `str_ingredient15` | `String` |  |
| `str_ingredient16` | `String` |  |
| `str_ingredient17` | `String` |  |
| `str_ingredient18` | `String` |  |
| `str_ingredient19` | `String` |  |
| `str_ingredient2` | `String` |  |
| `str_ingredient20` | `String` |  |
| `str_ingredient3` | `String` |  |
| `str_ingredient4` | `String` |  |
| `str_ingredient5` | `String` |  |
| `str_ingredient6` | `String` |  |
| `str_ingredient7` | `String` |  |
| `str_ingredient8` | `String` |  |
| `str_ingredient9` | `String` |  |
| `str_instruction` | `String` |  |
| `str_meal` | `String` |  |
| `str_meal_thumb` | `String` |  |
| `str_measure1` | `String` |  |
| `str_measure10` | `String` |  |
| `str_measure11` | `String` |  |
| `str_measure12` | `String` |  |
| `str_measure13` | `String` |  |
| `str_measure14` | `String` |  |
| `str_measure15` | `String` |  |
| `str_measure16` | `String` |  |
| `str_measure17` | `String` |  |
| `str_measure18` | `String` |  |
| `str_measure19` | `String` |  |
| `str_measure2` | `String` |  |
| `str_measure20` | `String` |  |
| `str_measure3` | `String` |  |
| `str_measure4` | `String` |  |
| `str_measure5` | `String` |  |
| `str_measure6` | `String` |  |
| `str_measure7` | `String` |  |
| `str_measure8` | `String` |  |
| `str_measure9` | `String` |  |
| `str_source` | `String` |  |
| `str_tag` | `String` |  |
| `str_youtube` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Randomselection records (raises on error).
randomselections = client.Randomselection.list
```


### Search

Create an instance: `search = client.Search`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `String` |  |
| `id_meal` | `String` |  |
| `str_area` | `String` |  |
| `str_category` | `String` |  |
| `str_creative_commons_confirmed` | `String` |  |
| `str_drink_alternate` | `String` |  |
| `str_image_source` | `String` |  |
| `str_ingredient1` | `String` |  |
| `str_ingredient10` | `String` |  |
| `str_ingredient11` | `String` |  |
| `str_ingredient12` | `String` |  |
| `str_ingredient13` | `String` |  |
| `str_ingredient14` | `String` |  |
| `str_ingredient15` | `String` |  |
| `str_ingredient16` | `String` |  |
| `str_ingredient17` | `String` |  |
| `str_ingredient18` | `String` |  |
| `str_ingredient19` | `String` |  |
| `str_ingredient2` | `String` |  |
| `str_ingredient20` | `String` |  |
| `str_ingredient3` | `String` |  |
| `str_ingredient4` | `String` |  |
| `str_ingredient5` | `String` |  |
| `str_ingredient6` | `String` |  |
| `str_ingredient7` | `String` |  |
| `str_ingredient8` | `String` |  |
| `str_ingredient9` | `String` |  |
| `str_instruction` | `String` |  |
| `str_meal` | `String` |  |
| `str_meal_thumb` | `String` |  |
| `str_measure1` | `String` |  |
| `str_measure10` | `String` |  |
| `str_measure11` | `String` |  |
| `str_measure12` | `String` |  |
| `str_measure13` | `String` |  |
| `str_measure14` | `String` |  |
| `str_measure15` | `String` |  |
| `str_measure16` | `String` |  |
| `str_measure17` | `String` |  |
| `str_measure18` | `String` |  |
| `str_measure19` | `String` |  |
| `str_measure2` | `String` |  |
| `str_measure20` | `String` |  |
| `str_measure3` | `String` |  |
| `str_measure4` | `String` |  |
| `str_measure5` | `String` |  |
| `str_measure6` | `String` |  |
| `str_measure7` | `String` |  |
| `str_measure8` | `String` |  |
| `str_measure9` | `String` |  |
| `str_source` | `String` |  |
| `str_tag` | `String` |  |
| `str_youtube` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Search records (raises on error).
searchs = client.Search.list
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

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

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── FreeMeal_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`FreeMeal_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
category = client.Category
category.list()

# category.data_get now returns the category data from the last list
# category.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
