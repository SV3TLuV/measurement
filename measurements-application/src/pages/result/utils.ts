import {SerializedError} from "@reduxjs/toolkit";
import {FetchBaseQueryError} from "@reduxjs/toolkit/query/react";
import {ApiError, ApiExceptionCodes, errors, ErrorResultObject} from "./models.ts";
import {ResultProps} from "antd";

export const responseToResult = ({ error }: { error: FetchBaseQueryError | SerializedError }): ErrorResultObject => {
    if (isErrorFetchBaseQueryError(error)) {
        const { data } = error

        if (isApiError(data)) {
            if (data.exceptionKey as ApiExceptionCodes !== undefined) {
                switch (data.exceptionKey) {
                    case ApiExceptionCodes.AlreadyExists:
                        return {
                            ...errors[data.exceptionKey],
                            title: data.message,
                        }
                    default:
                        return errors[data.exceptionKey]
                }
            }
        }
    }

    return errors[ApiExceptionCodes.Unknown]
}

export const getTextFromResultProps = (obj: ResultProps): string => {
    if (obj.subTitle !== undefined) {
        return `${obj.title} ${obj.subTitle}`
    }

    return `${obj.title}`
}

export const isApiError = (error: unknown): error is ApiError => {
    return (error as ApiError).exceptionKey !== undefined;
}

export const isErrorFetchBaseQueryError = (error: unknown): error is FetchBaseQueryError => {
    return (error as FetchBaseQueryError).data !== undefined;
}