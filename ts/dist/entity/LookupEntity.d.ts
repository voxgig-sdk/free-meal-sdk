import { FreeMealEntityBase } from '../FreeMealEntityBase';
import type { FreeMealSDK } from '../FreeMealSDK';
import type { Control } from '../types';
import type { Lookup, LookupListMatch } from '../FreeMealTypes';
declare class LookupEntity extends FreeMealEntityBase<Lookup> {
    constructor(client: FreeMealSDK, entopts: any);
    make(this: LookupEntity): LookupEntity;
    list(this: any, reqmatch?: LookupListMatch, ctrl?: Control): Promise<LookupEntity[]>;
}
export { LookupEntity };
