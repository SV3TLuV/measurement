import {Tree, TreeDataNode, TreeProps} from "antd";
import {useAppDispatch, useTypedSelector} from "../../../hooks/redux.ts";
import {Key, useEffect, useMemo, useState} from "react";
import {useGetMeQuery, useGetUserObjectsQuery} from "../../users/api/userApi.ts";
import {getTreeMenuHeight} from "../utils/getTreeMenuHeight.ts";
import {getObjectTitle} from "../utils/getObjectTitle.ts";
import {TreeObjectType} from "../types/treeObjectType.ts";
import {selectTreeMenuItem} from "../stores/treeMenuSlice.ts";

type TreeMenuProps = {
    onLoading: (loading: boolean) => void
}

export const TreeMenu = (props: TreeMenuProps) => {
    const { onLoading } = props;

    const { selected } = useTypedSelector(state => state.menu)

    const selectedKeys  = useMemo(() => {
        return selected ? [ selected.id ] : []
    }, [selected])

    const [expandedKeys, setExpandedKeys] = useState<Key[]>([]);

    const { data: user } = useGetMeQuery()
    const { data: objects = [], isFetching } = useGetUserObjectsQuery(user?.id ?? 0, {
        skip: !user
    })
    
    const dispatch = useAppDispatch()

    const [height, setHeight] = useState<number>(getTreeMenuHeight())

    const treeData: TreeDataNode[] = objects.map(object => ({
        key: object.id,
        title: object.title,
        children: object.children?.map(city => ({
            key: city.id,
            title: city.title,
            children: city.children?.map(post => ({
                key: post.id,
                title: `станция ${getObjectTitle(post, true)}`,
            }))
        }))
    }))

    const handleSelect: TreeProps['onSelect'] = (keys) => {
        if (keys.length === 0) {
            dispatch(selectTreeMenuItem(null))
            return;
        }

        const id = Number(keys[0]);

        const findItem = (id: number): { type: TreeObjectType } | undefined => {
            for (const lab of objects) {
                if (lab.id === id) {
                    return { type: 'Laboratory' };
                }
                for (const city of lab?.children ?? []) {
                    if (city.id === id) {
                        return { type: 'City' };
                    }
                    const post = city?.children?.find(post => post.id === id);
                    if (post) {
                        return { type: 'Post' };
                    }
                }
            }
        };

        const item = findItem(id);
        if (!item) {
            return;
        }

        dispatch(selectTreeMenuItem({
            id: id,
            type: item.type,
        }))
    }

    const handleExpand: TreeProps['onExpand'] = (keys) => {
        setExpandedKeys(keys)
    }

    useEffect(() => {
        const handleResize = () => {
            if (window.innerHeight !== height) {
                setHeight(getTreeMenuHeight());
            }
        };

        window.addEventListener('resize', handleResize);

        return () => window.removeEventListener('resize', handleResize);
    }, [height])

    useEffect(() => {
        onLoading(isFetching)
    }, [isFetching, onLoading]);

    useEffect(() => {
        if (selected) {
            const relevantIds = new Set<number>();

            objects.forEach(object => {
                object.children?.forEach(city => {
                    city.children?.forEach(post => {
                        if (selectedKeys.includes(post.id) || selectedKeys.includes(city.id)) {
                            relevantIds.add(object.id);
                            relevantIds.add(city.id);
                        }
                    });
                });
            });

            setExpandedKeys(prev => ([...prev, ...Array.from(relevantIds).concat(selectedKeys)]))
        }
    }, [objects]);

    return (
        <Tree
            showLine
            treeData={treeData}
            onSelect={handleSelect}
            selectedKeys={selectedKeys}
            onExpand={handleExpand}
            expandedKeys={expandedKeys}
            height={height}
        />
    )
}
