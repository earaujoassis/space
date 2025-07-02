import React from 'react'
import { connect } from 'react-redux'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import * as actions from '@actions'

import './style.css'

const extractMessage = (error) => {
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

const toast = ({ internalSetToastDisplay, displayToast, success, error }) => {
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
            onClick={(e) => {
              e.preventDefault()
              internalSetToastDisplay()
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

const mapStateToProps = (state) => {
  return {
    success: state.root.success,
    error: state.root.error,
    displayToast: state.root.displayToast,
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    internalSetToastDisplay: () => dispatch(actions.internalSetToastDisplay()),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(toast)
