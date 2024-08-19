import {useState} from "react";

type useStepsReturn = {
    step: number
    next: () => void
    previous: () => void
    reset: () => void
}

export const useSteps = (): useStepsReturn => {
    const [step, setStep] = useState<number>(0)

    const next = () => setStep(prev => prev + 1)
    const previous = () => setStep(prev => prev - 1)
    const reset = () => setStep(0)

    return {step, next, previous, reset}
}