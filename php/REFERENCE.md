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
| `idCategory` | `string` | No | Unique category identifier |
| `strCategory` | `string` | No | Category name |
| `strCategoryDescription` | `string` | No | Category description |
| `strCategoryThumb` | `string` | No | URL to category thumbnail image |

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
| `idMeal` | `string` | No | Unique meal identifier |
| `strMeal` | `string` | No | Meal name |
| `strMealThumb` | `string` | No | URL to meal thumbnail image |

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
| `strArea` | `string` | No |  |
| `strCategory` | `string` | No |  |
| `strIngredient` | `string` | No |  |

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

