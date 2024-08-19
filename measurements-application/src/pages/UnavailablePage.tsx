import {Button, Result, Spin, Statistic} from "antd";
import dayjs from "dayjs";
import classNames from "classnames";
import {useGetAppIsOnlineQuery} from "../features/app/api/appApi.ts";
import './UnavailablePage.scss';

const { Countdown } = Statistic;

const checkInterval = 15000

export const UnavailablePage = () => {
    const { isFetching, refetch } = useGetAppIsOnlineQuery(undefined, {
        pollingInterval: checkInterval,
    })

    const deadline = dayjs().add(checkInterval, 'milliseconds').valueOf();

    const handleCheckNow = async () => await refetch()

    return (
        <div className={classNames('unavailable-page')}>
            {isFetching ? (
                <Spin size='large' />
            ) : (
                <Result
                    status='500'
                    subTitle={(
                        <div>
                            Извините, сервис временно недоступен.
                            <br />
                            Автоматическая проверка через
                        </div>
                    )}
                    extra={(
                        <div>
                            <div className={classNames('timer')}>
                                <Countdown
                                    value={deadline}
                                    format='ss'
                                />
                            </div>
                            <div className={classNames('actions')}>
                                <Button onClick={handleCheckNow}>
                                    Проверить сейчас
                                </Button>
                            </div>
                        </div>
                    )}
                />
            )}
        </div>
    )
}