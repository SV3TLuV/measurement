import { RefObject } from "react";
import {Position} from "../types/position.ts";
import {Size} from "../types/size.ts";

export const getCardPosition = (position: Position, size: Size, mapRef: RefObject<HTMLDivElement>): Position => {
    if (!mapRef.current) {
        return position;
    }

    const mapRect = mapRef.current.getBoundingClientRect();
    const mapWidth = mapRect.width;
    const mapHeight = mapRect.height;

    let x = position.x;
    let y = position.y;

    if (x + size.width > mapWidth) {
        x = mapWidth - size.width - 20;
    }

    if (y + size.height > mapHeight) {
        y = mapHeight - size.height - 20;
    }

    return {
        x: x,
        y: y
    }
}