import React, { createContext, useContext, useReducer } from 'react'

import { createSession, requestUpdate, createMagicSession } from './fetch'

const ACTIONS = {
  FETCH_START: 'FETCH_START',
  FETCH_SUCCESS: 'FETCH_SUCCESS',
  FETCH_ERROR: 'FETCH_ERROR',
  CLEAR_ERROR: 'CLEAR_ERROR',
  RESET: 'RESET',
}

const initialState = {
  payload: null,
  loading: false,
  success: false,
  error: null,
}

const appReducer = (state, action) => {
  switch (action.type) {
    case ACTIONS.FETCH_START:
      return {
        ...state,
        loading: true,
        success: false,
        payload: null,
        error: null,
      }

    case ACTIONS.FETCH_SUCCESS:
      return {
        ...state,
        loading: false,
        success: true,
        payload: action.payload,
        error: null,
      }

    case ACTIONS.FETCH_ERROR:
      return {
        ...state,
        loading: false,
        success: false,
        payload: null,
        error: action.payload,
      }

    case ACTIONS.CLEAR_ERROR:
      return {
        ...state,
        error: null,
      }

    case ACTIONS.RESET:
      return initialState

    default:
      return state
  }
}

const AppContext = createContext()

export const AppProvider = ({ children }) => {
  const [state, dispatch] = useReducer(appReducer, initialState)

  const actions = {
    signIn: async data => {
      dispatch({ type: ACTIONS.FETCH_START })

      try {
        const response = await createSession(data)

        if (!response.ok) {
          throw {
            response: response,
            error: new Error(
              `Error ${response.status}: ${response.statusText}`
            ),
          }
        }

        const payload = await response.json()
        dispatch({
          type: ACTIONS.FETCH_SUCCESS,
          payload: payload,
        })
      } catch ({ response, error }) {
        if (response.status === 400) {
          const payload = await response.json()
          dispatch({
            type: ACTIONS.FETCH_ERROR,
            payload: payload,
          })
        } else {
          dispatch({
            type: ACTIONS.FETCH_ERROR,
            payload: error.message,
          })
        }
      }
    },

    requestMagicLink: async data => {
      dispatch({ type: ACTIONS.FETCH_START })

      try {
        const response = await createMagicSession(data)

        if (!response.ok) {
          throw new Error(`Error ${response.status}: ${response.statusText}`)
        }

        dispatch({
          type: ACTIONS.FETCH_SUCCESS,
          payload: null,
        })
      } catch (error) {
        dispatch({
          type: ACTIONS.FETCH_ERROR,
          payload: error.message,
        })
      }
    },

    requestUpdate: async data => {
      dispatch({ type: ACTIONS.FETCH_START })

      try {
        const response = await requestUpdate(data)

        if (!response.ok) {
          throw new Error(`Error ${response.status}: ${response.statusText}`)
        }

        dispatch({
          type: ACTIONS.FETCH_SUCCESS,
          payload: null,
        })
      } catch (error) {
        dispatch({
          type: ACTIONS.FETCH_ERROR,
          payload: error.message,
        })
      }
    },

    clearError: () => {
      dispatch({ type: ACTIONS.CLEAR_ERROR })
    },

    reset: () => {
      dispatch({ type: ACTIONS.RESET })
    },
  }

  return (
    <AppContext.Provider value={{ state, actions }}>
      {children}
    </AppContext.Provider>
  )
}

export const useApp = () => {
  const context = useContext(AppContext)
  if (!context) {
    throw new Error('useApp must be called from the AppProvider')
  }
  return context
}
