<?php
declare(strict_types=1);

// FreeMeal SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

FreeMealUtility::setRegistrar(function (FreeMealUtility $u): void {
    $u->clean = [FreeMealClean::class, 'call'];
    $u->done = [FreeMealDone::class, 'call'];
    $u->make_error = [FreeMealMakeError::class, 'call'];
    $u->feature_add = [FreeMealFeatureAdd::class, 'call'];
    $u->feature_hook = [FreeMealFeatureHook::class, 'call'];
    $u->feature_init = [FreeMealFeatureInit::class, 'call'];
    $u->fetcher = [FreeMealFetcher::class, 'call'];
    $u->make_fetch_def = [FreeMealMakeFetchDef::class, 'call'];
    $u->make_context = [FreeMealMakeContext::class, 'call'];
    $u->make_options = [FreeMealMakeOptions::class, 'call'];
    $u->make_request = [FreeMealMakeRequest::class, 'call'];
    $u->make_response = [FreeMealMakeResponse::class, 'call'];
    $u->make_result = [FreeMealMakeResult::class, 'call'];
    $u->make_point = [FreeMealMakePoint::class, 'call'];
    $u->make_spec = [FreeMealMakeSpec::class, 'call'];
    $u->make_url = [FreeMealMakeUrl::class, 'call'];
    $u->param = [FreeMealParam::class, 'call'];
    $u->prepare_auth = [FreeMealPrepareAuth::class, 'call'];
    $u->prepare_body = [FreeMealPrepareBody::class, 'call'];
    $u->prepare_headers = [FreeMealPrepareHeaders::class, 'call'];
    $u->prepare_method = [FreeMealPrepareMethod::class, 'call'];
    $u->prepare_params = [FreeMealPrepareParams::class, 'call'];
    $u->prepare_path = [FreeMealPreparePath::class, 'call'];
    $u->prepare_query = [FreeMealPrepareQuery::class, 'call'];
    $u->result_basic = [FreeMealResultBasic::class, 'call'];
    $u->result_body = [FreeMealResultBody::class, 'call'];
    $u->result_headers = [FreeMealResultHeaders::class, 'call'];
    $u->transform_request = [FreeMealTransformRequest::class, 'call'];
    $u->transform_response = [FreeMealTransformResponse::class, 'call'];
});
