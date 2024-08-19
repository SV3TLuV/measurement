import {Period} from "../../../shared/enums/period.ts";
import {QueryWithPagination} from "../../../lib/api/models/queryWithPagination.ts";

export type GetMeasurementsQuery = QueryWithPagination & {
    objectId?: number | null,
    period?: Period | null,
    start?: string | null,
    end?: string | null,
}