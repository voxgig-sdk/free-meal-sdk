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
| `id_category` | `str` | No |  |
| `str_category` | `str` | No |  |
| `str_category_description` | `str` | No |  |
| `str_category_thumb` | `str` | No |  |

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
| `id_meal` | `str` | No |  |
| `str_meal` | `str` | No |  |
| `str_meal_thumb` | `str` | No |  |

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
| `date_modified` | `str` | No |  |
| `id_meal` | `str` | No |  |
| `str_area` | `str` | No |  |
| `str_category` | `str` | No |  |
| `str_creative_commons_confirmed` | `str` | No |  |
| `str_drink_alternate` | `str` | No |  |
| `str_image_source` | `str` | No |  |
| `str_ingredient1` | `str` | No |  |
| `str_ingredient10` | `str` | No |  |
| `str_ingredient11` | `str` | No |  |
| `str_ingredient12` | `str` | No |  |
| `str_ingredient13` | `str` | No |  |
| `str_ingredient14` | `str` | No |  |
| `str_ingredient15` | `str` | No |  |
| `str_ingredient16` | `str` | No |  |
| `str_ingredient17` | `str` | No |  |
| `str_ingredient18` | `str` | No |  |
| `str_ingredient19` | `str` | No |  |
| `str_ingredient2` | `str` | No |  |
| `str_ingredient20` | `str` | No |  |
| `str_ingredient3` | `str` | No |  |
| `str_ingredient4` | `str` | No |  |
| `str_ingredient5` | `str` | No |  |
| `str_ingredient6` | `str` | No |  |
| `str_ingredient7` | `str` | No |  |
| `str_ingredient8` | `str` | No |  |
| `str_ingredient9` | `str` | No |  |
| `str_instruction` | `str` | No |  |
| `str_meal` | `str` | No |  |
| `str_meal_thumb` | `str` | No |  |
| `str_measure1` | `str` | No |  |
| `str_measure10` | `str` | No |  |
| `str_measure11` | `str` | No |  |
| `str_measure12` | `str` | No |  |
| `str_measure13` | `str` | No |  |
| `str_measure14` | `str` | No |  |
| `str_measure15` | `str` | No |  |
| `str_measure16` | `str` | No |  |
| `str_measure17` | `str` | No |  |
| `str_measure18` | `str` | No |  |
| `str_measure19` | `str` | No |  |
| `str_measure2` | `str` | No |  |
| `str_measure20` | `str` | No |  |
| `str_measure3` | `str` | No |  |
| `str_measure4` | `str` | No |  |
| `str_measure5` | `str` | No |  |
| `str_measure6` | `str` | No |  |
| `str_measure7` | `str` | No |  |
| `str_measure8` | `str` | No |  |
| `str_measure9` | `str` | No |  |
| `str_source` | `str` | No |  |
| `str_tag` | `str` | No |  |
| `str_youtube` | `str` | No |  |

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
| `str_area` | `str` | No |  |
| `str_category` | `str` | No |  |
| `str_ingredient` | `str` | No |  |

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
| `date_modified` | `str` | No |  |
| `id_meal` | `str` | No |  |
| `str_area` | `str` | No |  |
| `str_category` | `str` | No |  |
| `str_creative_commons_confirmed` | `str` | No |  |
| `str_drink_alternate` | `str` | No |  |
| `str_image_source` | `str` | No |  |
| `str_ingredient1` | `str` | No |  |
| `str_ingredient10` | `str` | No |  |
| `str_ingredient11` | `str` | No |  |
| `str_ingredient12` | `str` | No |  |
| `str_ingredient13` | `str` | No |  |
| `str_ingredient14` | `str` | No |  |
| `str_ingredient15` | `str` | No |  |
| `str_ingredient16` | `str` | No |  |
| `str_ingredient17` | `str` | No |  |
| `str_ingredient18` | `str` | No |  |
| `str_ingredient19` | `str` | No |  |
| `str_ingredient2` | `str` | No |  |
| `str_ingredient20` | `str` | No |  |
| `str_ingredient3` | `str` | No |  |
| `str_ingredient4` | `str` | No |  |
| `str_ingredient5` | `str` | No |  |
| `str_ingredient6` | `str` | No |  |
| `str_ingredient7` | `str` | No |  |
| `str_ingredient8` | `str` | No |  |
| `str_ingredient9` | `str` | No |  |
| `str_instruction` | `str` | No |  |
| `str_meal` | `str` | No |  |
| `str_meal_thumb` | `str` | No |  |
| `str_measure1` | `str` | No |  |
| `str_measure10` | `str` | No |  |
| `str_measure11` | `str` | No |  |
| `str_measure12` | `str` | No |  |
| `str_measure13` | `str` | No |  |
| `str_measure14` | `str` | No |  |
| `str_measure15` | `str` | No |  |
| `str_measure16` | `str` | No |  |
| `str_measure17` | `str` | No |  |
| `str_measure18` | `str` | No |  |
| `str_measure19` | `str` | No |  |
| `str_measure2` | `str` | No |  |
| `str_measure20` | `str` | No |  |
| `str_measure3` | `str` | No |  |
| `str_measure4` | `str` | No |  |
| `str_measure5` | `str` | No |  |
| `str_measure6` | `str` | No |  |
| `str_measure7` | `str` | No |  |
| `str_measure8` | `str` | No |  |
| `str_measure9` | `str` | No |  |
| `str_source` | `str` | No |  |
| `str_tag` | `str` | No |  |
| `str_youtube` | `str` | No |  |

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
| `date_modified` | `str` | No |  |
| `id_meal` | `str` | No |  |
| `str_area` | `str` | No |  |
| `str_category` | `str` | No |  |
| `str_creative_commons_confirmed` | `str` | No |  |
| `str_drink_alternate` | `str` | No |  |
| `str_image_source` | `str` | No |  |
| `str_ingredient1` | `str` | No |  |
| `str_ingredient10` | `str` | No |  |
| `str_ingredient11` | `str` | No |  |
| `str_ingredient12` | `str` | No |  |
| `str_ingredient13` | `str` | No |  |
| `str_ingredient14` | `str` | No |  |
| `str_ingredient15` | `str` | No |  |
| `str_ingredient16` | `str` | No |  |
| `str_ingredient17` | `str` | No |  |
| `str_ingredient18` | `str` | No |  |
| `str_ingredient19` | `str` | No |  |
| `str_ingredient2` | `str` | No |  |
| `str_ingredient20` | `str` | No |  |
| `str_ingredient3` | `str` | No |  |
| `str_ingredient4` | `str` | No |  |
| `str_ingredient5` | `str` | No |  |
| `str_ingredient6` | `str` | No |  |
| `str_ingredient7` | `str` | No |  |
| `str_ingredient8` | `str` | No |  |
| `str_ingredient9` | `str` | No |  |
| `str_instruction` | `str` | No |  |
| `str_meal` | `str` | No |  |
| `str_meal_thumb` | `str` | No |  |
| `str_measure1` | `str` | No |  |
| `str_measure10` | `str` | No |  |
| `str_measure11` | `str` | No |  |
| `str_measure12` | `str` | No |  |
| `str_measure13` | `str` | No |  |
| `str_measure14` | `str` | No |  |
| `str_measure15` | `str` | No |  |
| `str_measure16` | `str` | No |  |
| `str_measure17` | `str` | No |  |
| `str_measure18` | `str` | No |  |
| `str_measure19` | `str` | No |  |
| `str_measure2` | `str` | No |  |
| `str_measure20` | `str` | No |  |
| `str_measure3` | `str` | No |  |
| `str_measure4` | `str` | No |  |
| `str_measure5` | `str` | No |  |
| `str_measure6` | `str` | No |  |
| `str_measure7` | `str` | No |  |
| `str_measure8` | `str` | No |  |
| `str_measure9` | `str` | No |  |
| `str_source` | `str` | No |  |
| `str_tag` | `str` | No |  |
| `str_youtube` | `str` | No |  |

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
| `date_modified` | `str` | No |  |
| `id_meal` | `str` | No |  |
| `str_area` | `str` | No |  |
| `str_category` | `str` | No |  |
| `str_creative_commons_confirmed` | `str` | No |  |
| `str_drink_alternate` | `str` | No |  |
| `str_image_source` | `str` | No |  |
| `str_ingredient1` | `str` | No |  |
| `str_ingredient10` | `str` | No |  |
| `str_ingredient11` | `str` | No |  |
| `str_ingredient12` | `str` | No |  |
| `str_ingredient13` | `str` | No |  |
| `str_ingredient14` | `str` | No |  |
| `str_ingredient15` | `str` | No |  |
| `str_ingredient16` | `str` | No |  |
| `str_ingredient17` | `str` | No |  |
| `str_ingredient18` | `str` | No |  |
| `str_ingredient19` | `str` | No |  |
| `str_ingredient2` | `str` | No |  |
| `str_ingredient20` | `str` | No |  |
| `str_ingredient3` | `str` | No |  |
| `str_ingredient4` | `str` | No |  |
| `str_ingredient5` | `str` | No |  |
| `str_ingredient6` | `str` | No |  |
| `str_ingredient7` | `str` | No |  |
| `str_ingredient8` | `str` | No |  |
| `str_ingredient9` | `str` | No |  |
| `str_instruction` | `str` | No |  |
| `str_meal` | `str` | No |  |
| `str_meal_thumb` | `str` | No |  |
| `str_measure1` | `str` | No |  |
| `str_measure10` | `str` | No |  |
| `str_measure11` | `str` | No |  |
| `str_measure12` | `str` | No |  |
| `str_measure13` | `str` | No |  |
| `str_measure14` | `str` | No |  |
| `str_measure15` | `str` | No |  |
| `str_measure16` | `str` | No |  |
| `str_measure17` | `str` | No |  |
| `str_measure18` | `str` | No |  |
| `str_measure19` | `str` | No |  |
| `str_measure2` | `str` | No |  |
| `str_measure20` | `str` | No |  |
| `str_measure3` | `str` | No |  |
| `str_measure4` | `str` | No |  |
| `str_measure5` | `str` | No |  |
| `str_measure6` | `str` | No |  |
| `str_measure7` | `str` | No |  |
| `str_measure8` | `str` | No |  |
| `str_measure9` | `str` | No |  |
| `str_source` | `str` | No |  |
| `str_tag` | `str` | No |  |
| `str_youtube` | `str` | No |  |

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
| `date_modified` | `str` | No |  |
| `id_meal` | `str` | No |  |
| `str_area` | `str` | No |  |
| `str_category` | `str` | No |  |
| `str_creative_commons_confirmed` | `str` | No |  |
| `str_drink_alternate` | `str` | No |  |
| `str_image_source` | `str` | No |  |
| `str_ingredient1` | `str` | No |  |
| `str_ingredient10` | `str` | No |  |
| `str_ingredient11` | `str` | No |  |
| `str_ingredient12` | `str` | No |  |
| `str_ingredient13` | `str` | No |  |
| `str_ingredient14` | `str` | No |  |
| `str_ingredient15` | `str` | No |  |
| `str_ingredient16` | `str` | No |  |
| `str_ingredient17` | `str` | No |  |
| `str_ingredient18` | `str` | No |  |
| `str_ingredient19` | `str` | No |  |
| `str_ingredient2` | `str` | No |  |
| `str_ingredient20` | `str` | No |  |
| `str_ingredient3` | `str` | No |  |
| `str_ingredient4` | `str` | No |  |
| `str_ingredient5` | `str` | No |  |
| `str_ingredient6` | `str` | No |  |
| `str_ingredient7` | `str` | No |  |
| `str_ingredient8` | `str` | No |  |
| `str_ingredient9` | `str` | No |  |
| `str_instruction` | `str` | No |  |
| `str_meal` | `str` | No |  |
| `str_meal_thumb` | `str` | No |  |
| `str_measure1` | `str` | No |  |
| `str_measure10` | `str` | No |  |
| `str_measure11` | `str` | No |  |
| `str_measure12` | `str` | No |  |
| `str_measure13` | `str` | No |  |
| `str_measure14` | `str` | No |  |
| `str_measure15` | `str` | No |  |
| `str_measure16` | `str` | No |  |
| `str_measure17` | `str` | No |  |
| `str_measure18` | `str` | No |  |
| `str_measure19` | `str` | No |  |
| `str_measure2` | `str` | No |  |
| `str_measure20` | `str` | No |  |
| `str_measure3` | `str` | No |  |
| `str_measure4` | `str` | No |  |
| `str_measure5` | `str` | No |  |
| `str_measure6` | `str` | No |  |
| `str_measure7` | `str` | No |  |
| `str_measure8` | `str` | No |  |
| `str_measure9` | `str` | No |  |
| `str_source` | `str` | No |  |
| `str_tag` | `str` | No |  |
| `str_youtube` | `str` | No |  |

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

