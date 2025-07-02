import React, { useEffect } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import EmailVerification from './emailVerification'
import Sessions from './sessions'

import './style.css'

const personal = ({
    fetchUserProfile,
    requestEmailVerification,
    fetchApplicationSessionsForUser,
    revokeApplicationSessionForUser,
    loading,
    application,
    user,
    sessions
}) => {
    useEffect(() => {
        fetchUserProfile(application.user_id, application.action_token)
        fetchApplicationSessionsForUser(application.user_id, application.action_token)
    }, [])

    useEffect(() => {
        if (sessions === undefined) {
            fetchApplicationSessionsForUser(application.user_id, application.action_token)
        }
    }, [sessions])

    let content = null

    if (loading.includes('user') || user === undefined) {
        content = (<SpinningSquare />)
    } else if (user && !user.error) {
        content = (
            <div className="globals__siblings globals__max">
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
                <div className="globals__children globals__overlay">
                    <EmailVerification
                        emailVerified={user.email_verified}
                        holder={user.email}
                        email={user.email}
                        requestEmailVerification={requestEmailVerification} />
                    <Sessions
                        sessions={sessions}
                        revokeApplicationSessionForUser={(id) => revokeApplicationSessionForUser(application.user_id, id, application.action_token)} />
                </div>
            </div>
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
        user: state.root.user,
        sessions: state.root.sessions
    }
}

const mapDispatchToProps = dispatch => {
    return {
        fetchUserProfile: (id, token) => dispatch(actions.fetchUserProfile(id, token)),
        requestEmailVerification: (holder, email) => dispatch(actions.requestEmailVerification(holder, email)),
        fetchApplicationSessionsForUser: (id, token) => dispatch(actions.fetchApplicationSessionsForUser(id, token)),
        revokeApplicationSessionForUser: (userId, sessionId, token) => dispatch(actions.revokeApplicationSessionForUser(userId, sessionId, token))
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(personal)
