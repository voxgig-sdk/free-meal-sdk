# FreeMeal PHP SDK Reference

Complete API reference for the FreeMeal PHP SDK.


## FreeMealSDK

### Constructor

```php
require_once __DIR__ . '/freemeal_sdk.php';

$client = new FreeMealSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `FreeMealSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = FreeMealSDK::test();
```


### Instance Methods

#### `Category($data = null)`

Create a new `CategoryEntity` instance. Pass `null` for no initial data.

#### `Filter($data = null)`

Create a new `FilterEntity` instance. Pass `null` for no initial data.

#### `Latest($data = null)`

Create a new `LatestEntity` instance. Pass `null` for no initial data.

#### `List($data = null)`

Create a new `ListEntity` instance. Pass `null` for no initial data.

#### `Lookup($data = null)`

Create a new `LookupEntity` instance. Pass `null` for no initial data.

#### `Random($data = null)`

Create a new `RandomEntity` instance. Pass `null` for no initial data.

#### `Randomselection($data = null)`

Create a new `RandomselectionEntity` instance. Pass `null` for no initial data.

#### `Search($data = null)`

Create a new `SearchEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): FreeMealUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## CategoryEntity

```php
$category = $client->Category();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_category` | `string` | No |  |
| `str_category` | `string` | No |  |
| `str_category_description` | `string` | No |  |
| `str_category_thumb` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Category()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CategoryEntity`

Create a new `CategoryEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FilterEntity

```php
$filter = $client->Filter();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_meal` | `string` | No |  |
| `str_meal` | `string` | No |  |
| `str_meal_thumb` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Filter()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FilterEntity`

Create a new `FilterEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## LatestEntity

```php
$latest = $client->Latest();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | `string` | No |  |
| `id_meal` | `string` | No |  |
| `str_area` | `string` | No |  |
| `str_category` | `string` | No |  |
| `str_creative_commons_confirmed` | `string` | No |  |
| `str_drink_alternate` | `string` | No |  |
| `str_image_source` | `string` | No |  |
| `str_ingredient1` | `string` | No |  |
| `str_ingredient10` | `string` | No |  |
| `str_ingredient11` | `string` | No |  |
| `str_ingredient12` | `string` | No |  |
| `str_ingredient13` | `string` | No |  |
| `str_ingredient14` | `string` | No |  |
| `str_ingredient15` | `string` | No |  |
| `str_ingredient16` | `string` | No |  |
| `str_ingredient17` | `string` | No |  |
| `str_ingredient18` | `string` | No |  |
| `str_ingredient19` | `string` | No |  |
| `str_ingredient2` | `string` | No |  |
| `str_ingredient20` | `string` | No |  |
| `str_ingredient3` | `string` | No |  |
| `str_ingredient4` | `string` | No |  |
| `str_ingredient5` | `string` | No |  |
| `str_ingredient6` | `string` | No |  |
| `str_ingredient7` | `string` | No |  |
| `str_ingredient8` | `string` | No |  |
| `str_ingredient9` | `string` | No |  |
| `str_instruction` | `string` | No |  |
| `str_meal` | `string` | No |  |
| `str_meal_thumb` | `string` | No |  |
| `str_measure1` | `string` | No |  |
| `str_measure10` | `string` | No |  |
| `str_measure11` | `string` | No |  |
| `str_measure12` | `string` | No |  |
| `str_measure13` | `string` | No |  |
| `str_measure14` | `string` | No |  |
| `str_measure15` | `string` | No |  |
| `str_measure16` | `string` | No |  |
| `str_measure17` | `string` | No |  |
| `str_measure18` | `string` | No |  |
| `str_measure19` | `string` | No |  |
| `str_measure2` | `string` | No |  |
| `str_measure20` | `string` | No |  |
| `str_measure3` | `string` | No |  |
| `str_measure4` | `string` | No |  |
| `str_measure5` | `string` | No |  |
| `str_measure6` | `string` | No |  |
| `str_measure7` | `string` | No |  |
| `str_measure8` | `string` | No |  |
| `str_measure9` | `string` | No |  |
| `str_source` | `string` | No |  |
| `str_tag` | `string` | No |  |
| `str_youtube` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Latest()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): LatestEntity`

Create a new `LatestEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ListEntity

```php
$list = $client->List();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `str_area` | `string` | No |  |
| `str_category` | `string` | No |  |
| `str_ingredient` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->List()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ListEntity`

Create a new `ListEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## LookupEntity

```php
$lookup = $client->Lookup();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | `string` | No |  |
| `id_meal` | `string` | No |  |
| `str_area` | `string` | No |  |
| `str_category` | `string` | No |  |
| `str_creative_commons_confirmed` | `string` | No |  |
| `str_drink_alternate` | `string` | No |  |
| `str_image_source` | `string` | No |  |
| `str_ingredient1` | `string` | No |  |
| `str_ingredient10` | `string` | No |  |
| `str_ingredient11` | `string` | No |  |
| `str_ingredient12` | `string` | No |  |
| `str_ingredient13` | `string` | No |  |
| `str_ingredient14` | `string` | No |  |
| `str_ingredient15` | `string` | No |  |
| `str_ingredient16` | `string` | No |  |
| `str_ingredient17` | `string` | No |  |
| `str_ingredient18` | `string` | No |  |
| `str_ingredient19` | `string` | No |  |
| `str_ingredient2` | `string` | No |  |
| `str_ingredient20` | `string` | No |  |
| `str_ingredient3` | `string` | No |  |
| `str_ingredient4` | `string` | No |  |
| `str_ingredient5` | `string` | No |  |
| `str_ingredient6` | `string` | No |  |
| `str_ingredient7` | `string` | No |  |
| `str_ingredient8` | `string` | No |  |
| `str_ingredient9` | `string` | No |  |
| `str_instruction` | `string` | No |  |
| `str_meal` | `string` | No |  |
| `str_meal_thumb` | `string` | No |  |
| `str_measure1` | `string` | No |  |
| `str_measure10` | `string` | No |  |
| `str_measure11` | `string` | No |  |
| `str_measure12` | `string` | No |  |
| `str_measure13` | `string` | No |  |
| `str_measure14` | `string` | No |  |
| `str_measure15` | `string` | No |  |
| `str_measure16` | `string` | No |  |
| `str_measure17` | `string` | No |  |
| `str_measure18` | `string` | No |  |
| `str_measure19` | `string` | No |  |
| `str_measure2` | `string` | No |  |
| `str_measure20` | `string` | No |  |
| `str_measure3` | `string` | No |  |
| `str_measure4` | `string` | No |  |
| `str_measure5` | `string` | No |  |
| `str_measure6` | `string` | No |  |
| `str_measure7` | `string` | No |  |
| `str_measure8` | `string` | No |  |
| `str_measure9` | `string` | No |  |
| `str_source` | `string` | No |  |
| `str_tag` | `string` | No |  |
| `str_youtube` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Lookup()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): LookupEntity`

Create a new `LookupEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RandomEntity

```php
$random = $client->Random();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | `string` | No |  |
| `id_meal` | `string` | No |  |
| `str_area` | `string` | No |  |
| `str_category` | `string` | No |  |
| `str_creative_commons_confirmed` | `string` | No |  |
| `str_drink_alternate` | `string` | No |  |
| `str_image_source` | `string` | No |  |
| `str_ingredient1` | `string` | No |  |
| `str_ingredient10` | `string` | No |  |
| `str_ingredient11` | `string` | No |  |
| `str_ingredient12` | `string` | No |  |
| `str_ingredient13` | `string` | No |  |
| `str_ingredient14` | `string` | No |  |
| `str_ingredient15` | `string` | No |  |
| `str_ingredient16` | `string` | No |  |
| `str_ingredient17` | `string` | No |  |
| `str_ingredient18` | `string` | No |  |
| `str_ingredient19` | `string` | No |  |
| `str_ingredient2` | `string` | No |  |
| `str_ingredient20` | `string` | No |  |
| `str_ingredient3` | `string` | No |  |
| `str_ingredient4` | `string` | No |  |
| `str_ingredient5` | `string` | No |  |
| `str_ingredient6` | `string` | No |  |
| `str_ingredient7` | `string` | No |  |
| `str_ingredient8` | `string` | No |  |
| `str_ingredient9` | `string` | No |  |
| `str_instruction` | `string` | No |  |
| `str_meal` | `string` | No |  |
| `str_meal_thumb` | `string` | No |  |
| `str_measure1` | `string` | No |  |
| `str_measure10` | `string` | No |  |
| `str_measure11` | `string` | No |  |
| `str_measure12` | `string` | No |  |
| `str_measure13` | `string` | No |  |
| `str_measure14` | `string` | No |  |
| `str_measure15` | `string` | No |  |
| `str_measure16` | `string` | No |  |
| `str_measure17` | `string` | No |  |
| `str_measure18` | `string` | No |  |
| `str_measure19` | `string` | No |  |
| `str_measure2` | `string` | No |  |
| `str_measure20` | `string` | No |  |
| `str_measure3` | `string` | No |  |
| `str_measure4` | `string` | No |  |
| `str_measure5` | `string` | No |  |
| `str_measure6` | `string` | No |  |
| `str_measure7` | `string` | No |  |
| `str_measure8` | `string` | No |  |
| `str_measure9` | `string` | No |  |
| `str_source` | `string` | No |  |
| `str_tag` | `string` | No |  |
| `str_youtube` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Random()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RandomEntity`

Create a new `RandomEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RandomselectionEntity

```php
$randomselection = $client->Randomselection();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | `string` | No |  |
| `id_meal` | `string` | No |  |
| `str_area` | `string` | No |  |
| `str_category` | `string` | No |  |
| `str_creative_commons_confirmed` | `string` | No |  |
| `str_drink_alternate` | `string` | No |  |
| `str_image_source` | `string` | No |  |
| `str_ingredient1` | `string` | No |  |
| `str_ingredient10` | `string` | No |  |
| `str_ingredient11` | `string` | No |  |
| `str_ingredient12` | `string` | No |  |
| `str_ingredient13` | `string` | No |  |
| `str_ingredient14` | `string` | No |  |
| `str_ingredient15` | `string` | No |  |
| `str_ingredient16` | `string` | No |  |
| `str_ingredient17` | `string` | No |  |
| `str_ingredient18` | `string` | No |  |
| `str_ingredient19` | `string` | No |  |
| `str_ingredient2` | `string` | No |  |
| `str_ingredient20` | `string` | No |  |
| `str_ingredient3` | `string` | No |  |
| `str_ingredient4` | `string` | No |  |
| `str_ingredient5` | `string` | No |  |
| `str_ingredient6` | `string` | No |  |
| `str_ingredient7` | `string` | No |  |
| `str_ingredient8` | `string` | No |  |
| `str_ingredient9` | `string` | No |  |
| `str_instruction` | `string` | No |  |
| `str_meal` | `string` | No |  |
| `str_meal_thumb` | `string` | No |  |
| `str_measure1` | `string` | No |  |
| `str_measure10` | `string` | No |  |
| `str_measure11` | `string` | No |  |
| `str_measure12` | `string` | No |  |
| `str_measure13` | `string` | No |  |
| `str_measure14` | `string` | No |  |
| `str_measure15` | `string` | No |  |
| `str_measure16` | `string` | No |  |
| `str_measure17` | `string` | No |  |
| `str_measure18` | `string` | No |  |
| `str_measure19` | `string` | No |  |
| `str_measure2` | `string` | No |  |
| `str_measure20` | `string` | No |  |
| `str_measure3` | `string` | No |  |
| `str_measure4` | `string` | No |  |
| `str_measure5` | `string` | No |  |
| `str_measure6` | `string` | No |  |
| `str_measure7` | `string` | No |  |
| `str_measure8` | `string` | No |  |
| `str_measure9` | `string` | No |  |
| `str_source` | `string` | No |  |
| `str_tag` | `string` | No |  |
| `str_youtube` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Randomselection()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RandomselectionEntity`

Create a new `RandomselectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date_modified` | `string` | No |  |
| `id_meal` | `string` | No |  |
| `str_area` | `string` | No |  |
| `str_category` | `string` | No |  |
| `str_creative_commons_confirmed` | `string` | No |  |
| `str_drink_alternate` | `string` | No |  |
| `str_image_source` | `string` | No |  |
| `str_ingredient1` | `string` | No |  |
| `str_ingredient10` | `string` | No |  |
| `str_ingredient11` | `string` | No |  |
| `str_ingredient12` | `string` | No |  |
| `str_ingredient13` | `string` | No |  |
| `str_ingredient14` | `string` | No |  |
| `str_ingredient15` | `string` | No |  |
| `str_ingredient16` | `string` | No |  |
| `str_ingredient17` | `string` | No |  |
| `str_ingredient18` | `string` | No |  |
| `str_ingredient19` | `string` | No |  |
| `str_ingredient2` | `string` | No |  |
| `str_ingredient20` | `string` | No |  |
| `str_ingredient3` | `string` | No |  |
| `str_ingredient4` | `string` | No |  |
| `str_ingredient5` | `string` | No |  |
| `str_ingredient6` | `string` | No |  |
| `str_ingredient7` | `string` | No |  |
| `str_ingredient8` | `string` | No |  |
| `str_ingredient9` | `string` | No |  |
| `str_instruction` | `string` | No |  |
| `str_meal` | `string` | No |  |
| `str_meal_thumb` | `string` | No |  |
| `str_measure1` | `string` | No |  |
| `str_measure10` | `string` | No |  |
| `str_measure11` | `string` | No |  |
| `str_measure12` | `string` | No |  |
| `str_measure13` | `string` | No |  |
| `str_measure14` | `string` | No |  |
| `str_measure15` | `string` | No |  |
| `str_measure16` | `string` | No |  |
| `str_measure17` | `string` | No |  |
| `str_measure18` | `string` | No |  |
| `str_measure19` | `string` | No |  |
| `str_measure2` | `string` | No |  |
| `str_measure20` | `string` | No |  |
| `str_measure3` | `string` | No |  |
| `str_measure4` | `string` | No |  |
| `str_measure5` | `string` | No |  |
| `str_measure6` | `string` | No |  |
| `str_measure7` | `string` | No |  |
| `str_measure8` | `string` | No |  |
| `str_measure9` | `string` | No |  |
| `str_source` | `string` | No |  |
| `str_tag` | `string` | No |  |
| `str_youtube` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Search()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new FreeMealSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

