package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewCategoryEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

var NewFilterEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

var NewLatestEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

var NewListEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

var NewLookupEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

var NewRandomEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

var NewRandomselectionEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

var NewSearchEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntity

