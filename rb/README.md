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
    puts "#{item["idCategory"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  latests = client.Latest.list()
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

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
latest = client.Latest.list()
puts latest
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
| `idCategory` | Unique category identifier |
| `strCategory` | Category name |
| `strCategoryDescription` | Category description |
| `strCategoryThumb` | URL to category thumbnail image |

Operations: List.

API path: `/categories.php`

#### Filter

| Field | Description |
| --- | --- |
| `idMeal` | Unique meal identifier |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |

Operations: List.

API path: `/filter.php`

#### Latest

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/latest.php`

#### List

| Field | Description |
| --- | --- |
| `strArea` |  |
| `strCategory` |  |
| `strIngredient` |  |

Operations: List.

API path: `/list.php`

#### Lookup

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/lookup.php`

#### Random

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/random.php`

#### Randomselection

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/randomselection.php`

#### Search

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

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
| `idCategory` | `String` | Unique category identifier |
| `strCategory` | `String` | Category name |
| `strCategoryDescription` | `String` | Category description |
| `strCategoryThumb` | `String` | URL to category thumbnail image |

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
| `idMeal` | `String` | Unique meal identifier |
| `strMeal` | `String` | Meal name |
| `strMealThumb` | `String` | URL to meal thumbnail image |

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
| `dateModified` | `String` |  |
| `idMeal` | `String` | Unique meal identifier |
| `strArea` | `String` | Meal area/region |
| `strCategory` | `String` | Meal category |
| `strCreativeCommonsConfirmed` | `String` |  |
| `strDrinkAlternate` | `String` |  |
| `strImageSource` | `String` |  |
| `strIngredient1` | `String` |  |
| `strIngredient10` | `String` |  |
| `strIngredient11` | `String` |  |
| `strIngredient12` | `String` |  |
| `strIngredient13` | `String` |  |
| `strIngredient14` | `String` |  |
| `strIngredient15` | `String` |  |
| `strIngredient16` | `String` |  |
| `strIngredient17` | `String` |  |
| `strIngredient18` | `String` |  |
| `strIngredient19` | `String` |  |
| `strIngredient2` | `String` |  |
| `strIngredient20` | `String` |  |
| `strIngredient3` | `String` |  |
| `strIngredient4` | `String` |  |
| `strIngredient5` | `String` |  |
| `strIngredient6` | `String` |  |
| `strIngredient7` | `String` |  |
| `strIngredient8` | `String` |  |
| `strIngredient9` | `String` |  |
| `strInstructions` | `String` | Cooking instructions |
| `strMeal` | `String` | Meal name |
| `strMealThumb` | `String` | URL to meal thumbnail image |
| `strMeasure1` | `String` |  |
| `strMeasure10` | `String` |  |
| `strMeasure11` | `String` |  |
| `strMeasure12` | `String` |  |
| `strMeasure13` | `String` |  |
| `strMeasure14` | `String` |  |
| `strMeasure15` | `String` |  |
| `strMeasure16` | `String` |  |
| `strMeasure17` | `String` |  |
| `strMeasure18` | `String` |  |
| `strMeasure19` | `String` |  |
| `strMeasure2` | `String` |  |
| `strMeasure20` | `String` |  |
| `strMeasure3` | `String` |  |
| `strMeasure4` | `String` |  |
| `strMeasure5` | `String` |  |
| `strMeasure6` | `String` |  |
| `strMeasure7` | `String` |  |
| `strMeasure8` | `String` |  |
| `strMeasure9` | `String` |  |
| `strSource` | `String` |  |
| `strTags` | `String` | Comma-separated tags |
| `strYoutube` | `String` | YouTube video URL |

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
| `strArea` | `String` |  |
| `strCategory` | `String` |  |
| `strIngredient` | `String` |  |

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
| `dateModified` | `String` |  |
| `idMeal` | `String` | Unique meal identifier |
| `strArea` | `String` | Meal area/region |
| `strCategory` | `String` | Meal category |
| `strCreativeCommonsConfirmed` | `String` |  |
| `strDrinkAlternate` | `String` |  |
| `strImageSource` | `String` |  |
| `strIngredient1` | `String` |  |
| `strIngredient10` | `String` |  |
| `strIngredient11` | `String` |  |
| `strIngredient12` | `String` |  |
| `strIngredient13` | `String` |  |
| `strIngredient14` | `String` |  |
| `strIngredient15` | `String` |  |
| `strIngredient16` | `String` |  |
| `strIngredient17` | `String` |  |
| `strIngredient18` | `String` |  |
| `strIngredient19` | `String` |  |
| `strIngredient2` | `String` |  |
| `strIngredient20` | `String` |  |
| `strIngredient3` | `String` |  |
| `strIngredient4` | `String` |  |
| `strIngredient5` | `String` |  |
| `strIngredient6` | `String` |  |
| `strIngredient7` | `String` |  |
| `strIngredient8` | `String` |  |
| `strIngredient9` | `String` |  |
| `strInstructions` | `String` | Cooking instructions |
| `strMeal` | `String` | Meal name |
| `strMealThumb` | `String` | URL to meal thumbnail image |
| `strMeasure1` | `String` |  |
| `strMeasure10` | `String` |  |
| `strMeasure11` | `String` |  |
| `strMeasure12` | `String` |  |
| `strMeasure13` | `String` |  |
| `strMeasure14` | `String` |  |
| `strMeasure15` | `String` |  |
| `strMeasure16` | `String` |  |
| `strMeasure17` | `String` |  |
| `strMeasure18` | `String` |  |
| `strMeasure19` | `String` |  |
| `strMeasure2` | `String` |  |
| `strMeasure20` | `String` |  |
| `strMeasure3` | `String` |  |
| `strMeasure4` | `String` |  |
| `strMeasure5` | `String` |  |
| `strMeasure6` | `String` |  |
| `strMeasure7` | `String` |  |
| `strMeasure8` | `String` |  |
| `strMeasure9` | `String` |  |
| `strSource` | `String` |  |
| `strTags` | `String` | Comma-separated tags |
| `strYoutube` | `String` | YouTube video URL |

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
| `dateModified` | `String` |  |
| `idMeal` | `String` | Unique meal identifier |
| `strArea` | `String` | Meal area/region |
| `strCategory` | `String` | Meal category |
| `strCreativeCommonsConfirmed` | `String` |  |
| `strDrinkAlternate` | `String` |  |
| `strImageSource` | `String` |  |
| `strIngredient1` | `String` |  |
| `strIngredient10` | `String` |  |
| `strIngredient11` | `String` |  |
| `strIngredient12` | `String` |  |
| `strIngredient13` | `String` |  |
| `strIngredient14` | `String` |  |
| `strIngredient15` | `String` |  |
| `strIngredient16` | `String` |  |
| `strIngredient17` | `String` |  |
| `strIngredient18` | `String` |  |
| `strIngredient19` | `String` |  |
| `strIngredient2` | `String` |  |
| `strIngredient20` | `String` |  |
| `strIngredient3` | `String` |  |
| `strIngredient4` | `String` |  |
| `strIngredient5` | `String` |  |
| `strIngredient6` | `String` |  |
| `strIngredient7` | `String` |  |
| `strIngredient8` | `String` |  |
| `strIngredient9` | `String` |  |
| `strInstructions` | `String` | Cooking instructions |
| `strMeal` | `String` | Meal name |
| `strMealThumb` | `String` | URL to meal thumbnail image |
| `strMeasure1` | `String` |  |
| `strMeasure10` | `String` |  |
| `strMeasure11` | `String` |  |
| `strMeasure12` | `String` |  |
| `strMeasure13` | `String` |  |
| `strMeasure14` | `String` |  |
| `strMeasure15` | `String` |  |
| `strMeasure16` | `String` |  |
| `strMeasure17` | `String` |  |
| `strMeasure18` | `String` |  |
| `strMeasure19` | `String` |  |
| `strMeasure2` | `String` |  |
| `strMeasure20` | `String` |  |
| `strMeasure3` | `String` |  |
| `strMeasure4` | `String` |  |
| `strMeasure5` | `String` |  |
| `strMeasure6` | `String` |  |
| `strMeasure7` | `String` |  |
| `strMeasure8` | `String` |  |
| `strMeasure9` | `String` |  |
| `strSource` | `String` |  |
| `strTags` | `String` | Comma-separated tags |
| `strYoutube` | `String` | YouTube video URL |

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
| `dateModified` | `String` |  |
| `idMeal` | `String` | Unique meal identifier |
| `strArea` | `String` | Meal area/region |
| `strCategory` | `String` | Meal category |
| `strCreativeCommonsConfirmed` | `String` |  |
| `strDrinkAlternate` | `String` |  |
| `strImageSource` | `String` |  |
| `strIngredient1` | `String` |  |
| `strIngredient10` | `String` |  |
| `strIngredient11` | `String` |  |
| `strIngredient12` | `String` |  |
| `strIngredient13` | `String` |  |
| `strIngredient14` | `String` |  |
| `strIngredient15` | `String` |  |
| `strIngredient16` | `String` |  |
| `strIngredient17` | `String` |  |
| `strIngredient18` | `String` |  |
| `strIngredient19` | `String` |  |
| `strIngredient2` | `String` |  |
| `strIngredient20` | `String` |  |
| `strIngredient3` | `String` |  |
| `strIngredient4` | `String` |  |
| `strIngredient5` | `String` |  |
| `strIngredient6` | `String` |  |
| `strIngredient7` | `String` |  |
| `strIngredient8` | `String` |  |
| `strIngredient9` | `String` |  |
| `strInstructions` | `String` | Cooking instructions |
| `strMeal` | `String` | Meal name |
| `strMealThumb` | `String` | URL to meal thumbnail image |
| `strMeasure1` | `String` |  |
| `strMeasure10` | `String` |  |
| `strMeasure11` | `String` |  |
| `strMeasure12` | `String` |  |
| `strMeasure13` | `String` |  |
| `strMeasure14` | `String` |  |
| `strMeasure15` | `String` |  |
| `strMeasure16` | `String` |  |
| `strMeasure17` | `String` |  |
| `strMeasure18` | `String` |  |
| `strMeasure19` | `String` |  |
| `strMeasure2` | `String` |  |
| `strMeasure20` | `String` |  |
| `strMeasure3` | `String` |  |
| `strMeasure4` | `String` |  |
| `strMeasure5` | `String` |  |
| `strMeasure6` | `String` |  |
| `strMeasure7` | `String` |  |
| `strMeasure8` | `String` |  |
| `strMeasure9` | `String` |  |
| `strSource` | `String` |  |
| `strTags` | `String` | Comma-separated tags |
| `strYoutube` | `String` | YouTube video URL |

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
| `dateModified` | `String` |  |
| `idMeal` | `String` | Unique meal identifier |
| `strArea` | `String` | Meal area/region |
| `strCategory` | `String` | Meal category |
| `strCreativeCommonsConfirmed` | `String` |  |
| `strDrinkAlternate` | `String` |  |
| `strImageSource` | `String` |  |
| `strIngredient1` | `String` |  |
| `strIngredient10` | `String` |  |
| `strIngredient11` | `String` |  |
| `strIngredient12` | `String` |  |
| `strIngredient13` | `String` |  |
| `strIngredient14` | `String` |  |
| `strIngredient15` | `String` |  |
| `strIngredient16` | `String` |  |
| `strIngredient17` | `String` |  |
| `strIngredient18` | `String` |  |
| `strIngredient19` | `String` |  |
| `strIngredient2` | `String` |  |
| `strIngredient20` | `String` |  |
| `strIngredient3` | `String` |  |
| `strIngredient4` | `String` |  |
| `strIngredient5` | `String` |  |
| `strIngredient6` | `String` |  |
| `strIngredient7` | `String` |  |
| `strIngredient8` | `String` |  |
| `strIngredient9` | `String` |  |
| `strInstructions` | `String` | Cooking instructions |
| `strMeal` | `String` | Meal name |
| `strMealThumb` | `String` | URL to meal thumbnail image |
| `strMeasure1` | `String` |  |
| `strMeasure10` | `String` |  |
| `strMeasure11` | `String` |  |
| `strMeasure12` | `String` |  |
| `strMeasure13` | `String` |  |
| `strMeasure14` | `String` |  |
| `strMeasure15` | `String` |  |
| `strMeasure16` | `String` |  |
| `strMeasure17` | `String` |  |
| `strMeasure18` | `String` |  |
| `strMeasure19` | `String` |  |
| `strMeasure2` | `String` |  |
| `strMeasure20` | `String` |  |
| `strMeasure3` | `String` |  |
| `strMeasure4` | `String` |  |
| `strMeasure5` | `String` |  |
| `strMeasure6` | `String` |  |
| `strMeasure7` | `String` |  |
| `strMeasure8` | `String` |  |
| `strMeasure9` | `String` |  |
| `strSource` | `String` |  |
| `strTags` | `String` | Comma-separated tags |
| `strYoutube` | `String` | YouTube video URL |

#### Example: List

```ruby
# list returns an Array of Search records (raises on error).
searchs = client.Search.list
```

## Features

This SDK ships 4 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`ratelimit`](#ratelimit) | Client-side rate limiting via a token bucket |
| [`retry`](#retry) | Automatic retry of transient failures with exponential backoff |
| [`test`](#test) | In-memory mock transport for testing without a live server |
| [`timeout`](#timeout) | Per-request timeout with transport abort |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### ratelimit

Client-side rate limiting via a token bucket.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Automatic retry of transient failures with exponential backoff.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Per-request timeout with transport abort.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


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

- **RatelimitFeature**: Client-side rate limiting via a token bucket
- **RetryFeature**: Automatic retry of transient failures with exponential backoff
- **TestFeature**: In-memory mock transport for testing without a live server
- **TimeoutFeature**: Per-request timeout with transport abort

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
latest = client.Latest
latest.list()

# latest.data_get now returns the latest data from the last list
# latest.match_get returns the last match criteria
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
