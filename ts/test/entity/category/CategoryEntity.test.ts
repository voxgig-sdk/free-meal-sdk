

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


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('CategoryEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when FREE_MEAL_TEST_LIVE=TRUE.
  afterEach(liveDelay('FREE_MEAL_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = FreeMealSDK.test()
    const ent = testsdk.Category()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.FREE_MEAL_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'category.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"idCategory","req":false,"short":"Unique category identifier","type":"`$STRING`","index$":0},{"active":true,"name":"strCategory","req":false,"short":"Category name","type":"`$STRING`","index$":1},{"active":true,"name":"strCategoryDescription","req":false,"short":"Category description","type":"`$STRING`","index$":2},{"active":true,"name":"strCategoryThumb","req":false,"short":"URL to category thumbnail image","type":"`$STRING`","index$":3}],"name":"category","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /categories.php","json":"{\"operationId\":\"listCategories\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"categories\":{\"items\":{\"properties\":{\"idCategory\":{\"description\":\"Unique category identifier\",\"type\":\"string\"},\"strCategory\":{\"description\":\"Category name\",\"type\":\"string\"},\"strCategoryDescription\":{\"description\":\"Category description\",\"type\":\"string\"},\"strCategoryThumb\":{\"description\":\"URL to category thumbnail image\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"Successful response\"}},\"securitySchemes\":{\"ApiKeyAuth\":{\"description\":\"API key embedded in the URL path. Use '1' for testing, or get a premium key for production.\",\"in\":\"path\",\"name\":\"api_key\",\"type\":\"apiKey\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/categories.php","segments":[{"lit":"categories.php"}],"select":{},"transform":{"req":"`reqdata`","res":"`body.categories`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"category","name__orig":"category","Name":"Category","name_":"category","name-":"category","NAME":"CATEGORY","index$":0}, {"active":true,"entity":"category","key$":"BasicCategoryFlow","kind":"basic","name":"BasicCategoryFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"category_ref01"}}],"index$":0}]}, 'Category')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let category_ref01_data = Object.values(setup.data.existing.category)[0] as any

    // LIST
    const category_ref01_ent = client.Category()
    const category_ref01_match: any = {}

    const category_ref01_list = (await category_ref01_ent.list(category_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/category/CategoryTestData.json')

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
    ['category01','category02','category03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'FREE_MEAL_TEST_CATEGORY_ENTID': idmap,
    'FREE_MEAL_TEST_LIVE': 'FALSE',
    'FREE_MEAL_TEST_EXPLAIN': 'FALSE',
    'FREE_MEAL_APIKEY': '',
  })

  idmap = env['FREE_MEAL_TEST_CATEGORY_ENTID']

  const live = 'TRUE' === env.FREE_MEAL_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['FREE_MEAL_TEST_CATEGORY_ENTID']
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
  
