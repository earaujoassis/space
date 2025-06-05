import React, { useEffect, useState } from 'react'

import Row from '@core/components/Row.jsx'
import Columns from '@core/components/Columns.jsx'
import PasswordInput from '@core/components/PasswordInput.jsx'
import SuccessBox from '@core/components/SuccessBox.jsx'

import { extractDataForm } from '@core/utils/forms'

import { useApp, AppProvider } from '../context/useApp'

const UpdatePassword = ({ onSuccess }) => {
    const { state, actions } = useApp()
    const [errorMessage, setErrorMessage] = useState(null)

    const attemptPassworUpdate = (target) => {
        const data = extractDataForm(target, ['new_password', 'password_confirmation'])
        data.append('_', state.server['action_token'])
        actions.updatePassword(data)
    }

    useEffect(() => {
        actions.loadServerData()

        return () => {
            actions.reset()
        }
    }, [])

    useEffect(() => {
        if (state.success) {
            onSuccess()
        }
    }, [state.success])

    useEffect(() => {
        if (state.error && state.error._message && state.error.error) {
            setErrorMessage(`Error: ${state.error.error}`)
        } else if (state.error) {
            setErrorMessage('Something unexpected happened')
        }
    }, [state.error])

    return (
        <div className="middle-box plain resource-owner-password">
            <Row>
                <Columns className="small-12">
                    <form
                        className="form-common"
                        action="."
                        method="patch"
                        onSubmit={(e) => {
                            e.preventDefault()
                            attemptPassworUpdate(e.target)
                        }}>
                        <p>Update your password with the required fields below</p>
                        <PasswordInput placeholder="New password" name="new_password" />
                        <PasswordInput placeholder="Confirm password" name="password_confirmation" />
                        <button type="submit"
                            className="button expand"
                            disabled={false}>Update password</button>
                    </form>
                    {errorMessage ? <p className="error-message">{errorMessage}</p> : null}
                </Columns>
            </Row>
        </div>
    )
}

const OwnerUpdate = () => {
    const [hasUpdated, setHasUpdated] = useState(false)

    return (
        <div className="resource-owner-update">
            <Row>
                <Columns className="small-12">
                    {hasUpdated === true ? (
                        <SuccessBox>
                            <p>Password updated sucessfully!</p>
                            <p>Get <a href="/">back to the application</a>.</p>
                        </SuccessBox>
                    ) : (
                        <UpdatePassword onSuccess={() => setHasUpdated(true)} />
                    )}
                </Columns>
            </Row>
        </div>
    )
}

const wrapper = () => {
    return (
        <AppProvider>
            <OwnerUpdate />
        </AppProvider>
    )
}

export default wrapper
