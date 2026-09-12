import { FreeMealEntityBase } from '../FreeMealEntityBase';
import type { FreeMealSDK } from '../FreeMealSDK';
import type { Control } from '../types';
import type { Random, RandomListMatch } from '../FreeMealTypes';
declare class RandomEntity extends FreeMealEntityBase<Random> {
    constructor(client: FreeMealSDK, entopts: any);
    make(this: RandomEntity): RandomEntity;
    list(this: any, reqmatch?: RandomListMatch, ctrl?: Control): Promise<RandomEntity[]>;
}
export { RandomEntity };
