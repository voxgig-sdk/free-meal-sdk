<?php
declare(strict_types=1);

// FreeMeal SDK configuration

class FreeMealConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "FreeMeal",
                "slug" => "free-meal",
                "version" => "0.0.1",
                "target" => "php",
            ],
            "feature" => [
                "ratelimit" => [
          'options' => [
            'active' => false,
            'burst' => 5,
            'rate' => 5,
          ],
          'optspec' => [
            'now' => '`$FUNCTION`',
            'sleep' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'wrap',
        ],
                "retry" => [
          'options' => [
            'active' => false,
            'factor' => 2,
            'maxDelay' => 2000,
            'minDelay' => 50,
            'retries' => 2,
            'statuses' => [
              408,
              425,
              429,
              500,
              502,
              503,
              504,
            ],
          ],
          'optspec' => [
            'jitter' => '`$BOOLEAN`',
            'sleep' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'wrap',
        ],
                "test" => [
          'options' => [
            'active' => false,
          ],
          'optspec' => [
            'entity' => '`$MAP`',
            'net' => '`$MAP`',
          ],
          'strict' => false,
          'transport' => 'base',
        ],
                "timeout" => [
          'options' => [
            'active' => false,
            'ms' => 30000,
          ],
          'optspec' => [
            'clearTimer' => '`$FUNCTION`',
            'setTimer' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'wrap',
        ],
            ],
            "options" => [
                "base" => "https://www.themealdb.com/api/json/v1/1",
                "auth" => [
                    "prefix" => "",
                    "in" => "path",
                    "name" => "api_key",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "category" => [],
                    "filter" => [],
                    "latest" => [],
                    "list" => [],
                    "lookup" => [],
                    "random" => [],
                    "randomselection" => [],
                    "search" => [],
                ],
            ],
            "entity" => [
        'category' => [
          'fields' => [
            [
              'name' => 'idCategory',
              'title' => 'Id Category',
              'type' => '`$STRING`',
              'short' => 'Unique category identifier',
            ],
            [
              'name' => 'strCategory',
              'title' => 'Str Category',
              'type' => '`$STRING`',
              'short' => 'Category name',
            ],
            [
              'name' => 'strCategoryDescription',
              'title' => 'Str Category Description',
              'type' => '`$STRING`',
              'short' => 'Category description',
            ],
            [
              'name' => 'strCategoryThumb',
              'title' => 'Str Category Thumb',
              'type' => '`$STRING`',
              'short' => 'URL to category thumbnail image',
            ],
          ],
          'name' => 'category',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/categories.php',
                  'segments' => [
                    [
                      'lit' => 'categories.php',
                    ],
                  ],
                  'parts' => [
                    'categories.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.categories`',
                  ],
                  'args' => [],
                  'select' => [],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'filter' => [
          'fields' => [
            [
              'name' => 'idMeal',
              'title' => 'Id Meal',
              'type' => '`$STRING`',
              'short' => 'Unique meal identifier',
            ],
            [
              'name' => 'strMeal',
              'title' => 'Str Meal',
              'type' => '`$STRING`',
              'short' => 'Meal name',
            ],
            [
              'name' => 'strMealThumb',
              'title' => 'Str Meal Thumb',
              'type' => '`$STRING`',
              'short' => 'URL to meal thumbnail image',
            ],
          ],
          'name' => 'filter',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/filter.php',
                  'segments' => [
                    [
                      'lit' => 'filter.php',
                    ],
                  ],
                  'parts' => [
                    'filter.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.meals`',
                  ],
                  'args' => [
                    'query' => [
                      [
                        'name' => 'a',
                        'orig' => 'a',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'Canadian',
                      ],
                      [
                        'name' => 'c',
                        'orig' => 'c',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'Seafood',
                      ],
                      [
                        'name' => 'i',
                        'orig' => 'i',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'chicken_breast',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'a',
                      'c',
                      'i',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'latest' => [
          'fields' => [
            [
              'name' => 'dateModified',
              'title' => 'Date Modified',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'idMeal',
              'title' => 'Id Meal',
              'type' => '`$STRING`',
              'short' => 'Unique meal identifier',
            ],
            [
              'name' => 'strArea',
              'title' => 'Str Area',
              'type' => '`$STRING`',
              'short' => 'Meal area/region',
            ],
            [
              'name' => 'strCategory',
              'title' => 'Str Category',
              'type' => '`$STRING`',
              'short' => 'Meal category',
            ],
            [
              'name' => 'strCreativeCommonsConfirmed',
              'title' => 'Str Creative Commons Confirmed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strDrinkAlternate',
              'title' => 'Str Drink Alternate',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strImageSource',
              'title' => 'Str Image Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient1',
              'title' => 'Str Ingredient1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient10',
              'title' => 'Str Ingredient10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient11',
              'title' => 'Str Ingredient11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient12',
              'title' => 'Str Ingredient12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient13',
              'title' => 'Str Ingredient13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient14',
              'title' => 'Str Ingredient14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient15',
              'title' => 'Str Ingredient15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient16',
              'title' => 'Str Ingredient16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient17',
              'title' => 'Str Ingredient17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient18',
              'title' => 'Str Ingredient18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient19',
              'title' => 'Str Ingredient19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient2',
              'title' => 'Str Ingredient2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient20',
              'title' => 'Str Ingredient20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient3',
              'title' => 'Str Ingredient3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient4',
              'title' => 'Str Ingredient4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient5',
              'title' => 'Str Ingredient5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient6',
              'title' => 'Str Ingredient6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient7',
              'title' => 'Str Ingredient7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient8',
              'title' => 'Str Ingredient8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient9',
              'title' => 'Str Ingredient9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strInstructions',
              'title' => 'Str Instructions',
              'type' => '`$STRING`',
              'short' => 'Cooking instructions',
            ],
            [
              'name' => 'strMeal',
              'title' => 'Str Meal',
              'type' => '`$STRING`',
              'short' => 'Meal name',
            ],
            [
              'name' => 'strMealThumb',
              'title' => 'Str Meal Thumb',
              'type' => '`$STRING`',
              'short' => 'URL to meal thumbnail image',
            ],
            [
              'name' => 'strMeasure1',
              'title' => 'Str Measure1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure10',
              'title' => 'Str Measure10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure11',
              'title' => 'Str Measure11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure12',
              'title' => 'Str Measure12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure13',
              'title' => 'Str Measure13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure14',
              'title' => 'Str Measure14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure15',
              'title' => 'Str Measure15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure16',
              'title' => 'Str Measure16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure17',
              'title' => 'Str Measure17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure18',
              'title' => 'Str Measure18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure19',
              'title' => 'Str Measure19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure2',
              'title' => 'Str Measure2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure20',
              'title' => 'Str Measure20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure3',
              'title' => 'Str Measure3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure4',
              'title' => 'Str Measure4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure5',
              'title' => 'Str Measure5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure6',
              'title' => 'Str Measure6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure7',
              'title' => 'Str Measure7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure8',
              'title' => 'Str Measure8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure9',
              'title' => 'Str Measure9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strSource',
              'title' => 'Str Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strTags',
              'title' => 'Str Tags',
              'type' => '`$STRING`',
              'short' => 'Comma-separated tags',
            ],
            [
              'name' => 'strYoutube',
              'title' => 'Str Youtube',
              'type' => '`$STRING`',
              'short' => 'YouTube video URL',
            ],
          ],
          'name' => 'latest',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/latest.php',
                  'segments' => [
                    [
                      'lit' => 'latest.php',
                    ],
                  ],
                  'parts' => [
                    'latest.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.meals`',
                  ],
                  'args' => [],
                  'select' => [],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'list' => [
          'fields' => [
            [
              'name' => 'strArea',
              'title' => 'Str Area',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strCategory',
              'title' => 'Str Category',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient',
              'title' => 'Str Ingredient',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'list',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/list.php',
                  'segments' => [
                    [
                      'lit' => 'list.php',
                    ],
                  ],
                  'parts' => [
                    'list.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.meals`',
                  ],
                  'args' => [
                    'query' => [
                      [
                        'name' => 'a',
                        'orig' => 'a',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'c',
                        'orig' => 'c',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'i',
                        'orig' => 'i',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'a',
                      'c',
                      'i',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'lookup' => [
          'fields' => [
            [
              'name' => 'dateModified',
              'title' => 'Date Modified',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'idMeal',
              'title' => 'Id Meal',
              'type' => '`$STRING`',
              'short' => 'Unique meal identifier',
            ],
            [
              'name' => 'strArea',
              'title' => 'Str Area',
              'type' => '`$STRING`',
              'short' => 'Meal area/region',
            ],
            [
              'name' => 'strCategory',
              'title' => 'Str Category',
              'type' => '`$STRING`',
              'short' => 'Meal category',
            ],
            [
              'name' => 'strCreativeCommonsConfirmed',
              'title' => 'Str Creative Commons Confirmed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strDrinkAlternate',
              'title' => 'Str Drink Alternate',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strImageSource',
              'title' => 'Str Image Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient1',
              'title' => 'Str Ingredient1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient10',
              'title' => 'Str Ingredient10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient11',
              'title' => 'Str Ingredient11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient12',
              'title' => 'Str Ingredient12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient13',
              'title' => 'Str Ingredient13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient14',
              'title' => 'Str Ingredient14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient15',
              'title' => 'Str Ingredient15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient16',
              'title' => 'Str Ingredient16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient17',
              'title' => 'Str Ingredient17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient18',
              'title' => 'Str Ingredient18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient19',
              'title' => 'Str Ingredient19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient2',
              'title' => 'Str Ingredient2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient20',
              'title' => 'Str Ingredient20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient3',
              'title' => 'Str Ingredient3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient4',
              'title' => 'Str Ingredient4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient5',
              'title' => 'Str Ingredient5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient6',
              'title' => 'Str Ingredient6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient7',
              'title' => 'Str Ingredient7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient8',
              'title' => 'Str Ingredient8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient9',
              'title' => 'Str Ingredient9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strInstructions',
              'title' => 'Str Instructions',
              'type' => '`$STRING`',
              'short' => 'Cooking instructions',
            ],
            [
              'name' => 'strMeal',
              'title' => 'Str Meal',
              'type' => '`$STRING`',
              'short' => 'Meal name',
            ],
            [
              'name' => 'strMealThumb',
              'title' => 'Str Meal Thumb',
              'type' => '`$STRING`',
              'short' => 'URL to meal thumbnail image',
            ],
            [
              'name' => 'strMeasure1',
              'title' => 'Str Measure1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure10',
              'title' => 'Str Measure10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure11',
              'title' => 'Str Measure11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure12',
              'title' => 'Str Measure12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure13',
              'title' => 'Str Measure13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure14',
              'title' => 'Str Measure14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure15',
              'title' => 'Str Measure15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure16',
              'title' => 'Str Measure16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure17',
              'title' => 'Str Measure17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure18',
              'title' => 'Str Measure18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure19',
              'title' => 'Str Measure19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure2',
              'title' => 'Str Measure2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure20',
              'title' => 'Str Measure20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure3',
              'title' => 'Str Measure3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure4',
              'title' => 'Str Measure4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure5',
              'title' => 'Str Measure5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure6',
              'title' => 'Str Measure6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure7',
              'title' => 'Str Measure7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure8',
              'title' => 'Str Measure8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure9',
              'title' => 'Str Measure9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strSource',
              'title' => 'Str Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strTags',
              'title' => 'Str Tags',
              'type' => '`$STRING`',
              'short' => 'Comma-separated tags',
            ],
            [
              'name' => 'strYoutube',
              'title' => 'Str Youtube',
              'type' => '`$STRING`',
              'short' => 'YouTube video URL',
            ],
          ],
          'name' => 'lookup',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/lookup.php',
                  'segments' => [
                    [
                      'lit' => 'lookup.php',
                    ],
                  ],
                  'parts' => [
                    'lookup.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.meals`',
                  ],
                  'args' => [
                    'query' => [
                      [
                        'name' => 'i',
                        'orig' => 'i',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'reqd' => true,
                        'example' => '52772',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'i',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'random' => [
          'fields' => [
            [
              'name' => 'dateModified',
              'title' => 'Date Modified',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'idMeal',
              'title' => 'Id Meal',
              'type' => '`$STRING`',
              'short' => 'Unique meal identifier',
            ],
            [
              'name' => 'strArea',
              'title' => 'Str Area',
              'type' => '`$STRING`',
              'short' => 'Meal area/region',
            ],
            [
              'name' => 'strCategory',
              'title' => 'Str Category',
              'type' => '`$STRING`',
              'short' => 'Meal category',
            ],
            [
              'name' => 'strCreativeCommonsConfirmed',
              'title' => 'Str Creative Commons Confirmed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strDrinkAlternate',
              'title' => 'Str Drink Alternate',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strImageSource',
              'title' => 'Str Image Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient1',
              'title' => 'Str Ingredient1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient10',
              'title' => 'Str Ingredient10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient11',
              'title' => 'Str Ingredient11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient12',
              'title' => 'Str Ingredient12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient13',
              'title' => 'Str Ingredient13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient14',
              'title' => 'Str Ingredient14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient15',
              'title' => 'Str Ingredient15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient16',
              'title' => 'Str Ingredient16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient17',
              'title' => 'Str Ingredient17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient18',
              'title' => 'Str Ingredient18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient19',
              'title' => 'Str Ingredient19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient2',
              'title' => 'Str Ingredient2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient20',
              'title' => 'Str Ingredient20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient3',
              'title' => 'Str Ingredient3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient4',
              'title' => 'Str Ingredient4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient5',
              'title' => 'Str Ingredient5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient6',
              'title' => 'Str Ingredient6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient7',
              'title' => 'Str Ingredient7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient8',
              'title' => 'Str Ingredient8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient9',
              'title' => 'Str Ingredient9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strInstructions',
              'title' => 'Str Instructions',
              'type' => '`$STRING`',
              'short' => 'Cooking instructions',
            ],
            [
              'name' => 'strMeal',
              'title' => 'Str Meal',
              'type' => '`$STRING`',
              'short' => 'Meal name',
            ],
            [
              'name' => 'strMealThumb',
              'title' => 'Str Meal Thumb',
              'type' => '`$STRING`',
              'short' => 'URL to meal thumbnail image',
            ],
            [
              'name' => 'strMeasure1',
              'title' => 'Str Measure1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure10',
              'title' => 'Str Measure10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure11',
              'title' => 'Str Measure11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure12',
              'title' => 'Str Measure12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure13',
              'title' => 'Str Measure13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure14',
              'title' => 'Str Measure14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure15',
              'title' => 'Str Measure15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure16',
              'title' => 'Str Measure16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure17',
              'title' => 'Str Measure17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure18',
              'title' => 'Str Measure18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure19',
              'title' => 'Str Measure19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure2',
              'title' => 'Str Measure2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure20',
              'title' => 'Str Measure20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure3',
              'title' => 'Str Measure3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure4',
              'title' => 'Str Measure4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure5',
              'title' => 'Str Measure5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure6',
              'title' => 'Str Measure6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure7',
              'title' => 'Str Measure7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure8',
              'title' => 'Str Measure8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure9',
              'title' => 'Str Measure9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strSource',
              'title' => 'Str Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strTags',
              'title' => 'Str Tags',
              'type' => '`$STRING`',
              'short' => 'Comma-separated tags',
            ],
            [
              'name' => 'strYoutube',
              'title' => 'Str Youtube',
              'type' => '`$STRING`',
              'short' => 'YouTube video URL',
            ],
          ],
          'name' => 'random',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/random.php',
                  'segments' => [
                    [
                      'lit' => 'random.php',
                    ],
                  ],
                  'parts' => [
                    'random.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.meals`',
                  ],
                  'args' => [],
                  'select' => [],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'randomselection' => [
          'fields' => [
            [
              'name' => 'dateModified',
              'title' => 'Date Modified',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'idMeal',
              'title' => 'Id Meal',
              'type' => '`$STRING`',
              'short' => 'Unique meal identifier',
            ],
            [
              'name' => 'strArea',
              'title' => 'Str Area',
              'type' => '`$STRING`',
              'short' => 'Meal area/region',
            ],
            [
              'name' => 'strCategory',
              'title' => 'Str Category',
              'type' => '`$STRING`',
              'short' => 'Meal category',
            ],
            [
              'name' => 'strCreativeCommonsConfirmed',
              'title' => 'Str Creative Commons Confirmed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strDrinkAlternate',
              'title' => 'Str Drink Alternate',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strImageSource',
              'title' => 'Str Image Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient1',
              'title' => 'Str Ingredient1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient10',
              'title' => 'Str Ingredient10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient11',
              'title' => 'Str Ingredient11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient12',
              'title' => 'Str Ingredient12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient13',
              'title' => 'Str Ingredient13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient14',
              'title' => 'Str Ingredient14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient15',
              'title' => 'Str Ingredient15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient16',
              'title' => 'Str Ingredient16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient17',
              'title' => 'Str Ingredient17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient18',
              'title' => 'Str Ingredient18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient19',
              'title' => 'Str Ingredient19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient2',
              'title' => 'Str Ingredient2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient20',
              'title' => 'Str Ingredient20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient3',
              'title' => 'Str Ingredient3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient4',
              'title' => 'Str Ingredient4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient5',
              'title' => 'Str Ingredient5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient6',
              'title' => 'Str Ingredient6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient7',
              'title' => 'Str Ingredient7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient8',
              'title' => 'Str Ingredient8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient9',
              'title' => 'Str Ingredient9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strInstructions',
              'title' => 'Str Instructions',
              'type' => '`$STRING`',
              'short' => 'Cooking instructions',
            ],
            [
              'name' => 'strMeal',
              'title' => 'Str Meal',
              'type' => '`$STRING`',
              'short' => 'Meal name',
            ],
            [
              'name' => 'strMealThumb',
              'title' => 'Str Meal Thumb',
              'type' => '`$STRING`',
              'short' => 'URL to meal thumbnail image',
            ],
            [
              'name' => 'strMeasure1',
              'title' => 'Str Measure1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure10',
              'title' => 'Str Measure10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure11',
              'title' => 'Str Measure11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure12',
              'title' => 'Str Measure12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure13',
              'title' => 'Str Measure13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure14',
              'title' => 'Str Measure14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure15',
              'title' => 'Str Measure15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure16',
              'title' => 'Str Measure16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure17',
              'title' => 'Str Measure17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure18',
              'title' => 'Str Measure18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure19',
              'title' => 'Str Measure19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure2',
              'title' => 'Str Measure2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure20',
              'title' => 'Str Measure20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure3',
              'title' => 'Str Measure3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure4',
              'title' => 'Str Measure4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure5',
              'title' => 'Str Measure5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure6',
              'title' => 'Str Measure6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure7',
              'title' => 'Str Measure7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure8',
              'title' => 'Str Measure8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure9',
              'title' => 'Str Measure9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strSource',
              'title' => 'Str Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strTags',
              'title' => 'Str Tags',
              'type' => '`$STRING`',
              'short' => 'Comma-separated tags',
            ],
            [
              'name' => 'strYoutube',
              'title' => 'Str Youtube',
              'type' => '`$STRING`',
              'short' => 'YouTube video URL',
            ],
          ],
          'name' => 'randomselection',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/randomselection.php',
                  'segments' => [
                    [
                      'lit' => 'randomselection.php',
                    ],
                  ],
                  'parts' => [
                    'randomselection.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.meals`',
                  ],
                  'args' => [],
                  'select' => [],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'search' => [
          'fields' => [
            [
              'name' => 'dateModified',
              'title' => 'Date Modified',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'idMeal',
              'title' => 'Id Meal',
              'type' => '`$STRING`',
              'short' => 'Unique meal identifier',
            ],
            [
              'name' => 'strArea',
              'title' => 'Str Area',
              'type' => '`$STRING`',
              'short' => 'Meal area/region',
            ],
            [
              'name' => 'strCategory',
              'title' => 'Str Category',
              'type' => '`$STRING`',
              'short' => 'Meal category',
            ],
            [
              'name' => 'strCreativeCommonsConfirmed',
              'title' => 'Str Creative Commons Confirmed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strDrinkAlternate',
              'title' => 'Str Drink Alternate',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strImageSource',
              'title' => 'Str Image Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient1',
              'title' => 'Str Ingredient1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient10',
              'title' => 'Str Ingredient10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient11',
              'title' => 'Str Ingredient11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient12',
              'title' => 'Str Ingredient12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient13',
              'title' => 'Str Ingredient13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient14',
              'title' => 'Str Ingredient14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient15',
              'title' => 'Str Ingredient15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient16',
              'title' => 'Str Ingredient16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient17',
              'title' => 'Str Ingredient17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient18',
              'title' => 'Str Ingredient18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient19',
              'title' => 'Str Ingredient19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient2',
              'title' => 'Str Ingredient2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient20',
              'title' => 'Str Ingredient20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient3',
              'title' => 'Str Ingredient3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient4',
              'title' => 'Str Ingredient4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient5',
              'title' => 'Str Ingredient5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient6',
              'title' => 'Str Ingredient6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient7',
              'title' => 'Str Ingredient7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient8',
              'title' => 'Str Ingredient8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strIngredient9',
              'title' => 'Str Ingredient9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strInstructions',
              'title' => 'Str Instructions',
              'type' => '`$STRING`',
              'short' => 'Cooking instructions',
            ],
            [
              'name' => 'strMeal',
              'title' => 'Str Meal',
              'type' => '`$STRING`',
              'short' => 'Meal name',
            ],
            [
              'name' => 'strMealThumb',
              'title' => 'Str Meal Thumb',
              'type' => '`$STRING`',
              'short' => 'URL to meal thumbnail image',
            ],
            [
              'name' => 'strMeasure1',
              'title' => 'Str Measure1',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure10',
              'title' => 'Str Measure10',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure11',
              'title' => 'Str Measure11',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure12',
              'title' => 'Str Measure12',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure13',
              'title' => 'Str Measure13',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure14',
              'title' => 'Str Measure14',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure15',
              'title' => 'Str Measure15',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure16',
              'title' => 'Str Measure16',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure17',
              'title' => 'Str Measure17',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure18',
              'title' => 'Str Measure18',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure19',
              'title' => 'Str Measure19',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure2',
              'title' => 'Str Measure2',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure20',
              'title' => 'Str Measure20',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure3',
              'title' => 'Str Measure3',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure4',
              'title' => 'Str Measure4',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure5',
              'title' => 'Str Measure5',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure6',
              'title' => 'Str Measure6',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure7',
              'title' => 'Str Measure7',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure8',
              'title' => 'Str Measure8',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strMeasure9',
              'title' => 'Str Measure9',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strSource',
              'title' => 'Str Source',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'strTags',
              'title' => 'Str Tags',
              'type' => '`$STRING`',
              'short' => 'Comma-separated tags',
            ],
            [
              'name' => 'strYoutube',
              'title' => 'Str Youtube',
              'type' => '`$STRING`',
              'short' => 'YouTube video URL',
            ],
          ],
          'name' => 'search',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/search.php',
                  'segments' => [
                    [
                      'lit' => 'search.php',
                    ],
                  ],
                  'parts' => [
                    'search.php',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.meals`',
                  ],
                  'args' => [
                    'query' => [
                      [
                        'name' => 'f',
                        'orig' => 'f',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'a',
                      ],
                      [
                        'name' => 's',
                        'orig' => 's',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'Arrabiata',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'f',
                      's',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return FreeMealFeatures::make_feature($name);
    }
}
