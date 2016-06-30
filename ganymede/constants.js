export const ActionTypes = ['SUCCESS', 'ERROR', 'SETUP_APP']
    .reduce(function(obj, str){ obj[str] = str; return obj; }, {});
