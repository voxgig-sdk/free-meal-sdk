import { FreeMealEntityBase } from '../FreeMealEntityBase';
import type { FreeMealSDK } from '../FreeMealSDK';
import type { Control } from '../types';
import type { Filter, FilterListMatch } from '../FreeMealTypes';
declare class FilterEntity extends FreeMealEntityBase<Filter> {
    constructor(client: FreeMealSDK, entopts: any);
    make(this: FilterEntity): FilterEntity;
    list(this: any, reqmatch?: FilterListMatch, ctrl?: Control): Promise<FilterEntity[]>;
}
export { FilterEntity };
