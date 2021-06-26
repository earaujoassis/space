import * as actionTypes from '@core/constants/internal';

const initialState = {
  loading: [],
  data: {},
  tokenExpired: false
};

const addLoadingContext = (state, action) => {
  const loading = JSON.parse(JSON.stringify(state.loading));
  loading.push(action.context);
  return Object.assign({}, state, { loading });
};

const removeLoadingContext = (state, action) => {
  const loading = JSON.parse(JSON.stringify(state.loading));
  return Object.assign({}, state, { loading: loading.filter(element => element !== action.context) });
};

const loadData = (state) => {
  if (document.getElementById('data')) {
    const data = JSON.parse(document.getElementById('data').innerHTML);
    if (data['feature.gates']) {
      const features = data['feature.gates'];
      delete data['feature.gates'];
      data.features = Object.entries(features).filter(([k, v]) => v === true).map(([k, v]) => k);
    }
    return Object.assign({}, state, { data });
  }

  return state;
};

const errorHandler = (state, action) => {
  if (action.error.response && action.error.response.status === 401) {
    return Object.assign({}, state, { tokenExpired: true });
  }

  return state;
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.INTERNAL_LOADING_CONTEXT_ADD: return addLoadingContext(state, action);
    case actionTypes.INTERNAL_LOADING_CONTEXT_REMOVE: return removeLoadingContext(state, action);
    case actionTypes.INTERNAL_LOAD_DATA: return loadData(state, action);
    case actionTypes.INTERNAL_ERROR_HANDLING: return errorHandler(state, action);
    default: return state;
  }
};

export default reducer;
