<?php
declare(strict_types=1);

// FreeMeal SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class FreeMealFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new FreeMealBaseFeature();
            case "test":
                return new FreeMealTestFeature();
            default:
                return new FreeMealBaseFeature();
        }
    }
}
