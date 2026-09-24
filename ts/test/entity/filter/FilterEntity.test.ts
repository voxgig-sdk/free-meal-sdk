

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { FreeMealSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('FilterEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when FREE_MEAL_TEST_LIVE=TRUE.
  afterEach(liveDelay('FREE_MEAL_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = FreeMealSDK.test()
    const ent = testsdk.Filter()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.FREE_MEAL_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'filter.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"idMeal":{"a":true,"h":"Id Meal","n":"idMeal","r":false,"sh":"Unique meal identifier","t":"`$STRING`","key$":"idMeal","index$":0},"strMeal":{"a":true,"h":"Str Meal","n":"strMeal","r":false,"sh":"Meal name","t":"`$STRING`","key$":"strMeal","index$":1},"strMealThumb":{"a":true,"h":"Str Meal Thumb","n":"strMealThumb","r":false,"sh":"URL to meal thumbnail image","t":"`$STRING`","key$":"strMealThumb","index$":2}},"name":"filter","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /filter.php","source":"openapi3","version":2},"g":{"query":[{"a":true,"ex":"Canadian","k":"query","n":"a","or":"a","r":false,"t":"`$STRING`","index$":0},{"a":true,"ex":"Seafood","k":"query","n":"c","or":"c","r":false,"t":"`$STRING`","index$":1},{"a":true,"ex":"chicken_breast","k":"query","n":"i","or":"i","r":false,"t":"`$STRING`","index$":2}]},"k":"http","m":"GET","o":"/filter.php","q":{"exist":["a","c","i"]},"r":{},"s":[{"lit":"filter.php"}],"t":{"req":"`reqdata`","res":"`body.meals`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"filter","name__orig":"filter","Name":"Filter","name_":"filter","name-":"filter","NAME":"FILTER","index$":1}, {"active":true,"entity":"filter","key$":"BasicFilterFlow","kind":"basic","name":"BasicFilterFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"filter_ref01"}}],"index$":0}]}, 'Filter', {"GET /filter.php":{"protocol":"http","operationId":"filterMeals","responses":{"200":{"description":"Successful response","content":{"application/json":{"schema":{"type":"object","properties":{"meals":{"items":{"properties":{"idMeal":{"description":"Unique meal identifier","type":"string","key$":"idMeal"},"strMeal":{"description":"Meal name","type":"string","key$":"strMeal"},"strMealThumb":{"description":"URL to meal thumbnail image","type":"string","key$":"strMealThumb"}},"type":"object","x-ref":"#/components/schemas/MealSummary","index$":0},"key$":"meals","type":"array"}}}}}}},"parameters":[{"name":"i","in":"query","description":"Filter by ingredient (comma-separated for multi-ingredient with Premium API)","required":false,"schema":{"type":"string","example":"chicken_breast"},"index$":0},{"name":"c","in":"query","description":"Filter by category","required":false,"schema":{"type":"string","example":"Seafood"},"index$":1},{"name":"a","in":"query","description":"Filter by area","required":false,"schema":{"type":"string","example":"Canadian"},"index$":2}],"securitySource":"unspecified","securitySchemes":{"ApiKeyAuth":{"type":"apiKey","in":"path","name":"api_key","description":"API key embedded in the URL path. Use '1' for testing, or get a premium key for production."}}}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let filter_ref01_data = Object.values(setup.data.existing.filter)[0] as any

    // LIST
    const filter_ref01_ent = client.Filter()
    const filter_ref01_match: any = {}

    const filter_ref01_list = (await filter_ref01_ent.list(filter_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/filter/FilterTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = FreeMealSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['filter01','filter02','filter03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'FREE_MEAL_TEST_FILTER_ENTID': idmap,
    'FREE_MEAL_TEST_LIVE': 'FALSE',
    'FREE_MEAL_TEST_EXPLAIN': 'FALSE',
    'FREE_MEAL_APIKEY': '',
  })

  idmap = env['FREE_MEAL_TEST_FILTER_ENTID']

  const live = 'TRUE' === env.FREE_MEAL_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['FREE_MEAL_TEST_FILTER_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new FreeMealSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.FREE_MEAL_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.FREE_MEAL_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
