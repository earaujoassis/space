import { ActionTypes } from '../constants'
import { ActionCreator, errorHandler, processResponse, processData, successHandler } from '../../core/actions/base'
import SpaceApi from '../../core/utils/SpaceApi'

class UsersActionFactory {
    signIn(data) {
        let action = new ActionCreator()
        action.setUUID()
        action.dispatch({type: ActionTypes.SEND_DATA})
        SpaceApi
            .createUser(data)
            .then(processResponse)
            .then(processData)
            .then(successHandler)
            .catch(errorHandler)
        return action
    }
}

const UsersActions = new UsersActionFactory()

export default UsersActions
