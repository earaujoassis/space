import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NewEmail from './newEmail'
import Emails from './emails'

const personal = ({
    fetchUserProfile,
    fetchEmails,
    requestEmailVerification,
    addEmail,
    loading,
    application,
    emails,
    user
}) => {
    const [requestedVerification, setRequestedVerification] = useState([])
    let content = null

    useEffect(() => {
        fetchUserProfile(application.user_id, application.action_token)
        fetchEmails(application.action_token)
    }, [])

    useEffect(() => {
        if (emails === undefined) {
            fetchEmails(application.action_token)
        }
    }, [emails])

    if (
        loading.includes('email') ||
        loading.includes('user') ||
        emails === undefined ||
        user === undefined
    ) {
        content = (<SpinningSquare />)
    } else if (emails) {
        const primaryEmail = {
            verified: user.email_verified,
            address: user.email,
            primary: true
        }
        const emailsComplete = [primaryEmail, ...emails]
        content = (
            <>
                <p>Emails you can use to receive notifications. For authentication you can only use your primary e-mail.</p>
                <Emails
                    emails={emailsComplete}
                    requestedVerification={requestedVerification}
                    setRequestedVerification={setRequestedVerification}
                    requestEmailVerification={(email) => requestEmailVerification(user.email, email)} />
                <NewEmail addEmail={(data) => addEmail(data, application.action_token)} />
            </>
        )
    }

    return (
        <>
            <h2>Emails</h2>
            <div className="emails-root">
                {content}
            </div>
        </>
    )
}

const mapStateToProps = state => {
    return {
        loading: state.root.loading,
        application: state.root.application,
        emails: state.root.emails,
        user: state.root.user
    }
}

const mapDispatchToProps = dispatch => {
    return {
        fetchUserProfile: (id, token) => dispatch(actions.fetchUserProfile(id, token)),
        fetchEmails: (token) => dispatch(actions.fetchEmails(token)),
        requestEmailVerification: (holder, email) => dispatch(actions.requestEmailVerification(holder, email)),
        addEmail: (data, token) => dispatch(actions.addEmail(data, token))
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(personal)
