import Actions from './Actions';

const status = response => {
  if (response.status >= 200 && response.status < 300) {
    return Promise.resolve(response);
  }
  else if (response.status === 400) {
    return Promise.reject(response.json());
  }
  else {
    return Promise.reject(response.status);
  }
}

const json = response => response.json()

export class API {

  constructor(apiURL) {
    this.url = apiURL;
    this.token = null;
  }

  setToken(token) {
    this.token = token;
  }

  makeRequest(path, method='GET') {
    let context = {
      method: method,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'X-LP-Token': this.token,
      },
    };
    return fetch(`${this.url}/${path}`, context);
  }

  getSongs() {
    return this.makeRequest('songs/')
      .then(status)
      .then(json)
      .then((data) => {
        Actions.setSongs(data);
      });
  }

  getSong(id) {
    return this.makeRequest(`songs/${id}/`)
      .then(status)
      .then(json)
      .then((song) => {
        Actions.updateSong(song);
      });
  }

};

export default new API('/api');
