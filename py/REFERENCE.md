# FreeMeal Python SDK Reference

Complete API reference for the FreeMeal Python SDK.


## FreeMealSDK

### Constructor

```python
from freemeal_sdk import FreeMealSDK

client = FreeMealSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `FreeMealSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = FreeMealSDK.test()
```


### Instance Methods

#### `Category(data=None)`

Create a new `CategoryEntity` instance. Pass `None` for no initial data.

#### `Filter(data=None)`

Create a new `FilterEntity` instance. Pass `None` for no initial data.

#### `Latest(data=None)`

Create a new `LatestEntity` instance. Pass `None` for no initial data.

#### `List(data=None)`

Create a new `ListEntity` instance. Pass `None` for no initial data.

#### `Lookup(data=None)`

Create a new `LookupEntity` instance. Pass `None` for no initial data.

#### `Random(data=None)`

Create a new `RandomEntity` instance. Pass `None` for no initial data.

#### `Randomselection(data=None)`

Create a new `RandomselectionEntity` instance. Pass `None` for no initial data.

#### `Search(data=None)`

Create a new `SearchEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## CategoryEntity

```python
category = client.Category()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `idCategory` | `str` | No |  |
| `strCategory` | `str` | No |  |
| `strCategoryDescription` | `str` | No |  |
| `strCategoryThumb` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Category().list()
for category in results:
    print(category)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CategoryEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FilterEntity

```python
filter = client.Filter()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `idMeal` | `str` | No |  |
| `strMeal` | `str` | No |  |
| `strMealThumb` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Filter().list()
for filter in results:
    print(filter)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FilterEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LatestEntity

```python
latest = client.Latest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `str` | No |  |
| `idMeal` | `str` | No |  |
| `strArea` | `str` | No |  |
| `strCategory` | `str` | No |  |
| `strCreativeCommonsConfirmed` | `str` | No |  |
| `strDrinkAlternate` | `str` | No |  |
| `strImageSource` | `str` | No |  |
| `strIngredient1` | `str` | No |  |
| `strIngredient10` | `str` | No |  |
| `strIngredient11` | `str` | No |  |
| `strIngredient12` | `str` | No |  |
| `strIngredient13` | `str` | No |  |
| `strIngredient14` | `str` | No |  |
| `strIngredient15` | `str` | No |  |
| `strIngredient16` | `str` | No |  |
| `strIngredient17` | `str` | No |  |
| `strIngredient18` | `str` | No |  |
| `strIngredient19` | `str` | No |  |
| `strIngredient2` | `str` | No |  |
| `strIngredient20` | `str` | No |  |
| `strIngredient3` | `str` | No |  |
| `strIngredient4` | `str` | No |  |
| `strIngredient5` | `str` | No |  |
| `strIngredient6` | `str` | No |  |
| `strIngredient7` | `str` | No |  |
| `strIngredient8` | `str` | No |  |
| `strIngredient9` | `str` | No |  |
| `strInstructions` | `str` | No |  |
| `strMeal` | `str` | No |  |
| `strMealThumb` | `str` | No |  |
| `strMeasure1` | `str` | No |  |
| `strMeasure10` | `str` | No |  |
| `strMeasure11` | `str` | No |  |
| `strMeasure12` | `str` | No |  |
| `strMeasure13` | `str` | No |  |
| `strMeasure14` | `str` | No |  |
| `strMeasure15` | `str` | No |  |
| `strMeasure16` | `str` | No |  |
| `strMeasure17` | `str` | No |  |
| `strMeasure18` | `str` | No |  |
| `strMeasure19` | `str` | No |  |
| `strMeasure2` | `str` | No |  |
| `strMeasure20` | `str` | No |  |
| `strMeasure3` | `str` | No |  |
| `strMeasure4` | `str` | No |  |
| `strMeasure5` | `str` | No |  |
| `strMeasure6` | `str` | No |  |
| `strMeasure7` | `str` | No |  |
| `strMeasure8` | `str` | No |  |
| `strMeasure9` | `str` | No |  |
| `strSource` | `str` | No |  |
| `strTags` | `str` | No |  |
| `strYoutube` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Latest().list()
for latest in results:
    print(latest)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LatestEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListEntity

```python
list = client.List()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `strArea` | `str` | No |  |
| `strCategory` | `str` | No |  |
| `strIngredient` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.List().list()
for list in results:
    print(list)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LookupEntity

```python
lookup = client.Lookup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `str` | No |  |
| `idMeal` | `str` | No |  |
| `strArea` | `str` | No |  |
| `strCategory` | `str` | No |  |
| `strCreativeCommonsConfirmed` | `str` | No |  |
| `strDrinkAlternate` | `str` | No |  |
| `strImageSource` | `str` | No |  |
| `strIngredient1` | `str` | No |  |
| `strIngredient10` | `str` | No |  |
| `strIngredient11` | `str` | No |  |
| `strIngredient12` | `str` | No |  |
| `strIngredient13` | `str` | No |  |
| `strIngredient14` | `str` | No |  |
| `strIngredient15` | `str` | No |  |
| `strIngredient16` | `str` | No |  |
| `strIngredient17` | `str` | No |  |
| `strIngredient18` | `str` | No |  |
| `strIngredient19` | `str` | No |  |
| `strIngredient2` | `str` | No |  |
| `strIngredient20` | `str` | No |  |
| `strIngredient3` | `str` | No |  |
| `strIngredient4` | `str` | No |  |
| `strIngredient5` | `str` | No |  |
| `strIngredient6` | `str` | No |  |
| `strIngredient7` | `str` | No |  |
| `strIngredient8` | `str` | No |  |
| `strIngredient9` | `str` | No |  |
| `strInstructions` | `str` | No |  |
| `strMeal` | `str` | No |  |
| `strMealThumb` | `str` | No |  |
| `strMeasure1` | `str` | No |  |
| `strMeasure10` | `str` | No |  |
| `strMeasure11` | `str` | No |  |
| `strMeasure12` | `str` | No |  |
| `strMeasure13` | `str` | No |  |
| `strMeasure14` | `str` | No |  |
| `strMeasure15` | `str` | No |  |
| `strMeasure16` | `str` | No |  |
| `strMeasure17` | `str` | No |  |
| `strMeasure18` | `str` | No |  |
| `strMeasure19` | `str` | No |  |
| `strMeasure2` | `str` | No |  |
| `strMeasure20` | `str` | No |  |
| `strMeasure3` | `str` | No |  |
| `strMeasure4` | `str` | No |  |
| `strMeasure5` | `str` | No |  |
| `strMeasure6` | `str` | No |  |
| `strMeasure7` | `str` | No |  |
| `strMeasure8` | `str` | No |  |
| `strMeasure9` | `str` | No |  |
| `strSource` | `str` | No |  |
| `strTags` | `str` | No |  |
| `strYoutube` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Lookup().list()
for lookup in results:
    print(lookup)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LookupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RandomEntity

```python
random = client.Random()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `str` | No |  |
| `idMeal` | `str` | No |  |
| `strArea` | `str` | No |  |
| `strCategory` | `str` | No |  |
| `strCreativeCommonsConfirmed` | `str` | No |  |
| `strDrinkAlternate` | `str` | No |  |
| `strImageSource` | `str` | No |  |
| `strIngredient1` | `str` | No |  |
| `strIngredient10` | `str` | No |  |
| `strIngredient11` | `str` | No |  |
| `strIngredient12` | `str` | No |  |
| `strIngredient13` | `str` | No |  |
| `strIngredient14` | `str` | No |  |
| `strIngredient15` | `str` | No |  |
| `strIngredient16` | `str` | No |  |
| `strIngredient17` | `str` | No |  |
| `strIngredient18` | `str` | No |  |
| `strIngredient19` | `str` | No |  |
| `strIngredient2` | `str` | No |  |
| `strIngredient20` | `str` | No |  |
| `strIngredient3` | `str` | No |  |
| `strIngredient4` | `str` | No |  |
| `strIngredient5` | `str` | No |  |
| `strIngredient6` | `str` | No |  |
| `strIngredient7` | `str` | No |  |
| `strIngredient8` | `str` | No |  |
| `strIngredient9` | `str` | No |  |
| `strInstructions` | `str` | No |  |
| `strMeal` | `str` | No |  |
| `strMealThumb` | `str` | No |  |
| `strMeasure1` | `str` | No |  |
| `strMeasure10` | `str` | No |  |
| `strMeasure11` | `str` | No |  |
| `strMeasure12` | `str` | No |  |
| `strMeasure13` | `str` | No |  |
| `strMeasure14` | `str` | No |  |
| `strMeasure15` | `str` | No |  |
| `strMeasure16` | `str` | No |  |
| `strMeasure17` | `str` | No |  |
| `strMeasure18` | `str` | No |  |
| `strMeasure19` | `str` | No |  |
| `strMeasure2` | `str` | No |  |
| `strMeasure20` | `str` | No |  |
| `strMeasure3` | `str` | No |  |
| `strMeasure4` | `str` | No |  |
| `strMeasure5` | `str` | No |  |
| `strMeasure6` | `str` | No |  |
| `strMeasure7` | `str` | No |  |
| `strMeasure8` | `str` | No |  |
| `strMeasure9` | `str` | No |  |
| `strSource` | `str` | No |  |
| `strTags` | `str` | No |  |
| `strYoutube` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Random().list()
for random in results:
    print(random)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RandomEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RandomselectionEntity

```python
randomselection = client.Randomselection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `str` | No |  |
| `idMeal` | `str` | No |  |
| `strArea` | `str` | No |  |
| `strCategory` | `str` | No |  |
| `strCreativeCommonsConfirmed` | `str` | No |  |
| `strDrinkAlternate` | `str` | No |  |
| `strImageSource` | `str` | No |  |
| `strIngredient1` | `str` | No |  |
| `strIngredient10` | `str` | No |  |
| `strIngredient11` | `str` | No |  |
| `strIngredient12` | `str` | No |  |
| `strIngredient13` | `str` | No |  |
| `strIngredient14` | `str` | No |  |
| `strIngredient15` | `str` | No |  |
| `strIngredient16` | `str` | No |  |
| `strIngredient17` | `str` | No |  |
| `strIngredient18` | `str` | No |  |
| `strIngredient19` | `str` | No |  |
| `strIngredient2` | `str` | No |  |
| `strIngredient20` | `str` | No |  |
| `strIngredient3` | `str` | No |  |
| `strIngredient4` | `str` | No |  |
| `strIngredient5` | `str` | No |  |
| `strIngredient6` | `str` | No |  |
| `strIngredient7` | `str` | No |  |
| `strIngredient8` | `str` | No |  |
| `strIngredient9` | `str` | No |  |
| `strInstructions` | `str` | No |  |
| `strMeal` | `str` | No |  |
| `strMealThumb` | `str` | No |  |
| `strMeasure1` | `str` | No |  |
| `strMeasure10` | `str` | No |  |
| `strMeasure11` | `str` | No |  |
| `strMeasure12` | `str` | No |  |
| `strMeasure13` | `str` | No |  |
| `strMeasure14` | `str` | No |  |
| `strMeasure15` | `str` | No |  |
| `strMeasure16` | `str` | No |  |
| `strMeasure17` | `str` | No |  |
| `strMeasure18` | `str` | No |  |
| `strMeasure19` | `str` | No |  |
| `strMeasure2` | `str` | No |  |
| `strMeasure20` | `str` | No |  |
| `strMeasure3` | `str` | No |  |
| `strMeasure4` | `str` | No |  |
| `strMeasure5` | `str` | No |  |
| `strMeasure6` | `str` | No |  |
| `strMeasure7` | `str` | No |  |
| `strMeasure8` | `str` | No |  |
| `strMeasure9` | `str` | No |  |
| `strSource` | `str` | No |  |
| `strTags` | `str` | No |  |
| `strYoutube` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Randomselection().list()
for randomselection in results:
    print(randomselection)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RandomselectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SearchEntity

```python
search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dateModified` | `str` | No |  |
| `idMeal` | `str` | No |  |
| `strArea` | `str` | No |  |
| `strCategory` | `str` | No |  |
| `strCreativeCommonsConfirmed` | `str` | No |  |
| `strDrinkAlternate` | `str` | No |  |
| `strImageSource` | `str` | No |  |
| `strIngredient1` | `str` | No |  |
| `strIngredient10` | `str` | No |  |
| `strIngredient11` | `str` | No |  |
| `strIngredient12` | `str` | No |  |
| `strIngredient13` | `str` | No |  |
| `strIngredient14` | `str` | No |  |
| `strIngredient15` | `str` | No |  |
| `strIngredient16` | `str` | No |  |
| `strIngredient17` | `str` | No |  |
| `strIngredient18` | `str` | No |  |
| `strIngredient19` | `str` | No |  |
| `strIngredient2` | `str` | No |  |
| `strIngredient20` | `str` | No |  |
| `strIngredient3` | `str` | No |  |
| `strIngredient4` | `str` | No |  |
| `strIngredient5` | `str` | No |  |
| `strIngredient6` | `str` | No |  |
| `strIngredient7` | `str` | No |  |
| `strIngredient8` | `str` | No |  |
| `strIngredient9` | `str` | No |  |
| `strInstructions` | `str` | No |  |
| `strMeal` | `str` | No |  |
| `strMealThumb` | `str` | No |  |
| `strMeasure1` | `str` | No |  |
| `strMeasure10` | `str` | No |  |
| `strMeasure11` | `str` | No |  |
| `strMeasure12` | `str` | No |  |
| `strMeasure13` | `str` | No |  |
| `strMeasure14` | `str` | No |  |
| `strMeasure15` | `str` | No |  |
| `strMeasure16` | `str` | No |  |
| `strMeasure17` | `str` | No |  |
| `strMeasure18` | `str` | No |  |
| `strMeasure19` | `str` | No |  |
| `strMeasure2` | `str` | No |  |
| `strMeasure20` | `str` | No |  |
| `strMeasure3` | `str` | No |  |
| `strMeasure4` | `str` | No |  |
| `strMeasure5` | `str` | No |  |
| `strMeasure6` | `str` | No |  |
| `strMeasure7` | `str` | No |  |
| `strMeasure8` | `str` | No |  |
| `strMeasure9` | `str` | No |  |
| `strSource` | `str` | No |  |
| `strTags` | `str` | No |  |
| `strYoutube` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Search().list()
for search in results:
    print(search)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = FreeMealSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

