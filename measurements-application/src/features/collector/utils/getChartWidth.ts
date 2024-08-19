export const getChartWidth = () => {
    const padding = 136;
    const result = window.innerWidth - padding;
    return result > 800 ? 800 : result;
}