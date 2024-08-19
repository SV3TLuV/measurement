export const getObjectMapHeight = (): number => {
    const headerHeight = 64
    const tabHeight = 56
    return window.innerHeight - headerHeight - tabHeight
}
