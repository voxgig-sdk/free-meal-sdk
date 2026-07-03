package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/free-meal-sdk/go"
	"github.com/voxgig-sdk/free-meal-sdk/go/core"

	vs "github.com/voxgig-sdk/free-meal-sdk/go/utility/struct"
)

func TestRandomselectionEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Randomselection(nil)
		if ent == nil {
			t.Fatal("expected non-nil RandomselectionEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := randomselectionBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "randomselection." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set FREEMEAL_TEST_RANDOMSELECTION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		randomselectionRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.randomselection", setup.data)))
		var randomselectionRef01Data map[string]any
		if len(randomselectionRef01DataRaw) > 0 {
			randomselectionRef01Data = core.ToMapAny(randomselectionRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = randomselectionRef01Data

		// LIST
		randomselectionRef01Ent := client.Randomselection(nil)
		randomselectionRef01Match := map[string]any{}

		randomselectionRef01ListResult, err := randomselectionRef01Ent.List(randomselectionRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, randomselectionRef01ListOk := randomselectionRef01ListResult.([]any)
		if !randomselectionRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", randomselectionRef01ListResult)
		}

	})
}

func randomselectionBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "randomselection", "RandomselectionTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read randomselection test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse randomselection test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"randomselection01", "randomselection02", "randomselection03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("FREEMEAL_TEST_RANDOMSELECTION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"FREEMEAL_TEST_RANDOMSELECTION_ENTID": idmap,
		"FREEMEAL_TEST_LIVE":      "FALSE",
		"FREEMEAL_TEST_EXPLAIN":   "FALSE",
		"FREEMEAL_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["FREEMEAL_TEST_RANDOMSELECTION_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["FREEMEAL_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["FREEMEAL_APIKEY"],
			},
			extra,
		})
		client = sdk.NewFreeMealSDK(core.ToMapAny(mergedOpts))
	}

	live := env["FREEMEAL_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["FREEMEAL_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
