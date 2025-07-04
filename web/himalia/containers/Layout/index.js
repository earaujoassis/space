import React, { useEffect } from 'react'
import { ErrorBoundary } from 'react-error-boundary'
import { Outlet } from 'react-router-dom'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'
import Menu from '@components/Menu'
import Toast from '@components/Toast'

import Error from '@containers/Error'

import './style.css'

const layout = ({
  fetchWorkspace,
  fetchUserProfile,
  loading,
  application,
  user,
}) => {
  useEffect(() => {
    fetchWorkspace()
  }, [])

  useEffect(() => {
    if (
      application !== undefined &&
      application.user_id &&
      !loading.includes('user') &&
      user === undefined
    ) {
      fetchUserProfile(application.user_id)
    }
  }, [application, user])

  let outlet = <Outlet />
  const applicationErrorContent = (
    <Error type="bug">
      <p>An unexpected error happened</p>
    </Error>
  )

  if (
    loading.includes('application') ||
    loading.includes('user') ||
    application === undefined ||
    user === undefined
  ) {
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
    user: state.root.user,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchWorkspace: () => dispatch(actions.fetchWorkspace()),
    fetchUserProfile: id => dispatch(actions.fetchUserProfile(id)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(layout)
