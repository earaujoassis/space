import * as actionTypes from './types'

export const internalSetToastDisplay = (displayToast = false) => {
    return {
        type: actionTypes.INTERNAL_DISPLAY_TOAST,
        displayToast
    }
}
