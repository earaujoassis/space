export const ActionTypes = ['SUCCESS', 'ERROR', 'SEND_DATA'].reduce(function(obj, str){ obj[str] = str; return obj }, {})
