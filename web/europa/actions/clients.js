/* eslint-disable quote-props */
import fetch from '@core/actions/fetch';
import { addLoadingContext, removeLoadingContext, errorHandler } from '@core/actions/internal';
import { clientRecordStart, clientRecordSuccess, clientRecordError } from '@core/actions/clients';

const CONTEXT = 'clients';

export const fetchClients = (token) => {
  return dispatch => {
    dispatch(clientRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.get('clients/', { headers: { 'Authorization': `Bearer ${token}` } })
      .then(response => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(clientRecordSuccess(response.data));
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(clientRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};

export const createClient = (token, data, callbackSuccess = () => {}) => {
  return dispatch => {
    dispatch(clientRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.post('clients/create', data, { headers: { 'Authorization': `Bearer ${token}` } })
      .then(() => {
        dispatch(fetchClients(token));
        callbackSuccess();
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(clientRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};

export const updateClient = (token, id, data, callbackSuccess = () => {}) => {
  return dispatch => {
    dispatch(clientRecordStart());
    dispatch(addLoadingContext(CONTEXT));
    fetch.patch(`clients/${id}/profile`, data, { headers: { 'Authorization': `Bearer ${token}` } })
      .then(() => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(fetchClients(token));
        callbackSuccess();
      })
      .catch(error => {
        dispatch(removeLoadingContext(CONTEXT));
        dispatch(clientRecordError(error));
        dispatch(errorHandler(error));
      });
  };
};
