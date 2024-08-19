import classNames from "classnames";
import {PostsTable} from "../features/objects/components/PostsTable.tsx";
import './PostsPage.scss';

export const PostsPage = () => {
    return (
        <div className={classNames('posts-page')}>
            <PostsTable/>
        </div>
    )
}