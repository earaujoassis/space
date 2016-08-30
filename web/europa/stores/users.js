import { ActionTypes } from '../../core/constants'
import dispatcher from '../../core/dispatcher'
import Store from '../../core/stores/base'

var _state = {}
var _setupData = {}

class UserStoreBase extends Store {
    constructor() {
        super()
        this.dispatchToken = dispatcher.register(function(action) {
            switch (action.type) {
                case ActionTypes.SUCCESS:
                    UserStore.setCommons(action)
                    UserStore.emitChange()
                    break

                case ActionTypes.ERROR:
                    UserStore.setCommons(action)
                    UserStore.emitChange()
                    break
            }
        });
    }

    getState() {
        return _state
    }

    getActionToken() {
        return _setupData['action_token']
    }

    getUserId() {
        return _setupData['user_id']
    }

    loadData() {
        if (document.getElementById("data")) {
            _setupData = JSON.parse(document.getElementById("data").innerHTML)
        }
    }

    setCommons(action) {
        _state.error = action.error || null
        _state.payload = action.payload || null
        _state.type = action.type || null
        _state.actionUUID = action.actionUUID || action.payload.actionUUID || null
    }
}

const UserStore = new UserStoreBase();

export default UserStore;
