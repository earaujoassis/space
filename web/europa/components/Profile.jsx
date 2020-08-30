import React, { useEffect, useState } from 'react'

import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'
import { Entry } from '../../core/components/Form.jsx'

import UserStore from '../stores/users'
import UsersActions from '../actions/users'

const profile = () => {
    const [storeState, setStoreState] = useState({isLoading: true})
    const [updatePasswordRequested, setUpdatePasswordRequested] = useState(false)
    const [secretCodesRequested, setSecretCodesRequested] = useState(false)

    useEffect(() => {
        let updateLocalStoreState = () => {
            if (UserStore.success()) {
                let state = Object.assign({isLoading: false}, UserStore.getState().payload || {})
                setStoreState(state)
            }
        }

        UserStore.addChangeListener(updateLocalStoreState)
        UsersActions.fetchProfile()

        return function cleanup() {
            UserStore.removeChangeListener(updateLocalStoreState)
        }
    }, [])

    const { user, isLoading } = storeState
    let inputKey

    if (isLoading || user === undefined) {
        return (
            <Row>
                <Columns className="small-offset-1 small-10 end">
                    <p className="text-center">Loading...</p>
                </Columns>
            </Row>
        )
    }

    return (
        <Row>
            <Columns className="small-offset-1 small-10 end">
                <Row className="profile">
                    <Columns className="small-12">
                        <Entry field="Name" value={`${user.first_name} ${user.last_name}`} />
                        <Entry field="Username" value={user.username} />
                        <Entry field="Email" value={user.email} />
                        <Entry field="Timezone" value={user.timezone_identifier} />
                        {!UserStore.isCurrentUserAdmin() && UserStore.isFeatureGateActive('user.adminify') && (
                            <Row className="profile-entry">
                                <Columns className="columns small-11 small-offset-1 field">
                                    <input className="thin-input"
                                        ref={(r) => inputKey = r}
                                        type="text"
                                        name="application_key"
                                        placeholder="Application Key"
                                        required />
                                    <button className="button-anchor"
                                        onClick={() => UsersActions.adminify(inputKey.value) }>
                                        Make me an admin
                                    </button>
                                </Columns>
                            </Row>
                        )}
                        {UserStore.isCurrentUserAdmin() && (
                            <Entry field="Role" value="Administrator" />
                        )}
                    </Columns>
                </Row>
                <Row className="profile-actions">
                    <Columns className="small-12">
                        {updatePasswordRequested ? (
                            <p>You should receive an e-mail message in the next few minutes.</p>
                        ) : (
                            <p><button
                                onClick={(e) => {
                                    e.preventDefault()
                                    let formData = new FormData()
                                    formData.append('request_type', 'password')
                                    formData.append('holder', user.username)
                                    UsersActions.requestUpdate(formData)
                                    setUpdatePasswordRequested(true)
                                }}
                                className="button-anchor">Update password</button></p>
                        )}
                        {secretCodesRequested ? (
                            <p>You should receive an e-mail message in the next few minutes.</p>
                        ) : (
                            <p><button
                                onClick={(e) => {
                                    e.preventDefault()
                                    let formData = new FormData()
                                    formData.append('request_type', 'secrets')
                                    formData.append('holder', user.username)
                                    UsersActions.requestUpdate(formData)
                                    setSecretCodesRequested(true)
                                }}
                                className="button-anchor">Recreate recovery code and secret code generator</button></p>
                        )}
                    </Columns>
                </Row>
            </Columns>
        </Row>
    )
}

export default profile
