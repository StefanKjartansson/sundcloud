import AppDispatcher from './Dispatcher';
import Service from './Service';
import {SET_SONGS} from './Constants';

export default {
  setSongs: data => {
    AppDispatcher.dispatch({
      actionType: SET_SONGS,
      data: data,
    });
  },
  refresh: () => {
    Service.getSongs();
  },
  updateSong: song => {
    AppDispatcher.dispatch({
      actionType: SET_SONG,
      data: song,
    });
  },
};
