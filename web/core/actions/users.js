import * as actionTypes from '@core/constants/users';

export const userRecordStart = () => {
  return {
    type: actionTypes.USER_RECORD_START
  };
};

export const userRecordSuccess = (data) => {
  return {
    type: actionTypes.USER_RECORD_SUCCESS,
    data
  };
};

export const userRecordError = (error) => {
  return {
    type: actionTypes.USER_RECORD_ERROR,
    error: error
  };
};
