import {useNavigate} from "react-router-dom";
import {FetchBaseQueryError} from "@reduxjs/toolkit/query/react";
import {SerializedError} from "@reduxjs/toolkit";
import {getTextFromResultProps, responseToResult} from "../utils.ts";
import {ApiExceptionCodes, SuccessResultObject} from "../models.ts";
import {message} from "antd";
import {useResultDataContext} from "../../../context/result-data-context/hooks.ts";
import {RoutePaths} from "../../../lib/router/enums/routePaths.ts";

type Response<T> = T | { data: T } | { error: FetchBaseQueryError | SerializedError }

export type useShowResultFuncProps<T> = {
    response: Response<T>,
    showError?: boolean,
    showOnlyError?: boolean,
    messageDuration?: number,
    getSuccessMessage?: (data: T) => SuccessResultObject | undefined
}

export const useShowResult = () => {
    const context = useResultDataContext()
    const navigate = useNavigate()

    return async function <T>(
        {
            response,
            showError = true,
            showOnlyError = false,
            messageDuration = 3,
            getSuccessMessage = () => undefined
        }: useShowResultFuncProps<T>) {

        if (typeof response === 'object' && response) {
            if (showError && 'error' in response && response.error !== undefined) {
                const result = responseToResult(response)

                switch (result.exceptionCode) {
                    case ApiExceptionCodes.Blocked:
                        navigate(`/${RoutePaths.Result}`, { state: result })
                        break
                    default:
                        await message.error(getTextFromResultProps(result), messageDuration)
                        break
                }
            } else if (response) {
                const data = 'data' in response ? response.data : response as T;
                const successMessage = getSuccessMessage(data)

                if (!showOnlyError && successMessage && context.setExtra) {
                    switch (successMessage.type) {
                        case 'result':
                            context.setExtra(successMessage.extra)
                            successMessage.extra = null
                            navigate(`/${RoutePaths.Result}`, { state: successMessage })
                            break
                        case 'message':
                            await message.success(getTextFromResultProps(successMessage), messageDuration)
                            break
                    }
                }
            }
        }
    }
}