<?php
declare(strict_types=1);

// FreeMeal SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class FreeMealMakeContext
{
    public static function call(array $ctxmap, ?FreeMealContext $basectx): FreeMealContext
    {
        return new FreeMealContext($ctxmap, $basectx);
    }
}
