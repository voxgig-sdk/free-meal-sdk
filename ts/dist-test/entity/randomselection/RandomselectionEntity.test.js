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
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('RandomselectionEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when FREE_MEAL_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('FREE_MEAL_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.FreeMealSDK.test();
        const ent = testsdk.Randomselection();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.FREE_MEAL_TEST_LIVE;
        for (const op of ['list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'randomselection.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "dateModified": { "a": true, "h": "Date Modified", "n": "dateModified", "r": false, "t": "`$STRING`", "key$": "dateModified", "index$": 0 }, "idMeal": { "a": true, "h": "Id Meal", "n": "idMeal", "r": false, "sh": "Unique meal identifier", "t": "`$STRING`", "key$": "idMeal", "index$": 1 }, "strArea": { "a": true, "h": "Str Area", "n": "strArea", "r": false, "sh": "Meal area/region", "t": "`$STRING`", "key$": "strArea", "index$": 2 }, "strCategory": { "a": true, "h": "Str Category", "n": "strCategory", "r": false, "sh": "Meal category", "t": "`$STRING`", "key$": "strCategory", "index$": 3 }, "strCreativeCommonsConfirmed": { "a": true, "h": "Str Creative Commons Confirmed", "n": "strCreativeCommonsConfirmed", "r": false, "t": "`$STRING`", "key$": "strCreativeCommonsConfirmed", "index$": 4 }, "strDrinkAlternate": { "a": true, "h": "Str Drink Alternate", "n": "strDrinkAlternate", "r": false, "t": "`$STRING`", "key$": "strDrinkAlternate", "index$": 5 }, "strImageSource": { "a": true, "h": "Str Image Source", "n": "strImageSource", "r": false, "t": "`$STRING`", "key$": "strImageSource", "index$": 6 }, "strIngredient1": { "a": true, "h": "Str Ingredient1", "n": "strIngredient1", "r": false, "t": "`$STRING`", "key$": "strIngredient1", "index$": 7 }, "strIngredient10": { "a": true, "h": "Str Ingredient10", "n": "strIngredient10", "r": false, "t": "`$STRING`", "key$": "strIngredient10", "index$": 8 }, "strIngredient11": { "a": true, "h": "Str Ingredient11", "n": "strIngredient11", "r": false, "t": "`$STRING`", "key$": "strIngredient11", "index$": 9 }, "strIngredient12": { "a": true, "h": "Str Ingredient12", "n": "strIngredient12", "r": false, "t": "`$STRING`", "key$": "strIngredient12", "index$": 10 }, "strIngredient13": { "a": true, "h": "Str Ingredient13", "n": "strIngredient13", "r": false, "t": "`$STRING`", "key$": "strIngredient13", "index$": 11 }, "strIngredient14": { "a": true, "h": "Str Ingredient14", "n": "strIngredient14", "r": false, "t": "`$STRING`", "key$": "strIngredient14", "index$": 12 }, "strIngredient15": { "a": true, "h": "Str Ingredient15", "n": "strIngredient15", "r": false, "t": "`$STRING`", "key$": "strIngredient15", "index$": 13 }, "strIngredient16": { "a": true, "h": "Str Ingredient16", "n": "strIngredient16", "r": false, "t": "`$STRING`", "key$": "strIngredient16", "index$": 14 }, "strIngredient17": { "a": true, "h": "Str Ingredient17", "n": "strIngredient17", "r": false, "t": "`$STRING`", "key$": "strIngredient17", "index$": 15 }, "strIngredient18": { "a": true, "h": "Str Ingredient18", "n": "strIngredient18", "r": false, "t": "`$STRING`", "key$": "strIngredient18", "index$": 16 }, "strIngredient19": { "a": true, "h": "Str Ingredient19", "n": "strIngredient19", "r": false, "t": "`$STRING`", "key$": "strIngredient19", "index$": 17 }, "strIngredient2": { "a": true, "h": "Str Ingredient2", "n": "strIngredient2", "r": false, "t": "`$STRING`", "key$": "strIngredient2", "index$": 18 }, "strIngredient20": { "a": true, "h": "Str Ingredient20", "n": "strIngredient20", "r": false, "t": "`$STRING`", "key$": "strIngredient20", "index$": 19 }, "strIngredient3": { "a": true, "h": "Str Ingredient3", "n": "strIngredient3", "r": false, "t": "`$STRING`", "key$": "strIngredient3", "index$": 20 }, "strIngredient4": { "a": true, "h": "Str Ingredient4", "n": "strIngredient4", "r": false, "t": "`$STRING`", "key$": "strIngredient4", "index$": 21 }, "strIngredient5": { "a": true, "h": "Str Ingredient5", "n": "strIngredient5", "r": false, "t": "`$STRING`", "key$": "strIngredient5", "index$": 22 }, "strIngredient6": { "a": true, "h": "Str Ingredient6", "n": "strIngredient6", "r": false, "t": "`$STRING`", "key$": "strIngredient6", "index$": 23 }, "strIngredient7": { "a": true, "h": "Str Ingredient7", "n": "strIngredient7", "r": false, "t": "`$STRING`", "key$": "strIngredient7", "index$": 24 }, "strIngredient8": { "a": true, "h": "Str Ingredient8", "n": "strIngredient8", "r": false, "t": "`$STRING`", "key$": "strIngredient8", "index$": 25 }, "strIngredient9": { "a": true, "h": "Str Ingredient9", "n": "strIngredient9", "r": false, "t": "`$STRING`", "key$": "strIngredient9", "index$": 26 }, "strInstructions": { "a": true, "h": "Str Instructions", "n": "strInstructions", "r": false, "sh": "Cooking instructions", "t": "`$STRING`", "key$": "strInstructions", "index$": 27 }, "strMeal": { "a": true, "h": "Str Meal", "n": "strMeal", "r": false, "sh": "Meal name", "t": "`$STRING`", "key$": "strMeal", "index$": 28 }, "strMealThumb": { "a": true, "h": "Str Meal Thumb", "n": "strMealThumb", "r": false, "sh": "URL to meal thumbnail image", "t": "`$STRING`", "key$": "strMealThumb", "index$": 29 }, "strMeasure1": { "a": true, "h": "Str Measure1", "n": "strMeasure1", "r": false, "t": "`$STRING`", "key$": "strMeasure1", "index$": 30 }, "strMeasure10": { "a": true, "h": "Str Measure10", "n": "strMeasure10", "r": false, "t": "`$STRING`", "key$": "strMeasure10", "index$": 31 }, "strMeasure11": { "a": true, "h": "Str Measure11", "n": "strMeasure11", "r": false, "t": "`$STRING`", "key$": "strMeasure11", "index$": 32 }, "strMeasure12": { "a": true, "h": "Str Measure12", "n": "strMeasure12", "r": false, "t": "`$STRING`", "key$": "strMeasure12", "index$": 33 }, "strMeasure13": { "a": true, "h": "Str Measure13", "n": "strMeasure13", "r": false, "t": "`$STRING`", "key$": "strMeasure13", "index$": 34 }, "strMeasure14": { "a": true, "h": "Str Measure14", "n": "strMeasure14", "r": false, "t": "`$STRING`", "key$": "strMeasure14", "index$": 35 }, "strMeasure15": { "a": true, "h": "Str Measure15", "n": "strMeasure15", "r": false, "t": "`$STRING`", "key$": "strMeasure15", "index$": 36 }, "strMeasure16": { "a": true, "h": "Str Measure16", "n": "strMeasure16", "r": false, "t": "`$STRING`", "key$": "strMeasure16", "index$": 37 }, "strMeasure17": { "a": true, "h": "Str Measure17", "n": "strMeasure17", "r": false, "t": "`$STRING`", "key$": "strMeasure17", "index$": 38 }, "strMeasure18": { "a": true, "h": "Str Measure18", "n": "strMeasure18", "r": false, "t": "`$STRING`", "key$": "strMeasure18", "index$": 39 }, "strMeasure19": { "a": true, "h": "Str Measure19", "n": "strMeasure19", "r": false, "t": "`$STRING`", "key$": "strMeasure19", "index$": 40 }, "strMeasure2": { "a": true, "h": "Str Measure2", "n": "strMeasure2", "r": false, "t": "`$STRING`", "key$": "strMeasure2", "index$": 41 }, "strMeasure20": { "a": true, "h": "Str Measure20", "n": "strMeasure20", "r": false, "t": "`$STRING`", "key$": "strMeasure20", "index$": 42 }, "strMeasure3": { "a": true, "h": "Str Measure3", "n": "strMeasure3", "r": false, "t": "`$STRING`", "key$": "strMeasure3", "index$": 43 }, "strMeasure4": { "a": true, "h": "Str Measure4", "n": "strMeasure4", "r": false, "t": "`$STRING`", "key$": "strMeasure4", "index$": 44 }, "strMeasure5": { "a": true, "h": "Str Measure5", "n": "strMeasure5", "r": false, "t": "`$STRING`", "key$": "strMeasure5", "index$": 45 }, "strMeasure6": { "a": true, "h": "Str Measure6", "n": "strMeasure6", "r": false, "t": "`$STRING`", "key$": "strMeasure6", "index$": 46 }, "strMeasure7": { "a": true, "h": "Str Measure7", "n": "strMeasure7", "r": false, "t": "`$STRING`", "key$": "strMeasure7", "index$": 47 }, "strMeasure8": { "a": true, "h": "Str Measure8", "n": "strMeasure8", "r": false, "t": "`$STRING`", "key$": "strMeasure8", "index$": 48 }, "strMeasure9": { "a": true, "h": "Str Measure9", "n": "strMeasure9", "r": false, "t": "`$STRING`", "key$": "strMeasure9", "index$": 49 }, "strSource": { "a": true, "h": "Str Source", "n": "strSource", "r": false, "t": "`$STRING`", "key$": "strSource", "index$": 50 }, "strTags": { "a": true, "h": "Str Tags", "n": "strTags", "r": false, "sh": "Comma-separated tags", "t": "`$STRING`", "key$": "strTags", "index$": 51 }, "strYoutube": { "a": true, "h": "Str Youtube", "n": "strYoutube", "r": false, "sh": "YouTube video URL", "t": "`$STRING`", "key$": "strYoutube", "index$": 52 } }, "name": "randomselection", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /randomselection.php", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "GET", "o": "/randomselection.php", "q": {}, "r": {}, "s": [{ "lit": "randomselection.php" }], "t": { "req": "`reqdata`", "res": "`body.meals`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [] }, "key$": "randomselection", "name__orig": "randomselection", "Name": "Randomselection", "name_": "randomselection", "name-": "randomselection", "NAME": "RANDOMSELECTION", "index$": 6 }, { "active": true, "entity": "randomselection", "key$": "BasicRandomselectionFlow", "kind": "basic", "name": "BasicRandomselectionFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "randomselection_ref01" } }], "index$": 0 }] }, 'Randomselection', { "GET /randomselection.php": { "protocol": "http", "operationId": "getRandomSelection", "responses": { "200": { "description": "Successful response", "content": { "application/json": { "schema": { "type": "object", "properties": { "meals": { "items": { "properties": { "dateModified": { "nullable": true, "type": "string", "key$": "dateModified" }, "idMeal": { "description": "Unique meal identifier", "type": "string", "key$": "idMeal" }, "strArea": { "description": "Meal area/region", "type": "string", "key$": "strArea" }, "strCategory": { "description": "Meal category", "type": "string", "key$": "strCategory" }, "strCreativeCommonsConfirmed": { "nullable": true, "type": "string", "key$": "strCreativeCommonsConfirmed" }, "strDrinkAlternate": { "nullable": true, "type": "string", "key$": "strDrinkAlternate" }, "strImageSource": { "nullable": true, "type": "string", "key$": "strImageSource" }, "strIngredient1": { "type": "string", "key$": "strIngredient1" }, "strIngredient10": { "type": "string", "key$": "strIngredient10" }, "strIngredient11": { "type": "string", "key$": "strIngredient11" }, "strIngredient12": { "type": "string", "key$": "strIngredient12" }, "strIngredient13": { "type": "string", "key$": "strIngredient13" }, "strIngredient14": { "type": "string", "key$": "strIngredient14" }, "strIngredient15": { "type": "string", "key$": "strIngredient15" }, "strIngredient16": { "type": "string", "key$": "strIngredient16" }, "strIngredient17": { "type": "string", "key$": "strIngredient17" }, "strIngredient18": { "type": "string", "key$": "strIngredient18" }, "strIngredient19": { "type": "string", "key$": "strIngredient19" }, "strIngredient2": { "type": "string", "key$": "strIngredient2" }, "strIngredient20": { "type": "string", "key$": "strIngredient20" }, "strIngredient3": { "type": "string", "key$": "strIngredient3" }, "strIngredient4": { "type": "string", "key$": "strIngredient4" }, "strIngredient5": { "type": "string", "key$": "strIngredient5" }, "strIngredient6": { "type": "string", "key$": "strIngredient6" }, "strIngredient7": { "type": "string", "key$": "strIngredient7" }, "strIngredient8": { "type": "string", "key$": "strIngredient8" }, "strIngredient9": { "type": "string", "key$": "strIngredient9" }, "strInstructions": { "description": "Cooking instructions", "type": "string", "key$": "strInstructions" }, "strMeal": { "description": "Meal name", "type": "string", "key$": "strMeal" }, "strMealThumb": { "description": "URL to meal thumbnail image", "type": "string", "key$": "strMealThumb" }, "strMeasure1": { "type": "string", "key$": "strMeasure1" }, "strMeasure10": { "type": "string", "key$": "strMeasure10" }, "strMeasure11": { "type": "string", "key$": "strMeasure11" }, "strMeasure12": { "type": "string", "key$": "strMeasure12" }, "strMeasure13": { "type": "string", "key$": "strMeasure13" }, "strMeasure14": { "type": "string", "key$": "strMeasure14" }, "strMeasure15": { "type": "string", "key$": "strMeasure15" }, "strMeasure16": { "type": "string", "key$": "strMeasure16" }, "strMeasure17": { "type": "string", "key$": "strMeasure17" }, "strMeasure18": { "type": "string", "key$": "strMeasure18" }, "strMeasure19": { "type": "string", "key$": "strMeasure19" }, "strMeasure2": { "type": "string", "key$": "strMeasure2" }, "strMeasure20": { "type": "string", "key$": "strMeasure20" }, "strMeasure3": { "type": "string", "key$": "strMeasure3" }, "strMeasure4": { "type": "string", "key$": "strMeasure4" }, "strMeasure5": { "type": "string", "key$": "strMeasure5" }, "strMeasure6": { "type": "string", "key$": "strMeasure6" }, "strMeasure7": { "type": "string", "key$": "strMeasure7" }, "strMeasure8": { "type": "string", "key$": "strMeasure8" }, "strMeasure9": { "type": "string", "key$": "strMeasure9" }, "strSource": { "nullable": true, "type": "string", "key$": "strSource" }, "strTags": { "description": "Comma-separated tags", "nullable": true, "type": "string", "key$": "strTags" }, "strYoutube": { "description": "YouTube video URL", "type": "string", "key$": "strYoutube" } }, "type": "object", "x-ref": "#/components/schemas/Meal", "index$": 0 }, "key$": "meals", "maxItems": 10, "type": "array" } } } } } }, "403": { "description": "Forbidden - Premium API key required" } }, "parameters": [], "security": [{ "ApiKeyAuth": [] }], "securitySource": "operation", "securitySchemes": { "ApiKeyAuth": { "type": "apiKey", "in": "path", "name": "api_key", "description": "API key embedded in the URL path. Use '1' for testing, or get a premium key for production." } } } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let randomselection_ref01_data = Object.values(setup.data.existing.randomselection)[0];
        // LIST
        const randomselection_ref01_ent = client.Randomselection();
        const randomselection_ref01_match = {};
        const randomselection_ref01_list = (await randomselection_ref01_ent.list(randomselection_ref01_match)).map((e) => e.data());
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/randomselection/RandomselectionTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.FreeMealSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['randomselection01', 'randomselection02', 'randomselection03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'FREE_MEAL_TEST_RANDOMSELECTION_ENTID': idmap,
        'FREE_MEAL_TEST_LIVE': 'FALSE',
        'FREE_MEAL_TEST_EXPLAIN': 'FALSE',
        'FREE_MEAL_APIKEY': '',
    });
    idmap = env['FREE_MEAL_TEST_RANDOMSELECTION_ENTID'];
    const live = 'TRUE' === env.FREE_MEAL_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['FREE_MEAL_TEST_RANDOMSELECTION_ENTID'];
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
//# sourceMappingURL=RandomselectionEntity.test.js.map