import {
  WORKSPACE_RECORD_START,
  WORKSPACE_RECORD_SUCCESS,
  WORKSPACE_RECORD_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const workspaceRecordStart = () => {
  return {
    type: WORKSPACE_RECORD_START,
  }
}

export const workspaceRecordSuccess = data => {
  return {
    type: WORKSPACE_RECORD_SUCCESS,
    workspace: data.workspace,
  }
}

export const workspaceRecordError = error => {
  return {
    type: WORKSPACE_RECORD_ERROR,
    error: error,
  }
}

export const fetchWorkspace = () => {
  return dispatch => {
    dispatch(workspaceRecordStart())
    fetch
      .get('users/me/workspace')
      .then(response => {
        dispatch(workspaceRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(workspaceRecordError(error))
        dispatch(toastError(error))
      })
  }
}
