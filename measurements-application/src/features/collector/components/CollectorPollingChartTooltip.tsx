import classNames from "classnames";
import {TooltipProps} from "recharts";
import {ValueType} from "recharts/types/component/DefaultTooltipContent";
import {tickFormatter} from "../utils/tickFormatter.ts";

export const CollectorPollingChartTooltip = (props: TooltipProps<ValueType, string>) => {
    const { active, payload, label } = props

    if (!active || !payload) {
        return null
    }

    return (
        <div className={classNames('custom-tooltip')}>
            <p className="label">
                {`Дата: ${tickFormatter(label)}`}
            </p>
            <p>
                {`Получено измерений: ${payload[0].value}`}
            </p>
            <p>
                {`Опрошено станций: ${payload[0].payload?.postCount}`}
            </p>
            <p>
                {`Длительность опроса: ${payload[0].payload?.duration}`}
            </p>
        </div>
    )
}