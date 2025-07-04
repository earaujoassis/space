import {
  CLIENT_RECORD_START,
  CLIENT_RECORD_SUCCESS,
  CLIENT_RECORD_ERROR,
  CLIENT_RECORD_STALE,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const clientRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const clientRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.clients || state.data,
    stale: action.stale,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const clientRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const clientRecordStale = state => {
  const updateState = {
    stale: true,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case CLIENT_RECORD_START:
      return clientRecordStart(state, action)
    case CLIENT_RECORD_SUCCESS:
      return clientRecordSuccess(state, action)
    case CLIENT_RECORD_ERROR:
      return clientRecordError(state, action)
    case CLIENT_RECORD_STALE:
      return clientRecordStale(state, action)
    default:
      return state
  }
}

export default reducer
