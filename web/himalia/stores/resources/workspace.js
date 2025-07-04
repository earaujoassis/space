import {
  WORKSPACE_RECORD_START,
  WORKSPACE_RECORD_SUCCESS,
  WORKSPACE_RECORD_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const workspaceRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const workspaceRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.workspace,
    stale: false,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const workspaceRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case WORKSPACE_RECORD_START:
      return workspaceRecordStart(state, action)
    case WORKSPACE_RECORD_SUCCESS:
      return workspaceRecordSuccess(state, action)
    case WORKSPACE_RECORD_ERROR:
      return workspaceRecordError(state, action)
    default:
      return state
  }
}

export default reducer
