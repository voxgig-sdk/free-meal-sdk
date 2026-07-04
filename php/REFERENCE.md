# FreeMeal PHP SDK Reference

Complete API reference for the FreeMeal PHP SDK.


## FreeMealSDK

### Constructor

```php
require_once __DIR__ . '/free-meal_sdk.php';

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

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

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
$category = $client->category();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_category` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_category_description` | ``$STRING`` | No |  |
| `str_category_thumb` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->category()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): CategoryEntity`

Create a new `CategoryEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## FilterEntity

```php
$filter = $client->filter();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id_meal` | ``$STRING`` | No |  |
| `str_meal` | ``$STRING`` | No |  |
| `str_meal_thumb` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->filter()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): FilterEntity`

Create a new `FilterEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## LatestEntity

```php
$latest = $client->latest();
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

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->latest()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): LatestEntity`

Create a new `LatestEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ListEntity

```php
$list = $client->list();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `str_area` | ``$STRING`` | No |  |
| `str_category` | ``$STRING`` | No |  |
| `str_ingredient` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->list()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ListEntity`

Create a new `ListEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## LookupEntity

```php
$lookup = $client->lookup();
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

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->lookup()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): LookupEntity`

Create a new `LookupEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## RandomEntity

```php
$random = $client->random();
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

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->random()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): RandomEntity`

Create a new `RandomEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## RandomselectionEntity

```php
$randomselection = $client->randomselection();
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

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->randomselection()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): RandomselectionEntity`

Create a new `RandomselectionEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->search();
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

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->search()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `getName(): string`

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

