import {
  USER_REQUEST_START,
  USER_REQUEST_SUCCESS,
  USER_REQUEST_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const userRequestStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const userRequestSuccess = state => {
  NProgress.done()
  const updateState = {
    stale: false,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const userRequestError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case USER_REQUEST_START:
      return userRequestStart(state, action)
    case USER_REQUEST_SUCCESS:
      return userRequestSuccess(state, action)
    case USER_REQUEST_ERROR:
      return userRequestError(state, action)
    default:
      return state
  }
}

export default reducer
