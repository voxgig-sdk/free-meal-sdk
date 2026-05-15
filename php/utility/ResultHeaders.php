<?php
declare(strict_types=1);

// FreeMeal SDK utility: result_headers

class FreeMealResultHeaders
{
    public static function call(FreeMealContext $ctx): ?FreeMealResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
