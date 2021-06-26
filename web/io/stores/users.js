import { ActionTypes } from '../../core/constants';
import dispatcher from '../../core/dispatcher';
import Store from '../../core/stores/base';

const _state = {};

class UserStoreBase extends Store {
  constructor () {
    super();
    this.dispatchToken = dispatcher.register(function (action) {
      switch (action.type) {
        case ActionTypes.SUCCESS:
          UserStore.setCommons(action);
          UserStore.emitChange();
          break;

        case ActionTypes.ERROR:
          UserStore.setCommons(action);
          UserStore.emitChange();
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

const UserStore = new UserStoreBase();

export default UserStore;
