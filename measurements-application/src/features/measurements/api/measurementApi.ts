import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import {buildUrlArguments} from "../../../utils/buildUrlArguments.ts";
import HTTPMethod from "http-method-enum";
import {PagedList} from "../../../lib/api/models/pagedList.ts";
import {Measurement} from "../types/measurement.ts";
import {GetMeasurementsQuery} from "../types/getMeasurementsQuery.ts";
import {getFileNameFromHeaders} from "../../../lib/api/utils/getFileNameFromHeaders.ts";
import {FetchBaseQueryError} from "@reduxjs/toolkit/query/react";
import {downloadFile} from "../../../utils/downloadFile.ts";
import {baseQuery} from "../../../lib/api/utils/baseQuery.ts";
import {ExportMeasurementsQuery} from "../types/exportMeasurementsQuery.ts";

export const measurementApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getMeasurements: builder.query<PagedList<Measurement>, GetMeasurementsQuery>({
            query: query => ({
                url: `measurements?${buildUrlArguments(query ?? {
                    page: 1,
                    pageSize: 50
                })}`,
                method: HTTPMethod.GET
            }),
            providesTags: result => [
                ...(result?.items ?? []).map(({id}) => ({type: ApiTags.Measurement, id} as const)),
            ],
        }),
        getLastMeasurement: builder.query<Measurement | undefined, number>({
            query: query => ({
                url: `/measurements?${buildUrlArguments({
                    objectId: query,
                    page: 1,
                    pageSize: 1
                })}`,
                method: HTTPMethod.GET,
            }),
            transformResponse: (response: PagedList<Measurement>): Measurement | undefined => {
                const items = response.items ?? [];
                if (items && items.length > 0) {
                    return items[0]
                }
                return undefined
            },
            providesTags: () => [
                {type: ApiTags.Measurement}
            ],
        }),
        downloadMeasurements: builder.mutation<void, ExportMeasurementsQuery>({
            queryFn: async (query, api, extraOptions) => {
                const response = await baseQuery({
                    url: `measurements/export?${buildUrlArguments(query)}`,
                    method: HTTPMethod.GET,
                    responseHandler: response => response.blob()
                }, api, extraOptions)

                const headers = response!.meta?.response?.headers

                if (headers) {
                    const fileName = getFileNameFromHeaders(headers)

                    if (fileName) {
                        const blob = response.data as Blob
                        downloadFile(fileName, blob)
                    }
                }

                return { error: response.error as FetchBaseQueryError }
            },
        }),
    })
})

export const {
    useGetMeasurementsQuery,
    useGetLastMeasurementQuery,
    useDownloadMeasurementsMutation,
} = measurementApi