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

#### `Test() *FreeMealSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *FreeMealSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
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
fmt.Println(category.GetName()) // "category"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `idCategory` | `string` | No | Unique category identifier |
| `strCategory` | `string` | No | Category name |
| `strCategoryDescription` | `string` | No | Category description |
| `strCategoryThumb` | `string` | No | URL to category thumbnail image |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Category(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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
fmt.Println(filter.GetName()) // "filter"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `idMeal` | `string` | No | Unique meal identifier |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Filter(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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
fmt.Println(latest.GetName()) // "latest"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Latest(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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
fmt.Println(list.GetName()) // "list"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `strArea` | `string` | No |  |
| `strCategory` | `string` | No |  |
| `strIngredient` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.List(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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
fmt.Println(lookup.GetName()) // "lookup"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Lookup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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
fmt.Println(random.GetName()) // "random"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Random(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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
fmt.Println(randomselection.GetName()) // "randomselection"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Randomselection(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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
fmt.Println(search.GetName()) // "search"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
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

