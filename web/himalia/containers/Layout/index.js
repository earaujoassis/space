import React, { useEffect } from 'react'
import { Outlet } from 'react-router-dom'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'
import Menu from '@components/Menu'
import Toast from '@components/Toast'

import './style.css'

const layout = ({ fetchWorkspace, loading, application }) => {
  useEffect(() => {
    fetchWorkspace()
  }, [])

  let outlet = <Outlet />
  let loadingStatus = false

  if (loading.includes('application') || application === undefined) {
    loadingStatus = true
    outlet = <SpinningSquare />
  } else if (application && application.error) {
    outlet = (
      <>
        <div className="error-illustration">
          <img width="300" src="/public/imgs/illustration.server_error.svg" />
        </div>
        <div className="globals__error-message">
          <h2>Oh, crap</h2>
          <p>Bad server! The server has just found an error.</p>
        </div>
      </>
    )
  }

  return (
    <div role="main" className="layout-root">
      <div className="layout-root__menu-container">
        <Menu isUserAdmin={application && application.user_is_admin} />
      </div>
      <div
        className={`layout-root__corpus-container ${
          loadingStatus ? 'loading' : 'loaded'
        }`}
      >
        {outlet}
        <Toast />
      </div>
    </div>
  )
}

const mapStateToProps = (state) => {
  return {
    loading: state.root.loading,
    application: state.root.application,
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    fetchWorkspace: () => dispatch(actions.fetchWorkspace()),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(layout)
