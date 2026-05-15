<?php
declare(strict_types=1);

// FreeMeal SDK utility: feature_add

class FreeMealFeatureAdd
{
    public static function call(FreeMealContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
