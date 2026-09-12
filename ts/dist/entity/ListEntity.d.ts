import { FreeMealEntityBase } from '../FreeMealEntityBase';
import type { FreeMealSDK } from '../FreeMealSDK';
import type { Control } from '../types';
import type { List, ListListMatch } from '../FreeMealTypes';
declare class ListEntity extends FreeMealEntityBase<List> {
    constructor(client: FreeMealSDK, entopts: any);
    make(this: ListEntity): ListEntity;
    list(this: any, reqmatch?: ListListMatch, ctrl?: Control): Promise<ListEntity[]>;
}
export { ListEntity };
