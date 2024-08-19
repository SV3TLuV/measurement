import {jwtDecode} from "jwt-decode";

export const getUserNameFromToken = (token: string): string | undefined => {
    const result = jwtDecode(token)
    return result.sub
}