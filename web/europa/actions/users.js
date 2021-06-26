import { ActionTypes } from '../../core/constants';
import { ActionCreator, processResponse, processData, processHandler, processHandlerClojure } from '../../core/actions/base';
import SpaceApi from '../../core/utils/spaceApi';

import UserStore from '../stores/users';

const actionProxy = (name) => {
  const token = UserStore.getActionToken();
  const id = UserStore.getUserId();
  const action = new ActionCreator();

  action.setUUID();
  UserStore.associateAction(action.actionID());
  action.dispatch({ type: ActionTypes.SEND_DATA });
  return SpaceApi[name](id, token)
    .then(processResponse)
    .then(processData)
    .then(processHandlerClojure(action));
};

class UsersActionFactory {
  fetchProfile () {
    return actionProxy('fetchProfile');
  }

  fetchActiveClients () {
    return actionProxy('fetchActiveClients');
  }

  requestUpdate (data) {
    const action = new ActionCreator();
    action.setUUID();
    action.dispatch({ type: ActionTypes.SEND_DATA });
    return SpaceApi.requestUpdate(data)
      .then(processResponse)
      .then(processData)
      .then(processHandler);
  }

  adminify (key) {
    const token = UserStore.getActionToken();
    const action = new ActionCreator();
    const data = new FormData();
    data.append('application_key', key);
    data.append('user_id', UserStore.getUserId());

    action.setUUID();
    UserStore.associateAction(action.actionID());
    action.dispatch({ type: ActionTypes.SEND_DATA });
    return SpaceApi.adminify(token, data)
      .then(processResponse)
      .then(processData)
      .then(processHandlerClojure(action))
      .then(() => UsersActions.fetchProfile());
  }

  revokeActiveClient (key) {
    const token = UserStore.getActionToken();
    const id = UserStore.getUserId();
    const action = new ActionCreator();

    action.setUUID();
    UserStore.associateAction(action.actionID());
    action.dispatch({ type: ActionTypes.SEND_DATA });
    return SpaceApi.revokeActiveClient(id, key, token)
      .then(processResponse)
      .then(processData)
      .then(processHandlerClojure(action))
      .then(() => UsersActions.fetchActiveClients());
  }
}

const UsersActions = new UsersActionFactory();

export default UsersActions;
