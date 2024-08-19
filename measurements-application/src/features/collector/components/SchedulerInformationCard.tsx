import classNames from "classnames";
import {Card, Col, Progress, Row, Statistic, Typography} from "antd";
import {
    useGetCollectorInformationQuery,
    useGetCollectorStateQuery,
    useGetCollectorStatisticQuery
} from "../api/collectorApi.ts";
import {CollectorPollingChart} from "./CollectorPollingChart.tsx";
import {useCallback, useEffect, useState} from "react";
import './SchedulerInformationCard.scss';
import {convertToStringFormat} from "../../../utils/convertToStringFormat.ts";
import {convertToLocalTime} from "../../../utils/convertToLocalTime.ts";
import dayjs from "dayjs";

const { Countdown } = Statistic

export const SchedulerInformationCard = () => {
    const { data: information, refetch } = useGetCollectorInformationQuery()
    const {data: statistics = []} = useGetCollectorStatisticQuery()
    const {data: state} = useGetCollectorStateQuery()

    const [isFinish, setIsFinish] = useState(true)
    const [value, setValue] = useState(0)

    const getLastPollingDate = useCallback(() => {
        if (!information || !information.lastPollingDateTime) {
            return 'Нет данных.'
        }

        return convertToStringFormat(convertToLocalTime(information.lastPollingDateTime))
    }, [information])

    const handleFinish = () => setIsFinish(true)

    useEffect(() => {
        if (state && state.status === 'pending' && isFinish) {
            refetch()
        }
    }, [state, isFinish]);

    useEffect(() => {
        setTimeout(() => {
            if (information && information.untilNextPolling) {
                const milliseconds = convertToLocalTime(information.untilNextPolling).valueOf()
                const isOldest = milliseconds <= dayjs().valueOf()

                if (!isOldest) {
                    setValue(convertToLocalTime(information.untilNextPolling).valueOf())
                    setIsFinish(false)
                }
            }
        }, 0)
    }, [information]);

    return (
        <Card
            className={classNames('statistic-card')}
            title="Статистика"
        >
            {information && (
                <>
                    <Row
                        className={classNames('statistic-card-base-info')}
                        align="middle"
                        gutter={16}
                    >
                        <Col flex='auto'>
                            <Statistic
                                title='Опрашивается станций:'
                                value={information.listenedPostCount}
                                suffix={`/ ${information.postCount}`}
                            />
                        </Col>
                    </Row>
                    <Row
                        className={classNames('statistic-card-collecting-info')}
                        align='middle'
                        gutter={16}
                    >
                        <Col
                            className={classNames('last-collecting-info')}
                            flex='none'
                        >
                            <Statistic
                                title="Последний сбор данных"
                                valueRender={() => (
                                    <span className={classNames('ant-collector-information-content-value')}>
                                        {getLastPollingDate()}
                                    </span>
                                )}
                            />
                        </Col>
                        <Col
                            className={classNames('next-collecting-info')}
                            flex='none'
                        >
                            {!isFinish ? (
                                <Countdown
                                    title="Следующий сбор через"
                                    value={value}
                                    format="HH:mm:ss"
                                    onFinish={handleFinish}
                                />
                            ) : (
                                <Statistic
                                    title="Следующий сбор через"
                                    valueRender={() => (
                                        <span className={classNames('ant-collector-information-content-value')}>
                                                {'Нет данных.'}
                                            </span>
                                    )}
                                />
                            )}
                        </Col>
                        <Col
                            className={classNames('last-collecting-info')}
                            flex='none'
                        >
                            <Statistic
                                title="Интервал опроса"
                                valueRender={() => (
                                    <span className={classNames('ant-collector-information-content-value')}>
                                        {`${information.pollingInterval / 60} минут`}
                                    </span>
                                )}
                            />
                        </Col>
                    </Row>
                    <Row
                        className={classNames('statistic-card-scheduler-status')}
                        align="middle"
                        gutter={16}
                    >
                        <Col flex='none'>
                            <Typography.Text>
                                Статус: {state?.status === 'collecting' ? 'Сбор данных' : 'Ожидание'}
                            </Typography.Text>
                        </Col>
                        {state?.status === 'collecting' && (
                            <Col flex='auto'>
                                <Progress
                                    percent={state?.pollingPercent}
                                    strokeColor={{
                                        '0%': '#108ee9',
                                        '100%': '#87d068',
                                    }}
                                />
                            </Col>
                        )}
                    </Row>
                    {statistics.length > 0 && (
                        <Row
                            className={classNames('statistic-card-chart')}
                            align="middle"
                            gutter={16}
                        >
                            <CollectorPollingChart statistics={statistics}/>
                        </Row>
                    )}
                </>
            )}
        </Card>
    )
}