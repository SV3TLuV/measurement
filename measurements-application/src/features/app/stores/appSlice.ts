import {createSlice} from "@reduxjs/toolkit";

type AppState = {
    available: boolean,
}

const appSlice = createSlice({
    name: "app",
    initialState: {
        available: false,
    } as AppState,
    reducers: {
        available: (state) => {
            state.available = true;
        },
        unavailable: (state) => {
            state.available = false;
        },
    }
})

export const {
    available,
    unavailable,
} = appSlice.actions
export default appSlice.reducer