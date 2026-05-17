package voxgigfreemealsdk

import (
	"github.com/voxgig-sdk/free-meal-sdk/go/core"
	"github.com/voxgig-sdk/free-meal-sdk/go/entity"
	"github.com/voxgig-sdk/free-meal-sdk/go/feature"
	_ "github.com/voxgig-sdk/free-meal-sdk/go/utility"
)

// Type aliases preserve external API.
type FreeMealSDK = core.FreeMealSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type FreeMealEntity = core.FreeMealEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type FreeMealError = core.FreeMealError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewCategoryEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewCategoryEntity(client, entopts)
	}
	core.NewFilterEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewFilterEntity(client, entopts)
	}
	core.NewLatestEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewLatestEntity(client, entopts)
	}
	core.NewListEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewListEntity(client, entopts)
	}
	core.NewLookupEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewLookupEntity(client, entopts)
	}
	core.NewRandomEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewRandomEntity(client, entopts)
	}
	core.NewRandomselectionEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewRandomselectionEntity(client, entopts)
	}
	core.NewSearchEntityFunc = func(client *core.FreeMealSDK, entopts map[string]any) core.FreeMealEntity {
		return entity.NewSearchEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewFreeMealSDK = core.NewFreeMealSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
