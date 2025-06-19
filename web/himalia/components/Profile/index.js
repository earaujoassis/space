import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

const processedRequestMessage = (
    <p>You should receive an e-mail message in the next few minutes.</p>
)

const personal = ({ fetchUserProfile, requestEmailVerification, loading, application, user }) => {
    useEffect(() => {
        fetchUserProfile(application.user_id, application.action_token)
    }, [])

    let content = null

    const [emailVerificationRequested, setEmailVerificationRequested] = useState(false)

    const requestEmailVerificationMessage = emailVerificationRequested ? (processedRequestMessage) : (
        <p>
            <button
                onClick={() => {
                    requestEmailVerification(user.username)
                    setEmailVerificationRequested(true)
                }}
                className="button-anchor">
                Request e-mail verification
            </button>
        </p>
    )

    if (loading.includes('user') || user === undefined) {
        content = (<SpinningSquare />)
    } else if (user && !user.error) {
        content = (
            <>
                <div className="globals__siblings">
                    <div className="globals__children">
                        <div className="globals__input-wrapper">
                            <label htmlFor="personal__full-name">Full name</label>
                            <input className="read-only" disabled id="personal__full-name" value={`${user.first_name} ${user.last_name}`} type="text" />
                        </div>
                        <div className="globals__input-wrapper">
                            <label htmlFor="personal__username">Username</label>
                            <input className="read-only" disabled id="personal__username" value={user.username} type="text" />
                        </div>
                        <div className="globals__input-wrapper">
                            <label htmlFor="personal__email">Email</label>
                            <input className="read-only" disabled id="personal__email" value={user.email} type="text" />
                        </div>
                        <div className="globals__input-wrapper">
                            <label htmlFor="personal__role">Role</label>
                            <input className="read-only" disabled id="personal__role" value={user.is_admin ? 'Administrator' : 'Member' } type="text" />
                        </div>
                        <div className="globals__input-wrapper">
                            <label htmlFor="personal__timezone">Timezone</label>
                            <input className="read-only" disabled id="personal__timezone" value={user.timezone_identifier} type="text" />
                        </div>
                    </div>
                    <div className="globals__children">
                        {user.email_verified ? null : (
                            <>
                                <div className="globals__warning-box profile__email-verification">
                                    <p>Your e-mail is not verified. Please request your e-mail verification to avoid any disruption in your account access.</p>
                                    {requestEmailVerificationMessage}
                                </div>
                            </>
                        )}
                    </div>
                </div>
            </>
        )
    }

    return (
        <>
            <h2>Personal information</h2>
            <div className="personal-root">
                {content}
            </div>
        </>
    )
}

const mapStateToProps = state => {
    return {
        loading: state.root.loading,
        application: state.root.application,
        user: state.root.user
    }
}

const mapDispatchToProps = dispatch => {
    return {
        fetchUserProfile: (id, token) => dispatch(actions.fetchUserProfile(id, token)),
        requestEmailVerification: (username) => dispatch(actions.requestEmailVerification(username))
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(personal)
