# FreeMeal Golang SDK Reference

Complete API reference for the FreeMeal Golang SDK.


## FreeMealSDK

### Constructor

```go
func NewFreeMealSDK(options map[string]any) *FreeMealSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `TestSDK(testopts, sdkopts map[string]any) *FreeMealSDK`

Create a test client with mock features active. Both arguments may be `nil`.

```go
client := sdk.TestSDK(nil, nil)
```


### Instance Methods

#### `Category(data map[string]any) FreeMealEntity`

Create a new `Category` entity instance. Pass `nil` for no initial data.

#### `Filter(data map[string]any) FreeMealEntity`

Create a new `Filter` entity instance. Pass `nil` for no initial data.

#### `Latest(data map[string]any) FreeMealEntity`

Create a new `Latest` entity instance. Pass `nil` for no initial data.

#### `List(data map[string]any) FreeMealEntity`

Create a new `List` entity instance. Pass `nil` for no initial data.

#### `Lookup(data map[string]any) FreeMealEntity`

Create a new `Lookup` entity instance. Pass `nil` for no initial data.

#### `Random(data map[string]any) FreeMealEntity`

Create a new `Random` entity instance. Pass `nil` for no initial data.

#### `Randomselection(data map[string]any) FreeMealEntity`

Create a new `Randomselection` entity instance. Pass `nil` for no initial data.

#### `Search(data map[string]any) FreeMealEntity`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## CategoryEntity

```go
category := client.Category(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_category` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_category_description` | ``$STRING`` | No |  |
| `str_category_thumb` | ``$STRING`` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Category(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CategoryEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FilterEntity

```go
filter := client.Filter(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_meal` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Filter(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FilterEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## LatestEntity

```go
latest := client.Latest(nil)
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Latest(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LatestEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ListEntity

```go
list := client.List(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_ingredient` | ``$STRING`` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.List(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ListEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## LookupEntity

```go
lookup := client.Lookup(nil)
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Lookup(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LookupEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RandomEntity

```go
random := client.Random(nil)
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Random(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RandomEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RandomselectionEntity

```go
randomselection := client.Randomselection(nil)
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Randomselection(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RandomselectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SearchEntity

```go
search := client.Search(nil)
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewFreeMealSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

