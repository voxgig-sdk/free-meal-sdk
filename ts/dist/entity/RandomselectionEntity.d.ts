import { FreeMealEntityBase } from '../FreeMealEntityBase';
import type { FreeMealSDK } from '../FreeMealSDK';
import type { Control } from '../types';
import type { Randomselection, RandomselectionListMatch } from '../FreeMealTypes';
declare class RandomselectionEntity extends FreeMealEntityBase<Randomselection> {
    constructor(client: FreeMealSDK, entopts: any);
    make(this: RandomselectionEntity): RandomselectionEntity;
    list(this: any, reqmatch?: RandomselectionListMatch, ctrl?: Control): Promise<RandomselectionEntity[]>;
}
export { RandomselectionEntity };
