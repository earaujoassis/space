import { ActionTypes } from '../../core/constants'
import { ActionCreator, processResponse, processData, processHandler } from '../../core/actions/base'
import SpaceApi from '../../core/utils/space-api'

class UsersActionFactory {
    signUp(data) {
        let action = new ActionCreator()
        action.setUUID()
        action.dispatch({type: ActionTypes.SEND_DATA})
        SpaceApi
            .createUser(data)
            .then(processResponse)
            .then(processData)
            .then(processHandler)
        return action
    }
}

const UsersActions = new UsersActionFactory()

export default UsersActions
