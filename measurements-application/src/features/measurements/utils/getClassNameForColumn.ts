import {Measurement} from "../types/measurement.ts";
import {Quality} from "../../quality/types/quality.ts";
import {parseToFloatSafely} from "../../../utils/parseToFloatSafely.ts";

export const getClassNameForColumn = (key: string, value: string, measurement: Measurement, qualities: Quality[]): string => {
    const classNames = []

    if (key.startsWith("v")) {
        const qualityName = key.replace("v", "q") as keyof Measurement
        const rowQualities = measurement[qualityName]

        if (typeof rowQualities === "string") {
            const ids = rowQualities.split(",").map(value => Number(value))

            const quality = qualities.filter(quality => ids.includes(quality.id))
                .sort((a, b) => a.priority - b.priority)
                .pop()

            if (quality) {
                classNames.push(qualityToClassName(quality))
            }
        }
    }

    const numberValue = parseToFloatSafely(value)

    if (key === "v202932" && numberValue) {
        if (numberValue >= 0.080) {
            classNames.push("dangerous")
        }

        if (numberValue > 0.008) {
            classNames.push("warning")
        }
    }

    return classNames.join(" ")
}

function qualityToClassName(quality: Quality): string {
    switch (quality.id) {
        case 1:
            return "no_data"
        case 2:
            return "adjustment"
        case 3:
            return "reliable"
        case 4:
            return "doubtful"
        case 5:
            return "out_of_range"
        case 6:
            return "rejected"
        case 7:
            return "calculation"
    }

    return ""
}