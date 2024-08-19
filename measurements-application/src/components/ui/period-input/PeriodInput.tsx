import {useEffect, useState} from 'react';
import { InputNumber, Row, Col } from 'antd';
import {getYearFromNumber} from "../../../utils/getYearFromNumber.ts";
import {getMonthFromNumber} from "../../../utils/getMonthFromNumber.ts";

type PeriodInputProps = {
    id?: string;
    value?: number;
    onChange?: (value: number | undefined) => void;
}

export const PeriodInput = (props: PeriodInputProps) => {
    const { id, value, onChange } = props;

    const [year, setYear] = useState<number>(0);
    const [month, setMonth] = useState<number>(
        value ? Math.floor(value % 12) : 0);

    const getTotalMonths = (year: number, month: number) => (year * 12) + month

    const handleChangeYear = (value: number) => {
        let newYear = value;
        let newMonth = month;

        if (newYear < 0) {
            newYear = 0;
            newMonth = 0;
        }

        if (year !== newYear || month !== newMonth) {
            setYear(newYear)
            setMonth(newMonth)
            onChange?.(getTotalMonths(newYear, newMonth))
        }
    }

    const handleChangeMonth = (value: number) => {
        let newYear = year;
        let newMonth = value;

        if (newMonth < 0 && newYear > 0) {
            newYear -= 1;
            newMonth = 11;
        }

        if (newMonth < 0) {
            newMonth = 0;
        }

        if (newMonth > 11) {
            newYear += 1;
            newMonth = 0;
        }

        if (year !== newYear || month !== newMonth) {
            setYear(newYear)
            setMonth(newMonth)
            onChange?.(getTotalMonths(newYear, newMonth))
        }
    }

    useEffect(() => {
        if (value) {
            setYear(Math.floor(value / 12))
            setMonth(Math.floor(value % 12))
        }
    }, [value]);

    return (
        <span id={id}>
            <Row align="middle" gutter={16}>
                <Col>
                    <InputNumber
                        min={0}
                        value={year}
                        onChange={value => handleChangeYear(value ?? 0)}
                    />
                </Col>
                <Col>
                    <span>{getYearFromNumber(year)}</span>
                </Col>
                <Col>
                    <InputNumber
                        value={month}
                        onChange={value => handleChangeMonth(value ?? 0)}
                    />
                </Col>
                <Col>
                    <span>{getMonthFromNumber(month)}</span>
                </Col>
            </Row>
        </span>
    );
}