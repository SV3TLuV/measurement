import * as signalR from "@microsoft/signalr"
import {baseApi} from "../api/baseApi.ts";
import {ApiTags} from "../api/enums/apiTags.ts";
import {AppDispatch} from "../store";
import {PollingInformation} from "../../features/collector/types/pollingInformation.ts";
import {setPollingInfo} from "../../features/app/stores/appSlice.ts";
import {User} from "../../features/users/types/user.ts";
import {Roles} from "../../shared/enums/roles.ts";
import {UserDisabledEvent} from "./models/userDisabledEvent.ts";
import {authApi} from "../../features/auth/api/authApi.ts";
import {passwordChangedEvent} from "./models/passwordChangedEvent.ts";
import {UserDeletedEvent} from "./models/userDeletedEvent.ts";

export const startSignalR = (dispatch: AppDispatch, user?: User | undefined) => {
    if (!user) {
        return null
    }

    const connection = new signalR.HubConnectionBuilder()
        .withAutomaticReconnect()
        .withStatefulReconnect()
        .withUrl(`${import.meta.env.VITE_API_URL}/api/events`)
        .build();

    connection.start()
        .then(() => console.log('SignalR Connected'))
        .catch(err => console.error('SignalR Connection Error: ', err));

    registerMeasurementCallbacks(connection, dispatch)
    registerPostCallbacks(connection, dispatch)

    if (user.role.id === Roles.Admin) {
        registerConfigurationCallbacks(connection, dispatch)
        registerSchedulerCallbacks(connection, dispatch)
        registerUserCallbacks(connection, dispatch, user)
    }

    return connection
};

const registerMeasurementCallbacks = (connection: signalR.HubConnection, dispatch: AppDispatch) => {
    connection.on('MeasurementUpdated', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.Measurement }
        ]))
    });
}

const registerPostCallbacks = (connection: signalR.HubConnection, dispatch: AppDispatch) => {
    connection.on('PostPollingStatusChanged', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.Object },
            { type: ApiTags.Configuration },
            { type: ApiTags.Scheduler },
        ]))
    });

    connection.on('PostFindNew', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.Object },
            { type: ApiTags.City },
            { type: ApiTags.Laboratory },
            { type: ApiTags.Configuration },
            { type: ApiTags.Scheduler },
        ]))
    });
}

const registerUserCallbacks = (connection: signalR.HubConnection, dispatch: AppDispatch, user: User) => {
    connection.on('UserCreated', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.User }
        ]))
    });

    connection.on('UserDeleted', (event: UserDeletedEvent) => {
        if (user.id === event.id) {
            dispatch(authApi.endpoints.logout.initiate())
        } else {
            dispatch(baseApi.util.invalidateTags([
                { type: ApiTags.User }
            ]))
        }
    });

    connection.on('UserUpdated', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.User }
        ]))
    });

    connection.on('UserEnabled', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.User }
        ]))
    });

    connection.on('UserDisabled', (event: UserDisabledEvent) => {
        if (user.id === event.id) {
            dispatch(authApi.endpoints.logout.initiate())
        } else {
            dispatch(baseApi.util.invalidateTags([
                { type: ApiTags.User }
            ]))
        }
    });

    connection.on('PasswordChanged', (event: passwordChangedEvent) => {
        if (event.logout) {
            dispatch(authApi.endpoints.logout.initiate())
        }
    })
}

const registerConfigurationCallbacks = (connection: signalR.HubConnection, dispatch: AppDispatch) => {
    connection.on('ConfigurationUpdated', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.Configuration },
            { type: ApiTags.Scheduler },
        ]))
    });
}

const registerSchedulerCallbacks = (connection: signalR.HubConnection, dispatch: AppDispatch) => {
    connection.on('SchedulerOnlineStatusChanged', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.Scheduler },
        ]))
    });

    connection.on('SchedulerInformationUpdated', () => {
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.Scheduler },
        ]))
    });

    connection.on('SchedulerCollecting', (information: PollingInformation) => {
        dispatch(setPollingInfo(information))
    });

    connection.on('SchedulerCollectingEnd', (information: PollingInformation) => {
        dispatch(setPollingInfo(information))
        dispatch(baseApi.util.invalidateTags([
            { type: ApiTags.Scheduler },
        ]))
    });
}