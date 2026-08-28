# FreeMeal Ruby SDK Reference

Complete API reference for the FreeMeal Ruby SDK.


## FreeMealSDK

### Constructor

```ruby
require_relative 'FreeMeal_sdk'

client = FreeMealSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `FreeMealSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = FreeMealSDK.test
```


### Instance Methods

#### `Category(data = nil)`

Create a new `Category` entity instance. Pass `nil` for no initial data.

#### `Filter(data = nil)`

Create a new `Filter` entity instance. Pass `nil` for no initial data.

#### `Latest(data = nil)`

Create a new `Latest` entity instance. Pass `nil` for no initial data.

#### `List(data = nil)`

Create a new `List` entity instance. Pass `nil` for no initial data.

#### `Lookup(data = nil)`

Create a new `Lookup` entity instance. Pass `nil` for no initial data.

#### `Random(data = nil)`

Create a new `Random` entity instance. Pass `nil` for no initial data.

#### `Randomselection(data = nil)`

Create a new `Randomselection` entity instance. Pass `nil` for no initial data.

#### `Search(data = nil)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## CategoryEntity

```ruby
category = client.Category
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `idCategory` | `String` | No | Unique category identifier |
| `strCategory` | `String` | No | Category name |
| `strCategoryDescription` | `String` | No | Category description |
| `strCategoryThumb` | `String` | No | URL to category thumbnail image |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Category.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CategoryEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FilterEntity

```ruby
filter = client.Filter
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `idMeal` | `String` | No | Unique meal identifier |
| `strMeal` | `String` | No | Meal name |
| `strMealThumb` | `String` | No | URL to meal thumbnail image |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Filter.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FilterEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## LatestEntity

```ruby
latest = client.Latest
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `String` | No |  |
| `idMeal` | `String` | No | Unique meal identifier |
| `strArea` | `String` | No | Meal area/region |
| `strCategory` | `String` | No | Meal category |
| `strCreativeCommonsConfirmed` | `String` | No |  |
| `strDrinkAlternate` | `String` | No |  |
| `strImageSource` | `String` | No |  |
| `strIngredient1` | `String` | No |  |
| `strIngredient10` | `String` | No |  |
| `strIngredient11` | `String` | No |  |
| `strIngredient12` | `String` | No |  |
| `strIngredient13` | `String` | No |  |
| `strIngredient14` | `String` | No |  |
| `strIngredient15` | `String` | No |  |
| `strIngredient16` | `String` | No |  |
| `strIngredient17` | `String` | No |  |
| `strIngredient18` | `String` | No |  |
| `strIngredient19` | `String` | No |  |
| `strIngredient2` | `String` | No |  |
| `strIngredient20` | `String` | No |  |
| `strIngredient3` | `String` | No |  |
| `strIngredient4` | `String` | No |  |
| `strIngredient5` | `String` | No |  |
| `strIngredient6` | `String` | No |  |
| `strIngredient7` | `String` | No |  |
| `strIngredient8` | `String` | No |  |
| `strIngredient9` | `String` | No |  |
| `strInstructions` | `String` | No | Cooking instructions |
| `strMeal` | `String` | No | Meal name |
| `strMealThumb` | `String` | No | URL to meal thumbnail image |
| `strMeasure1` | `String` | No |  |
| `strMeasure10` | `String` | No |  |
| `strMeasure11` | `String` | No |  |
| `strMeasure12` | `String` | No |  |
| `strMeasure13` | `String` | No |  |
| `strMeasure14` | `String` | No |  |
| `strMeasure15` | `String` | No |  |
| `strMeasure16` | `String` | No |  |
| `strMeasure17` | `String` | No |  |
| `strMeasure18` | `String` | No |  |
| `strMeasure19` | `String` | No |  |
| `strMeasure2` | `String` | No |  |
| `strMeasure20` | `String` | No |  |
| `strMeasure3` | `String` | No |  |
| `strMeasure4` | `String` | No |  |
| `strMeasure5` | `String` | No |  |
| `strMeasure6` | `String` | No |  |
| `strMeasure7` | `String` | No |  |
| `strMeasure8` | `String` | No |  |
| `strMeasure9` | `String` | No |  |
| `strSource` | `String` | No |  |
| `strTags` | `String` | No | Comma-separated tags |
| `strYoutube` | `String` | No | YouTube video URL |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Latest.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `LatestEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ListEntity

```ruby
list = client.List
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `strArea` | `String` | No |  |
| `strCategory` | `String` | No |  |
| `strIngredient` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.List.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ListEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## LookupEntity

```ruby
lookup = client.Lookup
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `String` | No |  |
| `idMeal` | `String` | No | Unique meal identifier |
| `strArea` | `String` | No | Meal area/region |
| `strCategory` | `String` | No | Meal category |
| `strCreativeCommonsConfirmed` | `String` | No |  |
| `strDrinkAlternate` | `String` | No |  |
| `strImageSource` | `String` | No |  |
| `strIngredient1` | `String` | No |  |
| `strIngredient10` | `String` | No |  |
| `strIngredient11` | `String` | No |  |
| `strIngredient12` | `String` | No |  |
| `strIngredient13` | `String` | No |  |
| `strIngredient14` | `String` | No |  |
| `strIngredient15` | `String` | No |  |
| `strIngredient16` | `String` | No |  |
| `strIngredient17` | `String` | No |  |
| `strIngredient18` | `String` | No |  |
| `strIngredient19` | `String` | No |  |
| `strIngredient2` | `String` | No |  |
| `strIngredient20` | `String` | No |  |
| `strIngredient3` | `String` | No |  |
| `strIngredient4` | `String` | No |  |
| `strIngredient5` | `String` | No |  |
| `strIngredient6` | `String` | No |  |
| `strIngredient7` | `String` | No |  |
| `strIngredient8` | `String` | No |  |
| `strIngredient9` | `String` | No |  |
| `strInstructions` | `String` | No | Cooking instructions |
| `strMeal` | `String` | No | Meal name |
| `strMealThumb` | `String` | No | URL to meal thumbnail image |
| `strMeasure1` | `String` | No |  |
| `strMeasure10` | `String` | No |  |
| `strMeasure11` | `String` | No |  |
| `strMeasure12` | `String` | No |  |
| `strMeasure13` | `String` | No |  |
| `strMeasure14` | `String` | No |  |
| `strMeasure15` | `String` | No |  |
| `strMeasure16` | `String` | No |  |
| `strMeasure17` | `String` | No |  |
| `strMeasure18` | `String` | No |  |
| `strMeasure19` | `String` | No |  |
| `strMeasure2` | `String` | No |  |
| `strMeasure20` | `String` | No |  |
| `strMeasure3` | `String` | No |  |
| `strMeasure4` | `String` | No |  |
| `strMeasure5` | `String` | No |  |
| `strMeasure6` | `String` | No |  |
| `strMeasure7` | `String` | No |  |
| `strMeasure8` | `String` | No |  |
| `strMeasure9` | `String` | No |  |
| `strSource` | `String` | No |  |
| `strTags` | `String` | No | Comma-separated tags |
| `strYoutube` | `String` | No | YouTube video URL |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Lookup.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `LookupEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RandomEntity

```ruby
random = client.Random
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `String` | No |  |
| `idMeal` | `String` | No | Unique meal identifier |
| `strArea` | `String` | No | Meal area/region |
| `strCategory` | `String` | No | Meal category |
| `strCreativeCommonsConfirmed` | `String` | No |  |
| `strDrinkAlternate` | `String` | No |  |
| `strImageSource` | `String` | No |  |
| `strIngredient1` | `String` | No |  |
| `strIngredient10` | `String` | No |  |
| `strIngredient11` | `String` | No |  |
| `strIngredient12` | `String` | No |  |
| `strIngredient13` | `String` | No |  |
| `strIngredient14` | `String` | No |  |
| `strIngredient15` | `String` | No |  |
| `strIngredient16` | `String` | No |  |
| `strIngredient17` | `String` | No |  |
| `strIngredient18` | `String` | No |  |
| `strIngredient19` | `String` | No |  |
| `strIngredient2` | `String` | No |  |
| `strIngredient20` | `String` | No |  |
| `strIngredient3` | `String` | No |  |
| `strIngredient4` | `String` | No |  |
| `strIngredient5` | `String` | No |  |
| `strIngredient6` | `String` | No |  |
| `strIngredient7` | `String` | No |  |
| `strIngredient8` | `String` | No |  |
| `strIngredient9` | `String` | No |  |
| `strInstructions` | `String` | No | Cooking instructions |
| `strMeal` | `String` | No | Meal name |
| `strMealThumb` | `String` | No | URL to meal thumbnail image |
| `strMeasure1` | `String` | No |  |
| `strMeasure10` | `String` | No |  |
| `strMeasure11` | `String` | No |  |
| `strMeasure12` | `String` | No |  |
| `strMeasure13` | `String` | No |  |
| `strMeasure14` | `String` | No |  |
| `strMeasure15` | `String` | No |  |
| `strMeasure16` | `String` | No |  |
| `strMeasure17` | `String` | No |  |
| `strMeasure18` | `String` | No |  |
| `strMeasure19` | `String` | No |  |
| `strMeasure2` | `String` | No |  |
| `strMeasure20` | `String` | No |  |
| `strMeasure3` | `String` | No |  |
| `strMeasure4` | `String` | No |  |
| `strMeasure5` | `String` | No |  |
| `strMeasure6` | `String` | No |  |
| `strMeasure7` | `String` | No |  |
| `strMeasure8` | `String` | No |  |
| `strMeasure9` | `String` | No |  |
| `strSource` | `String` | No |  |
| `strTags` | `String` | No | Comma-separated tags |
| `strYoutube` | `String` | No | YouTube video URL |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Random.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RandomEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RandomselectionEntity

```ruby
randomselection = client.Randomselection
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `String` | No |  |
| `idMeal` | `String` | No | Unique meal identifier |
| `strArea` | `String` | No | Meal area/region |
| `strCategory` | `String` | No | Meal category |
| `strCreativeCommonsConfirmed` | `String` | No |  |
| `strDrinkAlternate` | `String` | No |  |
| `strImageSource` | `String` | No |  |
| `strIngredient1` | `String` | No |  |
| `strIngredient10` | `String` | No |  |
| `strIngredient11` | `String` | No |  |
| `strIngredient12` | `String` | No |  |
| `strIngredient13` | `String` | No |  |
| `strIngredient14` | `String` | No |  |
| `strIngredient15` | `String` | No |  |
| `strIngredient16` | `String` | No |  |
| `strIngredient17` | `String` | No |  |
| `strIngredient18` | `String` | No |  |
| `strIngredient19` | `String` | No |  |
| `strIngredient2` | `String` | No |  |
| `strIngredient20` | `String` | No |  |
| `strIngredient3` | `String` | No |  |
| `strIngredient4` | `String` | No |  |
| `strIngredient5` | `String` | No |  |
| `strIngredient6` | `String` | No |  |
| `strIngredient7` | `String` | No |  |
| `strIngredient8` | `String` | No |  |
| `strIngredient9` | `String` | No |  |
| `strInstructions` | `String` | No | Cooking instructions |
| `strMeal` | `String` | No | Meal name |
| `strMealThumb` | `String` | No | URL to meal thumbnail image |
| `strMeasure1` | `String` | No |  |
| `strMeasure10` | `String` | No |  |
| `strMeasure11` | `String` | No |  |
| `strMeasure12` | `String` | No |  |
| `strMeasure13` | `String` | No |  |
| `strMeasure14` | `String` | No |  |
| `strMeasure15` | `String` | No |  |
| `strMeasure16` | `String` | No |  |
| `strMeasure17` | `String` | No |  |
| `strMeasure18` | `String` | No |  |
| `strMeasure19` | `String` | No |  |
| `strMeasure2` | `String` | No |  |
| `strMeasure20` | `String` | No |  |
| `strMeasure3` | `String` | No |  |
| `strMeasure4` | `String` | No |  |
| `strMeasure5` | `String` | No |  |
| `strMeasure6` | `String` | No |  |
| `strMeasure7` | `String` | No |  |
| `strMeasure8` | `String` | No |  |
| `strMeasure9` | `String` | No |  |
| `strSource` | `String` | No |  |
| `strTags` | `String` | No | Comma-separated tags |
| `strYoutube` | `String` | No | YouTube video URL |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Randomselection.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RandomselectionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SearchEntity

```ruby
search = client.Search
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `String` | No |  |
| `idMeal` | `String` | No | Unique meal identifier |
| `strArea` | `String` | No | Meal area/region |
| `strCategory` | `String` | No | Meal category |
| `strCreativeCommonsConfirmed` | `String` | No |  |
| `strDrinkAlternate` | `String` | No |  |
| `strImageSource` | `String` | No |  |
| `strIngredient1` | `String` | No |  |
| `strIngredient10` | `String` | No |  |
| `strIngredient11` | `String` | No |  |
| `strIngredient12` | `String` | No |  |
| `strIngredient13` | `String` | No |  |
| `strIngredient14` | `String` | No |  |
| `strIngredient15` | `String` | No |  |
| `strIngredient16` | `String` | No |  |
| `strIngredient17` | `String` | No |  |
| `strIngredient18` | `String` | No |  |
| `strIngredient19` | `String` | No |  |
| `strIngredient2` | `String` | No |  |
| `strIngredient20` | `String` | No |  |
| `strIngredient3` | `String` | No |  |
| `strIngredient4` | `String` | No |  |
| `strIngredient5` | `String` | No |  |
| `strIngredient6` | `String` | No |  |
| `strIngredient7` | `String` | No |  |
| `strIngredient8` | `String` | No |  |
| `strIngredient9` | `String` | No |  |
| `strInstructions` | `String` | No | Cooking instructions |
| `strMeal` | `String` | No | Meal name |
| `strMealThumb` | `String` | No | URL to meal thumbnail image |
| `strMeasure1` | `String` | No |  |
| `strMeasure10` | `String` | No |  |
| `strMeasure11` | `String` | No |  |
| `strMeasure12` | `String` | No |  |
| `strMeasure13` | `String` | No |  |
| `strMeasure14` | `String` | No |  |
| `strMeasure15` | `String` | No |  |
| `strMeasure16` | `String` | No |  |
| `strMeasure17` | `String` | No |  |
| `strMeasure18` | `String` | No |  |
| `strMeasure19` | `String` | No |  |
| `strMeasure2` | `String` | No |  |
| `strMeasure20` | `String` | No |  |
| `strMeasure3` | `String` | No |  |
| `strMeasure4` | `String` | No |  |
| `strMeasure5` | `String` | No |  |
| `strMeasure6` | `String` | No |  |
| `strMeasure7` | `String` | No |  |
| `strMeasure8` | `String` | No |  |
| `strMeasure9` | `String` | No |  |
| `strSource` | `String` | No |  |
| `strTags` | `String` | No | Comma-separated tags |
| `strYoutube` | `String` | No | YouTube video URL |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Search.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = FreeMealSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

