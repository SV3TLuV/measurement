import {Spin} from "antd";
import classNames from "classnames";
import './Loading.scss';

type LoadingProps = {
    loading: boolean
}

export const Loading = (props: LoadingProps) => {
    const { loading } = props

    if (!loading) {
        return null;
    }

    return (
        <div className={classNames('loading-overlay')}>
            <Spin size="large"/>
        </div>
    )
}