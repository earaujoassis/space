export const ActionTypes = ['SUCCESS', 'ERROR', 'SEND_DATA'].reduce((obj, str) => { obj[str] = str; return obj; }, {});
