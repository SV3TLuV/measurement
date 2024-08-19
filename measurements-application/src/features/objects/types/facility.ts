import {ObjectType} from "./objectType.ts";

export type Facility = {
    id: number
    title: string
    address: string | null
    lat: number | null
    lon: number | null
    type: ObjectType
    parentId: number | null
    children: Facility[] | null
    lastPollingDateTime: Date | null
    isListened: boolean | null
    laboratory: string | null
    city: string | null
    status: string | null
}