import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {Period} from "../../../shared/enums/period.ts";
import {Dayjs} from "dayjs";

type MeasurementFiltersSlice = {
    period: Period | null,
    startDate: Dayjs | null,
    endDate: Dayjs | null,
}

const measurementFiltersSlice = createSlice({
    name: "measurementFilters",
    initialState: {
        period: Period.All
    } as MeasurementFiltersSlice,
    reducers: {
        setPeriod: (state, action: PayloadAction<Period | null>) => {
            state.period = action.payload
        },
        setStartDate: (state, action: PayloadAction<Dayjs | null>) => {
            state.startDate = action.payload
        },
        setEndDate: (state, action: PayloadAction<Dayjs | null>) => {
            state.endDate = action.payload
        }
    }
})

export const {
    setPeriod,
    setStartDate,
    setEndDate
} = measurementFiltersSlice.actions
export default measurementFiltersSlice.reducer