import { ActionTypes } from '../constants'
import ActionCreator from '../../core/actions/base'
import SpaceApi from '../../core/utils/SpaceApi'

class UsersActionFactory {
    signIn(data) {
        let action = new ActionCreator()
        action.setUUID()
        action.dispatch({type: ActionTypes.SEND_DATA})
        SpaceApi
            .createUser(data)
            .catch((error) => {
                let action = new ActionCreator()
                action.setUUID()
                action.dispatch({type: ActionTypes.ERROR, payload: error})
                return action
            })
            .then((response) => {
                let action = new ActionCreator()
                action.setUUID()
                action.dispatch({type: ActionTypes.SUCCESS, payload: response})
                return action
            })
        return action
    }
}

const UsersActions = new UsersActionFactory()

export default UsersActions
