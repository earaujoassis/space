/* eslint-disable quote-props */
import fetch from '@core/actions/fetch';
import { addLoadingContext, removeLoadingContext, errorHandler } from '@core/actions/internal';
import { userRecordStart, userRecordSuccess, userRecordError } from '@core/actions/users';

const CONTEXT = 'users';

export const requestUserUpdate = (data) => {
  return dispatch => {
    dispatch(userRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.post('users/update/request', data)
      .then(response => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordSuccess(response.data));
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};

export const adminify = (token, data) => {
  return dispatch => {
    dispatch(userRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.patch('users/update/adminify', data, { headers: { 'Authorization': `Bearer ${token}` } })
      .then(response => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordSuccess(response.data));
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};

export const fetchProfile = (token, id) => {
  return dispatch => {
    dispatch(userRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.get(`users/${id}/profile`, { headers: { 'Authorization': `Bearer ${token}` } })
      .then(response => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordSuccess(response.data));
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};

export const fetchActiveClients = (token, id) => {
  return dispatch => {
    dispatch(userRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.get(`users/${id}/clients`, { headers: { 'Authorization': `Bearer ${token}` } })
      .then(response => {
        dispatch(userRecordSuccess(response.data));
        dispatch(fetchProfile(token, id));
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};

export const revokeActiveClient = (token, id, key) => {
  return dispatch => {
    dispatch(userRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.delete(`users/${id}/clients/${key}/revoke`, { headers: { 'Authorization': `Bearer ${token}` } })
      .then(() => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(fetchActiveClients(token, id));
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(userRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};
