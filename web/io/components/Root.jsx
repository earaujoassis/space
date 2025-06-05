import React, { useEffect, useState } from 'react'

import { AppProvider, useApp } from '../context/useApp'

import Blocked from './Blocked.jsx'
import SignUp from './SignUp.jsx'
import Success from './Success.jsx'

const isFeatureActive = (key, data) => {
    return data['feature.gates'] && data['feature.gates'][key]
}

const Root = () => {
    const { actions, state } = useApp()
    const [validationFailed, setValidationFailed] = useState(false)

    useEffect(() => {
        actions.loadServerData()
        if (state.error && state.error.user) {
            setValidationFailed(true)
        }
    }, [state.error])

    if (!state.server) {
        return <></>
    }

    if (!isFeatureActive('user.create', state.server)) {
        return (
            <>
                <Blocked />
            </>
        )
    }
    if (state.payload && state.payload.recover_secret && state.payload.code_secret_image) {
        return (
            <>
                <Success
                    codeSecretImage={state.payload.code_secret_image}
                    recoverSecret={state.payload.recover_secret}
                />
            </>
        )
    } else {
        return (
            <>
                <SignUp validationFailed={validationFailed} />
            </>
        )
    }
}

const wrapper = () => {
    return (
        <AppProvider>
            <Root />
        </AppProvider>
    )
}

export default wrapper
