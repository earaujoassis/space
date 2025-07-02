import * as actionTypes from '@actions/types'

const initialState = {
    loading: [],
    displayToast: false,
    application: undefined,
    user: undefined,
    emails: undefined,
    sessions: undefined,
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
    return Object.assign({}, state, {
        loading: addLoading(state, 'application')
    })
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
    return Object.assign({}, state, {
        loading: addLoading(state, 'user'),
        stateSignal: 'user_record_start'
    })
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

const userRequestStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, {
        loading: addLoading(state, 'user'),
        stateSignal: 'user_request_start'
    })
}

const userRequestSuccess = (state) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'user'),
        success: true,
        error: null,
        stateSignal: 'user_request_success'
    })
}

const userRequestError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'user'),
        displayToast: true,
        success: false,
        error: action.error,
        stateSignal: 'user_request_error'
    })
}

const emailRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, {
        loading: addLoading(state, 'email'),
        stateSignal: 'email_record_start'
    })
}

const emailRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'email'),
        success: true,
        error: null,
        emails: action.emails,
        stateSignal: 'email_record_success'
    })
}

const emailRecordError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'email'),
        displayToast: true,
        success: false,
        error: action.error,
        emails: action.emails,
        stateSignal: 'email_record_error'
    })
}

const sessionRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, {
        loading: addLoading(state, 'session'),
        stateSignal: 'session_record_start'
    })
}

const sessionRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'session'),
        success: true,
        error: null,
        sessions: action.sessions,
        stateSignal: 'session_record_success'
    })
}

const sessionRecordError = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'session'),
        displayToast: true,
        success: false,
        error: action.error,
        sessions: action.sessions,
        stateSignal: 'session_record_error'
    })
}

const clientRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, {
        loading: addLoading(state, 'client'),
        stateSignal: 'client_record_start'
    })
}

const clientRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'client'),
        success: true,
        error: null,
        clients: action.clients,
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
        clients: action.clients,
        stateSignal: 'client_record_error'
    })
}

const serviceRecordStart = (state) => {
    NProgress.start()
    return Object.assign({}, state, {
        loading: addLoading(state, 'service'),
        stateSignal: 'service_record_start'
    })
}

const serviceRecordSuccess = (state, action) => {
    NProgress.done()
    return Object.assign({}, state, {
        loading: reduceLoading(state, 'service'),
        success: true,
        error: null,
        services: action.services,
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
        services: action.services,
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
    case actionTypes.USER_REQUEST_START: return userRequestStart(state, action)
    case actionTypes.USER_REQUEST_SUCCESS: return userRequestSuccess(state, action)
    case actionTypes.USER_REQUEST_ERROR: return userRequestError(state, action)
    case actionTypes.EMAIL_RECORD_START: return emailRecordStart(state, action)
    case actionTypes.EMAIL_RECORD_SUCCESS: return emailRecordSuccess(state, action)
    case actionTypes.EMAIL_RECORD_ERROR: return emailRecordError(state, action)
    case actionTypes.SESSION_RECORD_START: return sessionRecordStart(state, action)
    case actionTypes.SESSION_RECORD_SUCCESS: return sessionRecordSuccess(state, action)
    case actionTypes.SESSION_RECORD_ERROR: return sessionRecordError(state, action)
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
