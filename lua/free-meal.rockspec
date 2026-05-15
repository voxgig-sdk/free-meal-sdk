package = "voxgig-sdk-free-meal"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/free-meal-sdk.git"
}
description = {
  summary = "FreeMeal SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["free-meal_sdk"] = "free-meal_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
