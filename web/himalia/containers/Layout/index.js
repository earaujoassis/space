import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { ErrorBoundary } from 'react-error-boundary'
import { Outlet } from 'react-router-dom'

import { fetchWorkspace, fetchUserProfile } from '@actions'
import SpinningSquare from '@ui/SpinningSquare'
import Menu from '@components/Menu'
import Toast from '@components/Toast'

import Error from '@containers/Error'

import './style.css'

const layout = () => {
  const loading = useSelector(state => state.root.loading)
  const application = useSelector(state => state.root.application)
  const user = useSelector(state => state.root.user)

  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchWorkspace())
  }, [])

  useEffect(() => {
    if (
      application !== undefined &&
      application.user_id &&
      !loading.includes('user') &&
      user === undefined
    ) {
      dispatch(fetchUserProfile(application.user_id))
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

export default layout
