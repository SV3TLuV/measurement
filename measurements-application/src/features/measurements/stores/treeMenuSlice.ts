import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {TreeItemInfo} from "../types/treeItemInfo.ts";

type TreeMenuSlice = {
    selected: TreeItemInfo | null
}

const treeMenuSlice = createSlice({
    name: "treeMenu",
    initialState: {
        selected: null
    } as TreeMenuSlice,
    reducers: {
        selectTreeMenuItem: (state, action: PayloadAction<TreeItemInfo | null>) => {
            state.selected = action.payload
        },
    }
})

export const {
    selectTreeMenuItem
} = treeMenuSlice.actions
export default treeMenuSlice.reducer