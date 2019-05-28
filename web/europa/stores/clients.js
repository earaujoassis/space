import { ActionTypes } from '../../core/constants'
import dispatcher from '../../core/dispatcher'
import Store from '../../core/stores/base'

var _state = {}
var _actions = new Set()

class ClientStoreBase extends Store {
    constructor() {
        super()
        this.dispatchToken = dispatcher.register(function(action) {
            switch (action.type) {
            case ActionTypes.SUCCESS:
                if (_actions.has(action.actionUUID)) {
                    ClientStore.setCommons(action)
                    ClientStore.emitChange()
                }
                break

            case ActionTypes.ERROR:
                if (_actions.has(action.actionUUID)) {
                    ClientStore.setCommons(action)
                    ClientStore.emitChange()
                }
                break
            }
        })
    }

    getState() {
        return _state
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

const ClientStore = new ClientStoreBase()

export default ClientStore
