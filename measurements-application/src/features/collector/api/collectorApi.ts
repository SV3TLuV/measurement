import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import HTTPMethod from "http-method-enum";
import {CollectorInformation} from "../types/collectorInformation.ts";
import {CollectorState} from "../types/collectorState.ts";
import {PollingStatistic} from "../types/pollingStatistic.ts";
import {buildUrlArguments} from "../../../utils/buildUrlArguments.ts";

export const collectorApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getCollectorInformation: builder.query<CollectorInformation, void | undefined>({
            query: () => ({
                url: `collector/information`,
                method: HTTPMethod.GET
            }),
            providesTags: () => [
                {type: ApiTags.Scheduler}
            ],
        }),
        getCollectorState: builder.query<CollectorState, void | undefined>({
            query: () => ({
                url: `collector/state`,
                method: HTTPMethod.GET
            }),
            providesTags: () => [
                {type: ApiTags.Scheduler}
            ]
        }),
        getCollectorStatistic: builder.query<PollingStatistic[], void | undefined>({
            query: () => ({
                url: `collector/statistics?${buildUrlArguments({
                    page: 1,
                    pageSize: 50
                })}`,
                method: HTTPMethod.GET,
            }),
            providesTags: () => [
                {type: ApiTags.Scheduler}
            ]
        })
    })
})

export const {
    useGetCollectorStateQuery,
    useGetCollectorInformationQuery,
    useGetCollectorStatisticQuery,
} = collectorApi