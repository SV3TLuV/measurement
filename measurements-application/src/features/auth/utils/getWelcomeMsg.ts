import {getUserNameFromToken} from "../../../utils/getUserNameFromToken.ts";

export function getWelcomeMsg(token: string): string {
    const login = getUserNameFromToken(token)
    if (login) {
        return `Успешный вход, ${login}!`
    }
    return `Добро пожаловать!`
}
