import {ResultProps} from "antd";

export type Dictionary<T> = {
    [key: string]: T
}

export type ApiError = {
    status: number,
    exceptionKey: number,
    message: string,
}

export type ErrorResultObject = ResultProps & {
    exceptionCode: ApiExceptionCodes
}

export type SuccessResultObject = ResultProps & {
    type: 'result' | 'message',
}

export type ResultObject = ErrorResultObject | SuccessResultObject

export enum ApiExceptionCodes {
    Unknown = 0,
    InternalServerError = 1,
    Blocked = 2,
    IncorrectLoginOrPassword = 3,
    Unauthorized = 4,
    BadRequest = 5,
    NotFound = 6,
    AlreadyExists = 7,
}

export const errors: Dictionary<ErrorResultObject> = {
    [ApiExceptionCodes.Unknown]: {
        status: 'error',
        title: 'Неизвестная ошибка.',
        subTitle: 'Повторите попытку позже.',
        exceptionCode: ApiExceptionCodes.Unknown,
    },
    [ApiExceptionCodes.InternalServerError]: {
        status: '500',
        title: 'Внутренняя ошибка сервера.',
        subTitle: 'Произошла ошибка. Повторите попытку позже.',
        exceptionCode: ApiExceptionCodes.InternalServerError,
    },
    [ApiExceptionCodes.Blocked]: {
        status: 'error',
        title: 'Ваш аккаунт был заблокирован.',
        exceptionCode: ApiExceptionCodes.Blocked,
    },
    [ApiExceptionCodes.IncorrectLoginOrPassword]: {
        status: 'error',
        title: 'Неверный логин или пароль.',
        exceptionCode: ApiExceptionCodes.IncorrectLoginOrPassword,
    },
    [ApiExceptionCodes.Unauthorized]: {
        status: 'error',
        title: 'Неавторизован.',
        exceptionCode: ApiExceptionCodes.Unauthorized,
    },
    [ApiExceptionCodes.BadRequest]: {
        status: 'error',
        title: 'Некорректный запрос.',
        exceptionCode: ApiExceptionCodes.BadRequest,
    },
    [ApiExceptionCodes.NotFound]: {
        status: 'error',
        title: 'Запрашиваемый ресур не найден.',
        exceptionCode: ApiExceptionCodes.NotFound,
    },
    [ApiExceptionCodes.AlreadyExists]: {
        status: 'error',
        title: 'Уже существует.',
        exceptionCode: ApiExceptionCodes.AlreadyExists,
    },
}
