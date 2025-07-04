import {
  SESSION_RECORD_START,
  SESSION_RECORD_SUCCESS,
  SESSION_RECORD_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const sessionRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const sessionRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.sessions || state.data,
    stale: action.stale,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const sessionRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case SESSION_RECORD_START:
      return sessionRecordStart(state, action)
    case SESSION_RECORD_SUCCESS:
      return sessionRecordSuccess(state, action)
    case SESSION_RECORD_ERROR:
      return sessionRecordError(state, action)
    default:
      return state
  }
}

export default reducer
