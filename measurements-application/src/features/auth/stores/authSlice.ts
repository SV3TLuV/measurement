import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {AuthResponse} from "../types/authResponse.ts";

type AuthSlice = {
    isAuthorized: boolean
    accessToken: string | null
    refreshToken: string | null
}

const authSlice = createSlice({
    name: "authorization",
    initialState: {
        isAuthorized: false,
        accessToken: null,
        refreshToken: null
    } as AuthSlice,
    reducers: {
        login: (state, action: PayloadAction<AuthResponse>) => {
            state.accessToken = action.payload.accessToken
            state.refreshToken = action.payload.refreshToken
            state.isAuthorized = true
        },
        logout: (state) => {
            state.accessToken = null
            state.refreshToken = null
            state.isAuthorized = false
        }
    }
})

export const {
    login,
    logout
} = authSlice.actions
export default authSlice.reducer