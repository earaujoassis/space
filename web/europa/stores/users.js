import { ActionTypes } from '../../core/constants'
import dispatcher from '../../core/dispatcher'
import Store from '../../core/stores/base'

var _state = {}
var _setupData = {}
var _actions = new Set()

class UserStoreBase extends Store {
    constructor() {
        super()
        this.dispatchToken = dispatcher.register(function(action) {
            switch (action.type) {
            case ActionTypes.SUCCESS:
                if (_actions.has(action.actionUUID)) {
                    UserStore.setCommons(action)
                    UserStore.emitChange()
                }
                break

            case ActionTypes.ERROR:
                if (_actions.has(action.actionUUID)) {
                    UserStore.setCommons(action)
                    UserStore.emitChange()
                }
                break
            }
        })
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

    isFeatureGateActive(key) {
        return _setupData['feature.gates'] && _setupData['feature.gates'][key]
    }

    isCurrentUserAdmin() {
        return _state.payload && _state.payload.user && _state.payload.user.is_admin !== undefined ?
            _state.payload.user.is_admin : _setupData['user_is_admin'] === true
    }

    loadData() {
        if (document.getElementById('data')) {
            _setupData = JSON.parse(document.getElementById('data').innerHTML)
        }
    }

    setCommons(action) {
        _state.error = action.error || null
        _state.payload = action.payload || null
        _state.type = action.type || null
        _state.actionUUID = action.actionUUID || action.payload.actionUUID || null
    }

    associateAction(token) {
        _actions.add(token)
    }
}

const UserStore = new UserStoreBase()

export default UserStore
