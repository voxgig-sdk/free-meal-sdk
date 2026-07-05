# FreeMeal Python SDK



The Python SDK for the FreeMeal API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Category()` — each
carrying a small, uniform set of operations (`list`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/free-meal-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from freemeal_sdk import FreeMealSDK

client = FreeMealSDK({
    "apikey": os.environ.get("FREE_MEAL_APIKEY"),
})
```

### 2. List category records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    categorys = client.Category().list()
    for category in categorys:
        print(category)
except Exception as err:
    print(f"list failed: {err}")
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    categorys = client.Category().list()
    print(categorys)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = FreeMealSDK.test()

# Entity ops return the bare record and raise on error.
category = client.Category().list()
# category contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = FreeMealSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
FREE_MEAL_TEST_LIVE=TRUE
FREE_MEAL_APIKEY=<your-key>
```

Then run:

```bash
cd py && pytest test/
```


## Reference

### FreeMealSDK

```python
from freemeal_sdk import FreeMealSDK

client = FreeMealSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = FreeMealSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### FreeMealSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `Category` | `(data) -> CategoryEntity` | Create a Category entity instance. |
| `Filter` | `(data) -> FilterEntity` | Create a Filter entity instance. |
| `Latest` | `(data) -> LatestEntity` | Create a Latest entity instance. |
| `List` | `(data) -> ListEntity` | Create a List entity instance. |
| `Lookup` | `(data) -> LookupEntity` | Create a Lookup entity instance. |
| `Random` | `(data) -> RandomEntity` | Create a Random entity instance. |
| `Randomselection` | `(data) -> RandomselectionEntity` | Create a Randomselection entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the bare result data (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### Category

| Field | Description |
| --- | --- |
| `id_category` |  |
| `str_category` |  |
| `str_category_description` |  |
| `str_category_thumb` |  |

Operations: List.

API path: `/categories.php`

#### Filter

| Field | Description |
| --- | --- |
| `id_meal` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |

Operations: List.

API path: `/filter.php`

#### Latest

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/latest.php`

#### List

| Field | Description |
| --- | --- |
| `str_area` |  |
| `str_category` |  |
| `str_ingredient` |  |

Operations: List.

API path: `/list.php`

#### Lookup

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/lookup.php`

#### Random

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/random.php`

#### Randomselection

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/randomselection.php`

#### Search

| Field | Description |
| --- | --- |
| `date_modified` |  |
| `id_meal` |  |
| `str_area` |  |
| `str_category` |  |
| `str_creative_commons_confirmed` |  |
| `str_drink_alternate` |  |
| `str_image_source` |  |
| `str_ingredient1` |  |
| `str_ingredient10` |  |
| `str_ingredient11` |  |
| `str_ingredient12` |  |
| `str_ingredient13` |  |
| `str_ingredient14` |  |
| `str_ingredient15` |  |
| `str_ingredient16` |  |
| `str_ingredient17` |  |
| `str_ingredient18` |  |
| `str_ingredient19` |  |
| `str_ingredient2` |  |
| `str_ingredient20` |  |
| `str_ingredient3` |  |
| `str_ingredient4` |  |
| `str_ingredient5` |  |
| `str_ingredient6` |  |
| `str_ingredient7` |  |
| `str_ingredient8` |  |
| `str_ingredient9` |  |
| `str_instruction` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |
| `str_measure1` |  |
| `str_measure10` |  |
| `str_measure11` |  |
| `str_measure12` |  |
| `str_measure13` |  |
| `str_measure14` |  |
| `str_measure15` |  |
| `str_measure16` |  |
| `str_measure17` |  |
| `str_measure18` |  |
| `str_measure19` |  |
| `str_measure2` |  |
| `str_measure20` |  |
| `str_measure3` |  |
| `str_measure4` |  |
| `str_measure5` |  |
| `str_measure6` |  |
| `str_measure7` |  |
| `str_measure8` |  |
| `str_measure9` |  |
| `str_source` |  |
| `str_tag` |  |
| `str_youtube` |  |

Operations: List.

API path: `/search.php`



## Entities


### Category

Create an instance: `category = client.Category()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_category` | `str` |  |
| `str_category` | `str` |  |
| `str_category_description` | `str` |  |
| `str_category_thumb` | `str` |  |

#### Example: List

```python
categorys = client.Category().list()
```


### Filter

Create an instance: `filter = client.Filter()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_meal` | `str` |  |
| `str_meal` | `str` |  |
| `str_meal_thumb` | `str` |  |

#### Example: List

```python
filters = client.Filter().list()
```


### Latest

Create an instance: `latest = client.Latest()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `str` |  |
| `id_meal` | `str` |  |
| `str_area` | `str` |  |
| `str_category` | `str` |  |
| `str_creative_commons_confirmed` | `str` |  |
| `str_drink_alternate` | `str` |  |
| `str_image_source` | `str` |  |
| `str_ingredient1` | `str` |  |
| `str_ingredient10` | `str` |  |
| `str_ingredient11` | `str` |  |
| `str_ingredient12` | `str` |  |
| `str_ingredient13` | `str` |  |
| `str_ingredient14` | `str` |  |
| `str_ingredient15` | `str` |  |
| `str_ingredient16` | `str` |  |
| `str_ingredient17` | `str` |  |
| `str_ingredient18` | `str` |  |
| `str_ingredient19` | `str` |  |
| `str_ingredient2` | `str` |  |
| `str_ingredient20` | `str` |  |
| `str_ingredient3` | `str` |  |
| `str_ingredient4` | `str` |  |
| `str_ingredient5` | `str` |  |
| `str_ingredient6` | `str` |  |
| `str_ingredient7` | `str` |  |
| `str_ingredient8` | `str` |  |
| `str_ingredient9` | `str` |  |
| `str_instruction` | `str` |  |
| `str_meal` | `str` |  |
| `str_meal_thumb` | `str` |  |
| `str_measure1` | `str` |  |
| `str_measure10` | `str` |  |
| `str_measure11` | `str` |  |
| `str_measure12` | `str` |  |
| `str_measure13` | `str` |  |
| `str_measure14` | `str` |  |
| `str_measure15` | `str` |  |
| `str_measure16` | `str` |  |
| `str_measure17` | `str` |  |
| `str_measure18` | `str` |  |
| `str_measure19` | `str` |  |
| `str_measure2` | `str` |  |
| `str_measure20` | `str` |  |
| `str_measure3` | `str` |  |
| `str_measure4` | `str` |  |
| `str_measure5` | `str` |  |
| `str_measure6` | `str` |  |
| `str_measure7` | `str` |  |
| `str_measure8` | `str` |  |
| `str_measure9` | `str` |  |
| `str_source` | `str` |  |
| `str_tag` | `str` |  |
| `str_youtube` | `str` |  |

#### Example: List

```python
latests = client.Latest().list()
```


### List

Create an instance: `list = client.List()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `str_area` | `str` |  |
| `str_category` | `str` |  |
| `str_ingredient` | `str` |  |

#### Example: List

```python
lists = client.List().list()
```


### Lookup

Create an instance: `lookup = client.Lookup()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `str` |  |
| `id_meal` | `str` |  |
| `str_area` | `str` |  |
| `str_category` | `str` |  |
| `str_creative_commons_confirmed` | `str` |  |
| `str_drink_alternate` | `str` |  |
| `str_image_source` | `str` |  |
| `str_ingredient1` | `str` |  |
| `str_ingredient10` | `str` |  |
| `str_ingredient11` | `str` |  |
| `str_ingredient12` | `str` |  |
| `str_ingredient13` | `str` |  |
| `str_ingredient14` | `str` |  |
| `str_ingredient15` | `str` |  |
| `str_ingredient16` | `str` |  |
| `str_ingredient17` | `str` |  |
| `str_ingredient18` | `str` |  |
| `str_ingredient19` | `str` |  |
| `str_ingredient2` | `str` |  |
| `str_ingredient20` | `str` |  |
| `str_ingredient3` | `str` |  |
| `str_ingredient4` | `str` |  |
| `str_ingredient5` | `str` |  |
| `str_ingredient6` | `str` |  |
| `str_ingredient7` | `str` |  |
| `str_ingredient8` | `str` |  |
| `str_ingredient9` | `str` |  |
| `str_instruction` | `str` |  |
| `str_meal` | `str` |  |
| `str_meal_thumb` | `str` |  |
| `str_measure1` | `str` |  |
| `str_measure10` | `str` |  |
| `str_measure11` | `str` |  |
| `str_measure12` | `str` |  |
| `str_measure13` | `str` |  |
| `str_measure14` | `str` |  |
| `str_measure15` | `str` |  |
| `str_measure16` | `str` |  |
| `str_measure17` | `str` |  |
| `str_measure18` | `str` |  |
| `str_measure19` | `str` |  |
| `str_measure2` | `str` |  |
| `str_measure20` | `str` |  |
| `str_measure3` | `str` |  |
| `str_measure4` | `str` |  |
| `str_measure5` | `str` |  |
| `str_measure6` | `str` |  |
| `str_measure7` | `str` |  |
| `str_measure8` | `str` |  |
| `str_measure9` | `str` |  |
| `str_source` | `str` |  |
| `str_tag` | `str` |  |
| `str_youtube` | `str` |  |

#### Example: List

```python
lookups = client.Lookup().list()
```


### Random

Create an instance: `random = client.Random()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `str` |  |
| `id_meal` | `str` |  |
| `str_area` | `str` |  |
| `str_category` | `str` |  |
| `str_creative_commons_confirmed` | `str` |  |
| `str_drink_alternate` | `str` |  |
| `str_image_source` | `str` |  |
| `str_ingredient1` | `str` |  |
| `str_ingredient10` | `str` |  |
| `str_ingredient11` | `str` |  |
| `str_ingredient12` | `str` |  |
| `str_ingredient13` | `str` |  |
| `str_ingredient14` | `str` |  |
| `str_ingredient15` | `str` |  |
| `str_ingredient16` | `str` |  |
| `str_ingredient17` | `str` |  |
| `str_ingredient18` | `str` |  |
| `str_ingredient19` | `str` |  |
| `str_ingredient2` | `str` |  |
| `str_ingredient20` | `str` |  |
| `str_ingredient3` | `str` |  |
| `str_ingredient4` | `str` |  |
| `str_ingredient5` | `str` |  |
| `str_ingredient6` | `str` |  |
| `str_ingredient7` | `str` |  |
| `str_ingredient8` | `str` |  |
| `str_ingredient9` | `str` |  |
| `str_instruction` | `str` |  |
| `str_meal` | `str` |  |
| `str_meal_thumb` | `str` |  |
| `str_measure1` | `str` |  |
| `str_measure10` | `str` |  |
| `str_measure11` | `str` |  |
| `str_measure12` | `str` |  |
| `str_measure13` | `str` |  |
| `str_measure14` | `str` |  |
| `str_measure15` | `str` |  |
| `str_measure16` | `str` |  |
| `str_measure17` | `str` |  |
| `str_measure18` | `str` |  |
| `str_measure19` | `str` |  |
| `str_measure2` | `str` |  |
| `str_measure20` | `str` |  |
| `str_measure3` | `str` |  |
| `str_measure4` | `str` |  |
| `str_measure5` | `str` |  |
| `str_measure6` | `str` |  |
| `str_measure7` | `str` |  |
| `str_measure8` | `str` |  |
| `str_measure9` | `str` |  |
| `str_source` | `str` |  |
| `str_tag` | `str` |  |
| `str_youtube` | `str` |  |

#### Example: List

```python
randoms = client.Random().list()
```


### Randomselection

Create an instance: `randomselection = client.Randomselection()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `str` |  |
| `id_meal` | `str` |  |
| `str_area` | `str` |  |
| `str_category` | `str` |  |
| `str_creative_commons_confirmed` | `str` |  |
| `str_drink_alternate` | `str` |  |
| `str_image_source` | `str` |  |
| `str_ingredient1` | `str` |  |
| `str_ingredient10` | `str` |  |
| `str_ingredient11` | `str` |  |
| `str_ingredient12` | `str` |  |
| `str_ingredient13` | `str` |  |
| `str_ingredient14` | `str` |  |
| `str_ingredient15` | `str` |  |
| `str_ingredient16` | `str` |  |
| `str_ingredient17` | `str` |  |
| `str_ingredient18` | `str` |  |
| `str_ingredient19` | `str` |  |
| `str_ingredient2` | `str` |  |
| `str_ingredient20` | `str` |  |
| `str_ingredient3` | `str` |  |
| `str_ingredient4` | `str` |  |
| `str_ingredient5` | `str` |  |
| `str_ingredient6` | `str` |  |
| `str_ingredient7` | `str` |  |
| `str_ingredient8` | `str` |  |
| `str_ingredient9` | `str` |  |
| `str_instruction` | `str` |  |
| `str_meal` | `str` |  |
| `str_meal_thumb` | `str` |  |
| `str_measure1` | `str` |  |
| `str_measure10` | `str` |  |
| `str_measure11` | `str` |  |
| `str_measure12` | `str` |  |
| `str_measure13` | `str` |  |
| `str_measure14` | `str` |  |
| `str_measure15` | `str` |  |
| `str_measure16` | `str` |  |
| `str_measure17` | `str` |  |
| `str_measure18` | `str` |  |
| `str_measure19` | `str` |  |
| `str_measure2` | `str` |  |
| `str_measure20` | `str` |  |
| `str_measure3` | `str` |  |
| `str_measure4` | `str` |  |
| `str_measure5` | `str` |  |
| `str_measure6` | `str` |  |
| `str_measure7` | `str` |  |
| `str_measure8` | `str` |  |
| `str_measure9` | `str` |  |
| `str_source` | `str` |  |
| `str_tag` | `str` |  |
| `str_youtube` | `str` |  |

#### Example: List

```python
randomselections = client.Randomselection().list()
```


### Search

Create an instance: `search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `str` |  |
| `id_meal` | `str` |  |
| `str_area` | `str` |  |
| `str_category` | `str` |  |
| `str_creative_commons_confirmed` | `str` |  |
| `str_drink_alternate` | `str` |  |
| `str_image_source` | `str` |  |
| `str_ingredient1` | `str` |  |
| `str_ingredient10` | `str` |  |
| `str_ingredient11` | `str` |  |
| `str_ingredient12` | `str` |  |
| `str_ingredient13` | `str` |  |
| `str_ingredient14` | `str` |  |
| `str_ingredient15` | `str` |  |
| `str_ingredient16` | `str` |  |
| `str_ingredient17` | `str` |  |
| `str_ingredient18` | `str` |  |
| `str_ingredient19` | `str` |  |
| `str_ingredient2` | `str` |  |
| `str_ingredient20` | `str` |  |
| `str_ingredient3` | `str` |  |
| `str_ingredient4` | `str` |  |
| `str_ingredient5` | `str` |  |
| `str_ingredient6` | `str` |  |
| `str_ingredient7` | `str` |  |
| `str_ingredient8` | `str` |  |
| `str_ingredient9` | `str` |  |
| `str_instruction` | `str` |  |
| `str_meal` | `str` |  |
| `str_meal_thumb` | `str` |  |
| `str_measure1` | `str` |  |
| `str_measure10` | `str` |  |
| `str_measure11` | `str` |  |
| `str_measure12` | `str` |  |
| `str_measure13` | `str` |  |
| `str_measure14` | `str` |  |
| `str_measure15` | `str` |  |
| `str_measure16` | `str` |  |
| `str_measure17` | `str` |  |
| `str_measure18` | `str` |  |
| `str_measure19` | `str` |  |
| `str_measure2` | `str` |  |
| `str_measure20` | `str` |  |
| `str_measure3` | `str` |  |
| `str_measure4` | `str` |  |
| `str_measure5` | `str` |  |
| `str_measure6` | `str` |  |
| `str_measure7` | `str` |  |
| `str_measure8` | `str` |  |
| `str_measure9` | `str` |  |
| `str_source` | `str` |  |
| `str_tag` | `str` |  |
| `str_youtube` | `str` |  |

#### Example: List

```python
searchs = client.Search().list()
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── freemeal_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`freemeal_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
category = client.Category()
category.list()

# category.data_get() now returns the category data from the last list
# category.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
