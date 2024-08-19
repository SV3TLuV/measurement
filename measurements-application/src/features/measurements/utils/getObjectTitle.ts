import {Facility} from "../../objects/types/facility.ts";

export const getObjectTitle = (post: Facility, withAddress: boolean = false) => {
    let label = `№ ${post.title}`

    if (withAddress && post.address) {
        label += ` (${post.address})`
    }

    return label
}