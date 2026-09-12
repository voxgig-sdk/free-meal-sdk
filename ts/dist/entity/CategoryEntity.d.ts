import { FreeMealEntityBase } from '../FreeMealEntityBase';
import type { FreeMealSDK } from '../FreeMealSDK';
import type { Control } from '../types';
import type { Category, CategoryListMatch } from '../FreeMealTypes';
declare class CategoryEntity extends FreeMealEntityBase<Category> {
    constructor(client: FreeMealSDK, entopts: any);
    make(this: CategoryEntity): CategoryEntity;
    list(this: any, reqmatch?: CategoryListMatch, ctrl?: Control): Promise<CategoryEntity[]>;
}
export { CategoryEntity };
