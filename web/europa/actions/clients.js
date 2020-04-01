import { ActionTypes } from '../../core/constants'
import { ActionCreator, processResponse, processData, processHandlerClojure } from '../../core/actions/base'
import SpaceApi from '../../core/utils/spaceApi'

import ClientStore from '../stores/clients'
import UserStore from '../stores/users'

class ClientsActionFactory {
    createClient(data) {
        let token = UserStore.getActionToken()
        let action = new ActionCreator()

        action.setUUID()
        ClientStore.associateAction(action.actionID())
        action.dispatch({type: ActionTypes.SEND_DATA})
        return SpaceApi['createClient'](token, data)
            .then(processResponse)
            .then(processData)
            .then(processHandlerClojure(action))
    }

    updateClient(id, data) {
        let token = UserStore.getActionToken()
        let action = new ActionCreator()

        action.setUUID()
        ClientStore.associateAction(action.actionID())
        action.dispatch({type: ActionTypes.SEND_DATA})
        return SpaceApi['updateClient'](id, token, data)
            .then(processResponse)
            .then(processData)
            .then(processHandlerClojure(action))
    }

    fetchClients() {
        let token = UserStore.getActionToken()
        let action = new ActionCreator()

        action.setUUID()
        ClientStore.associateAction(action.actionID())
        action.dispatch({type: ActionTypes.SEND_DATA})
        return SpaceApi['fetchClients'](token)
            .then(processResponse)
            .then(processData)
            .then(processHandlerClojure(action))
    }
}

const ClientsActions = new ClientsActionFactory()

export default ClientsActions
