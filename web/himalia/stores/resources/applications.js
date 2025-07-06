import {
  APPLICATION_RECORD_START,
  APPLICATION_RECORD_SUCCESS,
  APPLICATION_RECORD_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const applicationRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const applicationRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.clients || state.data,
    stale: action.stale,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const applicationRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case APPLICATION_RECORD_START:
      return applicationRecordStart(state, action)
    case APPLICATION_RECORD_SUCCESS:
      return applicationRecordSuccess(state, action)
    case APPLICATION_RECORD_ERROR:
      return applicationRecordError(state, action)
    default:
      return state
  }
}

export default reducer
