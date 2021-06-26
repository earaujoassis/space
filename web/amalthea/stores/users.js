import { ActionTypes } from '../../core/constants';
import dispatcher from '../../core/dispatcher';
import Store from '../../core/stores/base';

const _state = {};
let _setupData = {};
const _actions = new Set();

class UserStoreBase extends Store {
  constructor () {
    super();
    this.dispatchToken = dispatcher.register(function (action) {
      switch (action.type) {
        case ActionTypes.SUCCESS:
          if (_actions.has(action.actionUUID)) {
            UserStore.setCommons(action);
            UserStore.emitChange();
          }
          break;

        case ActionTypes.ERROR:
          if (_actions.has(action.actionUUID)) {
            UserStore.setCommons(action);
            UserStore.emitChange();
          }
          break;
      }
    });
  }

  getState () {
    return _state;
  }

  getActionToken () {
    return _setupData.action_token;
  }

  loadData () {
    if (document.getElementById('data')) {
      _setupData = JSON.parse(document.getElementById('data').innerHTML);
    }
  }

  setCommons (action) {
    _state.error = action.error || null;
    _state.payload = action.payload || null;
    _state.type = action.type || null;
    _state.actionUUID = action.actionUUID || action.payload.actionUUID || null;
  }

  associateAction (token) {
    _actions.add(token);
  }
}

const UserStore = new UserStoreBase();

export default UserStore;
