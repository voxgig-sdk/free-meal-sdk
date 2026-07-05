# FreeMeal TypeScript SDK



The TypeScript SDK for the FreeMeal API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Category()` — each with a small set of operations (`list`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/free-meal-sdk/releases](https://github.com/voxgig-sdk/free-meal-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { FreeMealSDK } from '@voxgig-sdk/free-meal'

const client = new FreeMealSDK({
  apikey: process.env.FREE_MEAL_APIKEY,
})
```

### 2. List category records

`list()` resolves to an array of Category objects — iterate it directly:

```ts
const categorys = await client.Category().list()

for (const category of categorys) {
  console.log(category)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const categorys = await client.Category().list()
  console.log(categorys)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = FreeMealSDK.test()

const category = await client.Category().list()
// category is a bare entity populated with mock response data
console.log(category)
```

You can also use the instance method:

```ts
const client = new FreeMealSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Category()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new FreeMealSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```


## Reference

### FreeMealSDK

#### Constructor

```ts
new FreeMealSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Category(data?)` | `CategoryEntity` | Create a Category entity instance. |
| `Filter(data?)` | `FilterEntity` | Create a Filter entity instance. |
| `Latest(data?)` | `LatestEntity` | Create a Latest entity instance. |
| `List(data?)` | `ListEntity` | Create a List entity instance. |
| `Lookup(data?)` | `LookupEntity` | Create a Lookup entity instance. |
| `Random(data?)` | `RandomEntity` | Create a Random entity instance. |
| `Randomselection(data?)` | `RandomselectionEntity` | Create a Randomselection entity instance. |
| `Search(data?)` | `SearchEntity` | Create a Search entity instance. |
| `tester(testopts?, sdkopts?)` | `FreeMealSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `FreeMealSDK.test(testopts?, sdkopts?)` | `FreeMealSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): FreeMealSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### Category

| Field | Description |
| --- | --- |
| `id_category` |  |
| `str_category` |  |
| `str_category_description` |  |
| `str_category_thumb` |  |

Operations: list.

API path: `/categories.php`

#### Filter

| Field | Description |
| --- | --- |
| `id_meal` |  |
| `str_meal` |  |
| `str_meal_thumb` |  |

Operations: list.

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

Operations: list.

API path: `/latest.php`

#### List

| Field | Description |
| --- | --- |
| `str_area` |  |
| `str_category` |  |
| `str_ingredient` |  |

Operations: list.

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

Operations: list.

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

Operations: list.

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

Operations: list.

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

Operations: list.

API path: `/search.php`



## Entities


### Category

Create an instance: `const category = client.Category()`

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

```ts
const categorys = await client.Category().list()
```


### Filter

Create an instance: `const filter = client.Filter()`

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

```ts
const filters = await client.Filter().list()
```


### Latest

Create an instance: `const latest = client.Latest()`

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

```ts
const latests = await client.Latest().list()
```


### List

Create an instance: `const list = client.List()`

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

```ts
const lists = await client.List().list()
```


### Lookup

Create an instance: `const lookup = client.Lookup()`

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

```ts
const lookups = await client.Lookup().list()
```


### Random

Create an instance: `const random = client.Random()`

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

```ts
const randoms = await client.Random().list()
```


### Randomselection

Create an instance: `const randomselection = client.Randomselection()`

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

```ts
const randomselections = await client.Randomselection().list()
```


### Search

Create an instance: `const search = client.Search()`

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

```ts
const searchs = await client.Search().list()
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
free-meal/
├── src/
│   ├── FreeMealSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { FreeMealSDK } from '@voxgig-sdk/free-meal'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const category = client.Category()
await category.list()

// category.data() now returns the category data from the last `list`
// category.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
