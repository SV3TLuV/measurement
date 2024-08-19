import {GetMeasurementsQuery} from "./getMeasurementsQuery.ts";

export type ExportMeasurementsQuery =  GetMeasurementsQuery & {
    format: string
}