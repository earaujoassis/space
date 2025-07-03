import React, { useState, useEffect } from 'react'

import Row from '@core/components/Row.jsx'
import Columns from '@core/components/Columns.jsx'
import SuccessBox from '@core/components/SuccessBox.jsx'
import { getParameterByName } from '@core/utils/url'

import { useApp } from '../context/useApp'

const RequestForm = ({ type, onRequest }) => {
  const [lockedForm, setLockedForm] = useState(false)
  const [holder, setHolderValue] = useState('')

  const { actions } = useApp()

  useEffect(() => {
    const securityTimeoutID = setTimeout(() => {
      setLockedForm(true)
      clearTimeout(securityTimeoutID)
    }, 60000)

    return () => {
      clearTimeout(securityTimeoutID)
      actions.reset()
    }
  }, [])

  const handleRequest = e => {
    if (e) e.preventDefault()

    const formData = new FormData()
    formData.append('holder', holder)
    if (type === 'magic') {
      let next = getParameterByName('_')
      if (next && next) {
        formData.append('next', next)
      }
      formData.append('request_type', 'passwordless_signin')
      actions.requestMagicLink(formData)
    } else {
      formData.append('request_type', 'password')
      actions.requestUpdate(formData)
    }

    e.target.form.reset()
    setLockedForm(true)
    onRequest(true)
  }

  return (
    <div className="middle-box signin-content">
      <Row>
        <Columns className="small-12">
          <div className={`user-avatar ${type}`}></div>
          <form action="." method="post">
            <input
              type="text"
              name="holder"
              placeholder="Access holder"
              onChange={e => setHolderValue(e.target.value)}
              required
              disabled={lockedForm}
            />
            <button
              type="submit"
              className="button expand"
              onClick={e => handleRequest(e)}
              disabled={lockedForm}
            >
              {type === 'magic'
                ? 'Request Magic Link'
                : 'Request to update password'}
            </button>
          </form>
          <p className="upper-box">
            1<sub>min</sub> to make your request
          </p>
        </Columns>
      </Row>
    </div>
  )
}

const UserRequest = ({ type }) => {
  const [requestedFulfilled, setRequestFulfilment] = useState(false)

  return (
    <div>
      {requestedFulfilled === true ? (
        <SuccessBox>
          <p>
            If the account holder is valid and active, you should receive an
            e-mail message in the next few minutes.
          </p>
        </SuccessBox>
      ) : (
        <RequestForm type={type} onRequest={setRequestFulfilment} />
      )}
    </div>
  )
}

export default UserRequest
