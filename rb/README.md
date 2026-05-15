# FreeMeal Ruby SDK

The Ruby SDK for the FreeMeal API. Provides an entity-oriented interface using idiomatic Ruby conventions.


## Install
```bash
gem install free-meal-sdk
```

Or add to your `Gemfile`:

```ruby
gem "free-meal-sdk"
```

Then run:

```bash
bundle install
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "FreeMeal_sdk"

client = FreeMealSDK.new({
  "apikey" => ENV["FREE-MEAL_APIKEY"],
})
```

### 2. List categorys

```ruby
result, err = client.Category(nil).list(nil, nil)
raise err if err

if result.is_a?(Array)
  result.each do |item|
    d = item.data_get
    puts "#{d["id"]} #{d["name"]}"
  end
end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
raise err if err

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
end
```

### Prepare a request without sending it

```ruby
fetchdef, err = client.prepare({
  "path" => "/api/resource/{id}",
  "method" => "DELETE",
  "params" => { "id" => "example" },
})
raise err if err

puts fetchdef["url"]
puts fetchdef["method"]
puts fetchdef["headers"]
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = FreeMealSDK.test(nil, nil)

result, err = client.FreeMeal(nil).load(
  { "id" => "test01" }, nil
)
# result contains mock response data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = FreeMealSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
FREE-MEAL_TEST_LIVE=TRUE
FREE-MEAL_APIKEY=<your-key>
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### FreeMealSDK

```ruby
require_relative "FreeMeal_sdk"
client = FreeMealSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = FreeMealSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### FreeMealSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> [Hash, err]` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> [Hash, err]` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> [any, err]` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> [any, err]` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> [any, err]` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> [any, err]` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> [any, err]` | Remove an entity. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return `[any, err]`. The first value is a
`Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `false` and `err` contains the error value.

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

Create an instance: `const category = client.Category()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id_category` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_category_description` | ``$STRING`` |  |
| `str_category_thumb` | ``$STRING`` |  |

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
| `id_meal` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |

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
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

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
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_ingredient` | ``$STRING`` |  |

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
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

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
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

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
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

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
| `date_modified` | ``$STRING`` |  |
| `id_meal` | ``$STRING`` |  |
| `str_area` | ``$STRING`` |  |
| `str_category` | ``$STRING`` |  |
| `str_creative_commons_confirmed` | ``$STRING`` |  |
| `str_drink_alternate` | ``$STRING`` |  |
| `str_image_source` | ``$STRING`` |  |
| `str_ingredient1` | ``$STRING`` |  |
| `str_ingredient10` | ``$STRING`` |  |
| `str_ingredient11` | ``$STRING`` |  |
| `str_ingredient12` | ``$STRING`` |  |
| `str_ingredient13` | ``$STRING`` |  |
| `str_ingredient14` | ``$STRING`` |  |
| `str_ingredient15` | ``$STRING`` |  |
| `str_ingredient16` | ``$STRING`` |  |
| `str_ingredient17` | ``$STRING`` |  |
| `str_ingredient18` | ``$STRING`` |  |
| `str_ingredient19` | ``$STRING`` |  |
| `str_ingredient2` | ``$STRING`` |  |
| `str_ingredient20` | ``$STRING`` |  |
| `str_ingredient3` | ``$STRING`` |  |
| `str_ingredient4` | ``$STRING`` |  |
| `str_ingredient5` | ``$STRING`` |  |
| `str_ingredient6` | ``$STRING`` |  |
| `str_ingredient7` | ``$STRING`` |  |
| `str_ingredient8` | ``$STRING`` |  |
| `str_ingredient9` | ``$STRING`` |  |
| `str_instruction` | ``$STRING`` |  |
| `str_meal` | ``$STRING`` |  |
| `str_meal_thumb` | ``$STRING`` |  |
| `str_measure1` | ``$STRING`` |  |
| `str_measure10` | ``$STRING`` |  |
| `str_measure11` | ``$STRING`` |  |
| `str_measure12` | ``$STRING`` |  |
| `str_measure13` | ``$STRING`` |  |
| `str_measure14` | ``$STRING`` |  |
| `str_measure15` | ``$STRING`` |  |
| `str_measure16` | ``$STRING`` |  |
| `str_measure17` | ``$STRING`` |  |
| `str_measure18` | ``$STRING`` |  |
| `str_measure19` | ``$STRING`` |  |
| `str_measure2` | ``$STRING`` |  |
| `str_measure20` | ``$STRING`` |  |
| `str_measure3` | ``$STRING`` |  |
| `str_measure4` | ``$STRING`` |  |
| `str_measure5` | ``$STRING`` |  |
| `str_measure6` | ``$STRING`` |  |
| `str_measure7` | ``$STRING`` |  |
| `str_measure8` | ``$STRING`` |  |
| `str_measure9` | ``$STRING`` |  |
| `str_source` | ``$STRING`` |  |
| `str_tag` | ``$STRING`` |  |
| `str_youtube` | ``$STRING`` |  |

#### Example: List

```ts
const searchs = await client.Search().list()
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

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

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller as a second return value.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── FreeMeal_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`FreeMeal_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
moon = client.Moon
moon.load({ "planet_id" => "earth", "id" => "luna" })

# moon.data_get now returns the loaded moon data
# moon.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
