import {useState} from "react";

export const useDialog = (defaultValue: boolean = false) => {
    const [open, setOpen] = useState(defaultValue)
    const show = () => setOpen(true)
    const close = () => setOpen(false)
    return {open, show, close}
}