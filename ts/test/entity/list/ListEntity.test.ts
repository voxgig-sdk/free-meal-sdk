

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


describe('ListEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when FREE_MEAL_TEST_LIVE=TRUE.
  afterEach(liveDelay('FREE_MEAL_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = FreeMealSDK.test()
    const ent = testsdk.List()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.FREE_MEAL_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'list.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"strArea","req":false,"type":"`$STRING`","index$":0},{"active":true,"name":"strCategory","req":false,"type":"`$STRING`","index$":1},{"active":true,"name":"strIngredient","req":false,"type":"`$STRING`","index$":2}],"name":"list","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"query":[{"active":true,"kind":"query","name":"a","orig":"a","reqd":false,"type":"`$STRING`","index$":0},{"active":true,"kind":"query","name":"c","orig":"c","reqd":false,"type":"`$STRING`","index$":1},{"active":true,"kind":"query","name":"i","orig":"i","reqd":false,"type":"`$STRING`","index$":2}]},"contract":{"id":"GET /list.php","json":"{\"operationId\":\"listAttributes\",\"parameters\":[{\"description\":\"List all categories\",\"in\":\"query\",\"name\":\"c\",\"required\":false,\"schema\":{\"enum\":[\"list\"],\"type\":\"string\"}},{\"description\":\"List all areas\",\"in\":\"query\",\"name\":\"a\",\"required\":false,\"schema\":{\"enum\":[\"list\"],\"type\":\"string\"}},{\"description\":\"List all ingredients\",\"in\":\"query\",\"name\":\"i\",\"required\":false,\"schema\":{\"enum\":[\"list\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"meals\":{\"items\":{\"properties\":{\"strArea\":{\"type\":\"string\"},\"strCategory\":{\"type\":\"string\"},\"strIngredient\":{\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"Successful response\"}},\"securitySchemes\":{\"ApiKeyAuth\":{\"description\":\"API key embedded in the URL path. Use '1' for testing, or get a premium key for production.\",\"in\":\"path\",\"name\":\"api_key\",\"type\":\"apiKey\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/list.php","segments":[{"lit":"list.php"}],"select":{"exist":["a","c","i"]},"transform":{"req":"`reqdata`","res":"`body.meals`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"list","name__orig":"list","Name":"List","name_":"list","name-":"list","NAME":"LIST","index$":3}, {"active":true,"entity":"list","key$":"BasicListFlow","kind":"basic","name":"BasicListFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"list_ref01"}}],"index$":0}]}, 'List')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let list_ref01_data = Object.values(setup.data.existing.list)[0] as any

    // LIST
    const list_ref01_ent = client.List()
    const list_ref01_match: any = {}

    const list_ref01_list = (await list_ref01_ent.list(list_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/list/ListTestData.json')

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
    ['list01','list02','list03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'FREE_MEAL_TEST_LIST_ENTID': idmap,
    'FREE_MEAL_TEST_LIVE': 'FALSE',
    'FREE_MEAL_TEST_EXPLAIN': 'FALSE',
    'FREE_MEAL_APIKEY': '',
  })

  idmap = env['FREE_MEAL_TEST_LIST_ENTID']

  const live = 'TRUE' === env.FREE_MEAL_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['FREE_MEAL_TEST_LIST_ENTID']
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
  
