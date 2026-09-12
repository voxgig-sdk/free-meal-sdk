import { FreeMealEntityBase } from '../FreeMealEntityBase';
import type { FreeMealSDK } from '../FreeMealSDK';
import type { Control } from '../types';
import type { Latest, LatestListMatch } from '../FreeMealTypes';
declare class LatestEntity extends FreeMealEntityBase<Latest> {
    constructor(client: FreeMealSDK, entopts: any);
    make(this: LatestEntity): LatestEntity;
    list(this: any, reqmatch?: LatestListMatch, ctrl?: Control): Promise<LatestEntity[]>;
}
export { LatestEntity };
