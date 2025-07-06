import {
  USER_RECORD_START,
  USER_RECORD_SUCCESS,
  USER_RECORD_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const userRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const userRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.user,
    stale: false,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const userRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case USER_RECORD_START:
      return userRecordStart(state, action)
    case USER_RECORD_SUCCESS:
      return userRecordSuccess(state, action)
    case USER_RECORD_ERROR:
      return userRecordError(state, action)
    default:
      return state
  }
}

export default reducer
