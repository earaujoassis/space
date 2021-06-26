import * as actionTypes from '@core/constants/clients';

const initialState = {
  success: undefined,
  displayToast: undefined,
  error: undefined,
  clients: undefined
};

const clientRecordStart = (state, _action) => {
  // NProgress.start();
  return state;
};

const clientRecordSuccess = (state, action) => {
  // NProgress.done();
  return Object.assign({}, state, {
    success: true,
    displayToast: false,
    error: null,
    clients: action.data.clients
  });
};

const clientRecordError = (state, action) => {
  // NProgress.done();
  return Object.assign({}, state, {
    displayToast: true,
    success: false,
    error: action.error
  });
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.CLIENT_RECORD_START: return clientRecordStart(state, action);
    case actionTypes.CLIENT_RECORD_SUCCESS: return clientRecordSuccess(state, action);
    case actionTypes.CLIENT_RECORD_ERROR: return clientRecordError(state, action);
    default: return state;
  }
};

export default reducer;
