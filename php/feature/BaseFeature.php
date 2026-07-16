<?php
declare(strict_types=1);

// FreeMeal SDK base feature

class FreeMealBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(FreeMealContext $ctx, array $options): void {}
    public function PostConstruct(FreeMealContext $ctx): void {}
    public function PostConstructEntity(FreeMealContext $ctx): void {}
    public function SetData(FreeMealContext $ctx): void {}
    public function GetData(FreeMealContext $ctx): void {}
    public function GetMatch(FreeMealContext $ctx): void {}
    public function SetMatch(FreeMealContext $ctx): void {}
    public function PrePoint(FreeMealContext $ctx): void {}
    public function PreSpec(FreeMealContext $ctx): void {}
    public function PreRequest(FreeMealContext $ctx): void {}
    public function PreResponse(FreeMealContext $ctx): void {}
    public function PreResult(FreeMealContext $ctx): void {}
    public function PreDone(FreeMealContext $ctx): void {}
    public function PreUnexpected(FreeMealContext $ctx): void {}
}
