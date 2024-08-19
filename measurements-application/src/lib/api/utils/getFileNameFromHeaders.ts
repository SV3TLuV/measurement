export const getFileNameFromHeaders = (headers: Headers) => {
    const header = headers.get('content-disposition') ?? ''

    const regex = /filename[^*]?=\s*(?:"([^"]+)"|([^;]+))/i
    const matches = header.match(regex)

    if (matches) {
        return matches[1]
    }
}