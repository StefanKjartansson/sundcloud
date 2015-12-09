'use strict';

import _ from 'underscore';
import BaseStore from './BaseStore';
import {SET_SONGS, SET_SONG} from './Constants';


class Store extends BaseStore {

  constructor() {
    super();
    this.clear();
    this.subscribe(() => this._registerToActions.bind(this));
  }

  _registerToActions(action) {
    if (action.actionType === SET_SONGS) {
      this._data = action.data;
      this.emitChange();
    }
    if (action.actionType === SET_SONG) {
      let idx = _.findIndex(this._data, {id: action.data.id});
      this._data[idx] = action.data;
      this.emitChange();
    }
  }

  getById(id) {
    return _.findWhere(this._data, {id: id});
  }

  clear() {
    this._data = null;
  }

  get data() {
    return this._data;
  }

}

export default new Store();
