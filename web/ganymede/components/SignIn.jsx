import React, { useEffect, useState } from 'react'

import Row from '@core/components/Row.jsx'
import Columns from '@core/components/Columns.jsx'
import { useApp } from '../context/useApp'

import { getParameterByName } from '@core/utils/url'

const initialForm = () => ({ holder: '', password: '', passcode: '' })
const StepsOrder = ['access', 'secrecy', 'code']
const StepsData = {
    access: {
        className: 'unknown',
        name: 'holder',
        type: 'text',
        placeholder: 'Access holder',
        autocomplete: 'username',
        inputmode: 'text'
    },
    secrecy: {
        className: 'secrecy',
        name: 'password',
        type: 'password',
        placeholder: 'Passphrase',
        autocomplete: 'current-password',
        inputmode: 'text'
    },
    code: {
        className: 'code',
        name: 'passcode',
        type: 'text',
        placeholder: 'Code',
        autocomplete: 'one-time-code',
        inputmode: 'numeric'
    }
}

const SignIn = () => {
    const [currentStepIndex, setCurrentStepIndex] = useState(0)
    const [disabledSubmit, setDisabledSubmit] = useState(false)
    const [formData, setFormData] = useState(initialForm())
    const [securityTimeoutID, setSecurityTimeoutID] = useState(null)
    const [failed, setFailed] = useState(false)
    const [blocked, setBlocked] = useState(false)

    const { state, actions } = useApp()

    const stepName = StepsOrder[currentStepIndex]
    const step = StepsData[stepName]

    const setFormSubmitTimeout = (bool, delay) => {
        const securityTimeoutID = setTimeout(() => {
            if (!bool) {
                setCurrentStepIndex(0)
                setFailed(false)
                actions.clearError()
            }
            setDisabledSubmit(bool)
            setSecurityTimeoutID(null)
        }, delay)
        setSecurityTimeoutID(securityTimeoutID)
    }

    const updateStepValue = (e) => {
        formData[e.target.name] = e.target.value
        setFormData(formData)
    }

    const updateStep = (e) => {
        if (e) e.preventDefault()
        const form = Object.assign({}, formData)
        if (currentStepIndex + 1 < 3) {
            setFormData(form)
            setCurrentStepIndex(currentStepIndex + 1)
        } else {
            const signInData = new FormData()
            StepsOrder.forEach((stepKey) => {
                let name = StepsData[stepKey].name
                signInData.append(name, form[name])
            })
            actions.signIn(signInData)
            e.target.form.reset()
            setFormData(form)
            setDisabledSubmit(true)
        }
    }

    useEffect(() => {
        setFormSubmitTimeout(true, 2 * 60000)

        return () => {
            clearTimeout(securityTimeoutID)
            actions.reset()
        }
    }, [])

    useEffect(() => {
        if (state.success && state.payload) {
            const response = state.payload || {}
            const next = getParameterByName('_')
            let location = `${response.redirect_uri}?client_id=${response.client_id}&code=${response.code}&grant_type=${response.grant_type}&scope=${response.scope}&state=${response.state}`
            if (next && next) {
                location += `&_=${encodeURI(next)}`
            }
            window.location.href = location
        }
    }, [state.success, state.payload])

    useEffect(() => {
        if (state.error) {
            const error = state.error
            if (error.attempts === 'blocked') {
                setBlocked(true)
                setFormData(initialForm())
                setDisabledSubmit(true)
                setCurrentStepIndex(0)
                return
            }
            setFailed(true)
            setFormData(initialForm())
            setFormSubmitTimeout(false, 5 * 1000)
        }
    }, [state.error])

    return (
        <div className="middle-box signin-content">
            <Row>
                <Columns className="small-12">
                    <div className={`user-avatar ${step.className}`}></div>
                    {
                        failed ? (
                            <p className="error-message">Authentication failed</p>
                        ) : null
                    }
                    {
                        blocked ? (
                            <p className="error-message">Authentication blocked for user</p>
                        ) : null
                    }
                    <form action="." method="post">
                        {inputForStep(StepsData['access'], disabledSubmit, stepName === 'access', updateStepValue)}
                        {inputForStep(StepsData['secrecy'], disabledSubmit, stepName === 'secrecy', updateStepValue)}
                        {inputForStep(StepsData['code'], disabledSubmit, stepName === 'code', updateStepValue)}
                        <button type="submit" className="button expand"
                            disabled={disabledSubmit}
                            onClick={(e) => updateStep(e)}>Continue</button>
                    </form>
                    <p className="upper-box">2<sub>min</sub> to attempt a sign-in</p>
                </Columns>
            </Row>
        </div>
    )
}

const inputForStep = (step, disabledSubmit, display, updateStepValue) => {
    return (
        <input
            className={display ? 'active' : 'hidden'}
            type={step.type}
            name={step.name}
            autoComplete={step.autocomplete}
            placeholder={step.placeholder}
            inputMode={step.inputmode}
            onChange={(e) => updateStepValue(e)}
            required
            disabled={disabledSubmit}
        />
    )
}

export default SignIn
