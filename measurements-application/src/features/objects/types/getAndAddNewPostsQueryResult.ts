import {Facility} from "./facility.ts";

export type GetAndAddNewPostsQueryResult = {
    posts: Facility[]
    newPostsCount: number
    updatedPostsCount: number
}
