import { CategoryEntity } from './entity/CategoryEntity';
import { FilterEntity } from './entity/FilterEntity';
import { LatestEntity } from './entity/LatestEntity';
import { ListEntity } from './entity/ListEntity';
import { LookupEntity } from './entity/LookupEntity';
import { RandomEntity } from './entity/RandomEntity';
import { RandomselectionEntity } from './entity/RandomselectionEntity';
import { SearchEntity } from './entity/SearchEntity';
export type * from './FreeMealTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { FreeMealEntityBase } from './FreeMealEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class FreeMealSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    Category(entopts?: Record<string, any>): CategoryEntity;
    Filter(entopts?: Record<string, any>): FilterEntity;
    Latest(entopts?: Record<string, any>): LatestEntity;
    List(entopts?: Record<string, any>): ListEntity;
    Lookup(entopts?: Record<string, any>): LookupEntity;
    Random(entopts?: Record<string, any>): RandomEntity;
    Randomselection(entopts?: Record<string, any>): RandomselectionEntity;
    Search(entopts?: Record<string, any>): SearchEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): FreeMealSDK;
    tester(testopts?: any, sdkopts?: any): FreeMealSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof FreeMealSDK;
export { stdutil, config, BaseFeature, FreeMealEntityBase, FreeMealSDK, SDK, };
