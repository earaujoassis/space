import * as actionTypes from '@core/constants/clients';

export const clientRecordStart = () => {
  return {
    type: actionTypes.CLIENT_RECORD_START
  };
};

export const clientRecordSuccess = (data) => {
  return {
    type: actionTypes.CLIENT_RECORD_SUCCESS,
    data
  };
};

export const clientRecordError = (error) => {
  return {
    type: actionTypes.CLIENT_RECORD_ERROR,
    error: error
  };
};
