<?php

namespace Backlink;

final class Info
{
    public const VERSION = '0.1.0';
    public const WEBSITE = 'https://www.seedance-3.org';

    public static function getInfo(): array
    {
        return [
            'name' => 'seedance_3_org',
            'version' => self::VERSION,
            'website' => self::WEBSITE,
            'description' => 'Seedance 3.0 official website backlink helper package.',
        ];
    }

    public static function getPlatformUrl(): string
    {
        return self::WEBSITE;
    }
}
