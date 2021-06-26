import { ActionTypes } from '../../core/constants';
import { ActionCreator, processResponse, processData, processHandler } from '../../core/actions/base';
import SpaceApi from '../../core/utils/spaceApi';

class SessionsActionFactory {
  signIn (data) {
    const action = new ActionCreator();
    action.setUUID();
    action.dispatch({ type: ActionTypes.SEND_DATA });
    SpaceApi
      .createSession(data)
      .then(processResponse)
      .then(processData)
      .then(processHandler);
    return action.actionID();
  }

  requestMagicLink (data) {
    const action = new ActionCreator();
    action.setUUID();
    action.dispatch({ type: ActionTypes.SEND_DATA });
    SpaceApi
      .createMagicSession(data)
      .then(processResponse)
      .then(processData)
      .then(processHandler);
    return action.actionID();
  }

  requestUpdate (data) {
    const action = new ActionCreator();
    action.setUUID();
    action.dispatch({ type: ActionTypes.SEND_DATA });
    SpaceApi
      .requestUpdate(data)
      .then(processResponse)
      .then(processData)
      .then(processHandler);
    return action.actionID();
  }
}

const SessionsActions = new SessionsActionFactory();

export default SessionsActions;
