package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewCategoryEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

var NewFilterEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

var NewLatestEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

var NewListEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

var NewLookupEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

var NewRandomEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

var NewRandomselectionEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

var NewSearchEntityFunc func(client *FreeMealSDK, entopts map[string]any) FreeMealEntityBase

