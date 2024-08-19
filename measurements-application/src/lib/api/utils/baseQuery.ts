import {fetchBaseQuery} from "@reduxjs/toolkit/dist/query/react";
import {AppState} from "../../store";

export const baseQuery = fetchBaseQuery({
    baseUrl: `${import.meta.env.VITE_API_URL}/api`,
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Cache-Control': 'no-cache'
    },
    mode: "cors",
    prepareHeaders: (headers, { getState }) => {
        const accessToken = (getState() as AppState).auth.accessToken;

        if (accessToken) {
            headers.set('Authorization', `Bearer ${accessToken}`)
        }

        return headers
    }
})