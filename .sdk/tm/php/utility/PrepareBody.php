<?php
declare(strict_types=1);

// FreeMeal SDK utility: prepare_body

class FreeMealPrepareBody
{
    public static function call(FreeMealContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
