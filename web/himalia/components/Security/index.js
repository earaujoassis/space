import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'

import './style.css'

const processedRequestMessage = (
  <p>You should receive an e-mail message in the next few minutes.</p>
)

const security = ({
  fetchUserProfile,
  requestResetPassword,
  requestResetSecretCodes,
  becomeAdmin,
  loading,
  application,
  user,
}) => {
  useEffect(() => {
    fetchUserProfile(application.user_id, application.action_token)
  }, [])

  const [applicationKey, setApplicationKey] = useState('')
  const [resetPasswordRequested, setResetPasswordRequested] = useState(false)
  const [resetSecretCodesRequested, setResetSecretCodesRequested] =
    useState(false)

  const handleKeypressForAdminify = e => {
    if (e.key === 'Enter') {
      becomeAdmin(application.user_id, applicationKey, application.action_token)
    }
  }

  const becomeAdminUserBox =
    application && application.user_is_admin ? null : (
      <div className="globals__warning-box">
        <h3>Become an admin user</h3>
        <p>Using the application key, you can become an admin user.</p>
        <div className="globals__siblings">
          <div className="globals__input-wrapper">
            <label htmlFor="settings__application-key">Application key</label>
            <input
              value={applicationKey}
              onKeyDown={handleKeypressForAdminify}
              onChange={e => setApplicationKey(e.target.value)}
              id="settings__application-key"
              type="text"
            />
          </div>
          <div className="globals__input-wrapper"></div>
        </div>
        <p>
          <button
            onClick={() =>
              becomeAdmin(
                application.user_id,
                applicationKey,
                application.action_token
              )
            }
            className="button-anchor"
          >
            Confirm application key and become an admin
          </button>
        </p>
      </div>
    )

  const requestPasswordResetMessage = resetPasswordRequested ? (
    processedRequestMessage
  ) : (
    <p>
      <button
        onClick={() => {
          requestResetPassword(user.username)
          setResetPasswordRequested(true)
        }}
        className="button-anchor"
      >
        Request link to update password
      </button>
    </p>
  )

  const requestSecretCodesResetMessage = resetSecretCodesRequested ? (
    processedRequestMessage
  ) : (
    <p>
      <button
        onClick={() => {
          requestResetSecretCodes(user.username)
          setResetSecretCodesRequested(true)
        }}
        className="button-anchor"
      >
        Recreate recovery code and secret code generator (TOTP)
      </button>
    </p>
  )

  if (
    loading.includes('user') ||
    user === undefined ||
    application === undefined
  ) {
    return (
      <>
        <h2>Password &amp; Security</h2>
        <div className="security-root"></div>
      </>
    )
  }

  return (
    <>
      <h2>Password &amp; Security</h2>
      <div className="globals__siblings security-root">
        <div className="globals__warning-box">
          <h3>Update password through a magic link</h3>
          <p>
            Update your password through a magic link sent to your e-mail
            account. It will generate a temporary token so you can securely
            modify your account password.
          </p>
          {requestPasswordResetMessage}
        </div>
        <div className="globals__warning-box">
          <h3>Recreate recovery code and secret code generator</h3>
          <p>
            Through this request, you will recreate your account recovery code
            and the secret code generator used in your secondary factor
            authenticator app. This process is irreversible.
          </p>
          {requestSecretCodesResetMessage}
        </div>
        {becomeAdminUserBox}
      </div>
    </>
  )
}

const mapStateToProps = state => {
  return {
    loading: state.root.loading,
    application: state.root.application,
    user: state.root.user,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchUserProfile: (id, token) =>
      dispatch(actions.fetchUserProfile(id, token)),
    requestResetPassword: username =>
      dispatch(actions.requestResetPassword(username)),
    requestResetSecretCodes: username =>
      dispatch(actions.requestResetSecretCodes(username)),
    becomeAdmin: (id, key, token) =>
      dispatch(actions.becomeAdmin(id, key, token)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(security)
