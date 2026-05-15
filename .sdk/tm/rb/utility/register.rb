# FreeMeal SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

FreeMealUtility.registrar = ->(u) {
  u.clean = FreeMealUtilities::Clean
  u.done = FreeMealUtilities::Done
  u.make_error = FreeMealUtilities::MakeError
  u.feature_add = FreeMealUtilities::FeatureAdd
  u.feature_hook = FreeMealUtilities::FeatureHook
  u.feature_init = FreeMealUtilities::FeatureInit
  u.fetcher = FreeMealUtilities::Fetcher
  u.make_fetch_def = FreeMealUtilities::MakeFetchDef
  u.make_context = FreeMealUtilities::MakeContext
  u.make_options = FreeMealUtilities::MakeOptions
  u.make_request = FreeMealUtilities::MakeRequest
  u.make_response = FreeMealUtilities::MakeResponse
  u.make_result = FreeMealUtilities::MakeResult
  u.make_point = FreeMealUtilities::MakePoint
  u.make_spec = FreeMealUtilities::MakeSpec
  u.make_url = FreeMealUtilities::MakeUrl
  u.param = FreeMealUtilities::Param
  u.prepare_auth = FreeMealUtilities::PrepareAuth
  u.prepare_body = FreeMealUtilities::PrepareBody
  u.prepare_headers = FreeMealUtilities::PrepareHeaders
  u.prepare_method = FreeMealUtilities::PrepareMethod
  u.prepare_params = FreeMealUtilities::PrepareParams
  u.prepare_path = FreeMealUtilities::PreparePath
  u.prepare_query = FreeMealUtilities::PrepareQuery
  u.result_basic = FreeMealUtilities::ResultBasic
  u.result_body = FreeMealUtilities::ResultBody
  u.result_headers = FreeMealUtilities::ResultHeaders
  u.transform_request = FreeMealUtilities::TransformRequest
  u.transform_response = FreeMealUtilities::TransformResponse
}
