import { ActionTypes } from '../../core/constants'
import { ActionCreator, errorHandler, processResponse, processData, successHandler } from '../../core/actions/base'
import SpaceApi from '../../core/utils/SpaceApi'

class SessionsActionFactory {
    signIn(data) {
        let action = new ActionCreator()
        action.setUUID()
        action.dispatch({type: ActionTypes.SEND_DATA})
        SpaceApi
            .createSession(data)
            .then(processResponse)
            .then(processData)
            .then(successHandler)
            .catch(errorHandler)
        return action.actionID()
    }
};

const SessionsActions = new SessionsActionFactory()

export default SessionsActions
