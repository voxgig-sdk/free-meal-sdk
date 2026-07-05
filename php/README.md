# FreeMeal PHP SDK



The PHP SDK for the FreeMeal API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Category()` — with named operations (`list`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/free-meal-sdk/releases](https://github.com/voxgig-sdk/free-meal-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'freemeal_sdk.php';

$client = new FreeMealSDK([
    "apikey" => getenv("FREE_MEAL_APIKEY"),
]);
```

### 2. List category records

```php
try {
    // list() returns an array of Category records — iterate directly.
    $categorys = $client->Category()->list();
    foreach ($categorys as $item) {
        echo $item["id_category"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $categorys = $client->Category()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = FreeMealSDK::test();

// Entity ops return the bare mock record (throws on error).
$category = $client->Category()->list();
print_r($category);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new FreeMealSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
FREE_MEAL_TEST_LIVE=TRUE
FREE_MEAL_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### FreeMealSDK

```php
require_once 'freemeal_sdk.php';
$client = new FreeMealSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = FreeMealSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### FreeMealSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Category` | `($data): CategoryEntity` | Create a Category entity instance. |
| `Filter` | `($data): FilterEntity` | Create a Filter entity instance. |
| `Latest` | `($data): LatestEntity` | Create a Latest entity instance. |
| `List` | `($data): ListEntity` | Create a List entity instance. |
| `Lookup` | `($data): LookupEntity` | Create a Lookup entity instance. |
| `Random` | `($data): RandomEntity` | Create a Random entity instance. |
| `Randomselection` | `($data): RandomselectionEntity` | Create a Randomselection entity instance. |
| `Search` | `($data): SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$category = $client->Category();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_category` | `string` |  |
| `str_category` | `string` |  |
| `str_category_description` | `string` |  |
| `str_category_thumb` | `string` |  |

#### Example: List

```php
// list() returns an array of Category records (throws on error).
$categorys = $client->Category()->list();
```


### Filter

Create an instance: `$filter = $client->Filter();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_meal` | `string` |  |
| `str_meal` | `string` |  |
| `str_meal_thumb` | `string` |  |

#### Example: List

```php
// list() returns an array of Filter records (throws on error).
$filters = $client->Filter()->list();
```


### Latest

Create an instance: `$latest = $client->Latest();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `string` |  |
| `id_meal` | `string` |  |
| `str_area` | `string` |  |
| `str_category` | `string` |  |
| `str_creative_commons_confirmed` | `string` |  |
| `str_drink_alternate` | `string` |  |
| `str_image_source` | `string` |  |
| `str_ingredient1` | `string` |  |
| `str_ingredient10` | `string` |  |
| `str_ingredient11` | `string` |  |
| `str_ingredient12` | `string` |  |
| `str_ingredient13` | `string` |  |
| `str_ingredient14` | `string` |  |
| `str_ingredient15` | `string` |  |
| `str_ingredient16` | `string` |  |
| `str_ingredient17` | `string` |  |
| `str_ingredient18` | `string` |  |
| `str_ingredient19` | `string` |  |
| `str_ingredient2` | `string` |  |
| `str_ingredient20` | `string` |  |
| `str_ingredient3` | `string` |  |
| `str_ingredient4` | `string` |  |
| `str_ingredient5` | `string` |  |
| `str_ingredient6` | `string` |  |
| `str_ingredient7` | `string` |  |
| `str_ingredient8` | `string` |  |
| `str_ingredient9` | `string` |  |
| `str_instruction` | `string` |  |
| `str_meal` | `string` |  |
| `str_meal_thumb` | `string` |  |
| `str_measure1` | `string` |  |
| `str_measure10` | `string` |  |
| `str_measure11` | `string` |  |
| `str_measure12` | `string` |  |
| `str_measure13` | `string` |  |
| `str_measure14` | `string` |  |
| `str_measure15` | `string` |  |
| `str_measure16` | `string` |  |
| `str_measure17` | `string` |  |
| `str_measure18` | `string` |  |
| `str_measure19` | `string` |  |
| `str_measure2` | `string` |  |
| `str_measure20` | `string` |  |
| `str_measure3` | `string` |  |
| `str_measure4` | `string` |  |
| `str_measure5` | `string` |  |
| `str_measure6` | `string` |  |
| `str_measure7` | `string` |  |
| `str_measure8` | `string` |  |
| `str_measure9` | `string` |  |
| `str_source` | `string` |  |
| `str_tag` | `string` |  |
| `str_youtube` | `string` |  |

#### Example: List

```php
// list() returns an array of Latest records (throws on error).
$latests = $client->Latest()->list();
```


### List

Create an instance: `$list = $client->List();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `str_area` | `string` |  |
| `str_category` | `string` |  |
| `str_ingredient` | `string` |  |

#### Example: List

```php
// list() returns an array of List records (throws on error).
$lists = $client->List()->list();
```


### Lookup

Create an instance: `$lookup = $client->Lookup();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `string` |  |
| `id_meal` | `string` |  |
| `str_area` | `string` |  |
| `str_category` | `string` |  |
| `str_creative_commons_confirmed` | `string` |  |
| `str_drink_alternate` | `string` |  |
| `str_image_source` | `string` |  |
| `str_ingredient1` | `string` |  |
| `str_ingredient10` | `string` |  |
| `str_ingredient11` | `string` |  |
| `str_ingredient12` | `string` |  |
| `str_ingredient13` | `string` |  |
| `str_ingredient14` | `string` |  |
| `str_ingredient15` | `string` |  |
| `str_ingredient16` | `string` |  |
| `str_ingredient17` | `string` |  |
| `str_ingredient18` | `string` |  |
| `str_ingredient19` | `string` |  |
| `str_ingredient2` | `string` |  |
| `str_ingredient20` | `string` |  |
| `str_ingredient3` | `string` |  |
| `str_ingredient4` | `string` |  |
| `str_ingredient5` | `string` |  |
| `str_ingredient6` | `string` |  |
| `str_ingredient7` | `string` |  |
| `str_ingredient8` | `string` |  |
| `str_ingredient9` | `string` |  |
| `str_instruction` | `string` |  |
| `str_meal` | `string` |  |
| `str_meal_thumb` | `string` |  |
| `str_measure1` | `string` |  |
| `str_measure10` | `string` |  |
| `str_measure11` | `string` |  |
| `str_measure12` | `string` |  |
| `str_measure13` | `string` |  |
| `str_measure14` | `string` |  |
| `str_measure15` | `string` |  |
| `str_measure16` | `string` |  |
| `str_measure17` | `string` |  |
| `str_measure18` | `string` |  |
| `str_measure19` | `string` |  |
| `str_measure2` | `string` |  |
| `str_measure20` | `string` |  |
| `str_measure3` | `string` |  |
| `str_measure4` | `string` |  |
| `str_measure5` | `string` |  |
| `str_measure6` | `string` |  |
| `str_measure7` | `string` |  |
| `str_measure8` | `string` |  |
| `str_measure9` | `string` |  |
| `str_source` | `string` |  |
| `str_tag` | `string` |  |
| `str_youtube` | `string` |  |

#### Example: List

```php
// list() returns an array of Lookup records (throws on error).
$lookups = $client->Lookup()->list();
```


### Random

Create an instance: `$random = $client->Random();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `string` |  |
| `id_meal` | `string` |  |
| `str_area` | `string` |  |
| `str_category` | `string` |  |
| `str_creative_commons_confirmed` | `string` |  |
| `str_drink_alternate` | `string` |  |
| `str_image_source` | `string` |  |
| `str_ingredient1` | `string` |  |
| `str_ingredient10` | `string` |  |
| `str_ingredient11` | `string` |  |
| `str_ingredient12` | `string` |  |
| `str_ingredient13` | `string` |  |
| `str_ingredient14` | `string` |  |
| `str_ingredient15` | `string` |  |
| `str_ingredient16` | `string` |  |
| `str_ingredient17` | `string` |  |
| `str_ingredient18` | `string` |  |
| `str_ingredient19` | `string` |  |
| `str_ingredient2` | `string` |  |
| `str_ingredient20` | `string` |  |
| `str_ingredient3` | `string` |  |
| `str_ingredient4` | `string` |  |
| `str_ingredient5` | `string` |  |
| `str_ingredient6` | `string` |  |
| `str_ingredient7` | `string` |  |
| `str_ingredient8` | `string` |  |
| `str_ingredient9` | `string` |  |
| `str_instruction` | `string` |  |
| `str_meal` | `string` |  |
| `str_meal_thumb` | `string` |  |
| `str_measure1` | `string` |  |
| `str_measure10` | `string` |  |
| `str_measure11` | `string` |  |
| `str_measure12` | `string` |  |
| `str_measure13` | `string` |  |
| `str_measure14` | `string` |  |
| `str_measure15` | `string` |  |
| `str_measure16` | `string` |  |
| `str_measure17` | `string` |  |
| `str_measure18` | `string` |  |
| `str_measure19` | `string` |  |
| `str_measure2` | `string` |  |
| `str_measure20` | `string` |  |
| `str_measure3` | `string` |  |
| `str_measure4` | `string` |  |
| `str_measure5` | `string` |  |
| `str_measure6` | `string` |  |
| `str_measure7` | `string` |  |
| `str_measure8` | `string` |  |
| `str_measure9` | `string` |  |
| `str_source` | `string` |  |
| `str_tag` | `string` |  |
| `str_youtube` | `string` |  |

#### Example: List

```php
// list() returns an array of Random records (throws on error).
$randoms = $client->Random()->list();
```


### Randomselection

Create an instance: `$randomselection = $client->Randomselection();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `string` |  |
| `id_meal` | `string` |  |
| `str_area` | `string` |  |
| `str_category` | `string` |  |
| `str_creative_commons_confirmed` | `string` |  |
| `str_drink_alternate` | `string` |  |
| `str_image_source` | `string` |  |
| `str_ingredient1` | `string` |  |
| `str_ingredient10` | `string` |  |
| `str_ingredient11` | `string` |  |
| `str_ingredient12` | `string` |  |
| `str_ingredient13` | `string` |  |
| `str_ingredient14` | `string` |  |
| `str_ingredient15` | `string` |  |
| `str_ingredient16` | `string` |  |
| `str_ingredient17` | `string` |  |
| `str_ingredient18` | `string` |  |
| `str_ingredient19` | `string` |  |
| `str_ingredient2` | `string` |  |
| `str_ingredient20` | `string` |  |
| `str_ingredient3` | `string` |  |
| `str_ingredient4` | `string` |  |
| `str_ingredient5` | `string` |  |
| `str_ingredient6` | `string` |  |
| `str_ingredient7` | `string` |  |
| `str_ingredient8` | `string` |  |
| `str_ingredient9` | `string` |  |
| `str_instruction` | `string` |  |
| `str_meal` | `string` |  |
| `str_meal_thumb` | `string` |  |
| `str_measure1` | `string` |  |
| `str_measure10` | `string` |  |
| `str_measure11` | `string` |  |
| `str_measure12` | `string` |  |
| `str_measure13` | `string` |  |
| `str_measure14` | `string` |  |
| `str_measure15` | `string` |  |
| `str_measure16` | `string` |  |
| `str_measure17` | `string` |  |
| `str_measure18` | `string` |  |
| `str_measure19` | `string` |  |
| `str_measure2` | `string` |  |
| `str_measure20` | `string` |  |
| `str_measure3` | `string` |  |
| `str_measure4` | `string` |  |
| `str_measure5` | `string` |  |
| `str_measure6` | `string` |  |
| `str_measure7` | `string` |  |
| `str_measure8` | `string` |  |
| `str_measure9` | `string` |  |
| `str_source` | `string` |  |
| `str_tag` | `string` |  |
| `str_youtube` | `string` |  |

#### Example: List

```php
// list() returns an array of Randomselection records (throws on error).
$randomselections = $client->Randomselection()->list();
```


### Search

Create an instance: `$search = $client->Search();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date_modified` | `string` |  |
| `id_meal` | `string` |  |
| `str_area` | `string` |  |
| `str_category` | `string` |  |
| `str_creative_commons_confirmed` | `string` |  |
| `str_drink_alternate` | `string` |  |
| `str_image_source` | `string` |  |
| `str_ingredient1` | `string` |  |
| `str_ingredient10` | `string` |  |
| `str_ingredient11` | `string` |  |
| `str_ingredient12` | `string` |  |
| `str_ingredient13` | `string` |  |
| `str_ingredient14` | `string` |  |
| `str_ingredient15` | `string` |  |
| `str_ingredient16` | `string` |  |
| `str_ingredient17` | `string` |  |
| `str_ingredient18` | `string` |  |
| `str_ingredient19` | `string` |  |
| `str_ingredient2` | `string` |  |
| `str_ingredient20` | `string` |  |
| `str_ingredient3` | `string` |  |
| `str_ingredient4` | `string` |  |
| `str_ingredient5` | `string` |  |
| `str_ingredient6` | `string` |  |
| `str_ingredient7` | `string` |  |
| `str_ingredient8` | `string` |  |
| `str_ingredient9` | `string` |  |
| `str_instruction` | `string` |  |
| `str_meal` | `string` |  |
| `str_meal_thumb` | `string` |  |
| `str_measure1` | `string` |  |
| `str_measure10` | `string` |  |
| `str_measure11` | `string` |  |
| `str_measure12` | `string` |  |
| `str_measure13` | `string` |  |
| `str_measure14` | `string` |  |
| `str_measure15` | `string` |  |
| `str_measure16` | `string` |  |
| `str_measure17` | `string` |  |
| `str_measure18` | `string` |  |
| `str_measure19` | `string` |  |
| `str_measure2` | `string` |  |
| `str_measure20` | `string` |  |
| `str_measure3` | `string` |  |
| `str_measure4` | `string` |  |
| `str_measure5` | `string` |  |
| `str_measure6` | `string` |  |
| `str_measure7` | `string` |  |
| `str_measure8` | `string` |  |
| `str_measure9` | `string` |  |
| `str_source` | `string` |  |
| `str_tag` | `string` |  |
| `str_youtube` | `string` |  |

#### Example: List

```php
// list() returns an array of Search records (throws on error).
$searchs = $client->Search()->list();
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── freemeal_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`freemeal_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$category = $client->Category();
$category->list();

// $category->data_get() now returns the category data from the last list
// $category->match_get() returns the last match criteria
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
