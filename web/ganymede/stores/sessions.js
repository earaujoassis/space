import { ActionTypes } from '../../core/constants';
import dispatcher from '../../core/dispatcher';
import Store from '../../core/stores/base';

const _state = {};

class SessionStoreBase extends Store {
  constructor () {
    super();
    this.dispatchToken = dispatcher.register(function (action) {
      switch (action.type) {
        case ActionTypes.SUCCESS:
          SessionStore.setCommons(action);
          SessionStore.emitChange();
          break;

        case ActionTypes.ERROR:
          SessionStore.setCommons(action);
          SessionStore.emitChange();
          break;
      }
    });
  }

  getState () {
    return _state;
  }

  setCommons (action) {
    _state.error = action.error || null;
    _state.payload = action.payload || null;
    _state.type = action.type || null;
    _state.actionUUID = action.actionUUID || action.payload.actionUUID || null;
  }
}

const SessionStore = new SessionStoreBase();

export default SessionStore;
