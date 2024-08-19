import {CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis} from "recharts";
import {CollectorPollingChartTooltip} from "./CollectorPollingChartTooltip.tsx";
import {tickFormatter} from "../utils/tickFormatter.ts";
import {useEffect, useState} from "react";
import {getChartWidth} from "../utils/getChartWidth.ts";
import {PollingStatistic} from "../types/pollingStatistic.ts";

type SchedulerPollingChartProps = {
    statistics: PollingStatistic[]
}

export const CollectorPollingChart = (props: SchedulerPollingChartProps) => {
    const {statistics} = props
    const [width, setWidth] = useState(getChartWidth())

    useEffect(() => {
        const handleResize = () => {
            if (window.innerWidth !== width) {
                setWidth(getChartWidth())
            }
        };

        window.addEventListener('resize', handleResize);

        return () => window.removeEventListener('resize', handleResize);
    }, [width])

    return (
        <LineChart
            width={width}
            height={400}
            data={statistics}
        >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis
                dataKey="dateTime"
                tickFormatter={tickFormatter}
            />
            <YAxis />
            <Tooltip
                content={<CollectorPollingChartTooltip/>}
                labelFormatter={tickFormatter}
            />
            <Legend />
            <Line
                type='monotone'
                dataKey="receivedCount"
                name='Получено измерений'
                stroke='#82ca9d'
                activeDot={{
                    r: 10
                }}
            />
        </LineChart>
    )
}