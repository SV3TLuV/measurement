import {Position} from "../types/position.ts";
import {Card, Divider} from "antd";
import classNames from "classnames";
import './PostInfoCard.scss';
import {getObjectTitle} from "../utils/getObjectTitle.ts";
import {RefObject, useEffect, useRef, useState} from "react";
import {getCardPosition} from "../utils/getCardPosition.ts";
import {Measurement} from "../types/measurement.ts";
import {Facility} from "../../objects/types/facility.ts";
import {convertToLocalTime} from "../../../utils/convertToLocalTime.ts";
import {convertToStringFormat} from "../../../utils/convertToStringFormat.ts";

type PostInfoCardProps = {
    position: Position
    post: Facility
    mapRef: RefObject<HTMLDivElement>
    lastMeasurement: Measurement | undefined
}

export const PostInfoCard = (props: PostInfoCardProps) => {
    const { position, post, lastMeasurement, mapRef } = props;

    const postInfoCard = useRef<HTMLDivElement>(null)

    const [cardPosition, setCardPosition] =
        useState<Position>(getCardPosition(position, { width: 300, height: 0 }, mapRef))

    useEffect(() => {
        if (postInfoCard.current) {
            setCardPosition(getCardPosition({
                x: position.x,
                y: position.y,
            }, {
                width: postInfoCard.current.offsetWidth,
                height: postInfoCard.current.offsetHeight
            }, mapRef));
        }
    }, [position, postInfoCard]);

    return (
        <div
            className={classNames('post-info-card-container')}
            style={{
                left: `${cardPosition.x + 10}px`,
                top: `${cardPosition.y + 10}px`,
            }}
            ref={postInfoCard}
        >
            <Card style={{width: 300}}>
                <h3 className={classNames('post-info-card-title', 'no-margin')}>
                    {`станция ${getObjectTitle(post)}`}
                </h3>
                <Divider/>
                {lastMeasurement ? (
                    <>
                        <div className={classNames('measurement-grid')}>
                            <div className={classNames('measurement-row')}>
                                <p className={classNames('no-margin full-width')}>
                                    Данные за: {convertToStringFormat(convertToLocalTime(lastMeasurement.created))}
                                </p>
                            </div>
                            <div className={classNames('measurement-row')}>
                                <p className={classNames('no-margin')}>Дксд серы:</p>
                                <p className={classNames('no-margin')}>{lastMeasurement.v202918}</p>
                                <p className={classNames('no-margin')}>мг/м<sup>3</sup></p>
                            </div>
                            <div className={classNames('measurement-row')}>
                                <p className={classNames('no-margin')}>Угл. оксид:</p>
                                <p className={classNames('no-margin')}>{lastMeasurement.v202919}</p>
                                <p className={classNames('no-margin')}>мг/м<sup>3</sup></p>
                            </div>
                            <div className={classNames('measurement-row')}>
                                <p className={classNames('no-margin')}>Дксд азота:</p>
                                <p className={classNames('no-margin')}>{lastMeasurement.v202920}</p>
                                <p className={classNames('no-margin')}>мг/м<sup>3</sup></p>
                            </div>
                            <div className={classNames('measurement-row')}>
                                <p className={classNames('no-margin')}>Окс азота:</p>
                                <p className={classNames('no-margin')}>{lastMeasurement.v202921}</p>
                                <p className={classNames('no-margin')}>мг/м<sup>3</sup></p>
                            </div>
                            <div className={classNames('measurement-row')}>
                                <p className={classNames('no-margin')}>Сер. вод-д:</p>
                                <p className={classNames('no-margin')}>{lastMeasurement.v202932}</p>
                                <p className={classNames('no-margin')}>мг/м<sup>3</sup></p>
                            </div>
                        </div>
                    </>
                ) : (
                    <>Нет данных.</>
                )}
            </Card>
        </div>
    )
}