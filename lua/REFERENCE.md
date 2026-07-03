# FreeMeal Lua SDK Reference

Complete API reference for the FreeMeal Lua SDK.


## FreeMealSDK

### Constructor

```lua
local sdk = require("free-meal_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Category(data)`

Create a new `Category` entity instance. Pass `nil` for no initial data.

#### `Filter(data)`

Create a new `Filter` entity instance. Pass `nil` for no initial data.

#### `Latest(data)`

Create a new `Latest` entity instance. Pass `nil` for no initial data.

#### `List(data)`

Create a new `List` entity instance. Pass `nil` for no initial data.

#### `Lookup(data)`

Create a new `Lookup` entity instance. Pass `nil` for no initial data.

#### `Random(data)`

Create a new `Random` entity instance. Pass `nil` for no initial data.

#### `Randomselection(data)`

Create a new `Randomselection` entity instance. Pass `nil` for no initial data.

#### `Search(data)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## CategoryEntity

```lua
local category = client:Category(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_category` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_category_description` | ``$STRING`` | No |  |
| `str_category_thumb` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Category():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CategoryEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FilterEntity

```lua
local filter = client:Filter(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_meal` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Filter():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FilterEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## LatestEntity

```lua
local latest = client:Latest(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Latest():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LatestEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ListEntity

```lua
local list = client:List(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_ingredient` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:List():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## LookupEntity

```lua
local lookup = client:Lookup(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Lookup():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LookupEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RandomEntity

```lua
local random = client:Random(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Random():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RandomEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RandomselectionEntity

```lua
local randomselection = client:Randomselection(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Randomselection():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RandomselectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SearchEntity

```lua
local search = client:Search(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

