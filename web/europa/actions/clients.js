import { ActionTypes } from '../../core/constants'
import { ActionCreator, processResponse, processData, processHandlerClojure } from '../../core/actions/base'
import SpaceApi from '../../core/utils/SpaceApi'

import ClientStore from '../stores/clients'
import UserStore from '../stores/users'

class ClientsActionFactory {
    createClient(data) {
        let token = UserStore.getActionToken()
        let action = new ActionCreator()

        action.setUUID()
        ClientStore.associateAction(action.actionID())
        action.dispatch({type: ActionTypes.SEND_DATA})
        SpaceApi['createClient'](token, data)
            .then(processResponse)
            .then(processData)
            .then(processHandlerClojure(action)).then(() => {
                ClientsActions.fetchClients()
            })
        return action.actionID()
    }

    fetchClients() {
        let token = UserStore.getActionToken()
        let action = new ActionCreator()

        action.setUUID()
        ClientStore.associateAction(action.actionID())
        action.dispatch({type: ActionTypes.SEND_DATA})
        SpaceApi['fetchClients'](token)
            .then(processResponse)
            .then(processData)
            .then(processHandlerClojure(action))
        return action.actionID()
    }
}

const ClientsActions = new ClientsActionFactory()

export default ClientsActions
