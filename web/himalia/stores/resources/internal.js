import { INTERNAL_DISPLAY_TOAST, INTERNAL_TOAST_ERROR } from '@actions/types'

const initialState = {
  displayToast: false,
  error: undefined,
  stateSignal: undefined,
}

const internalSetToastDisplay = (state, action) => {
  return { ...state, displayToast: action.displayToast }
}

const internalSetToastError = (state, action) => {
  return { ...state, error: action.error, displayToast: true }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case INTERNAL_DISPLAY_TOAST:
      return internalSetToastDisplay(state, action)
    case INTERNAL_TOAST_ERROR:
      return internalSetToastError(state, action)
    default:
      return state
  }
}

export default reducer
