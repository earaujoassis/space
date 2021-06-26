import * as actionTypes from '@core/constants/users';

const initialState = {
  success: undefined,
  displayToast: undefined,
  error: undefined,
  user: undefined
};

const userRecordStart = (state, _action) => {
  // NProgress.start();
  return state;
};

const userRecordSuccess = (state, action) => {
  // NProgress.done();
  const nextState = {
    success: true,
    displayToast: false,
    error: null
  };
  if (action.data.user) {
    nextState.user = action.data.user;
  }
  if (action.data.clients) {
    nextState.clients = action.data.clients;
  }
  return Object.assign({}, state, nextState);
};

const userRecordError = (state, action) => {
  // NProgress.done();
  return Object.assign({}, state, {
    displayToast: true,
    success: false,
    error: action.error
  });
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.USER_RECORD_START: return userRecordStart(state, action);
    case actionTypes.USER_RECORD_SUCCESS: return userRecordSuccess(state, action);
    case actionTypes.USER_RECORD_ERROR: return userRecordError(state, action);
    default: return state;
  }
};

export default reducer;
