import { ActionTypes } from '../../core/constants';
import { ActionCreator, processResponse, processData, processHandlerClojure } from '../../core/actions/base';
import SpaceApi from '../../core/utils/spaceApi';

import UserStore from '../stores/users';

class UsersActionFactory {
  updatePassword (data) {
    const action = new ActionCreator();
    action.setUUID();
    UserStore.associateAction(action.actionID());
    action.dispatch({ type: ActionTypes.SEND_DATA });
    SpaceApi
      .updatePassword(data)
      .then(processResponse)
      .then(processData)
      .then(processHandlerClojure(action));
    return action;
  }
}

const UsersActions = new UsersActionFactory();

export default UsersActions;
