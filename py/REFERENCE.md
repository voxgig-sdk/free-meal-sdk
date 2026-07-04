# FreeMeal Python SDK Reference

Complete API reference for the FreeMeal Python SDK.


## FreeMealSDK

### Constructor

```python
from free-meal_sdk import FreeMealSDK

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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.category.list({})
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
filter = client.filter
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_meal` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.filter.list({})
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.latest.list({})
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
list = client.list
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_ingredient` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.list.list({})
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.lookup.list({})
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.random.list({})
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.randomselection.list({})
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.search.list({})
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

