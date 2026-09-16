"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('SearchEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when FREE_MEAL_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('FREE_MEAL_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.FreeMealSDK.test();
        const ent = testsdk.Search();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.FREE_MEAL_TEST_LIVE;
        for (const op of ['list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'search.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "dateModified", "req": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "idMeal", "req": false, "short": "Unique meal identifier", "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "strArea", "req": false, "short": "Meal area/region", "type": "`$STRING`", "index$": 2 }, { "active": true, "name": "strCategory", "req": false, "short": "Meal category", "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "strCreativeCommonsConfirmed", "req": false, "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "strDrinkAlternate", "req": false, "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "strImageSource", "req": false, "type": "`$STRING`", "index$": 6 }, { "active": true, "name": "strIngredient1", "req": false, "type": "`$STRING`", "index$": 7 }, { "active": true, "name": "strIngredient10", "req": false, "type": "`$STRING`", "index$": 8 }, { "active": true, "name": "strIngredient11", "req": false, "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "strIngredient12", "req": false, "type": "`$STRING`", "index$": 10 }, { "active": true, "name": "strIngredient13", "req": false, "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "strIngredient14", "req": false, "type": "`$STRING`", "index$": 12 }, { "active": true, "name": "strIngredient15", "req": false, "type": "`$STRING`", "index$": 13 }, { "active": true, "name": "strIngredient16", "req": false, "type": "`$STRING`", "index$": 14 }, { "active": true, "name": "strIngredient17", "req": false, "type": "`$STRING`", "index$": 15 }, { "active": true, "name": "strIngredient18", "req": false, "type": "`$STRING`", "index$": 16 }, { "active": true, "name": "strIngredient19", "req": false, "type": "`$STRING`", "index$": 17 }, { "active": true, "name": "strIngredient2", "req": false, "type": "`$STRING`", "index$": 18 }, { "active": true, "name": "strIngredient20", "req": false, "type": "`$STRING`", "index$": 19 }, { "active": true, "name": "strIngredient3", "req": false, "type": "`$STRING`", "index$": 20 }, { "active": true, "name": "strIngredient4", "req": false, "type": "`$STRING`", "index$": 21 }, { "active": true, "name": "strIngredient5", "req": false, "type": "`$STRING`", "index$": 22 }, { "active": true, "name": "strIngredient6", "req": false, "type": "`$STRING`", "index$": 23 }, { "active": true, "name": "strIngredient7", "req": false, "type": "`$STRING`", "index$": 24 }, { "active": true, "name": "strIngredient8", "req": false, "type": "`$STRING`", "index$": 25 }, { "active": true, "name": "strIngredient9", "req": false, "type": "`$STRING`", "index$": 26 }, { "active": true, "name": "strInstructions", "req": false, "short": "Cooking instructions", "type": "`$STRING`", "index$": 27 }, { "active": true, "name": "strMeal", "req": false, "short": "Meal name", "type": "`$STRING`", "index$": 28 }, { "active": true, "name": "strMealThumb", "req": false, "short": "URL to meal thumbnail image", "type": "`$STRING`", "index$": 29 }, { "active": true, "name": "strMeasure1", "req": false, "type": "`$STRING`", "index$": 30 }, { "active": true, "name": "strMeasure10", "req": false, "type": "`$STRING`", "index$": 31 }, { "active": true, "name": "strMeasure11", "req": false, "type": "`$STRING`", "index$": 32 }, { "active": true, "name": "strMeasure12", "req": false, "type": "`$STRING`", "index$": 33 }, { "active": true, "name": "strMeasure13", "req": false, "type": "`$STRING`", "index$": 34 }, { "active": true, "name": "strMeasure14", "req": false, "type": "`$STRING`", "index$": 35 }, { "active": true, "name": "strMeasure15", "req": false, "type": "`$STRING`", "index$": 36 }, { "active": true, "name": "strMeasure16", "req": false, "type": "`$STRING`", "index$": 37 }, { "active": true, "name": "strMeasure17", "req": false, "type": "`$STRING`", "index$": 38 }, { "active": true, "name": "strMeasure18", "req": false, "type": "`$STRING`", "index$": 39 }, { "active": true, "name": "strMeasure19", "req": false, "type": "`$STRING`", "index$": 40 }, { "active": true, "name": "strMeasure2", "req": false, "type": "`$STRING`", "index$": 41 }, { "active": true, "name": "strMeasure20", "req": false, "type": "`$STRING`", "index$": 42 }, { "active": true, "name": "strMeasure3", "req": false, "type": "`$STRING`", "index$": 43 }, { "active": true, "name": "strMeasure4", "req": false, "type": "`$STRING`", "index$": 44 }, { "active": true, "name": "strMeasure5", "req": false, "type": "`$STRING`", "index$": 45 }, { "active": true, "name": "strMeasure6", "req": false, "type": "`$STRING`", "index$": 46 }, { "active": true, "name": "strMeasure7", "req": false, "type": "`$STRING`", "index$": 47 }, { "active": true, "name": "strMeasure8", "req": false, "type": "`$STRING`", "index$": 48 }, { "active": true, "name": "strMeasure9", "req": false, "type": "`$STRING`", "index$": 49 }, { "active": true, "name": "strSource", "req": false, "type": "`$STRING`", "index$": 50 }, { "active": true, "name": "strTags", "req": false, "short": "Comma-separated tags", "type": "`$STRING`", "index$": 51 }, { "active": true, "name": "strYoutube", "req": false, "short": "YouTube video URL", "type": "`$STRING`", "index$": 52 }], "name": "search", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": { "query": [{ "active": true, "example": "a", "kind": "query", "name": "f", "orig": "f", "reqd": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "example": "Arrabiata", "kind": "query", "name": "s", "orig": "s", "reqd": false, "type": "`$STRING`", "index$": 1 }] }, "contract": { "id": "GET /search.php", "json": "{\"operationId\":\"searchMeals\",\"parameters\":[{\"description\":\"Search meals by name\",\"in\":\"query\",\"name\":\"s\",\"required\":false,\"schema\":{\"example\":\"Arrabiata\",\"type\":\"string\"}},{\"description\":\"List all meals by first letter\",\"in\":\"query\",\"name\":\"f\",\"required\":false,\"schema\":{\"example\":\"a\",\"maxLength\":1,\"minLength\":1,\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"meals\":{\"items\":{\"properties\":{\"dateModified\":{\"nullable\":true,\"type\":\"string\"},\"idMeal\":{\"description\":\"Unique meal identifier\",\"type\":\"string\"},\"strArea\":{\"description\":\"Meal area/region\",\"type\":\"string\"},\"strCategory\":{\"description\":\"Meal category\",\"type\":\"string\"},\"strCreativeCommonsConfirmed\":{\"nullable\":true,\"type\":\"string\"},\"strDrinkAlternate\":{\"nullable\":true,\"type\":\"string\"},\"strImageSource\":{\"nullable\":true,\"type\":\"string\"},\"strIngredient1\":{\"type\":\"string\"},\"strIngredient10\":{\"type\":\"string\"},\"strIngredient11\":{\"type\":\"string\"},\"strIngredient12\":{\"type\":\"string\"},\"strIngredient13\":{\"type\":\"string\"},\"strIngredient14\":{\"type\":\"string\"},\"strIngredient15\":{\"type\":\"string\"},\"strIngredient16\":{\"type\":\"string\"},\"strIngredient17\":{\"type\":\"string\"},\"strIngredient18\":{\"type\":\"string\"},\"strIngredient19\":{\"type\":\"string\"},\"strIngredient2\":{\"type\":\"string\"},\"strIngredient20\":{\"type\":\"string\"},\"strIngredient3\":{\"type\":\"string\"},\"strIngredient4\":{\"type\":\"string\"},\"strIngredient5\":{\"type\":\"string\"},\"strIngredient6\":{\"type\":\"string\"},\"strIngredient7\":{\"type\":\"string\"},\"strIngredient8\":{\"type\":\"string\"},\"strIngredient9\":{\"type\":\"string\"},\"strInstructions\":{\"description\":\"Cooking instructions\",\"type\":\"string\"},\"strMeal\":{\"description\":\"Meal name\",\"type\":\"string\"},\"strMealThumb\":{\"description\":\"URL to meal thumbnail image\",\"type\":\"string\"},\"strMeasure1\":{\"type\":\"string\"},\"strMeasure10\":{\"type\":\"string\"},\"strMeasure11\":{\"type\":\"string\"},\"strMeasure12\":{\"type\":\"string\"},\"strMeasure13\":{\"type\":\"string\"},\"strMeasure14\":{\"type\":\"string\"},\"strMeasure15\":{\"type\":\"string\"},\"strMeasure16\":{\"type\":\"string\"},\"strMeasure17\":{\"type\":\"string\"},\"strMeasure18\":{\"type\":\"string\"},\"strMeasure19\":{\"type\":\"string\"},\"strMeasure2\":{\"type\":\"string\"},\"strMeasure20\":{\"type\":\"string\"},\"strMeasure3\":{\"type\":\"string\"},\"strMeasure4\":{\"type\":\"string\"},\"strMeasure5\":{\"type\":\"string\"},\"strMeasure6\":{\"type\":\"string\"},\"strMeasure7\":{\"type\":\"string\"},\"strMeasure8\":{\"type\":\"string\"},\"strMeasure9\":{\"type\":\"string\"},\"strSource\":{\"nullable\":true,\"type\":\"string\"},\"strTags\":{\"description\":\"Comma-separated tags\",\"nullable\":true,\"type\":\"string\"},\"strYoutube\":{\"description\":\"YouTube video URL\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"Successful response\"}},\"securitySchemes\":{\"ApiKeyAuth\":{\"description\":\"API key embedded in the URL path. Use '1' for testing, or get a premium key for production.\",\"in\":\"path\",\"name\":\"api_key\",\"type\":\"apiKey\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/search.php", "segments": [{ "lit": "search.php" }], "select": { "exist": ["f", "s"] }, "transform": { "req": "`reqdata`", "res": "`body.meals`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [] }, "key$": "search", "name__orig": "search", "Name": "Search", "name_": "search", "name-": "search", "NAME": "SEARCH", "index$": 7 }, { "active": true, "entity": "search", "key$": "BasicSearchFlow", "kind": "basic", "name": "BasicSearchFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "search_ref01" } }], "index$": 0 }] }, 'Search');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let search_ref01_data = Object.values(setup.data.existing.search)[0];
        // LIST
        const search_ref01_ent = client.Search();
        const search_ref01_match = {};
        const search_ref01_list = (await search_ref01_ent.list(search_ref01_match)).map((e) => e.data());
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/search/SearchTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.FreeMealSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['search01', 'search02', 'search03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'FREE_MEAL_TEST_SEARCH_ENTID': idmap,
        'FREE_MEAL_TEST_LIVE': 'FALSE',
        'FREE_MEAL_TEST_EXPLAIN': 'FALSE',
        'FREE_MEAL_APIKEY': '',
    });
    idmap = env['FREE_MEAL_TEST_SEARCH_ENTID'];
    const live = 'TRUE' === env.FREE_MEAL_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['FREE_MEAL_TEST_SEARCH_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.FreeMealSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
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
        ]));
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
    };
    return setup;
}
//# sourceMappingURL=SearchEntity.test.js.map