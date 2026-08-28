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
| `idCategory` | `string` | No | Unique category identifier |
| `strCategory` | `string` | No | Category name |
| `strCategoryDescription` | `string` | No | Category description |
| `strCategoryThumb` | `string` | No | URL to category thumbnail image |

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
| `idMeal` | `string` | No | Unique meal identifier |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |

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
| `dateModified` | `string` | No |  |
| `idMeal` | `string` | No | Unique meal identifier |
| `strArea` | `string` | No | Meal area/region |
| `strCategory` | `string` | No | Meal category |
| `strCreativeCommonsConfirmed` | `string` | No |  |
| `strDrinkAlternate` | `string` | No |  |
| `strImageSource` | `string` | No |  |
| `strIngredient1` | `string` | No |  |
| `strIngredient10` | `string` | No |  |
| `strIngredient11` | `string` | No |  |
| `strIngredient12` | `string` | No |  |
| `strIngredient13` | `string` | No |  |
| `strIngredient14` | `string` | No |  |
| `strIngredient15` | `string` | No |  |
| `strIngredient16` | `string` | No |  |
| `strIngredient17` | `string` | No |  |
| `strIngredient18` | `string` | No |  |
| `strIngredient19` | `string` | No |  |
| `strIngredient2` | `string` | No |  |
| `strIngredient20` | `string` | No |  |
| `strIngredient3` | `string` | No |  |
| `strIngredient4` | `string` | No |  |
| `strIngredient5` | `string` | No |  |
| `strIngredient6` | `string` | No |  |
| `strIngredient7` | `string` | No |  |
| `strIngredient8` | `string` | No |  |
| `strIngredient9` | `string` | No |  |
| `strInstructions` | `string` | No | Cooking instructions |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |
| `strMeasure1` | `string` | No |  |
| `strMeasure10` | `string` | No |  |
| `strMeasure11` | `string` | No |  |
| `strMeasure12` | `string` | No |  |
| `strMeasure13` | `string` | No |  |
| `strMeasure14` | `string` | No |  |
| `strMeasure15` | `string` | No |  |
| `strMeasure16` | `string` | No |  |
| `strMeasure17` | `string` | No |  |
| `strMeasure18` | `string` | No |  |
| `strMeasure19` | `string` | No |  |
| `strMeasure2` | `string` | No |  |
| `strMeasure20` | `string` | No |  |
| `strMeasure3` | `string` | No |  |
| `strMeasure4` | `string` | No |  |
| `strMeasure5` | `string` | No |  |
| `strMeasure6` | `string` | No |  |
| `strMeasure7` | `string` | No |  |
| `strMeasure8` | `string` | No |  |
| `strMeasure9` | `string` | No |  |
| `strSource` | `string` | No |  |
| `strTags` | `string` | No | Comma-separated tags |
| `strYoutube` | `string` | No | YouTube video URL |

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
| `strArea` | `string` | No |  |
| `strCategory` | `string` | No |  |
| `strIngredient` | `string` | No |  |

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
| `dateModified` | `string` | No |  |
| `idMeal` | `string` | No | Unique meal identifier |
| `strArea` | `string` | No | Meal area/region |
| `strCategory` | `string` | No | Meal category |
| `strCreativeCommonsConfirmed` | `string` | No |  |
| `strDrinkAlternate` | `string` | No |  |
| `strImageSource` | `string` | No |  |
| `strIngredient1` | `string` | No |  |
| `strIngredient10` | `string` | No |  |
| `strIngredient11` | `string` | No |  |
| `strIngredient12` | `string` | No |  |
| `strIngredient13` | `string` | No |  |
| `strIngredient14` | `string` | No |  |
| `strIngredient15` | `string` | No |  |
| `strIngredient16` | `string` | No |  |
| `strIngredient17` | `string` | No |  |
| `strIngredient18` | `string` | No |  |
| `strIngredient19` | `string` | No |  |
| `strIngredient2` | `string` | No |  |
| `strIngredient20` | `string` | No |  |
| `strIngredient3` | `string` | No |  |
| `strIngredient4` | `string` | No |  |
| `strIngredient5` | `string` | No |  |
| `strIngredient6` | `string` | No |  |
| `strIngredient7` | `string` | No |  |
| `strIngredient8` | `string` | No |  |
| `strIngredient9` | `string` | No |  |
| `strInstructions` | `string` | No | Cooking instructions |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |
| `strMeasure1` | `string` | No |  |
| `strMeasure10` | `string` | No |  |
| `strMeasure11` | `string` | No |  |
| `strMeasure12` | `string` | No |  |
| `strMeasure13` | `string` | No |  |
| `strMeasure14` | `string` | No |  |
| `strMeasure15` | `string` | No |  |
| `strMeasure16` | `string` | No |  |
| `strMeasure17` | `string` | No |  |
| `strMeasure18` | `string` | No |  |
| `strMeasure19` | `string` | No |  |
| `strMeasure2` | `string` | No |  |
| `strMeasure20` | `string` | No |  |
| `strMeasure3` | `string` | No |  |
| `strMeasure4` | `string` | No |  |
| `strMeasure5` | `string` | No |  |
| `strMeasure6` | `string` | No |  |
| `strMeasure7` | `string` | No |  |
| `strMeasure8` | `string` | No |  |
| `strMeasure9` | `string` | No |  |
| `strSource` | `string` | No |  |
| `strTags` | `string` | No | Comma-separated tags |
| `strYoutube` | `string` | No | YouTube video URL |

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
| `dateModified` | `string` | No |  |
| `idMeal` | `string` | No | Unique meal identifier |
| `strArea` | `string` | No | Meal area/region |
| `strCategory` | `string` | No | Meal category |
| `strCreativeCommonsConfirmed` | `string` | No |  |
| `strDrinkAlternate` | `string` | No |  |
| `strImageSource` | `string` | No |  |
| `strIngredient1` | `string` | No |  |
| `strIngredient10` | `string` | No |  |
| `strIngredient11` | `string` | No |  |
| `strIngredient12` | `string` | No |  |
| `strIngredient13` | `string` | No |  |
| `strIngredient14` | `string` | No |  |
| `strIngredient15` | `string` | No |  |
| `strIngredient16` | `string` | No |  |
| `strIngredient17` | `string` | No |  |
| `strIngredient18` | `string` | No |  |
| `strIngredient19` | `string` | No |  |
| `strIngredient2` | `string` | No |  |
| `strIngredient20` | `string` | No |  |
| `strIngredient3` | `string` | No |  |
| `strIngredient4` | `string` | No |  |
| `strIngredient5` | `string` | No |  |
| `strIngredient6` | `string` | No |  |
| `strIngredient7` | `string` | No |  |
| `strIngredient8` | `string` | No |  |
| `strIngredient9` | `string` | No |  |
| `strInstructions` | `string` | No | Cooking instructions |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |
| `strMeasure1` | `string` | No |  |
| `strMeasure10` | `string` | No |  |
| `strMeasure11` | `string` | No |  |
| `strMeasure12` | `string` | No |  |
| `strMeasure13` | `string` | No |  |
| `strMeasure14` | `string` | No |  |
| `strMeasure15` | `string` | No |  |
| `strMeasure16` | `string` | No |  |
| `strMeasure17` | `string` | No |  |
| `strMeasure18` | `string` | No |  |
| `strMeasure19` | `string` | No |  |
| `strMeasure2` | `string` | No |  |
| `strMeasure20` | `string` | No |  |
| `strMeasure3` | `string` | No |  |
| `strMeasure4` | `string` | No |  |
| `strMeasure5` | `string` | No |  |
| `strMeasure6` | `string` | No |  |
| `strMeasure7` | `string` | No |  |
| `strMeasure8` | `string` | No |  |
| `strMeasure9` | `string` | No |  |
| `strSource` | `string` | No |  |
| `strTags` | `string` | No | Comma-separated tags |
| `strYoutube` | `string` | No | YouTube video URL |

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
| `dateModified` | `string` | No |  |
| `idMeal` | `string` | No | Unique meal identifier |
| `strArea` | `string` | No | Meal area/region |
| `strCategory` | `string` | No | Meal category |
| `strCreativeCommonsConfirmed` | `string` | No |  |
| `strDrinkAlternate` | `string` | No |  |
| `strImageSource` | `string` | No |  |
| `strIngredient1` | `string` | No |  |
| `strIngredient10` | `string` | No |  |
| `strIngredient11` | `string` | No |  |
| `strIngredient12` | `string` | No |  |
| `strIngredient13` | `string` | No |  |
| `strIngredient14` | `string` | No |  |
| `strIngredient15` | `string` | No |  |
| `strIngredient16` | `string` | No |  |
| `strIngredient17` | `string` | No |  |
| `strIngredient18` | `string` | No |  |
| `strIngredient19` | `string` | No |  |
| `strIngredient2` | `string` | No |  |
| `strIngredient20` | `string` | No |  |
| `strIngredient3` | `string` | No |  |
| `strIngredient4` | `string` | No |  |
| `strIngredient5` | `string` | No |  |
| `strIngredient6` | `string` | No |  |
| `strIngredient7` | `string` | No |  |
| `strIngredient8` | `string` | No |  |
| `strIngredient9` | `string` | No |  |
| `strInstructions` | `string` | No | Cooking instructions |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |
| `strMeasure1` | `string` | No |  |
| `strMeasure10` | `string` | No |  |
| `strMeasure11` | `string` | No |  |
| `strMeasure12` | `string` | No |  |
| `strMeasure13` | `string` | No |  |
| `strMeasure14` | `string` | No |  |
| `strMeasure15` | `string` | No |  |
| `strMeasure16` | `string` | No |  |
| `strMeasure17` | `string` | No |  |
| `strMeasure18` | `string` | No |  |
| `strMeasure19` | `string` | No |  |
| `strMeasure2` | `string` | No |  |
| `strMeasure20` | `string` | No |  |
| `strMeasure3` | `string` | No |  |
| `strMeasure4` | `string` | No |  |
| `strMeasure5` | `string` | No |  |
| `strMeasure6` | `string` | No |  |
| `strMeasure7` | `string` | No |  |
| `strMeasure8` | `string` | No |  |
| `strMeasure9` | `string` | No |  |
| `strSource` | `string` | No |  |
| `strTags` | `string` | No | Comma-separated tags |
| `strYoutube` | `string` | No | YouTube video URL |

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
| `dateModified` | `string` | No |  |
| `idMeal` | `string` | No | Unique meal identifier |
| `strArea` | `string` | No | Meal area/region |
| `strCategory` | `string` | No | Meal category |
| `strCreativeCommonsConfirmed` | `string` | No |  |
| `strDrinkAlternate` | `string` | No |  |
| `strImageSource` | `string` | No |  |
| `strIngredient1` | `string` | No |  |
| `strIngredient10` | `string` | No |  |
| `strIngredient11` | `string` | No |  |
| `strIngredient12` | `string` | No |  |
| `strIngredient13` | `string` | No |  |
| `strIngredient14` | `string` | No |  |
| `strIngredient15` | `string` | No |  |
| `strIngredient16` | `string` | No |  |
| `strIngredient17` | `string` | No |  |
| `strIngredient18` | `string` | No |  |
| `strIngredient19` | `string` | No |  |
| `strIngredient2` | `string` | No |  |
| `strIngredient20` | `string` | No |  |
| `strIngredient3` | `string` | No |  |
| `strIngredient4` | `string` | No |  |
| `strIngredient5` | `string` | No |  |
| `strIngredient6` | `string` | No |  |
| `strIngredient7` | `string` | No |  |
| `strIngredient8` | `string` | No |  |
| `strIngredient9` | `string` | No |  |
| `strInstructions` | `string` | No | Cooking instructions |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |
| `strMeasure1` | `string` | No |  |
| `strMeasure10` | `string` | No |  |
| `strMeasure11` | `string` | No |  |
| `strMeasure12` | `string` | No |  |
| `strMeasure13` | `string` | No |  |
| `strMeasure14` | `string` | No |  |
| `strMeasure15` | `string` | No |  |
| `strMeasure16` | `string` | No |  |
| `strMeasure17` | `string` | No |  |
| `strMeasure18` | `string` | No |  |
| `strMeasure19` | `string` | No |  |
| `strMeasure2` | `string` | No |  |
| `strMeasure20` | `string` | No |  |
| `strMeasure3` | `string` | No |  |
| `strMeasure4` | `string` | No |  |
| `strMeasure5` | `string` | No |  |
| `strMeasure6` | `string` | No |  |
| `strMeasure7` | `string` | No |  |
| `strMeasure8` | `string` | No |  |
| `strMeasure9` | `string` | No |  |
| `strSource` | `string` | No |  |
| `strTags` | `string` | No | Comma-separated tags |
| `strYoutube` | `string` | No | YouTube video URL |

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

