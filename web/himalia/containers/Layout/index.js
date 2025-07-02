import React, { useEffect } from 'react'
import { ErrorBoundary } from "react-error-boundary";
import { Outlet } from 'react-router-dom'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'
import Menu from '@components/Menu'
import Toast from '@components/Toast'

import Error from '@containers/Error'

import './style.css'

const layout = ({ fetchWorkspace, loading, application }) => {
  useEffect(() => {
    fetchWorkspace()
  }, [])

  let outlet = <Outlet />
  const applicationErrorContent = (
    <Error type="bug">
      <p>An unexpected error happened</p>
    </Error>
  )

  if (loading.includes('application') || application === undefined) {
    outlet = <SpinningSquare />
  } else if (application && application.error) {
    outlet = (
      <>
        <Error>
          <p>Bad server! The server has just found an error.</p>
        </Error>
      </>
    )
  }

  return (
    <div role="main" className="layout-root">
      <div className="layout-root__menu-container">
        <Menu isUserAdmin={application && application.user_is_admin} />
      </div>
      <div className="layout-root__corpus-container">
        <ErrorBoundary fallback={applicationErrorContent}>
          {outlet}
          <Toast />
        </ErrorBoundary>
      </div>
    </div>
  )
}

const mapStateToProps = state => {
  return {
    loading: state.root.loading,
    application: state.root.application,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchWorkspace: () => dispatch(actions.fetchWorkspace()),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(layout)
