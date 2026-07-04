# FreeMeal Ruby SDK Reference

Complete API reference for the FreeMeal Ruby SDK.


## FreeMealSDK

### Constructor

```ruby
require_relative 'free-meal_sdk'

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
category = client.category
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_category` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_category_description` | ``$STRING`` | No |  |
| `str_category_thumb` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.category.list(nil)
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
filter = client.filter
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_meal` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.filter.list(nil)
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
latest = client.latest
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | ``$STRING`` | No |  |
| `id_meal` | ``$STRING`` | No |  |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_creative_commons_confirmed` | ``$STRING`` | No |  |
| `str_drink_alternate` | ``$STRING`` | No |  |
| `str_image_source` | ``$STRING`` | No |  |
| `str_ingredient1` | ``$STRING`` | No |  |
| `str_ingredient10` | ``$STRING`` | No |  |
| `str_ingredient11` | ``$STRING`` | No |  |
| `str_ingredient12` | ``$STRING`` | No |  |
| `str_ingredient13` | ``$STRING`` | No |  |
| `str_ingredient14` | ``$STRING`` | No |  |
| `str_ingredient15` | ``$STRING`` | No |  |
| `str_ingredient16` | ``$STRING`` | No |  |
| `str_ingredient17` | ``$STRING`` | No |  |
| `str_ingredient18` | ``$STRING`` | No |  |
| `str_ingredient19` | ``$STRING`` | No |  |
| `str_ingredient2` | ``$STRING`` | No |  |
| `str_ingredient20` | ``$STRING`` | No |  |
| `str_ingredient3` | ``$STRING`` | No |  |
| `str_ingredient4` | ``$STRING`` | No |  |
| `str_ingredient5` | ``$STRING`` | No |  |
| `str_ingredient6` | ``$STRING`` | No |  |
| `str_ingredient7` | ``$STRING`` | No |  |
| `str_ingredient8` | ``$STRING`` | No |  |
| `str_ingredient9` | ``$STRING`` | No |  |
| `str_instruction` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |
| `str_measure1` | ``$STRING`` | No |  |
| `str_measure10` | ``$STRING`` | No |  |
| `str_measure11` | ``$STRING`` | No |  |
| `str_measure12` | ``$STRING`` | No |  |
| `str_measure13` | ``$STRING`` | No |  |
| `str_measure14` | ``$STRING`` | No |  |
| `str_measure15` | ``$STRING`` | No |  |
| `str_measure16` | ``$STRING`` | No |  |
| `str_measure17` | ``$STRING`` | No |  |
| `str_measure18` | ``$STRING`` | No |  |
| `str_measure19` | ``$STRING`` | No |  |
| `str_measure2` | ``$STRING`` | No |  |
| `str_measure20` | ``$STRING`` | No |  |
| `str_measure3` | ``$STRING`` | No |  |
| `str_measure4` | ``$STRING`` | No |  |
| `str_measure5` | ``$STRING`` | No |  |
| `str_measure6` | ``$STRING`` | No |  |
| `str_measure7` | ``$STRING`` | No |  |
| `str_measure8` | ``$STRING`` | No |  |
| `str_measure9` | ``$STRING`` | No |  |
| `str_source` | ``$STRING`` | No |  |
| `str_tag` | ``$STRING`` | No |  |
| `str_youtube` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.latest.list(nil)
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
list = client.list
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_ingredient` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.list.list(nil)
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
lookup = client.lookup
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | ``$STRING`` | No |  |
| `id_meal` | ``$STRING`` | No |  |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_creative_commons_confirmed` | ``$STRING`` | No |  |
| `str_drink_alternate` | ``$STRING`` | No |  |
| `str_image_source` | ``$STRING`` | No |  |
| `str_ingredient1` | ``$STRING`` | No |  |
| `str_ingredient10` | ``$STRING`` | No |  |
| `str_ingredient11` | ``$STRING`` | No |  |
| `str_ingredient12` | ``$STRING`` | No |  |
| `str_ingredient13` | ``$STRING`` | No |  |
| `str_ingredient14` | ``$STRING`` | No |  |
| `str_ingredient15` | ``$STRING`` | No |  |
| `str_ingredient16` | ``$STRING`` | No |  |
| `str_ingredient17` | ``$STRING`` | No |  |
| `str_ingredient18` | ``$STRING`` | No |  |
| `str_ingredient19` | ``$STRING`` | No |  |
| `str_ingredient2` | ``$STRING`` | No |  |
| `str_ingredient20` | ``$STRING`` | No |  |
| `str_ingredient3` | ``$STRING`` | No |  |
| `str_ingredient4` | ``$STRING`` | No |  |
| `str_ingredient5` | ``$STRING`` | No |  |
| `str_ingredient6` | ``$STRING`` | No |  |
| `str_ingredient7` | ``$STRING`` | No |  |
| `str_ingredient8` | ``$STRING`` | No |  |
| `str_ingredient9` | ``$STRING`` | No |  |
| `str_instruction` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |
| `str_measure1` | ``$STRING`` | No |  |
| `str_measure10` | ``$STRING`` | No |  |
| `str_measure11` | ``$STRING`` | No |  |
| `str_measure12` | ``$STRING`` | No |  |
| `str_measure13` | ``$STRING`` | No |  |
| `str_measure14` | ``$STRING`` | No |  |
| `str_measure15` | ``$STRING`` | No |  |
| `str_measure16` | ``$STRING`` | No |  |
| `str_measure17` | ``$STRING`` | No |  |
| `str_measure18` | ``$STRING`` | No |  |
| `str_measure19` | ``$STRING`` | No |  |
| `str_measure2` | ``$STRING`` | No |  |
| `str_measure20` | ``$STRING`` | No |  |
| `str_measure3` | ``$STRING`` | No |  |
| `str_measure4` | ``$STRING`` | No |  |
| `str_measure5` | ``$STRING`` | No |  |
| `str_measure6` | ``$STRING`` | No |  |
| `str_measure7` | ``$STRING`` | No |  |
| `str_measure8` | ``$STRING`` | No |  |
| `str_measure9` | ``$STRING`` | No |  |
| `str_source` | ``$STRING`` | No |  |
| `str_tag` | ``$STRING`` | No |  |
| `str_youtube` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.lookup.list(nil)
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
random = client.random
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | ``$STRING`` | No |  |
| `id_meal` | ``$STRING`` | No |  |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_creative_commons_confirmed` | ``$STRING`` | No |  |
| `str_drink_alternate` | ``$STRING`` | No |  |
| `str_image_source` | ``$STRING`` | No |  |
| `str_ingredient1` | ``$STRING`` | No |  |
| `str_ingredient10` | ``$STRING`` | No |  |
| `str_ingredient11` | ``$STRING`` | No |  |
| `str_ingredient12` | ``$STRING`` | No |  |
| `str_ingredient13` | ``$STRING`` | No |  |
| `str_ingredient14` | ``$STRING`` | No |  |
| `str_ingredient15` | ``$STRING`` | No |  |
| `str_ingredient16` | ``$STRING`` | No |  |
| `str_ingredient17` | ``$STRING`` | No |  |
| `str_ingredient18` | ``$STRING`` | No |  |
| `str_ingredient19` | ``$STRING`` | No |  |
| `str_ingredient2` | ``$STRING`` | No |  |
| `str_ingredient20` | ``$STRING`` | No |  |
| `str_ingredient3` | ``$STRING`` | No |  |
| `str_ingredient4` | ``$STRING`` | No |  |
| `str_ingredient5` | ``$STRING`` | No |  |
| `str_ingredient6` | ``$STRING`` | No |  |
| `str_ingredient7` | ``$STRING`` | No |  |
| `str_ingredient8` | ``$STRING`` | No |  |
| `str_ingredient9` | ``$STRING`` | No |  |
| `str_instruction` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |
| `str_measure1` | ``$STRING`` | No |  |
| `str_measure10` | ``$STRING`` | No |  |
| `str_measure11` | ``$STRING`` | No |  |
| `str_measure12` | ``$STRING`` | No |  |
| `str_measure13` | ``$STRING`` | No |  |
| `str_measure14` | ``$STRING`` | No |  |
| `str_measure15` | ``$STRING`` | No |  |
| `str_measure16` | ``$STRING`` | No |  |
| `str_measure17` | ``$STRING`` | No |  |
| `str_measure18` | ``$STRING`` | No |  |
| `str_measure19` | ``$STRING`` | No |  |
| `str_measure2` | ``$STRING`` | No |  |
| `str_measure20` | ``$STRING`` | No |  |
| `str_measure3` | ``$STRING`` | No |  |
| `str_measure4` | ``$STRING`` | No |  |
| `str_measure5` | ``$STRING`` | No |  |
| `str_measure6` | ``$STRING`` | No |  |
| `str_measure7` | ``$STRING`` | No |  |
| `str_measure8` | ``$STRING`` | No |  |
| `str_measure9` | ``$STRING`` | No |  |
| `str_source` | ``$STRING`` | No |  |
| `str_tag` | ``$STRING`` | No |  |
| `str_youtube` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.random.list(nil)
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
randomselection = client.randomselection
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | ``$STRING`` | No |  |
| `id_meal` | ``$STRING`` | No |  |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_creative_commons_confirmed` | ``$STRING`` | No |  |
| `str_drink_alternate` | ``$STRING`` | No |  |
| `str_image_source` | ``$STRING`` | No |  |
| `str_ingredient1` | ``$STRING`` | No |  |
| `str_ingredient10` | ``$STRING`` | No |  |
| `str_ingredient11` | ``$STRING`` | No |  |
| `str_ingredient12` | ``$STRING`` | No |  |
| `str_ingredient13` | ``$STRING`` | No |  |
| `str_ingredient14` | ``$STRING`` | No |  |
| `str_ingredient15` | ``$STRING`` | No |  |
| `str_ingredient16` | ``$STRING`` | No |  |
| `str_ingredient17` | ``$STRING`` | No |  |
| `str_ingredient18` | ``$STRING`` | No |  |
| `str_ingredient19` | ``$STRING`` | No |  |
| `str_ingredient2` | ``$STRING`` | No |  |
| `str_ingredient20` | ``$STRING`` | No |  |
| `str_ingredient3` | ``$STRING`` | No |  |
| `str_ingredient4` | ``$STRING`` | No |  |
| `str_ingredient5` | ``$STRING`` | No |  |
| `str_ingredient6` | ``$STRING`` | No |  |
| `str_ingredient7` | ``$STRING`` | No |  |
| `str_ingredient8` | ``$STRING`` | No |  |
| `str_ingredient9` | ``$STRING`` | No |  |
| `str_instruction` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |
| `str_measure1` | ``$STRING`` | No |  |
| `str_measure10` | ``$STRING`` | No |  |
| `str_measure11` | ``$STRING`` | No |  |
| `str_measure12` | ``$STRING`` | No |  |
| `str_measure13` | ``$STRING`` | No |  |
| `str_measure14` | ``$STRING`` | No |  |
| `str_measure15` | ``$STRING`` | No |  |
| `str_measure16` | ``$STRING`` | No |  |
| `str_measure17` | ``$STRING`` | No |  |
| `str_measure18` | ``$STRING`` | No |  |
| `str_measure19` | ``$STRING`` | No |  |
| `str_measure2` | ``$STRING`` | No |  |
| `str_measure20` | ``$STRING`` | No |  |
| `str_measure3` | ``$STRING`` | No |  |
| `str_measure4` | ``$STRING`` | No |  |
| `str_measure5` | ``$STRING`` | No |  |
| `str_measure6` | ``$STRING`` | No |  |
| `str_measure7` | ``$STRING`` | No |  |
| `str_measure8` | ``$STRING`` | No |  |
| `str_measure9` | ``$STRING`` | No |  |
| `str_source` | ``$STRING`` | No |  |
| `str_tag` | ``$STRING`` | No |  |
| `str_youtube` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.randomselection.list(nil)
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
search = client.search
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | ``$STRING`` | No |  |
| `id_meal` | ``$STRING`` | No |  |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_creative_commons_confirmed` | ``$STRING`` | No |  |
| `str_drink_alternate` | ``$STRING`` | No |  |
| `str_image_source` | ``$STRING`` | No |  |
| `str_ingredient1` | ``$STRING`` | No |  |
| `str_ingredient10` | ``$STRING`` | No |  |
| `str_ingredient11` | ``$STRING`` | No |  |
| `str_ingredient12` | ``$STRING`` | No |  |
| `str_ingredient13` | ``$STRING`` | No |  |
| `str_ingredient14` | ``$STRING`` | No |  |
| `str_ingredient15` | ``$STRING`` | No |  |
| `str_ingredient16` | ``$STRING`` | No |  |
| `str_ingredient17` | ``$STRING`` | No |  |
| `str_ingredient18` | ``$STRING`` | No |  |
| `str_ingredient19` | ``$STRING`` | No |  |
| `str_ingredient2` | ``$STRING`` | No |  |
| `str_ingredient20` | ``$STRING`` | No |  |
| `str_ingredient3` | ``$STRING`` | No |  |
| `str_ingredient4` | ``$STRING`` | No |  |
| `str_ingredient5` | ``$STRING`` | No |  |
| `str_ingredient6` | ``$STRING`` | No |  |
| `str_ingredient7` | ``$STRING`` | No |  |
| `str_ingredient8` | ``$STRING`` | No |  |
| `str_ingredient9` | ``$STRING`` | No |  |
| `str_instruction` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |
| `str_measure1` | ``$STRING`` | No |  |
| `str_measure10` | ``$STRING`` | No |  |
| `str_measure11` | ``$STRING`` | No |  |
| `str_measure12` | ``$STRING`` | No |  |
| `str_measure13` | ``$STRING`` | No |  |
| `str_measure14` | ``$STRING`` | No |  |
| `str_measure15` | ``$STRING`` | No |  |
| `str_measure16` | ``$STRING`` | No |  |
| `str_measure17` | ``$STRING`` | No |  |
| `str_measure18` | ``$STRING`` | No |  |
| `str_measure19` | ``$STRING`` | No |  |
| `str_measure2` | ``$STRING`` | No |  |
| `str_measure20` | ``$STRING`` | No |  |
| `str_measure3` | ``$STRING`` | No |  |
| `str_measure4` | ``$STRING`` | No |  |
| `str_measure5` | ``$STRING`` | No |  |
| `str_measure6` | ``$STRING`` | No |  |
| `str_measure7` | ``$STRING`` | No |  |
| `str_measure8` | ``$STRING`` | No |  |
| `str_measure9` | ``$STRING`` | No |  |
| `str_source` | ``$STRING`` | No |  |
| `str_tag` | ``$STRING`` | No |  |
| `str_youtube` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.search.list(nil)
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

