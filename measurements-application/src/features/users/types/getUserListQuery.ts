import {QueryWithPagination} from "../../../lib/api/models/queryWithPagination.ts";
import {QueryWithSearch} from "../../../lib/api/models/queryWithSearch.ts";

export type GetUserListQuery = QueryWithPagination & QueryWithSearch & {
    roleIds: number[] | null
}