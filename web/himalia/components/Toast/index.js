import React from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import { internalSetToastDisplay } from '@actions'

import './style.css'

const extractMessage = error => {
  if (
    error &&
    error.response &&
    error.response.data &&
    error.response.data.error
  ) {
    return `Server message: ${error.response.data.error}.`
  } else {
    return 'Something unexpected happened.'
  }
}

const toast = () => {
  const displayToast = useSelector(state => state.root.displayToast)
  const error = useSelector(state => state.root.error)
  const success = useSelector(state => state.root.success)

  const dispatch = useDispatch()

  if (displayToast === true && success === false) {
    return (
      <div className="toast-root">
        <div className="toast-icon-box">
          <span className="toast-icon">
            <FontAwesomeIcon
              className="menu-root__icon"
              icon={icon({ name: 'exclamation-circle' })}
            />
          </span>
        </div>
        <div className="toast-body">
          <span className="toast-title">Error</span>
          <span className="toast-description">{extractMessage(error)}</span>
        </div>
        <div className="toast-close-box">
          <button
            onClick={e => {
              e.preventDefault()
              dispatch(internalSetToastDisplay())
            }}
            className="toast-close"
          >
            &times;
          </button>
        </div>
      </div>
    )
  }

  return null
}

export default toast
