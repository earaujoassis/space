import * as actionTypes from '@core/constants/internal';

export const addLoadingContext = (context) => {
  return {
    type: actionTypes.INTERNAL_LOADING_CONTEXT_ADD,
    context
  };
};

export const removeLoadingContext = (context) => {
  return {
    type: actionTypes.INTERNAL_LOADING_CONTEXT_REMOVE,
    context
  };
};

export const loadData = () => {
  return {
    type: actionTypes.INTERNAL_LOAD_DATA
  };
};

export const errorHandler = (error) => {
  return {
    type: actionTypes.INTERNAL_ERROR_HANDLING,
    error
  };
};
