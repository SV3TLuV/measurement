import classNames from "classnames";
import {Button, Card, Form, Input, Popconfirm, Select} from "antd";
import {useEffect, useState} from "react";
import {useShowResult} from "../../../pages/result/hooks/useShowResult.tsx";
import lodash, {isNumber} from "lodash";
import {useOptions} from "../hooks/useOptions.ts";
import {PeriodInput} from "../../../components/ui/period-input/PeriodInput.tsx";
import {useGetConfigurationQuery, useUpdateConfigurationMutation} from "../api/appApi.ts";
import {Configuration} from "../types/configuration.ts";
import {getMinuteFromNumber} from "../../../utils/getMinuteFromNumber.ts";
import {getDayFromNumber} from "../../../utils/getDayFromNumber.ts";
import './ConfigurationForm.scss';
import {getHourFromNumber} from "../../../utils/getHourFromNumber.ts";

const { Option } = Select

export const ConfigurationForm = () => {
    const [form] = Form.useForm<Configuration>();

    const [isChanged, setChanged] = useState<boolean>(false)
    const {
        options: minutes,
    } = useOptions<number>({
        defaultValues: [10, 20, 30, 40, 50, 60],
        getTitle: value => `${value} ${getMinuteFromNumber(value)}`
    })

    const {
        options: checkDays,
        reset: resetCheckDays,
        onSearch: onSearchCheckDays
    } = useOptions<number>({
        defaultValues: [5, 10, 15, 20, 25, 30],
        getTitle: value => `${value} ${getDayFromNumber(value)}`
    })

    const {
        options: disablingHours,
        reset: resetDisablingHours,
        onSearch: onSearchDisablingHours
    } = useOptions<number>({
        defaultValues: Array.from({ length: 24 }, (_, i) => i + 1),
        getTitle: value => `${value} ${getHourFromNumber(value)}`
    })

    const {
        options: disableDays,
        reset: resetDisableDays,
        onSearch: onSearchDisableDays
    } = useOptions<number>({
        defaultValues: [5, 10, 15, 20, 25, 30],
        getTitle: value => `${value} ${getDayFromNumber(value)}`
    })

    const {data: configuration} = useGetConfigurationQuery()
    const [updateConfiguration] = useUpdateConfigurationMutation()

    const showResult = useShowResult()

    const handleSave = async () => {
        try {
            const values = await form.validateFields();
            const config = {
                asoizaLogin: values.asoizaLogin,
                asoizaPassword: values.asoizaPassword,
                collectingInterval: values.collectingInterval * 60, // convert to minutes
                deletingInterval: values.deletingInterval * (3600 * 24), // convert to days
                deletingThreshold: values.deletingThreshold * 2592000, // convert to months
                disablingInterval: values.disablingInterval * 3600, // convert to hours
                disablingThreshold: values.disablingThreshold * (3600 * 24), // convert to days
            } as Configuration
            const response = await updateConfiguration(config)
            await showResult({
                response: response,
                getSuccessMessage: () => ({
                    type: 'message',
                    title: `Конфигурация успешно изменена!`
                })
            })
            if ("data" in response) {
                setChanged(false)
            }
        } catch { /* empty */ }
    };

    const handleChange = (_: unknown, values: Configuration) => {
        setChanged(!lodash.isEqual(values, configuration))
    }

    const resetFormToDefault = () => {
        if (configuration) {
            const config = {
                asoizaLogin: configuration.asoizaLogin,
                asoizaPassword: configuration.asoizaPassword,
                collectingInterval: configuration.collectingInterval / 60, // convert to minutes
                deletingInterval: configuration.deletingInterval / (3600 * 24), // convert to days
                deletingThreshold: configuration.deletingThreshold / 2592000, // convert to months
                disablingInterval: configuration.disablingInterval / 3600, // convert to hours
                disablingThreshold: configuration.disablingThreshold / (3600 * 24), // convert to days
            } as Configuration

            handleCheckDaysSearch(config.deletingInterval.toString())
            handleDisablingHoursSearch(config.disablingInterval.toString())
            handleDisableDaysSearch(config.disablingThreshold.toString())
            form.setFieldsValue(config)
        }
    }

    const handleCheckDaysSearch = (value: string) => {
        resetCheckDays()

        const day = parseInt(value)

        if (isNumber(day) && day > 0) {
            onSearchCheckDays(day)
        }
    }

    const handleDisablingHoursSearch = (value: string) => {
        resetDisablingHours()

        const hour = parseInt(value)

        if (isNumber(hour) && hour > 0) {
            onSearchDisablingHours(hour)
        }
    }

    const handleDisableDaysSearch = (value: string) => {
        resetDisableDays()

        const day = parseInt(value)

        if (isNumber(day) && day > 0) {
            onSearchDisableDays(day)
        }
    }

    useEffect(() => {
        resetFormToDefault()
    }, [configuration, form])
    
    return (
        <Form
            form={form}
            onValuesChange={handleChange}
            layout="vertical"
        >
            <Card
                title="Данные для входа в АСОИЗА+"
                className={classNames('auth-data-card')}
            >
                <Form.Item
                    name={'asoizaLogin'}
                    label='Логин'
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, введите логин!'
                        },
                        {
                            max: 100,
                            message: 'Логин не должен превышать 100 символов!'
                        },
                    ]}
                >
                    <Input
                        count={{
                            show: true,
                            max: 100,
                        }}
                    />
                </Form.Item>
                <Form.Item
                    name={'asoizaPassword'}
                    label='Пароль'
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, введите пароль!'
                        },
                        {
                            max: 100,
                            message: 'Пароль не должен превышать 100 символов!'
                        }
                    ]}
                >
                    <Input.Password
                        count={{
                            show: true,
                            max: 100,
                        }}
                    />
                </Form.Item>
            </Card>
            <Card
                className={classNames('data-collection-card')}
                title="Cбор данных"
            >
                <Form.Item
                    name='collectingInterval'
                    label="Интервал сбора (в минутах)"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, укажите интервал опроса!'
                        },
                        () => ({
                            validator(_, value) {
                                if (value) {
                                    if (value === 0) {
                                        return Promise.reject(new Error('Интервал не может быть 0!'));
                                    } else if (value < 0) {
                                        return Promise.reject(new Error('Интервал не может быть отрицательным!'));
                                    } else if (value > 1440) {
                                        return Promise.reject(new Error('Интервал не может быть больше 24 часов (1440 минут)!'));
                                    }
                                }

                                return Promise.resolve();
                            },
                        }),
                    ]}
                >
                    <Select
                        placeholder="Выберите интервал"
                        showSearch
                    >
                        {minutes.map(option => (
                            <Option
                                key={option.key}
                                value={option.value}
                            >
                                {option.title}
                            </Option>
                        ))}
                    </Select>
                </Form.Item>
                <Form.Item
                    name='disablingInterval'
                    label="Интервал проверки активности (в часах)"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, укажите интервал опроса!'
                        },
                        () => ({
                            validator(_, value) {
                                if (value) {
                                    if (value === 0) {
                                        return Promise.reject(new Error('Интервал не может быть 0!'));
                                    } else if (value < 0) {
                                        return Promise.reject(new Error('Интервал не может быть отрицательным!'));
                                    } else if (value > 30) {
                                        return Promise.reject(new Error('Интервал не может быть больше 30 дней!'));
                                    }
                                }

                                return Promise.resolve();
                            },
                        }),
                    ]}
                >
                    <Select
                        placeholder="Выберите интервал или укажите свой"
                        showSearch
                        onSearch={handleDisablingHoursSearch}
                    >
                        {disablingHours.map(option => (
                            <Option
                                key={option.key}
                                value={option.value}
                            >
                                {option.title}
                            </Option>
                        ))}
                    </Select>
                </Form.Item>
                <Form.Item
                    name='disablingThreshold'
                    label="Отключать через (нет данных в днях)"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, укажите интервал опроса!'
                        },
                        () => ({
                            validator(_, value) {
                                if (value) {
                                    if (value === 0) {
                                        return Promise.reject(new Error('Интервал не может быть 0!'));
                                    } else if (value < 0) {
                                        return Promise.reject(new Error('Интервал не может быть отрицательным!'));
                                    } else if (value > 30) {
                                        return Promise.reject(new Error('Интервал не может быть больше 30 дней!'));
                                    }
                                }

                                return Promise.resolve();
                            },
                        }),
                    ]}
                >
                    <Select
                        placeholder="Выберите интервал или укажите свой"
                        showSearch
                        onSearch={handleDisableDaysSearch}
                    >
                        {disableDays.map(option => (
                            <Option
                                key={option.key}
                                value={option.value}
                            >
                                {option.title}
                            </Option>
                        ))}
                    </Select>
                </Form.Item>
            </Card>
            <Card
                className={classNames('delete-data-card')}
                title="Хранение данных"
            >
                <Form.Item
                    name='deletingInterval'
                    label="Интервал проверки (в днях)"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, укажите интервал опроса!'
                        },
                        () => ({
                            validator(_, value) {
                                if (value) {
                                    if (value === 0) {
                                        return Promise.reject(new Error('Интервал не может быть 0!'));
                                    } else if (value < 0) {
                                        return Promise.reject(new Error('Интервал не может быть отрицательным!'));
                                    } else if (value > 30) {
                                        return Promise.reject(new Error('Интервал не может быть больше 30 дней!'));
                                    }
                                }

                                return Promise.resolve();
                            },
                        }),
                    ]}
                >
                    <Select
                        placeholder="Выберите интервал или укажите свой"
                        showSearch
                        onSearch={handleCheckDaysSearch}
                    >
                        {checkDays.map(option => (
                            <Option
                                key={option.key}
                                value={option.value}
                            >
                                {option.title}
                            </Option>
                        ))}
                    </Select>
                </Form.Item>
                <Form.Item
                    name='deletingThreshold'
                    label="Срок хранения"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, заполните это поле!'
                        },
                        () => ({
                            validator(_, value) {
                                if (value) {
                                    if (value > 120) {
                                        return Promise.reject(new Error('Срок хранения не может быть больше 10 лет!'));
                                    }
                                }

                                return Promise.resolve();
                            },
                        }),
                    ]}
                >
                    <PeriodInput />
                </Form.Item>
            </Card>
            <Card className={classNames('card-actions')}>
                <Popconfirm
                    title='Изменение конфигурации'
                    description={
                        `Внимание, сервис будет перезапущен! Вы уверены что хотите изменить конфигурацию?`
                    }
                    okText='Да'
                    cancelText='Отмена'
                    onConfirm={handleSave}
                    onCancel={resetFormToDefault}
                >
                    <Button
                        className={classNames('button-center')}
                        disabled={!isChanged}
                    >
                        Сохранить
                    </Button>
                </Popconfirm>
            </Card>
        </Form>
    )
}