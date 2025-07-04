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
  const workspace = useSelector(state => state.workspace.data)
  const workspaceError = useSelector(state => state.workspace.error)
  const user = useSelector(state => state.user.data)
  const userError = useSelector(state => state.user.error)
  const loadingUser = useSelector(state => state.user.loading)

  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchWorkspace())
  }, [])

  useEffect(() => {
    if (
      workspace !== undefined &&
      workspace.user_id &&
      user === undefined &&
      !loadingUser
    ) {
      dispatch(fetchUserProfile(workspace.user_id))
    }
  }, [workspace, user])

  const workspaceErrorContent = (
    <Error type="bug">
      <p>An unexpected error happened</p>
    </Error>
  )

  let outlet

  if ((workspace === undefined && user === undefined) || loadingUser) {
    outlet = <SpinningSquare />
  } else if (workspaceError || userError) {
    outlet = (
      <>
        <Error>
          <p>Bad server! The server has just found an error.</p>
        </Error>
      </>
    )
  } else {
    outlet = <Outlet />
  }

  return (
    <div role="main" className="layout-root">
      <div className="layout-root__menu-container">
        <Menu isUserAdmin={workspace && workspace.user_is_admin} />
      </div>
      <div className="layout-root__corpus-container">
        <ErrorBoundary fallback={workspaceErrorContent}>
          {outlet}
          <Toast />
        </ErrorBoundary>
      </div>
    </div>
  )
}

export default layout
