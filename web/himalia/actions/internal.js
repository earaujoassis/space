import { INTERNAL_DISPLAY_TOAST, INTERNAL_TOAST_ERROR } from './types'

export const internalSetToastDisplay = (displayToast = false) => {
  return {
    type: INTERNAL_DISPLAY_TOAST,
    displayToast,
  }
}

export const internalSetToastError = error => {
  return {
    type: INTERNAL_TOAST_ERROR,
    error,
  }
}

export const toastError = error => {
  return dispatch => {
    dispatch(internalSetToastError(error))
  }
}
