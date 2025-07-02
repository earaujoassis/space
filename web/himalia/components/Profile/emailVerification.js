import React, { useState } from 'react'

const processedRequestMessage = (
    <p>You should receive an e-mail message in the next few minutes.</p>
)

const emailVerification = ({ requestEmailVerification, holder, email, emailVerified }) => {
    if (emailVerified) {
        return null
    }

    const [emailVerificationRequested, setEmailVerificationRequested] = useState(false)

    const requestEmailVerificationMessage = emailVerificationRequested ? (processedRequestMessage) : (
        <p>
            <button
                onClick={() => {
                    requestEmailVerification(holder, email)
                    setEmailVerificationRequested(true)
                }}
                className="button-anchor">
                Request e-mail verification
            </button>
        </p>
    )

    return (
        <>
            <h3 className="globals__subheader">Email verification</h3>
            <div className="globals__warning-box profile__email-verification">
                <p>Your e-mail is not verified. Please request your e-mail verification to avoid any disruption in your account access.</p>
                {requestEmailVerificationMessage}
            </div>
        </>
    )
}

export default emailVerification
