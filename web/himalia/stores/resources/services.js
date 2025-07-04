import {
  SERVICE_RECORD_START,
  SERVICE_RECORD_SUCCESS,
  SERVICE_RECORD_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const serviceRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const serviceRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.services || state.data,
    stale: action.stale,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const serviceRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case SERVICE_RECORD_START:
      return serviceRecordStart(state, action)
    case SERVICE_RECORD_SUCCESS:
      return serviceRecordSuccess(state, action)
    case SERVICE_RECORD_ERROR:
      return serviceRecordError(state, action)
    default:
      return state
  }
}

export default reducer
