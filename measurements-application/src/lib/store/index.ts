import {combineReducers, configureStore} from "@reduxjs/toolkit";
import {persistReducer} from "redux-persist";
import storage from "redux-persist/es/storage";
import persistStore from "redux-persist/es/persistStore";
import {FLUSH, PAUSE, PERSIST, PURGE, REGISTER, REHYDRATE} from "redux-persist/es/constants";
import {baseApi} from "../api/baseApi.ts";
import {authApi} from "../../features/auth/api/authApi.ts";
import authSlice from "../../features/auth/stores/authSlice.ts";
import appSlice from "../../features/app/stores/appSlice.ts";
import treeMenuSlice from "../../features/measurements/stores/treeMenuSlice.ts";
import measurementFiltersSlice from "../../features/measurements/stores/measurementFiltersSlice.ts";
import userFormSlice from "../../features/users/stores/userFormSlice.ts";

const rootReducer = combineReducers({
    [baseApi.reducerPath]: baseApi.reducer,
    [authApi.reducerPath]: authApi.reducer,
    app: appSlice,
    auth: authSlice,
    menu: treeMenuSlice,
    userForm: userFormSlice,
    measurementFilters: measurementFiltersSlice,
});

const persistedReducer = persistReducer({
    key: 'root',
    storage,
    whitelist: ['auth'],
}, rootReducer);

export const setupStore = () => {
    return configureStore({
        reducer: persistedReducer,
        middleware: getDefaultMiddleware =>
            getDefaultMiddleware({
                serializableCheck: {
                    ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER]
                }
            }).concat([
                baseApi.middleware,
                authApi.middleware
            ]),
    });
}

export const store = setupStore()
export const persistor = persistStore(store)
export type AppState = ReturnType<typeof rootReducer>
export type AppDispatch = typeof store.dispatch