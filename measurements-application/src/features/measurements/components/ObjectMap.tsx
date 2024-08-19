import {RefObject, useEffect, useMemo, useState} from 'react';
import 'ol/ol.css';
import Map from 'ol/Map';
import View from 'ol/View';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';
import Feature from 'ol/Feature';
import Point from 'ol/geom/Point';
import Select from 'ol/interaction/Select';
import { fromLonLat } from 'ol/proj';
import VectorSource from 'ol/source/Vector';
import VectorLayer from 'ol/layer/Vector';
import { Icon, Style } from 'ol/style';
import classNames from "classnames";
import {getObjectMapHeight} from "../utils/getObjectMapHeight.ts";
import './ObjectMap.scss';
import {click} from "ol/events/condition";
import {useGetLastMeasurementQuery} from "../api/measurementApi.ts";
import {useAppDispatch, useTypedSelector} from "../../../hooks/redux.ts";
import {selectTreeMenuItem} from "../stores/treeMenuSlice.ts";
import {PostInfoCard} from "./PostInfoCard.tsx";
import {Position} from "../types/position.ts";
import {useGetPostQuery} from "../../objects/api/objectApi.ts";
import {useGetMeQuery, useGetUserObjectsQuery} from "../../users/api/userApi.ts";
import {Facility} from "../../objects/types/facility.ts";

type ObjectMapProps = {
    mapRef: RefObject<HTMLDivElement>
}

export const ObjectMap = (props: ObjectMapProps) => {
    const {mapRef} = props

    const [height, setHeight] = useState<number>(getObjectMapHeight())
    const {selected} = useTypedSelector(state => state.menu)

    const {data: user} = useGetMeQuery()
    const {data: objects = [] } = useGetUserObjectsQuery(user?.id ?? 0, {
        skip: !user,
    })

    const posts = useMemo(() => {
        const result: Facility[] = []
        for (const [, lab] of objects.entries()) {
            if (lab.children) {
                for (const [, city] of lab.children.entries()) {
                    if (city.children) {
                        result.push(...city.children)
                    }
                }
            }
        }
        return result
    }, [objects])

    const [position, setPosition] = useState<Position>({ x: 0, y: 0 })
    const [postId, setPostId] = useState<number>()
    const { data: post } = useGetPostQuery(postId ?? 0, {
        skip: !postId,
    })
    const { data: lastMeasurement } = useGetLastMeasurementQuery(postId ?? 0, {
        skip: !postId
    })

    const dispatch = useAppDispatch();

    useEffect(() => {
        if (!mapRef.current || posts.length === 0)
            return;

        const map = new Map({
            target: mapRef.current,
            layers: [
                new TileLayer({
                    source: new OSM(),
                }),
            ],
            view: new View({
                center: fromLonLat([
                    posts[0].lon!,
                    posts[0].lat!
                ]),
                zoom: 10,
            }),
        });

        const vectorSource = new VectorSource();

        posts.forEach(post => {
            const iconFeature = new Feature({
                geometry: new Point(fromLonLat([
                    post.lon!,
                    post.lat!
                ])),
            });

            const iconStyle = new Style({
                image: new Icon({
                    anchor: [0.5, 50],
                    anchorXUnits: 'fraction',
                    anchorYUnits: 'pixels',
                    src: `/icons/map-tower-icon.png`,
                    size: [50, 50],
                }),
            });

            iconFeature.setId(post.id)
            iconFeature.setStyle(iconStyle);
            vectorSource.addFeature(iconFeature);
        });

        const vectorLayer = new VectorLayer({
            source: vectorSource,
        });

        map.addLayer(vectorLayer);

        const select = new Select({
            condition: click,
            hitTolerance: 10,
            style: (feature) => feature.get('originalStyle')
        });

        map.addInteraction(select);

        select.on('select', (event) => {
            const selectedFeatures = event.selected;
            if (selectedFeatures.length > 0) {
                const feature = selectedFeatures[0];
                const featureId = Number(feature.getId());
                const isSelected = selected ? selected.id === featureId : false

                dispatch(selectTreeMenuItem(!isSelected ? {
                    id: Number(featureId),
                    type: 'Post'
                } : null))
            }
        });

        map.on('pointermove', (event) => {
            const pixel = map.getEventPixel(event.originalEvent);
            const feature = map.forEachFeatureAtPixel(pixel,
                (feature) => feature,
                {
                    hitTolerance: 10
                });

            if (feature) {
                const featureId = Number(feature.getId());

                setPostId(featureId)
                setPosition({
                    x: pixel[0],
                    y: pixel[1],
                });
            } else {
                setPostId(undefined)
            }
        });

        return () => map.setTarget(undefined);
    }, [posts]);

    useEffect(() => {
        const handleResize = () => {
            if (window.innerHeight !== height) {
                setHeight(getObjectMapHeight());
            }
        };

        window.addEventListener('resize', handleResize);

        return () => window.removeEventListener('resize', handleResize);
    }, [height])

    return (
        <div>
            <div
                ref={mapRef}
                className={classNames('object-map')}
                style={{height: `${height}px`}}
            />
            {postId && post && (
                <PostInfoCard
                    position={position}
                    post={post}
                    mapRef={mapRef}
                    lastMeasurement={lastMeasurement}
                />
            )}
        </div>
    );
}