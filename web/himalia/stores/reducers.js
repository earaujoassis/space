import * as actionTypes from '@actions/types'

const initialState = {
    loading: [],
    displayToast: false,
    application: undefined,
    user: undefined,
    clients: undefined,
    services: undefined,
    error: undefined,
    success: undefined,
    stateSignal: undefined
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

const internalSetToastDisplay = (state, action) => {
    return Object.assign({}, state, {
        displayToast: action.displayToast
    })
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
    return Object.assign({}, state, { loading: addLoading(state, 'user'), stateSignal: 'user_record_start' })
}

const userRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'user'),
        success: true,
        error: null,
        user: action.user || { error: true },
        stateSignal: 'user_record_success'
    })
}

const userRecordError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'user'),
        displayToast: true,
        success: false,
        error: action.error,
        user: { error: true },
        stateSignal: 'user_record_error'
    })
}

const clientRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, { loading: addLoading(state, 'client'), stateSignal: 'client_record_start' })
}

const clientRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'client'),
        success: true,
        error: null,
        clients: action.clients || new Array(),
        stateSignal: 'client_record_success'
    })
}

const clientRecordError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'client'),
        displayToast: true,
        success: false,
        error: action.error,
        clients: action.clients || new Array(),
        stateSignal: 'client_record_error'
    })
}

const serviceRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, { loading: addLoading(state, 'service'), stateSignal: 'service_record_start' })
}

const serviceRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'service'),
        success: true,
        error: null,
        services: action.services || new Array(),
        stateSignal: 'service_record_success'
    })
}

const serviceRecordError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'service'),
        displayToast: true,
        success: false,
        error: action.error,
        services: action.services || new Array(),
        stateSignal: 'service_record_error'
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
    case actionTypes.CLIENT_RECORD_START: return clientRecordStart(state, action)
    case actionTypes.CLIENT_RECORD_SUCCESS: return clientRecordSuccess(state, action)
    case actionTypes.CLIENT_RECORD_ERROR: return clientRecordError(state, action)
    case actionTypes.SERVICE_RECORD_START: return serviceRecordStart(state, action)
    case actionTypes.SERVICE_RECORD_SUCCESS: return serviceRecordSuccess(state, action)
    case actionTypes.SERVICE_RECORD_ERROR: return serviceRecordError(state, action)
    default: return state
    }
}

export default reducer
