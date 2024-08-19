import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {User} from "../types/user.ts";

type Action = "Create" | "Edit"

type UserFormSlice = {
    userData: User | null
    nextBtnEnabled: boolean
    step: number
    type: Action
}

const initialState: UserFormSlice = {
    userData: null,
    nextBtnEnabled: false,
    step: 0,
    type: "Create",
}

export const userFormSlice = createSlice({
    name: 'userForm',
    initialState: initialState,
    reducers: {
        nextStep: (state) => {
            if (state.step < 3) {
                state.step++
            }
        },
        previousStep: (state) => {
            if (state.step > 0) {
                state.step--
            }
        },
        reset: (state) => {
            state.userData = initialState.userData
            state.nextBtnEnabled = initialState.nextBtnEnabled
            state.step = initialState.step
            state.type = initialState.type
        },
        enableNextBtn: (state) => {
            state.nextBtnEnabled = true
        },
        disableNextBtn: (state) => {
            state.nextBtnEnabled = false
        },
        setUser: (state, payload: PayloadAction<User | null>) => {
            state.userData = payload.payload
        },
        setType: (state, payload: PayloadAction<Action>) => {
            state.type = payload.payload
        },
    }
})

export const {
    nextStep,
    previousStep,
    reset,
    setUser,
    setType,
    enableNextBtn,
    disableNextBtn,
} = userFormSlice.actions
export default userFormSlice.reducer