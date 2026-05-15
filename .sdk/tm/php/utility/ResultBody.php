<?php
declare(strict_types=1);

// FreeMeal SDK utility: result_body

class FreeMealResultBody
{
    public static function call(FreeMealContext $ctx): ?FreeMealResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
