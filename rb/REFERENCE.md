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
| `id_category` | `String` | No |  |
| `str_category` | `String` | No |  |
| `str_category_description` | `String` | No |  |
| `str_category_thumb` | `String` | No |  |

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
| `id_meal` | `String` | No |  |
| `str_meal` | `String` | No |  |
| `str_meal_thumb` | `String` | No |  |

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
| `date_modified` | `String` | No |  |
| `id_meal` | `String` | No |  |
| `str_area` | `String` | No |  |
| `str_category` | `String` | No |  |
| `str_creative_commons_confirmed` | `String` | No |  |
| `str_drink_alternate` | `String` | No |  |
| `str_image_source` | `String` | No |  |
| `str_ingredient1` | `String` | No |  |
| `str_ingredient10` | `String` | No |  |
| `str_ingredient11` | `String` | No |  |
| `str_ingredient12` | `String` | No |  |
| `str_ingredient13` | `String` | No |  |
| `str_ingredient14` | `String` | No |  |
| `str_ingredient15` | `String` | No |  |
| `str_ingredient16` | `String` | No |  |
| `str_ingredient17` | `String` | No |  |
| `str_ingredient18` | `String` | No |  |
| `str_ingredient19` | `String` | No |  |
| `str_ingredient2` | `String` | No |  |
| `str_ingredient20` | `String` | No |  |
| `str_ingredient3` | `String` | No |  |
| `str_ingredient4` | `String` | No |  |
| `str_ingredient5` | `String` | No |  |
| `str_ingredient6` | `String` | No |  |
| `str_ingredient7` | `String` | No |  |
| `str_ingredient8` | `String` | No |  |
| `str_ingredient9` | `String` | No |  |
| `str_instruction` | `String` | No |  |
| `str_meal` | `String` | No |  |
| `str_meal_thumb` | `String` | No |  |
| `str_measure1` | `String` | No |  |
| `str_measure10` | `String` | No |  |
| `str_measure11` | `String` | No |  |
| `str_measure12` | `String` | No |  |
| `str_measure13` | `String` | No |  |
| `str_measure14` | `String` | No |  |
| `str_measure15` | `String` | No |  |
| `str_measure16` | `String` | No |  |
| `str_measure17` | `String` | No |  |
| `str_measure18` | `String` | No |  |
| `str_measure19` | `String` | No |  |
| `str_measure2` | `String` | No |  |
| `str_measure20` | `String` | No |  |
| `str_measure3` | `String` | No |  |
| `str_measure4` | `String` | No |  |
| `str_measure5` | `String` | No |  |
| `str_measure6` | `String` | No |  |
| `str_measure7` | `String` | No |  |
| `str_measure8` | `String` | No |  |
| `str_measure9` | `String` | No |  |
| `str_source` | `String` | No |  |
| `str_tag` | `String` | No |  |
| `str_youtube` | `String` | No |  |

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
| `str_area` | `String` | No |  |
| `str_category` | `String` | No |  |
| `str_ingredient` | `String` | No |  |

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
| `date_modified` | `String` | No |  |
| `id_meal` | `String` | No |  |
| `str_area` | `String` | No |  |
| `str_category` | `String` | No |  |
| `str_creative_commons_confirmed` | `String` | No |  |
| `str_drink_alternate` | `String` | No |  |
| `str_image_source` | `String` | No |  |
| `str_ingredient1` | `String` | No |  |
| `str_ingredient10` | `String` | No |  |
| `str_ingredient11` | `String` | No |  |
| `str_ingredient12` | `String` | No |  |
| `str_ingredient13` | `String` | No |  |
| `str_ingredient14` | `String` | No |  |
| `str_ingredient15` | `String` | No |  |
| `str_ingredient16` | `String` | No |  |
| `str_ingredient17` | `String` | No |  |
| `str_ingredient18` | `String` | No |  |
| `str_ingredient19` | `String` | No |  |
| `str_ingredient2` | `String` | No |  |
| `str_ingredient20` | `String` | No |  |
| `str_ingredient3` | `String` | No |  |
| `str_ingredient4` | `String` | No |  |
| `str_ingredient5` | `String` | No |  |
| `str_ingredient6` | `String` | No |  |
| `str_ingredient7` | `String` | No |  |
| `str_ingredient8` | `String` | No |  |
| `str_ingredient9` | `String` | No |  |
| `str_instruction` | `String` | No |  |
| `str_meal` | `String` | No |  |
| `str_meal_thumb` | `String` | No |  |
| `str_measure1` | `String` | No |  |
| `str_measure10` | `String` | No |  |
| `str_measure11` | `String` | No |  |
| `str_measure12` | `String` | No |  |
| `str_measure13` | `String` | No |  |
| `str_measure14` | `String` | No |  |
| `str_measure15` | `String` | No |  |
| `str_measure16` | `String` | No |  |
| `str_measure17` | `String` | No |  |
| `str_measure18` | `String` | No |  |
| `str_measure19` | `String` | No |  |
| `str_measure2` | `String` | No |  |
| `str_measure20` | `String` | No |  |
| `str_measure3` | `String` | No |  |
| `str_measure4` | `String` | No |  |
| `str_measure5` | `String` | No |  |
| `str_measure6` | `String` | No |  |
| `str_measure7` | `String` | No |  |
| `str_measure8` | `String` | No |  |
| `str_measure9` | `String` | No |  |
| `str_source` | `String` | No |  |
| `str_tag` | `String` | No |  |
| `str_youtube` | `String` | No |  |

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
| `date_modified` | `String` | No |  |
| `id_meal` | `String` | No |  |
| `str_area` | `String` | No |  |
| `str_category` | `String` | No |  |
| `str_creative_commons_confirmed` | `String` | No |  |
| `str_drink_alternate` | `String` | No |  |
| `str_image_source` | `String` | No |  |
| `str_ingredient1` | `String` | No |  |
| `str_ingredient10` | `String` | No |  |
| `str_ingredient11` | `String` | No |  |
| `str_ingredient12` | `String` | No |  |
| `str_ingredient13` | `String` | No |  |
| `str_ingredient14` | `String` | No |  |
| `str_ingredient15` | `String` | No |  |
| `str_ingredient16` | `String` | No |  |
| `str_ingredient17` | `String` | No |  |
| `str_ingredient18` | `String` | No |  |
| `str_ingredient19` | `String` | No |  |
| `str_ingredient2` | `String` | No |  |
| `str_ingredient20` | `String` | No |  |
| `str_ingredient3` | `String` | No |  |
| `str_ingredient4` | `String` | No |  |
| `str_ingredient5` | `String` | No |  |
| `str_ingredient6` | `String` | No |  |
| `str_ingredient7` | `String` | No |  |
| `str_ingredient8` | `String` | No |  |
| `str_ingredient9` | `String` | No |  |
| `str_instruction` | `String` | No |  |
| `str_meal` | `String` | No |  |
| `str_meal_thumb` | `String` | No |  |
| `str_measure1` | `String` | No |  |
| `str_measure10` | `String` | No |  |
| `str_measure11` | `String` | No |  |
| `str_measure12` | `String` | No |  |
| `str_measure13` | `String` | No |  |
| `str_measure14` | `String` | No |  |
| `str_measure15` | `String` | No |  |
| `str_measure16` | `String` | No |  |
| `str_measure17` | `String` | No |  |
| `str_measure18` | `String` | No |  |
| `str_measure19` | `String` | No |  |
| `str_measure2` | `String` | No |  |
| `str_measure20` | `String` | No |  |
| `str_measure3` | `String` | No |  |
| `str_measure4` | `String` | No |  |
| `str_measure5` | `String` | No |  |
| `str_measure6` | `String` | No |  |
| `str_measure7` | `String` | No |  |
| `str_measure8` | `String` | No |  |
| `str_measure9` | `String` | No |  |
| `str_source` | `String` | No |  |
| `str_tag` | `String` | No |  |
| `str_youtube` | `String` | No |  |

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
| `date_modified` | `String` | No |  |
| `id_meal` | `String` | No |  |
| `str_area` | `String` | No |  |
| `str_category` | `String` | No |  |
| `str_creative_commons_confirmed` | `String` | No |  |
| `str_drink_alternate` | `String` | No |  |
| `str_image_source` | `String` | No |  |
| `str_ingredient1` | `String` | No |  |
| `str_ingredient10` | `String` | No |  |
| `str_ingredient11` | `String` | No |  |
| `str_ingredient12` | `String` | No |  |
| `str_ingredient13` | `String` | No |  |
| `str_ingredient14` | `String` | No |  |
| `str_ingredient15` | `String` | No |  |
| `str_ingredient16` | `String` | No |  |
| `str_ingredient17` | `String` | No |  |
| `str_ingredient18` | `String` | No |  |
| `str_ingredient19` | `String` | No |  |
| `str_ingredient2` | `String` | No |  |
| `str_ingredient20` | `String` | No |  |
| `str_ingredient3` | `String` | No |  |
| `str_ingredient4` | `String` | No |  |
| `str_ingredient5` | `String` | No |  |
| `str_ingredient6` | `String` | No |  |
| `str_ingredient7` | `String` | No |  |
| `str_ingredient8` | `String` | No |  |
| `str_ingredient9` | `String` | No |  |
| `str_instruction` | `String` | No |  |
| `str_meal` | `String` | No |  |
| `str_meal_thumb` | `String` | No |  |
| `str_measure1` | `String` | No |  |
| `str_measure10` | `String` | No |  |
| `str_measure11` | `String` | No |  |
| `str_measure12` | `String` | No |  |
| `str_measure13` | `String` | No |  |
| `str_measure14` | `String` | No |  |
| `str_measure15` | `String` | No |  |
| `str_measure16` | `String` | No |  |
| `str_measure17` | `String` | No |  |
| `str_measure18` | `String` | No |  |
| `str_measure19` | `String` | No |  |
| `str_measure2` | `String` | No |  |
| `str_measure20` | `String` | No |  |
| `str_measure3` | `String` | No |  |
| `str_measure4` | `String` | No |  |
| `str_measure5` | `String` | No |  |
| `str_measure6` | `String` | No |  |
| `str_measure7` | `String` | No |  |
| `str_measure8` | `String` | No |  |
| `str_measure9` | `String` | No |  |
| `str_source` | `String` | No |  |
| `str_tag` | `String` | No |  |
| `str_youtube` | `String` | No |  |

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
| `date_modified` | `String` | No |  |
| `id_meal` | `String` | No |  |
| `str_area` | `String` | No |  |
| `str_category` | `String` | No |  |
| `str_creative_commons_confirmed` | `String` | No |  |
| `str_drink_alternate` | `String` | No |  |
| `str_image_source` | `String` | No |  |
| `str_ingredient1` | `String` | No |  |
| `str_ingredient10` | `String` | No |  |
| `str_ingredient11` | `String` | No |  |
| `str_ingredient12` | `String` | No |  |
| `str_ingredient13` | `String` | No |  |
| `str_ingredient14` | `String` | No |  |
| `str_ingredient15` | `String` | No |  |
| `str_ingredient16` | `String` | No |  |
| `str_ingredient17` | `String` | No |  |
| `str_ingredient18` | `String` | No |  |
| `str_ingredient19` | `String` | No |  |
| `str_ingredient2` | `String` | No |  |
| `str_ingredient20` | `String` | No |  |
| `str_ingredient3` | `String` | No |  |
| `str_ingredient4` | `String` | No |  |
| `str_ingredient5` | `String` | No |  |
| `str_ingredient6` | `String` | No |  |
| `str_ingredient7` | `String` | No |  |
| `str_ingredient8` | `String` | No |  |
| `str_ingredient9` | `String` | No |  |
| `str_instruction` | `String` | No |  |
| `str_meal` | `String` | No |  |
| `str_meal_thumb` | `String` | No |  |
| `str_measure1` | `String` | No |  |
| `str_measure10` | `String` | No |  |
| `str_measure11` | `String` | No |  |
| `str_measure12` | `String` | No |  |
| `str_measure13` | `String` | No |  |
| `str_measure14` | `String` | No |  |
| `str_measure15` | `String` | No |  |
| `str_measure16` | `String` | No |  |
| `str_measure17` | `String` | No |  |
| `str_measure18` | `String` | No |  |
| `str_measure19` | `String` | No |  |
| `str_measure2` | `String` | No |  |
| `str_measure20` | `String` | No |  |
| `str_measure3` | `String` | No |  |
| `str_measure4` | `String` | No |  |
| `str_measure5` | `String` | No |  |
| `str_measure6` | `String` | No |  |
| `str_measure7` | `String` | No |  |
| `str_measure8` | `String` | No |  |
| `str_measure9` | `String` | No |  |
| `str_source` | `String` | No |  |
| `str_tag` | `String` | No |  |
| `str_youtube` | `String` | No |  |

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

