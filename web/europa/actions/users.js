import { ActionTypes } from '../../core/constants'
import { ActionCreator, processResponse, processData, processHandlerClojure } from '../../core/actions/base'
import SpaceApi from '../../core/utils/SpaceApi'

import UserStore from '../stores/users'

const actionProxy = (name) => {
    let token = UserStore.getActionToken()
    let id = UserStore.getUserId()
    let action = new ActionCreator()

    action.setUUID()
    UserStore.associateAction(action.actionID())
    action.dispatch({type: ActionTypes.SEND_DATA})
    SpaceApi[name](id, token)
        .then(processResponse)
        .then(processData)
        .then(processHandlerClojure(action))
    return action.actionID()
}

class UsersActionFactory {
    fetchProfile() {
        return actionProxy('fetchProfile')
    }

    fetchActiveClients() {
        return actionProxy('fetchActiveClients')
    }

    adminify(key) {
        let token = UserStore.getActionToken()
        let id = UserStore.getUserId()
        let action = new ActionCreator()
        let data = new FormData()

        action.setUUID()
        UserStore.associateAction(action.actionID())
        action.dispatch({type: ActionTypes.SEND_DATA})
        data.append('application_key', key)
        SpaceApi.adminify(id, token, data)
            .then(processResponse)
            .then(processData)
            .then(processHandlerClojure(action))
            .then(() => {
                UsersActions.fetchProfile()
            })
        return action.actionID()
    }

    revokeActiveClient(key) {
        let token = UserStore.getActionToken()
        let id = UserStore.getUserId()
        let action = new ActionCreator()

        action.setUUID()
        UserStore.associateAction(action.actionID())
        action.dispatch({type: ActionTypes.SEND_DATA})
        SpaceApi['revokeActiveClient'](id, key, token)
            .then(processResponse)
            .then(processData)
            .then(processHandlerClojure(action))
            .then(() => {
                UsersActions.fetchActiveClients()
            })
        return action.actionID()
    }
}

const UsersActions = new UsersActionFactory()

export default UsersActions
