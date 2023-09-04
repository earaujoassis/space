import * as actionTypes from '@actions/types'

const initialState = {
    loading: [],
    displayToast: false,
    application: undefined,
    user: undefined,
    error: undefined,
    success: undefined
}

const addLoading = (state, entity) => {
    const loading = JSON.parse(JSON.stringify(state.loading))
    loading.push(entity)
    return loading
}

const reduceLoading = (state, entity) => {
    const loading = JSON.parse(JSON.stringify(state.loading))
    return loading.filter(element => element !== entity)
}

const internalRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, { loading: addLoading(state, 'application') })
}

const internalRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'application'),
        success: true,
        error: null,
        application: action.application || { error: true }
    })
}

const internalRecordError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'application'),
        success: false,
        error: action.error,
        application: { error: true }
    })
}

const userRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, { loading: addLoading(state, 'user') })
}

const userRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'user'),
        success: true,
        error: null,
        user: action.user || state.user || { error: true }
    })
}

const userRecordError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'user'),
        displayToast: true,
        success: false,
        error: action.error,
        user: { error: true }
    })
}

const internalSetToastDisplay = (state, action) => {
    return Object.assign({}, state, {
        displayToast: action.displayToast
    })
}

const reducer = (state = initialState, action) => {
    switch (action.type) {
    case actionTypes.INTERNAL_DISPLAY_TOAST: return internalSetToastDisplay(state, action)
    case actionTypes.INTERNAL_RECORD_START: return internalRecordStart(state, action)
    case actionTypes.INTERNAL_RECORD_SUCCESS: return internalRecordSuccess(state, action)
    case actionTypes.INTERNAL_RECORD_ERROR: return internalRecordError(state, action)
    case actionTypes.USER_RECORD_START: return userRecordStart(state, action)
    case actionTypes.USER_RECORD_SUCCESS: return userRecordSuccess(state, action)
    case actionTypes.USER_RECORD_ERROR: return userRecordError(state, action)
    default: return state
    }
}

export default reducer
